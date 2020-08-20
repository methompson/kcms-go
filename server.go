package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"com.methompson/go-test/config"
	"com.methompson/go-test/graph"
	"com.methompson/go-test/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	configDat, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	config.ReadConfig(configDat)

	srv := handler.
		NewDefaultServer(
			generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}),
		)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
