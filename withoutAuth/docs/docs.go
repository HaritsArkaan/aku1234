// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tugaspendahuluans": {
            "get": {
                "description": "Get a list of TugasPendahuluan.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TugasPendahuluan"
                ],
                "summary": "Get all TugasPendahuluan.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TugasPendahuluan"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creating a new TugasPendahuluan.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TugasPendahuluan"
                ],
                "summary": "Create New TugasPendahuluan.",
                "parameters": [
                    {
                        "description": "the body to create a new TugasPendahuluan",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.TugasPendahuluanInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TugasPendahuluan"
                        }
                    }
                }
            }
        },
        "/tugaspendahuluans/id/{id}": {
            "get": {
                "description": "Get an TugasPendahuluan by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TugasPendahuluan"
                ],
                "summary": "Get TugasPendahuluan.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TugasPendahuluan id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TugasPendahuluan"
                        }
                    }
                }
            }
        },
        "/tugaspendahuluans/{id}": {
            "delete": {
                "description": "Delete a TugasPendahuluan by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TugasPendahuluan"
                ],
                "summary": "Delete one TugasPendahuluan.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "TugasPendahuluan id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update TugasPendahuluan by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TugasPendahuluan"
                ],
                "summary": "Update TugasPendahuluan.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "TugasPendahuluan id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update TugasPendahuluan",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.TugasPendahuluanInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TugasPendahuluan"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.TugasPendahuluanInput": {
            "type": "object",
            "properties": {
                "deadline": {
                    "type": "string"
                },
                "deskripsi": {
                    "type": "string"
                },
                "judul": {
                    "type": "string"
                },
                "kategori": {
                    "type": "string"
                },
                "subJudul": {
                    "type": "string"
                }
            }
        },
        "models.TugasPendahuluan": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deadline": {
                    "type": "string"
                },
                "deskripsi": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "judul": {
                    "type": "string"
                },
                "kategori": {
                    "type": "string"
                },
                "subJudul": {
                    "type": "string"
                }
            }
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
