// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "task-management",
    "version": "1.0.0"
  },
  "host": "example.swagger.io",
  "basePath": "/v1",
  "paths": {
    "/auth/login": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "AuthApi"
        ],
        "summary": "ログイン認証",
        "operationId": "postAuthLogin",
        "parameters": [
          {
            "description": "body",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthUserReq"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ログイン成功",
            "schema": {
              "$ref": "#/definitions/AuthReturnJWT"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "500": {
            "description": "Internal Server Error"
          },
          "503": {
            "description": "Service Unavailable"
          }
        }
      }
    },
    "/auth/register": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "AuthApi"
        ],
        "summary": "ユーザー登録",
        "operationId": "postAuthRegister",
        "parameters": [
          {
            "description": "body",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthUserReq"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ユーザー登録成功",
            "schema": {
              "$ref": "#/definitions/AuthReturnJWT"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "500": {
            "description": "Internal Server Error"
          },
          "503": {
            "description": "Service Unavailable"
          }
        }
      }
    }
  },
  "definitions": {
    "AuthReturnJWT": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "AuthUserReq": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "screen_name": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "task-management",
    "version": "1.0.0"
  },
  "host": "example.swagger.io",
  "basePath": "/v1",
  "paths": {
    "/auth/login": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "AuthApi"
        ],
        "summary": "ログイン認証",
        "operationId": "postAuthLogin",
        "parameters": [
          {
            "description": "body",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthUserReq"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ログイン成功",
            "schema": {
              "$ref": "#/definitions/AuthReturnJWT"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "500": {
            "description": "Internal Server Error"
          },
          "503": {
            "description": "Service Unavailable"
          }
        }
      }
    },
    "/auth/register": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "AuthApi"
        ],
        "summary": "ユーザー登録",
        "operationId": "postAuthRegister",
        "parameters": [
          {
            "description": "body",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AuthUserReq"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ユーザー登録成功",
            "schema": {
              "$ref": "#/definitions/AuthReturnJWT"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "500": {
            "description": "Internal Server Error"
          },
          "503": {
            "description": "Service Unavailable"
          }
        }
      }
    }
  },
  "definitions": {
    "AuthReturnJWT": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "AuthUserReq": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "screen_name": {
          "type": "string"
        }
      }
    }
  }
}`))
}
