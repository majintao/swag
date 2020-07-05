// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/testapi/get-bar": {
            "get": {
                "description": "get Bar",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show a account",
                "operationId": "get-bar",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Bar"
                        }
                    }
                }
            }
        },
        "/testapi/get-barmap": {
            "get": {
                "description": "get BarMap",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "get-bar-map",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.BarMap"
                        }
                    }
                }
            }
        },
        "/testapi/get-foo": {
            "get": {
                "description": "get Foo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "get-foo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Foo"
                        }
                    }
                }
            }
        },
        "/testapi/get-foobar": {
            "get": {
                "description": "get FooBar",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "get-foobar",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.FooBar"
                        }
                    }
                }
            }
        },
        "/testapi/get-foobar-pointer": {
            "get": {
                "description": "get FooBarPointer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "get-foobar-pointer",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.FooBarPointer"
                        }
                    }
                }
            }
        },
        "/testapi/get-foobarmap": {
            "get": {
                "description": "get FoorBarMap",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "get-foo-bar-map",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.FooBarMap"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Bar": {
            "type": "object",
            "properties": {
                "field2": {
                    "type": "string"
                }
            }
        },
        "api.BarMap": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#/definitions/api.Bar"
            }
        },
        "api.Foo": {
            "type": "object",
            "properties": {
                "field1": {
                    "type": "string"
                }
            }
        },
        "api.FooBar": {
            "type": "object",
            "properties": {
                "field1": {
                    "type": "string"
                },
                "field2": {
                    "type": "string"
                }
            }
        },
        "api.FooBarMap": {
            "type": "object",
            "properties": {
                "field3": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/api.MapValue"
                    }
                }
            }
        },
        "api.FooBarPointer": {
            "type": "object",
            "properties": {
                "field1": {
                    "type": "string"
                },
                "field2": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "api.MapValue": {
            "type": "object",
            "properties": {
                "field4": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "petstore.swagger.io",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server",
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