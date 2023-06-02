// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/meals": {
            "get": {
                "description": "食事を追加する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "getMeals",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/handler.getMealsResponse"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "食事を追加する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "postMeals",
                "parameters": [
                    {
                        "description": "食事情報",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.postMealRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.postMealResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/meals/{id}": {
            "get": {
                "description": "食事を追加する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "getMealsId",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.getMealsResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "put": {
                "description": "食事を追加する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "putMealsId",
                "parameters": [
                    {
                        "description": "食事情報",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.putMealsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.putMealsResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "食事を消去する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "deleteMealsId",
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/users/lgin": {
            "post": {
                "description": "ログイン",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "postUsersLogin",
                "parameters": [
                    {
                        "description": "ログイン情報",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.userLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.userLoginResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/users/signup": {
            "post": {
                "description": "サインアップ",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "operationId": "postUsersSignup",
                "parameters": [
                    {
                        "description": "サインイン情報",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.userSignupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.userSignupResponse"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.getMealsResponse": {
            "description": "食事を追加する",
            "type": "object",
            "properties": {
                "calories": {
                    "type": "number"
                },
                "carbs": {
                    "type": "number"
                },
                "fat": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "mealtype": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "protein": {
                    "type": "number"
                }
            }
        },
        "handler.postMealRequest": {
            "type": "object",
            "properties": {
                "calories": {
                    "type": "number"
                },
                "carbs": {
                    "type": "number"
                },
                "fat": {
                    "type": "number"
                },
                "mealtype": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "protein": {
                    "type": "number"
                }
            }
        },
        "handler.postMealResponse": {
            "type": "object",
            "properties": {
                "calories": {
                    "type": "number"
                },
                "carbs": {
                    "type": "number"
                },
                "fat": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "mealtype": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "protein": {
                    "type": "number"
                }
            }
        },
        "handler.putMealsRequest": {
            "type": "object",
            "properties": {
                "calories": {
                    "type": "number"
                },
                "carbs": {
                    "type": "number"
                },
                "fat": {
                    "type": "number"
                },
                "mealtype": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "protein": {
                    "type": "number"
                }
            }
        },
        "handler.putMealsResponse": {
            "type": "object",
            "properties": {
                "calories": {
                    "type": "number"
                },
                "carbs": {
                    "type": "number"
                },
                "fat": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "mealtype": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "protein": {
                    "type": "number"
                }
            }
        },
        "handler.userLoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handler.userLoginResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "handler.userSignupRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "handler.userSignupResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Nutrition App API",
	Description:      "API server for the Nutrition App",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}