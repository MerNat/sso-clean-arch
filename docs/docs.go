// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/sso/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SSO"
                ],
                "summary": "Login to service",
                "parameters": [
                    {
                        "description": "Users Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.User"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "email": {
                                            "type": "string"
                                        },
                                        "password": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/json.GenericResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/sso/register": {
            "post": {
                "description": "Creating a User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SSO"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "Users Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/json.GenericResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "message": {
                                            "type": "string"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "json.GenericResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
