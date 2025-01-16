package router

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/t-okuji/learn-huma/controller"
)

func NewAuthorRouter(api huma.API, ac controller.IAuthorController) {
	huma.Register(api, huma.Operation{
		OperationID: "get-authors",
		Method:      http.MethodGet,
		Path:        "/authors",
		Summary:     "Get Authors",
		Description: "Get Authors.",
		Tags:        []string{"Authors"},
	}, func(ctx context.Context, _ *struct {
	}) (*controller.AuthorsOutput, error) {
		return ac.ListAuthors(ctx)
	})
}