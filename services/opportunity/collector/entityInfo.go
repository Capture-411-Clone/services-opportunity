/*
Package entityInfo provides functions to interact with the SAM (System for Award Management) API to fetch and process entity information. It includes functionalities to construct API URLs, fetch data from the API, process pagination, and check the registration status of entities.
*/
package collector

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/Capture-411/services-opportunity/config"
)

const (
	baseURL    = "https://api.sam.gov/entity-information/v3/entities?api_key=%s"
	defaultURL = "https://api.sam.gov/entity-information/v3/entities?api_key=%s&registrationStatus=A&includeSections=integrityInformation,All&proceedingsData=YES"
)

var (
	apiKey      = config.AppConfig.SAMApiKey
	searchValRe = regexp.MustCompile(`[^a-zA-Z0-9-_ ]`)
	// parameters  = []string{
	// 	"legalBusinessName",
	// 	"dbaName",
	// 	"debtSubjectToOffset",
	// 	"exclusionStatusFlag",
	// 	"registrationDate",
	// 	"activationDate",
	// 	"updateDate",
	// 	"registrationExpirationDate",
	// 	"ueiCreationDate",
	// 	"purposeOfRegistrationCode",
	// 	"purposeOfRegistrationDesc",
	// 	"physicalAddressCity",
	// 	"physicalAddressCongressionalDistrict",
	// 	"physicalAddressCountryCode",
	// 	"physicalAddressProvinceOrStateCode",
	// 	"physicalAddressZipPostalCode",
	// 	"entityStructureCode",
	// 	"entityStructureDesc",
	// 	"organizationStructureCode",
	// 	"organizationStructureDesc",
	// 	"businessTypeCode",
	// 	"businessTypeDesc",
	// 	"sbaBusinessTypeCode",
	// 	"sbaBusinessTypeDesc",
	// 	"primaryNaics",
	// 	"naicsCode",
	// 	"naicsDesc",
	// 	"naicsLimitedSB",
	// 	"pscCode",
	// 	"pscDesc",
	// 	"stateOfIncorporationCode",
	// 	"stateOfIncorporationDesc",
	// 	"countryOfIncorporationCode",
	// 	"countryOfIncorporationDesc",
	// 	"servedDisasterStateCode",
	// 	"servedDisasterStateName",
	// 	"servedDisasterCountyCode",
	// 	"servedDisasterCountyName",
	// 	"servedDisasterMSA",
	// 	"proceedingsData",
	// 	"includeSections",
	// 	"format",
	// 	"emailId",
	// }
)

// Entity represents the response from the SAM API.
type Entity struct {
	TotalRecords int           `json:"totalRecords"`
	EntityData   []*EntityData `json:"entityData"`
	SelfLink     string        `json:"selfLink"`
	NextLink     string        `json:"nextLink"`
}

// EntityData holds the entity registration and core data
type EntityData struct {
	EntityRegistration *EntityRegistration `json:"entityRegistration"`
	CoreData           *CoreData           `json:"coreData"`
}

// EntityRegistration contains the registration details of an entity.
type EntityRegistration struct {
	SamRegistered             string `json:"samRegistered"`
	UeiSAM                    string `json:"ueiSAM"`
	CageCode                  string `json:"cageCode"`
	LegalBusinessName         string `json:"legalBusinessName"`
	PurposeOfRegistrationDesc string `json:"purposeOfRegistrationDesc"`
	RegisteredStatus          string `json:"registrationStatus"`
}

// CoreData is currently an empty struct, reserved for future use.
type CoreData struct {
}

// sanitizeInput sanitizes the input string by removing any characters that are not alphanumeric, hyphens, or underscores.
func sanitizeInput(input string) string {
	return searchValRe.ReplaceAllString(input, "")
}

// ConstructURL constructs the full API URL with the provided base URL, API key, and a map of parameters.
func ConstructURL(searchValue string, params map[string]string) (string, error) {

	if len(params) == 0 {
		parsedURL, err := url.Parse(fmt.Sprintf(defaultURL, apiKey))
		// Add the query parameters
		queryParams := parsedURL.Query()
		if searchValue != "" {
			queryParams.Set("q", sanitizeInput(searchValue))
		}
		parsedURL.RawQuery = queryParams.Encode()
		if err != nil {
			return "", err
		}
		return parsedURL.String(), nil
	} else {
		// Parse the base URL
		parsedURL, err := url.Parse(fmt.Sprintf(baseURL, apiKey))
		if err != nil {
			return "", err
		}
		queryParams := parsedURL.Query()
		for key, value := range params {
			queryParams.Set(key, value)
		}

		// Add the query parameters
		if searchValue != "" {
			queryParams.Set("q", sanitizeInput(searchValue))
		}
		parsedURL.RawQuery = queryParams.Encode()
		return parsedURL.String(), nil
	}

}

// fetchEntityDataFromURL fetches data from the given URL and returns the response as a byte
func fetchEntityDataFromURL(url string) ([]byte, error) {
	// Make the HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the HTTP status is in the 200-299 range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return data, nil
}

// unmarshalJson unmarshals the JSON data into an Entity struct.
func unmarshalJson(data []byte) (Entity, error) {

	var entity Entity
	if err := json.Unmarshal(data, &entity); err != nil {
		log.Printf("Error unmarshalling XML: %v", err)
		return Entity{}, fmt.Errorf("error unmarshalling XML: %v", err)
	}

	return entity, nil
}

// getEntityPageURLs generates a list of URLs for pagination based on the total number of records.
// func getEntityPageURLs(selfLink string, totalRecord int) ([]string, error) {
// 	// Parse the last URL to get the query parameters
// 	var links []string

// 	if totalRecord < 11 {
// 		links = append(links, selfLink)
// 		return links, nil
// 	}
// 	parsedURL, err := url.Parse(selfLink)
// 	if err != nil {
// 		return links, err
// 	}

