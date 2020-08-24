package main

import (
	"log"
	"net/http"
	"os"

	"com.methompson/go-test/graph"
	"com.methompson/go-test/graph/generated"
	"com.methompson/go-test/kcms"
	"com.methompson/go-test/kcms/headers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	cms := kcms.MakeKCMS()

	srv := handler.
		NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{
					Resolvers: &graph.Resolver{
						KCMS: cms,
					},
				},
			),
		)

	router := chi.NewRouter()

	router.Use(headers.JWTExtractor(cms))

	// fmt.Printf("Handler %T\n", srv)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
