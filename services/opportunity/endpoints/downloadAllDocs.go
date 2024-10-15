package endpoints

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

type DownloadAllDocumentsRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *DownloadAllDocumentsRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) DownloadAllDocuments(ctx context.Context, id string) (string, response.ErrorResponse) {
	byteBaseBaseUri := config.AppConfig.ByteBaseInternalDownloadBaseUri
	bytebaseInternalAuthKey := config.AppConfig.BytebaseInternalAuthKey

	Id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(Id)

	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", theID).Preload("Roles").First(&user).Error
	if err != nil {
		s.logger.With(ctx).Error("Failed to download file", err)
		return "", response.ErrorBadRequest(nil, "Failed to download file")
	}

	var opportunity models.Opportunity
	err = s.db.Preload("Documents").First(&opportunity, "id = ?", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred while finding the opportunity")
	}

	if len(opportunity.Documents) == 0 {
		s.logger.With(ctx).Error("No documents found for this opportunity")
		return "", response.ErrorNotFound(nil, "No documents found for this opportunity")
	}

	savedFilePaths := []string{}
	for _, doc := range opportunity.Documents {
		if !policy.CanAccessDocument(ctx, *doc, user) && !s.UserBoughtOpportunity(ctx, user, opportunity) {
			continue
		}

		downloadLink, err := url.JoinPath(byteBaseBaseUri, "internal/files/download", doc.FilePath)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}

		fmt.Println("downloadLink: ", downloadLink)
		client := &http.Client{}
		req, err := http.NewRequest("GET", downloadLink, nil)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}
		req.Header.Set("Authorization", bytebaseInternalAuthKey)

		resp, err := client.Do(req)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			s.logger.With(ctx).Errorf("Failed to download file with status code %d and message %s", resp.StatusCode, resp.Status)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}
		defer resp.Body.Close()

		filePath := path.Join(os.TempDir(), path.Base(doc.FilePath)) // Use base to prevent directory traversal issues
		out, err := os.Create(filePath)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}

		savedFilePaths = append(savedFilePaths, filePath)
	}

	if len(savedFilePaths) == 0 {
		s.logger.With(ctx).Error("No documents found for this opportunity for you to download or you have not bought this opportunity")
		return "", response.ErrorNotFound(nil, "No documents found for you to download in this opportunity or you have not bought this opportunity")
	}

	zipFilePath := path.Join(os.TempDir(), opportunity.SolicitationNumber+".zip")
	fmt.Println("zipFilePath: ", zipFilePath)
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.ErrorBadRequest(nil, "Failed to download file")
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Regular expression to match UUID at the start of the file name
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F-]+@`)

	for _, filePath := range savedFilePaths {
		file, err := os.Open(filePath)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}

		// Remove UUID from the file name
		baseName := path.Base(filePath)
		newName := uuidRegex.ReplaceAllString(baseName, "")

		header.Name = newName // Only use the base name of the file
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}

		_, err = io.Copy(writer, file)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorBadRequest(nil, "Failed to download file")
		}
	}

	// Make sure to close the zip writer to flush the contents to the disk before returning
	if err := zipWriter.Close(); err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.ErrorBadRequest(nil, "Failed to download file")
	}

	return zipFilePath, response.ErrorResponse{}
}

func (s *service) UserBoughtOpportunity(ctx context.Context, user models.User, opportunity models.Opportunity) bool {
	var orders []models.Order
	err := s.db.Model(models.Order{}).Where("user_id = ? AND opportunity_id = ? AND paid_at IS NOT NULL AND refunded_ait IS NULL", user.ID, opportunity.ID).Find(&orders).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Checking Prev Orders Failed: %v", err)
		return false
	}

	return len(orders) > 0
}
