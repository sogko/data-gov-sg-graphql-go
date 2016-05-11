package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/pressly/chi"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema"
	"github.com/unrolled/render"
	"golang.org/x/net/context"
)

var R *render.Render

func init() {
	R = render.New(render.Options{
		Directory:     "views",
		IsDevelopment: true,
		Extensions:    []string{".html"},
	})
}

func serveGraphQL(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// get query
	opts := handler.NewRequestOptions(r)

	// execute graphql query
	params := graphql.Params{
		Schema:         schema.Root,
		RequestString:  opts.Query,
		VariableValues: opts.Variables,
		OperationName:  opts.OperationName,
		Context:        ctx,
	}
	result := graphql.Do(params)

	// render result
	R.JSON(w, http.StatusOK, result)
}
func main() {
	r := chi.NewRouter()

	r.Handle("/graphql", serveGraphQL)
	r.FileServer("/", http.Dir("static"))

	http.ListenAndServe(":3000", r)
}
