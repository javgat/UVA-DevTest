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
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "DevTest",
    "title": "DevTest",
    "contact": {
      "email": "javigaton@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/DevTest",
  "paths": {
    "/accesstokens": {
      "post": {
        "description": "Tries to login, and gets a JWT auth token if successful",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Registers a new authorized connection token",
        "operationId": "Login",
        "parameters": [
          {
            "description": "User who is trying to generate a token",
            "name": "loginUser",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/LoginUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful authentication",
            "schema": {
              "$ref": "#/definitions/JWTJson"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequestError"
          },
          "410": {
            "$ref": "#/responses/GoneError"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/users": {
      "post": {
        "description": "Adds an user to the system",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user",
          "auth"
        ],
        "summary": "adds an user",
        "operationId": "RegisterUser",
        "parameters": [
          {
            "description": "User item to add",
            "name": "signinUser",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/SigninUser"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "user created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "$ref": "#/responses/BadRequestError"
          },
          "409": {
            "$ref": "#/responses/ConflictError"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "example": "Object not found"
        }
      }
    },
    "JWTJson": {
      "type": "object",
      "required": [
        "token"
      ],
      "properties": {
        "token": {
          "type": "string",
          "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
        }
      }
    },
    "LoginUser": {
      "type": "object",
      "required": [
        "loginid",
        "pass"
      ],
      "properties": {
        "loginid": {
          "type": "string",
          "example": "carlosg72 || carlos@mail.com"
        },
        "pass": {
          "type": "string",
          "example": "pass"
        }
      }
    },
    "SigninUser": {
      "type": "object",
      "required": [
        "username",
        "email",
        "pass"
      ],
      "properties": {
        "email": {
          "type": "string",
          "example": "carlos@mail.com"
        },
        "pass": {
          "type": "string",
          "example": "pass"
        },
        "username": {
          "type": "string",
          "example": "carlosg72"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "username",
        "email",
        "pwhash"
      ],
      "properties": {
        "email": {
          "type": "string",
          "example": "carlos@mail.com"
        },
        "pwhash": {
          "type": "string",
          "example": "e$ia9s7ATDGba39pakscAKs"
        },
        "username": {
          "type": "string",
          "example": "carlosg72"
        }
      }
    }
  },
  "responses": {
    "BadRequestError": {
      "description": "Incorrect Request, or invalida data",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "ConflictError": {
      "description": "A user with same username/email already exists",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "GoneError": {
      "description": "That user (password and name) does not exist",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "InternalServerError": {
      "description": "Internal error"
    }
  },
  "securityDefinitions": {
    "APIKeyHeader": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    },
    "APIKeyQueryParam": {
      "type": "apiKey",
      "name": "api_key",
      "in": "query"
    }
  },
  "tags": [
    {
      "description": "Operations related to users",
      "name": "user"
    },
    {
      "description": "Operations related to authentication",
      "name": "auth"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "DevTest",
    "title": "DevTest",
    "contact": {
      "email": "javigaton@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/DevTest",
  "paths": {
    "/accesstokens": {
      "post": {
        "description": "Tries to login, and gets a JWT auth token if successful",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Registers a new authorized connection token",
        "operationId": "Login",
        "parameters": [
          {
            "description": "User who is trying to generate a token",
            "name": "loginUser",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/LoginUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful authentication",
            "schema": {
              "$ref": "#/definitions/JWTJson"
            }
          },
          "400": {
            "description": "Incorrect Request, or invalida data",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "410": {
            "description": "That user (password and name) does not exist",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/users": {
      "post": {
        "description": "Adds an user to the system",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user",
          "auth"
        ],
        "summary": "adds an user",
        "operationId": "RegisterUser",
        "parameters": [
          {
            "description": "User item to add",
            "name": "signinUser",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/SigninUser"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "user created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Incorrect Request, or invalida data",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "409": {
            "description": "A user with same username/email already exists",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "example": "Object not found"
        }
      }
    },
    "JWTJson": {
      "type": "object",
      "required": [
        "token"
      ],
      "properties": {
        "token": {
          "type": "string",
          "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
        }
      }
    },
    "LoginUser": {
      "type": "object",
      "required": [
        "loginid",
        "pass"
      ],
      "properties": {
        "loginid": {
          "type": "string",
          "example": "carlosg72 || carlos@mail.com"
        },
        "pass": {
          "type": "string",
          "example": "pass"
        }
      }
    },
    "SigninUser": {
      "type": "object",
      "required": [
        "username",
        "email",
        "pass"
      ],
      "properties": {
        "email": {
          "type": "string",
          "example": "carlos@mail.com"
        },
        "pass": {
          "type": "string",
          "example": "pass"
        },
        "username": {
          "type": "string",
          "example": "carlosg72"
        }
      }
    },
    "User": {
      "type": "object",
      "required": [
        "username",
        "email",
        "pwhash"
      ],
      "properties": {
        "email": {
          "type": "string",
          "example": "carlos@mail.com"
        },
        "pwhash": {
          "type": "string",
          "example": "e$ia9s7ATDGba39pakscAKs"
        },
        "username": {
          "type": "string",
          "example": "carlosg72"
        }
      }
    }
  },
  "responses": {
    "BadRequestError": {
      "description": "Incorrect Request, or invalida data",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "ConflictError": {
      "description": "A user with same username/email already exists",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "GoneError": {
      "description": "That user (password and name) does not exist",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "InternalServerError": {
      "description": "Internal error"
    }
  },
  "securityDefinitions": {
    "APIKeyHeader": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    },
    "APIKeyQueryParam": {
      "type": "apiKey",
      "name": "api_key",
      "in": "query"
    }
  },
  "tags": [
    {
      "description": "Operations related to users",
      "name": "user"
    },
    {
      "description": "Operations related to authentication",
      "name": "auth"
    }
  ]
}`))
}
