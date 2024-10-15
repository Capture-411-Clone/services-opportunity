package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/constants"
	"github.com/Capture-411/services-opportunity/openapi"
	"github.com/Capture-411/services-opportunity/openengine"
	"github.com/Capture-411/services-opportunity/openengine/engine"
)

func main() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	exportPath := GetExportPathFromArgs()
	servicesFolder := GetInternalDirPathFromArgs()

	config.Load()

	errResponses := openapi.DataAndMessageSwaggerErrorResponses

	var (
		protocol     = config.AppConfig.Protocol
		hostUri      = config.AppConfig.HostUri
		externalPort = config.AppConfig.ExternalPort
		apiPrefix    = config.AppConfig.ApiPrefix
		apiVersion   = config.AppConfig.ApiVersion1
	)

	if protocol == "https" {
		externalPort = 443
	}

	oppoUrl := fmt.Sprintf("%s://%s:%d%s%s", protocol, hostUri, externalPort, apiPrefix, apiVersion)

	oppoServer := engine.ApiServer{Url: oppoUrl}

	swaggerUiConfig := engine.SwaggerUiConfig{
		ExportPath: exportPath,
		ServeURI:   "/apidocs",
		Title:      config.AppConfig.Name + " API",
	}

	_, err = openengine.NewPackage().
		AddServers(engine.ApiServers{oppoServer}).
		AddIgnoredPaths(constants.OpenApiIgnored).
		AddIgnoredPaths([]string{exportPath}).
		ParseTags(servicesFolder, []string{"static", "healthcheck", "protobuf", "welcome", "mid"}).
		ParseSchemas(rootPath).
		ParseEnums(rootPath).
		AddDefaultErrors(404, 400, 401, 500).
		AddErrorResponses(errResponses).
		ParsePaths(rootPath).
		AddSecuritySchemes(openapi.SecuritySchemes).
		ExportSwaggerUi(swaggerUiConfig).
		Generate(exportPath)

	if err != nil {
		log.Print("\n\n --> Generate Yaml Errors: ", err, "\n\n")
	}
}
