package main

import (
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/example/chat"
	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/websocket"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", c.Handler(handler.GraphQL(chat.NewExecutableSchema(chat.New()),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}))),
	)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
