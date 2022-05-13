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
        "/captcha": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC"
                ],
                "summary": "获取验证码",
                "responses": {
                    "200": {
                        "description": "获取验证码成功返回体",
                        "schema": {
                            "$ref": "#/definitions/response.Base"
                        }
                    }
                }
            }
        },
        "/goods": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块"
                ],
                "summary": "获取商品列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "当前页第几页",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页的大小",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取商品列表返回体",
                        "schema": {
                            "$ref": "#/definitions/response.Base"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RBAC"
                ],
                "summary": "用户登录功能",
                "parameters": [
                    {
                        "description": "登录需要的json参数",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取验证码成功返回体",
                        "schema": {
                            "$ref": "#/definitions/response.Base"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Base": {
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
        "services.LoginRequest": {
            "type": "object",
            "properties": {
                "captchaId": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "verifyCode": {
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
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "微服务 API",
	Description:      "技术栈 Gin + Go-Micro + Gorm + gRPC",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
