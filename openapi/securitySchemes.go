package openapi

import (
	"github.com/Capture-411/services-opportunity/openengine/engine"
	authtype "github.com/Capture-411/services-opportunity/openengine/engine/types/authType"
)

var SecuritySchemes = engine.SecuritySchemesTypes{
	Http: engine.HttpSecuritySchemesDict{
		"jwtBearerAuth": engine.HttpSecurityScheme{
			Type:         authtype.TypeHttp,
			Description:  "HTTP Bearer authentication",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	},
}
