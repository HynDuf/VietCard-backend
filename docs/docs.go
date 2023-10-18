// Code generated by swaggo/swag. DO NOT EDIT.

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
            "email": "hynduf@gmail.com"
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
        "/api/card/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create New Card",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Create New Card",
                "parameters": [
                    {
                        "description": "Create Card Request",
                        "name": "create_card_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CreateCardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/card/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Card Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "card"
                ],
                "summary": "Update Card Details",
                "parameters": [
                    {
                        "description": "Update Card Request",
                        "name": "update_card_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateCardResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/deck/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create New Deck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Create New Deck",
                "parameters": [
                    {
                        "description": "Create Deck Request",
                        "name": "create_deck_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateDeckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CreateDeckResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/deck/review-cards": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Deck With Review Cards Of Logged In User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Get Deck With Review Cards Of Logged In User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.DeckWithReviewCards"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/deck/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Deck Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "Update Deck Details",
                "parameters": [
                    {
                        "description": "Update Deck Request",
                        "name": "update_deck_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateDeckRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateDeckResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Log In",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Log In",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/refresh": {
            "post": {
                "description": "Refresh Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "description": "Refresh Token Request",
                        "name": "refresh_token_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.RefreshTokenResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/signup": {
            "post": {
                "description": "Sign Up",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Sign Up",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SignupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/user/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update User Details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update User Details",
                "parameters": [
                    {
                        "description": "Update User Request",
                        "name": "update_user_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.UpdateUserResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Card": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_review": {
                    "type": "string"
                },
                "next_review": {
                    "type": "string"
                },
                "num_reviews": {
                    "type": "integer"
                },
                "question": {
                    "type": "string"
                },
                "sm2_ef": {
                    "type": "number"
                },
                "sm2_i": {
                    "type": "integer"
                },
                "sm2_n": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "wrong_answers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "entity.Deck": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_global": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.DeckWithReviewCards": {
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Card"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_global": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "hashed_password": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "last_streak": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                },
                "max_cards_review": {
                    "type": "integer"
                },
                "max_new_cards_learn": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "streak": {
                    "type": "integer"
                },
                "xp": {
                    "type": "integer"
                },
                "xp_to_level_up": {
                    "type": "integer"
                }
            }
        },
        "request.CreateCardRequest": {
            "type": "object",
            "required": [
                "answer",
                "deck_id",
                "question",
                "wrong_answers"
            ],
            "properties": {
                "answer": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                },
                "wrong_answers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "request.CreateDeckRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "request.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "request.UpdateCardRequest": {
            "type": "object",
            "required": [
                "card_id"
            ],
            "properties": {
                "answer": {
                    "type": "string"
                },
                "card_id": {
                    "type": "string"
                },
                "deck_id": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                },
                "wrong_answers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "request.UpdateDeckRequest": {
            "type": "object",
            "required": [
                "deck_id"
            ],
            "properties": {
                "deck_id": {
                    "type": "string"
                },
                "is_global": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "request.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "hashed_password": {
                    "type": "string"
                },
                "max_cards_review": {
                    "type": "integer"
                },
                "max_new_cards_learn": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.CreateCardResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.CreateDeckResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.RefreshTokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.SignupResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "response.UpdateCardResponse": {
            "type": "object",
            "properties": {
                "card": {
                    "$ref": "#/definitions/entity.Card"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.UpdateDeckResponse": {
            "type": "object",
            "properties": {
                "deck": {
                    "$ref": "#/definitions/entity.Deck"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                },
                "user": {
                    "$ref": "#/definitions/entity.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "VietCard Backend API",
	Description:      "Backend server for VietCard application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
