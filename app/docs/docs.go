// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-03 18:46:23.57785204 +0200 CEST m=+0.026280661

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/models": {
            "get": {
                "description": "Query all avaiable room models.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models"
                ],
                "summary": "Query models",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.RoomModel"
                            }
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/models/{id}": {
            "get": {
                "description": "Query a single room model by id with containing sensors.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models"
                ],
                "summary": "Query room model",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "RoomModelID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.RoomModel"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sensors": {
            "get": {
                "description": "Query all avaiable sensors.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensors"
                ],
                "summary": "Query sensors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Sensor"
                            }
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sensors/{id}": {
            "get": {
                "description": "Query a single sensor by id with containing sensor data.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensors"
                ],
                "summary": "Query sensor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SensorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Sensor"
                        }
                    },
                    "400": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Data": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "sensorID": {
                    "type": "integer"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "main.RoomModel": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sensors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Sensor"
                    }
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "main.Sensor": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Data"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "measurementUnit": {
                    "type": "string"
                },
                "meshID": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "roomModelID": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.2",
	Host:        "visense.f4.htw-berlin.de:8080",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "vi-sense BIM API",
	Description: "This API provides information about 3D room models with associated sensors and their data.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}