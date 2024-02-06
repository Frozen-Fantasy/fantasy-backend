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
        "/auth/email/send-code": {
            "post": {
                "description": "Отправка письма с кодом для подтверждения email пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Отправка кода подтверждения",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.EmailInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Выход пользователя из системы",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Выход из системы",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.RefreshInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/auth/refresh-tokens": {
            "post": {
                "description": "Обновление access и refresh токенов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Обновление токенов",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.RefreshInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.Tokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Авторизация пользователя в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Авторизация",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.Tokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Регистрация нового пользователя в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Регистрация",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.SignUpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/user/check-email": {
            "post": {
                "description": "Существует ли уже пользователь с таким email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Использован ли данный email в сервисе",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.EmailInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/user/check-nickname": {
            "post": {
                "description": "Существует ли уже пользователь с таким nickname",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Использован ли данный nickname в сервисе",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.NicknameInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Получение пользовательской информации по access токену",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Получение информации о пользователе",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.UserInfoModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/user/password/change": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Смена пароля",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Смена пароля",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.ChangePasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/user/password/forgot": {
            "post": {
                "description": "Отправка письма с информацией о сбросе пароля",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Забыли пароль",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.EmailInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        },
        "/user/password/reset": {
            "patch": {
                "description": "Восстановление пароля",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Восстановление пароля",
                "parameters": [
                    {
                        "description": "Входные параметры",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.ResetPasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.StatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg_api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.ChangePasswordInput": {
            "type": "object",
            "required": [
                "newPassword",
                "oldPassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                },
                "oldPassword": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.EmailInput": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.NicknameInput": {
            "type": "object",
            "required": [
                "nickname"
            ],
            "properties": {
                "nickname": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 4
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.RefreshInput": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.ResetPasswordInput": {
            "type": "object",
            "required": [
                "hash",
                "newPassword"
            ],
            "properties": {
                "hash": {
                    "type": "string"
                },
                "newPassword": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.SignInInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.SignUpInput": {
            "type": "object",
            "required": [
                "code",
                "email",
                "nickname",
                "password"
            ],
            "properties": {
                "code": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "nickname": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 4
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.Tokens": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "expiresIn": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "github_com_Frozen-Fantasy_fantasy-backend_git_pkg_models_user.UserInfoModel": {
            "type": "object",
            "properties": {
                "coins": {
                    "type": "integer"
                },
                "dateRegistration": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "photoLink": {
                    "type": "string"
                },
                "profileID": {
                    "type": "string"
                }
            }
        },
        "pkg_api.Error": {
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
        "pkg_api.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "fantasy api doc",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
