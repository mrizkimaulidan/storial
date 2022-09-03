package category

import (
	"log"
	"net/http"

	"github.com/mrizkimaulidan/storial/internal/service/category"
	"github.com/mrizkimaulidan/storial/pkg/response"
)

type categoryHandler struct {
	categoryService category.CategoryService
	response        *response.Response
}

func (ch *categoryHandler) GetAllCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categoriesResponse, err := ch.categoryService.GetAll(r.Context())
		if err != nil {
			ch.handleErr(err).JSON(w)
			return
		}

		ch.response.SetCode(http.StatusOK).SetMessage("OK").SetData(categoriesResponse).JSON(w)
	})
}

func (ch *categoryHandler) handleErr(err error) *response.Response {
	switch err {
	//
	}

	log.Println("[ERROR]", err)
	return ch.response.SetCode(http.StatusInternalServerError).SetMessage("internal server error")
}

func NewHandler(cs category.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: cs,
		response:        new(response.Response),
	}
}
