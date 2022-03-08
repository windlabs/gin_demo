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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/check": {
            "get": {
                "description": "服务初始连接测试  接口描述",
                "summary": "服务连接校验  接口简介",
                "responses": {}
            }
        },
        "/hello": {
            "get": {
                "description": "服务初始连接测试  接口描述",
                "summary": "服务连接校验  接口简介",
                "responses": {}
            }
        },
        "/test": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "接口描述信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "首页接口摘要信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name value",
                        "name": "name",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200, \"message\":\"success\"}"
                    },
                    "400": {
                        "description": "{\"code\":400, \"message\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "{\"code\":404, \"message\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"code\":500, \"message\":\"success\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8085",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API12222",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
