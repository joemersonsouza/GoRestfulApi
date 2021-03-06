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
            "name": "GitHub",
            "url": "https://www.github.com/joemersonsouza"
        },
        "license": {
            "name": "BSD License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/item": {
            "post": {
                "description": "Add a new item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add a new item",
                "parameters": [
                    {
                        "description": "Item body",
                        "name": "Item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Something goes wrong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/item/:{id}": {
            "put": {
                "description": "change an existing item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Change an item",
                "parameters": [
                    {
                        "description": "Item ID",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Item body",
                        "name": "Item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Item not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Something goes wrong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/item/{id}": {
            "get": {
                "description": "get an item by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a specific item",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product Item",
                        "schema": {
                            "$ref": "#/definitions/models.ProductRequest"
                        }
                    },
                    "404": {
                        "description": "Item not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/items": {
            "get": {
                "description": "get items",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all stored items",
                "responses": {
                    "200": {
                        "description": "List of products",
                        "schema": {
                            "$ref": "#/definitions/models.ProductRequest"
                        }
                    },
                    "204": {
                        "description": "No content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ProductRequest": {
            "type": "object",
            "properties": {
                "discount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Boilerplate API in GO",
	Description:      "This is a sample RESTFULL API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
