package story

import (
	"errors"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
	model "github.com/mrizkimaulidan/storial/internal/model/story"
	"github.com/mrizkimaulidan/storial/internal/service/story"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/story"
	jwtpkg "github.com/mrizkimaulidan/storial/pkg/jwt"
	"github.com/mrizkimaulidan/storial/pkg/response"
)

type storyHandler struct {
	storyService story.StoryService
	response     *response.Response
}

func (sh *storyHandler) FilterByCategorySlug() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		categorySlug := vars["categorySlug"]
		filterType := r.URL.Query().Get("filter")

		categoriesRespone, err := sh.storyService.FilterStoryByCategorySlug(r.Context(), categorySlug, filterType)
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		sh.response.SetCode(http.StatusOK).SetMessage("OK").SetData(categoriesRespone).JSON(w)
	})
}

func (sh *storyHandler) Filter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filterType := r.URL.Query().Get("filter")

		storiesResponse, err := sh.storyService.FilterStory(r.Context(), filterType)
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		sh.response.SetCode(http.StatusOK).SetMessage("OK").SetData(storiesResponse).JSON(w)
	})
}

func (sh *storyHandler) GetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)

		storiesResponse, err := sh.storyService.GetAllStory(r.Context(), user.Id)
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		sh.response.SetCode(http.StatusOK).SetMessage("OK").SetData(storiesResponse).JSON(w)
	})
}

func (sh *storyHandler) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		request := model.DeleteStoryRequest{
			Id: vars["id"],
		}

		storyResponse, err := sh.storyService.RemoveStory(r.Context(), request)
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		sh.response.SetCode(http.StatusOK).SetMessage("OK").SetData(storyResponse).JSON(w)
	})
}

func (sh *storyHandler) Store() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, fileheader, _ := r.FormFile("cover")

		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)

		request := model.CreateStoryRequest{
			UserID:          user.Id,
			CategoryID:      r.PostFormValue("categoryId"),
			Title:           r.PostFormValue("title"),
			Description:     r.PostFormValue("description"),
			IsAdult:         r.PostFormValue("isAdult"),
			Cover:           file,
			CoverFileheader: fileheader,
			IsPublished:     r.PostFormValue("isPublished"),
		}

		storyResponse, err := sh.storyService.AddStory(r.Context(), request)
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		sh.response.SetCode(http.StatusCreated).SetMessage("OK").SetData(storyResponse).JSON(w)
	})
}

func (sh *storyHandler) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, fileheader, _ := r.FormFile("cover")

		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)
		vars := mux.Vars(r)

		request := model.UpdateStoryRequest{
			Slug:            vars["slug"],
			UserID:          user.Id,
			CategoryID:      r.PostFormValue("categoryId"),
			Title:           r.PostFormValue("title"),
			Description:     r.PostFormValue("description"),
			IsAdult:         r.PostFormValue("isAdult"),
			Cover:           file,
			CoverFileheader: fileheader,
			IsPublished:     r.PostFormValue("isPublished"),
		}

		storyResponse, err := sh.storyService.EditStory(r.Context(), request)
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		sh.response.SetCode(http.StatusOK).SetMessage("OK").SetData(storyResponse).JSON(w)
	})
}

func (sh *storyHandler) LoadImageCover() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		imageBytes, err := sh.storyService.LoadStoryImageCover(r.Context(), vars["filename"])
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		w.Write(imageBytes)
	})
}

func (sh *storyHandler) GetBySlug() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		storyResponse, err := sh.storyService.GetStoryBySlug(r.Context(), slug)
		if err != nil {
			sh.handleErr(err).JSON(w)
			return
		}

		sh.response.SetCode(http.StatusOK).SetMessage("OK").SetData(storyResponse).JSON(w)
	})
}

func (sh *storyHandler) handleErr(err error) *response.Response {
	switch {
	case errors.As(err, &validation.Errors{}):
		return sh.response.Error(err).SetCode(http.StatusBadRequest)
	case errors.Is(err, exception.ErrStoryNotFound):
		return sh.response.Error(err).SetCode(http.StatusNotFound)
	case errors.Is(err, exception.ErrCoverImageNotFound):
		return sh.response.Error(err).SetCode(http.StatusNotFound)
	}

	log.Println("[ERROR]", err)
	return sh.response.Error(err).SetCode(http.StatusInternalServerError).SetMessage("internal server error")
}

func NewHandler(storyService story.StoryService) StoryHandler {
	return &storyHandler{
		storyService: storyService,
		response:     new(response.Response),
	}
}
