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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for managing ToDo items",
    "title": "ToDo API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/todos": {
      "get": {
        "summary": "Get all ToDo items",
        "operationId": "getTodos",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Todo"
              }
            }
          }
        }
      },
      "post": {
        "summary": "Create a new ToDo item",
        "operationId": "createTodo",
        "parameters": [
          {
            "name": "todo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          }
        }
      }
    },
    "/todos/{id}": {
      "get": {
        "summary": "Get a ToDo item by ID",
        "operationId": "getTodo",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "put": {
        "summary": "Update a ToDo item by ID",
        "operationId": "updateTodo",
        "parameters": [
          {
            "name": "todo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Updated successfully",
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "delete": {
        "summary": "Delete a ToDo item by ID",
        "operationId": "deleteTodo",
        "responses": {
          "204": {
            "description": "No Content, deleted successfully"
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Todo": {
      "type": "object",
      "required": [
        "title",
        "completed"
      ],
      "properties": {
        "completed": {
          "type": "boolean",
          "example": false
        },
        "created_at": {
          "type": "string",
          "format": "date-time",
          "example": "2023-10-01T12:00:00Z"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "title": {
          "type": "string"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time",
          "example": "2023-10-01T12:00:00Z"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for managing ToDo items",
    "title": "ToDo API",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/todos": {
      "get": {
        "summary": "Get all ToDo items",
        "operationId": "getTodos",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Todo"
              }
            }
          }
        }
      },
      "post": {
        "summary": "Create a new ToDo item",
        "operationId": "createTodo",
        "parameters": [
          {
            "name": "todo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          }
        }
      }
    },
    "/todos/{id}": {
      "get": {
        "summary": "Get a ToDo item by ID",
        "operationId": "getTodo",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "put": {
        "summary": "Update a ToDo item by ID",
        "operationId": "updateTodo",
        "parameters": [
          {
            "name": "todo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Updated successfully",
            "schema": {
              "$ref": "#/definitions/Todo"
            }
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "delete": {
        "summary": "Delete a ToDo item by ID",
        "operationId": "deleteTodo",
        "responses": {
          "204": {
            "description": "No Content, deleted successfully"
          },
          "404": {
            "description": "Not Found"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Todo": {
      "type": "object",
      "required": [
        "title",
        "completed"
      ],
      "properties": {
        "completed": {
          "type": "boolean",
          "example": false
        },
        "created_at": {
          "type": "string",
          "format": "date-time",
          "example": "2023-10-01T12:00:00Z"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "title": {
          "type": "string"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time",
          "example": "2023-10-01T12:00:00Z"
        }
      }
    }
  }
}`))
}
