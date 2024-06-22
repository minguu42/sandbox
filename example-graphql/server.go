package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/sandbox/example-graphql/graph"
	"github.com/minguu42/sandbox/example-graphql/graph/services"
	"github.com/minguu42/sandbox/example-graphql/internal"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("mysql", "root:@tcp(localhost:13306)/maindb?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	boil.DebugMode = true

	service := services.New(db)

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{
		Srv:     service,
		Loaders: graph.NewLoaders(service),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
