// Package im GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package im

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
        "/info": {
            "get": {
                "description": "用户信息",
                "tags": [
                    "admin"
                ],
                "responses": {
                    "200": {
                        "description": "code==0请求成功，否则请求失败！",
                        "schema": {
                            "$ref": "#/definitions/jwt.User"
                        }
                    }
                }
            }
        },
        "/kick": {
            "post": {
                "description": "踢人下线",
                "tags": [
                    "im"
                ],
                "parameters": [
                    {
                        "description": "参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code==0请求成功，否则请求失败！",
                        "schema": {
                            "$ref": "#/definitions/resp.Ret"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户登录",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "description": "参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code==0请求成功，否则请求失败！",
                        "schema": {
                            "$ref": "#/definitions/service.LoginRes"
                        }
                    }
                }
            }
        },
        "/online": {
            "get": {
                "description": "在线用户",
                "tags": [
                    "im"
                ],
                "parameters": [
                    {
                        "description": "参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code==0请求成功，否则请求失败！",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.User"
                            }
                        }
                    }
                }
            }
        },
        "/sendMessage": {
            "post": {
                "description": "发送消息",
                "tags": [
                    "im"
                ],
                "parameters": [
                    {
                        "description": "参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code==0请求成功，否则请求失败！",
                        "schema": {
                            "$ref": "#/definitions/resp.Ret"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "jwt.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "roleID": {
                    "type": "integer"
                }
            }
        },
        "resp.Ret": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                },
                "serverTime": {
                    "type": "integer"
                }
            }
        },
        "service.LoginReq": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "service.LoginRes": {
            "type": "object",
            "properties": {
                "expire": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "service.Message": {
            "type": "object",
            "required": [
                "body",
                "cmd",
                "from",
                "to"
            ],
            "properties": {
                "body": {
                    "description": "消息内容"
                },
                "cmd": {
                    "type": "integer"
                },
                "from": {
                    "description": "发送者即用户id，必须保证一个唯一",
                    "type": "string"
                },
                "to": {
                    "description": "cmd==10x是表示用户id，cmd==20x是表示群gid",
                    "type": "string"
                }
            }
        },
        "service.User": {
            "type": "object",
            "properties": {
                "gid": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:20105",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "IM API",
	Description:      "This is api document",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}