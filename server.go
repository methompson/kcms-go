package main

import (
	"log"
	"net/http"
	"os"

	"com.methompson/kcms-go/graph"
	"com.methompson/kcms-go/graph/generated"
	"com.methompson/kcms-go/kcms"
	"com.methompson/kcms-go/kcms/headers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

var getEnvironmentVal = os.Getenv

func main() {
	cms, err := kcms.MakeKCMS()

	if err != nil {
		log.Panic("kcms not created", err)
	}

	err = makeServer(cms)

	if err != nil {
		log.Fatal(err)
	}
}

func makeServer(cms *kcms.KCMS) error {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					KCMS: cms,
				},
			},
		),
	)

	port := getEnvironmentVal("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(headers.JWTExtractor(cms))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, router)
}
