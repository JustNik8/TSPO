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
        "/goods": {
            "get": {
                "description": "Возвращает список всех товаров",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Получение списка товаров",
                "operationId": "get-all-goods",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Good"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавление товара",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Добавление товара",
                "operationId": "add-good",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            }
        },
        "/goods/{good_id}": {
            "get": {
                "description": "Возвращает товар по его айди",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Получение товара по его айди",
                "operationId": "get-good-by-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Good"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновление данных о товаре",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Обновление данных о товаре",
                "operationId": "update-good",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление товара по его айди",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "goods"
                ],
                "summary": "Удаление товара по его айди",
                "operationId": "delete-good",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Body"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Good": {
            "type": "object",
            "required": [
                "desc",
                "measure_unit",
                "name",
                "price"
            ],
            "properties": {
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "measure_unit": {
                    "$ref": "#/definitions/dto.MeasureUnit"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock_quantity": {
                    "type": "integer"
                }
            }
        },
        "dto.MeasureUnit": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "enum": [
                        "METER",
                        "KILOGRAM",
                        "LITER",
                        "PIECE"
                    ]
                }
            }
        },
        "response.Body": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "response.IDResponse": {
            "type": "object",
            "properties": {
                "id": {
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