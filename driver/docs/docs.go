// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Eniz Coban",
            "email": "zulunovember.c@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/driver/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a driver to mongodb with lon and lat values",
                "parameters": [
                    {
                        "description": "Coordinate",
                        "name": "coordinate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Location"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/driver/import": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Import coordinates in Coordinates.csv file to mongo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/driver/match": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "driver"
                ],
                "summary": "Find a driver close to given coordinates",
                "parameters": [
                    {
                        "description": "Coordinate",
                        "name": "coordinate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Location"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Location": {
            "type": "object",
            "required": [
                "lat",
                "lon"
            ],
            "properties": {
                "lat": {
                    "type": "number",
                    "maximum": 90,
                    "minimum": -90
                },
                "lon": {
                    "type": "number",
                    "maximum": 180,
                    "minimum": -180
                }
            }
        },
        "app.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "any"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Driver API",
	Description:      "Sample driver matching API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