// 	// Extract query parameters
// 	queryParams := parsedURL.Query()

// 	// Extract and print the original 'page' parameter value
// 	pageParam := queryParams.Get("page")
// 	if pageParam == "" {
// 		return links, err
// 	}

// 	page := 0
// 	pages := (totalRecord + 9) / 10
// 	for i := 1; i < pages; i++ {
// 		page += 1
// 		queryParams.Set("page", strconv.Itoa(page))

// 		// Reconstruct the URL with the updated query parameters
// 		parsedURL.RawQuery = queryParams.Encode()
// 		newURL := parsedURL.String()
// 		links = append(links, newURL)
// 	}
// 	return links, nil
// }

// processAllEntityPages processes multiple pages of data, logging entry details, and returning a list of Entity structs.
// func processAllEntityPages(entityURL string) ([]Entity, error) {
// 	pagesEntity := []Entity{}

// 	log.Println("Fetching URL:", entityURL)

// 	data, err := fetchEntityDataFromURL(entityURL)
// 	if err != nil {
// 		log.Printf("Error fetching data from URL %s: %v", entityURL, err)
// 		return []Entity{}, err
// 	}

// 	log.Println("Data fetched successfully, processing...")
// 	firstPageEntity, err := unmarshalJson(data)
// 	if err != nil {
// 		return []Entity{}, err
// 	}

// 	pagesEntity = append(pagesEntity, firstPageEntity)

// 	urls, err := getEntityPageURLs(firstPageEntity.SelfLink, firstPageEntity.TotalRecords)
// 	if err != nil {
// 		return []Entity{}, err
// 	}

// 	// For pagination
// 	var wg sync.WaitGroup
// 	resultChan := make(chan Entity)
// 	errorChan := make(chan error, len(urls))

// 	for _, url := range urls {
// 		wg.Add(1)
// 		go func(url string) {
// 			defer wg.Done()
// 			data, err := fetchEntityDataFromURL(url)
// 			if err != nil {
// 				errorChan <- fmt.Errorf("error fetching data from URL %s: %v", url, err)
// 				return
// 			}
// 			pagesEntity, err := unmarshalJson(data)
// 			if err != nil {
// 				errorChan <- fmt.Errorf("error processing data from URL %s: %v", url, err)
// 				return
// 			}
// 			resultChan <- pagesEntity
// 		}(url)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 		close(errorChan)
// 	}()

// 	for feed := range resultChan {
// 		pagesEntity = append(pagesEntity, feed)
// 	}

// 	if len(errorChan) > 0 {
// 		for err := range errorChan {
// 			log.Println(err)
// 		}
// 		return []Entity{}, fmt.Errorf("errors occurred during page processing")
// 	}

// 	log.Println("Atom feed processed successfully.")

// 	return pagesEntity, nil
// }

// IsRegisterActivated checks if the entity's registration is active based on the search value and parameters.
func IsRegisterActivated(searchValue string, params map[string]string) (bool, error) {
	constructedUrl, err := ConstructURL(searchValue, params)
	if err != nil {
		return false, err
	}
	entityData, err := fetchEntityDataFromURL(constructedUrl)
	if err != nil {
		return false, err
	}
	entity, err := unmarshalJson(entityData)
	if err != nil {
		return false, err
	}

	if len(entity.EntityData) == 0 {
		return false, nil
	}
	return entity.EntityData[0].EntityRegistration.RegisteredStatus == "Active", nil
}

// EntityInfo fetches data for a single entity and returns the value of the specified field from the specified section
func EntityInfo(searchValue string, params map[string]string, sectionName, fieldName string) (string, error) {
	url, err := ConstructURL(searchValue, params)
	if err != nil {
		return "", err
	}
	res, err := fetchEntityDataFromURL(url)
	if err != nil {
		return "", err
	}
	entity, err := unmarshalJson(res)
	if err != nil {
		return "", err
	}

	if len(entity.EntityData) == 0 {
		return "", fmt.Errorf("no entity data found")
	}

	// Function to recursively dereference pointers and find fields
	var findField func(reflect.Value, []string) (reflect.Value, error)
	findField = func(v reflect.Value, path []string) (reflect.Value, error) {
		for v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return reflect.Value{}, fmt.Errorf("nil pointer encountered")
			}
			v = v.Elem()
		}

		if len(path) == 0 {
			return v, nil
		}

		if v.Kind() != reflect.Struct {
			return reflect.Value{}, fmt.Errorf("value is not a struct")
		}

		field := v.FieldByName(path[0])
		if !field.IsValid() {
			return reflect.Value{}, fmt.Errorf("field %s not found", path[0])
		}

		return findField(field, path[1:])
	}

	// Split the section and field names
	path := append(strings.Split(sectionName, "."), fieldName)

	value, err := findField(reflect.ValueOf(entity.EntityData[0]), path)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", value.Interface()), nil
}

// SAMPLE USAGE
//
// func main() {
// 	ur := "https://api.sam.gov/entity-information/v3/entities?api_key=REPLACE_WITH_API_KEY&includeSections=integrityInformation%2CAll&proceedingsData=YES&page=0&size=10"
// 	fmt.Println(getEntityPageURLs(ur, 565))

// 	fmt.Println(ConstructURL("sert", map[string]string{}))

// 	params := map[string]string{
// 		"samRegistered":   "Yes",
// 		"includeSections": "integrityInformation,All",
// 	}

// 	fmt.Println(ConstructURL("MIGUEL", params))

// 	fmt.Println(IsRegisterActivated("MIGUEL", params))

// 	fmt.Println(EntityInfo("MIGUEL", params, "EntityRegistration", "LegalBusinessName"))
// }
