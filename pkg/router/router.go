package router

import (
  "github.com/99designs/gqlgen/graphql/handler"
  "github.com/99designs/gqlgen/graphql/playground"
  "github.com/gorilla/mux"
	"github.com/kieranroneill/valkyrie/pkg/application"
	"github.com/kieranroneill/valkyrie/pkg/graphql"
	_handler "github.com/kieranroneill/valkyrie/pkg/handler"
	"github.com/kieranroneill/valkyrie/pkg/middleware"
	"github.com/kieranroneill/valkyrie/pkg/resolver"
  "net/http"
)

func New(app *application.Application) *mux.Router {
	r := mux.NewRouter()

  r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	hlcksr := r.PathPrefix("/healthcheck").Subrouter()
  hlcksr.
		HandleFunc(
			"",
			middleware.ApplyMiddleware(
				_handler.CreateHealthcheckHandler(app),
				middleware.LogRequest(),
			),
		).
		Methods(http.MethodGet)
  gqlsr := r.PathPrefix("/graphql").Subrouter()
  gqlsr.
    Handle(
      "",
      handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
        Resolvers: &resolver.Resolver{
          Config: app.Config,
          Database: app.Database,
        },
      })),
    )

	return r
}
