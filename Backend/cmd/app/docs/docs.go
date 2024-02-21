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
        "/api/v1/games": {
            "post": {
                "description": "Start a new game",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Start a new game",
                "parameters": [
                    {
                        "description": "Players",
                        "name": "players",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateGameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.CreateGameResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/games/{gameId}/events": {
            "get": {
                "description": "Get game events",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "games"
                ],
                "summary": "Get game events",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Game ID",
                        "name": "gameId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.GameSSERequest"
                        }
                    }
                }
            }
        },
        "/api/v1/heartbeat": {
            "get": {
                "description": "Check if the server is alive",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "heartbeat"
                ],
                "summary": "Check if the server is alive",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/api/v1/player/{id}/player-cards/": {
            "get": {
                "description": "GetPlayerCardsByPlayerId",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "player_cards"
                ],
                "summary": "GetPlayerCards",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.PlayerCardsResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/player/{playerId}/transmit-intelligence": {
            "post": {
                "description": "Transmit an intelligence card",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Transmit intelligence",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "playerId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Card ID",
                        "name": "card_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PlayCardRequest"
                        }
                    },
                    {
                        "description": "Intelligence Type",
                        "name": "intelligence_type",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PlayCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.PlayCardResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/players/{playerId}/accept": {
            "post": {
                "description": "Decide accept card or not",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Accept Card",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "playerId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Accept",
                        "name": "accept",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AcceptCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.PlayCardResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/players/{playerId}/player-cards": {
            "post": {
                "description": "Play a card",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Play a card",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Player ID",
                        "name": "playerId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Card ID",
                        "name": "card_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PlayCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/request.PlayCardResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.GameSSERequest": {
            "type": "object",
            "properties": {
                "game_id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "request.AcceptCardRequest": {
            "type": "object",
            "properties": {
                "accept": {
                    "type": "boolean"
                }
            }
        },
        "request.CreateGameRequest": {
            "type": "object",
            "properties": {
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.PlayerInfo"
                    }
                }
            }
        },
        "request.CreateGameResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "request.PlayCardRequest": {
            "type": "object",
            "properties": {
                "card_id": {
                    "type": "integer"
                },
                "intelligence_type": {
                    "type": "integer"
                }
            }
        },
        "request.PlayCardResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "boolean"
                }
            }
        },
        "request.PlayerCardsResponse": {
            "type": "object",
            "properties": {
                "player_cards": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": true
                    }
                }
            }
        },
        "request.PlayerInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
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
	Version:          "",
	Host:             "127.0.0.1:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "The Message API",
	Description:      "This is an online version of the \"The Message\" board game backend API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
