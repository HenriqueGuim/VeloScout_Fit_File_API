// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/decode": {
            "post": {
                "description": "Decode a .fit file",
                "produces": [
                    "application/json"
                ],
                "summary": "Decode a .fit file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to decode",
                        "name": "FitFile",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DataFormat"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DataFormat": {
            "type": "object",
            "properties": {
                "average_cadence": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "distance": {
                    "type": "number"
                },
                "duration": {
                    "type": "integer"
                },
                "elevation_gain": {
                    "type": "integer"
                },
                "heart_rate": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "i_f": {
                    "type": "number"
                },
                "power": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "t_ss": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "VeloScoutFitAPI Swagger API Documentation",
	Description:      "This is an API for decoding .fit files",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
