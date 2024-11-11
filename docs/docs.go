// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/job": {
            "post": {
                "description": "Upload a file to create a new job with its filename stored in the database",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "job"
                ],
                "summary": "create a new job",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to be uploaded",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Job successfully created",
                        "schema": {
                            "$ref": "#/definitions/models.Job"
                        }
                    },
                    "400": {
                        "description": "Bad request, file missing in the form-data",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error, unable to create job",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Job": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "filename": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/models.JobStatus"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.JobStatus": {
            "type": "string",
            "enum": [
                "QUEUED",
                "PROCESSING",
                "COMPLETED",
                "FAILED"
            ],
            "x-enum-varnames": [
                "JobQueued",
                "JobProcessing",
                "JobCompleted",
                "JobFailed"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
