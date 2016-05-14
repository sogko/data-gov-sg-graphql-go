package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/pressly/chi"
	"github.com/sogko/data-gov-sg-graphql-go/lib/datagovsg"
	"github.com/sogko/data-gov-sg-graphql-go/lib/schema"
	"github.com/unrolled/render"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
)

var R *render.Render
var API_KEY string

var IP string
var PORT string

func init() {

	// Determine server IP
	IP = os.Getenv("OPENSHIFT_GO_IP")
	if PORT == "" {
		PORT = os.Getenv("DATAGOVSG_IP")
	}
	log.Println("IP", IP)

	// Determine server PORT
	PORT = os.Getenv("OPENSHIFT_GO_PORT")
	if PORT == "" {
		PORT = os.Getenv("DATAGOVSG_PORT")
	}
	if PORT == "" {
		PORT = "3000"
	}

	// Get data.gov.sg API key from env vars (required)
	API_KEY = os.Getenv("DATAGOVSG_API_KEY")
	if API_KEY == "" {
		panic("Set DATAGOVSG_API_KEY environment variable before running server")
	}
	log.Println("API key OK")

	R = render.New(render.Options{
		Directory:     "views",
		IsDevelopment: true,
		Extensions:    []string{".html"},
	})
}

func serveGraphQL(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// get query
	opts := handler.NewRequestOptions(r)

	// init and store data.gov.sg client
	ctx = context.WithValue(ctx, "client", datagovsg.NewClient(API_KEY))

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

	bind := fmt.Sprintf("%s:%s", IP, PORT)
	log.Println("Starting server at", bind)

	http.ListenAndServe(bind, r)
}
