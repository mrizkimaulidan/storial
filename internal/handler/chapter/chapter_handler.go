package chapter

import (
	"errors"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
	model "github.com/mrizkimaulidan/storial/internal/model/chapter"
	"github.com/mrizkimaulidan/storial/internal/service/chapter"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/chapter"
	storyexception "github.com/mrizkimaulidan/storial/pkg/exception/story"
	jwtpkg "github.com/mrizkimaulidan/storial/pkg/jwt"
	"github.com/mrizkimaulidan/storial/pkg/response"
)

type chapterHandler struct {
	chapterService chapter.ChapterService
	response       *response.Response
}

func (ch *chapterHandler) LikeChapter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		storyID := vars["storyId"]
		chapterID := vars["chapterId"]
		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)

		likedResponse, err := ch.chapterService.LikeChapter(r.Context(), storyID, chapterID, user.Id)
		if err != nil {
			ch.handleErr(err).JSON(w)
			return
		}

		ch.response.SetCode(http.StatusOK).SetMessage("OK").SetData(likedResponse).JSON(w)
	})
}

func (ch *chapterHandler) AddChapter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)

		request := model.CreateChapterRequest{
			UserID:        user.Id,
			StorySlug:     vars["storySlug"],
			Title:         r.PostFormValue("title"),
			Body:          r.PostFormValue("body"),
			AuthorComment: r.PostFormValue("authorComment"),
			IsPublished:   r.PostFormValue("isPublished"),
		}

		chapterResponse, err := ch.chapterService.AddChapter(r.Context(), request)
		if err != nil {
			ch.handleErr(err).JSON(w)
			return
		}

		ch.response.SetCode(http.StatusCreated).SetMessage("OK").SetData(chapterResponse).JSON(w)
	})
}

func (ch *chapterHandler) EditChapter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)

		request := model.UpdateChapterRequest{
			UserID:        user.Id,
			StorySlug:     vars["storySlug"],
			ChapterSlug:   vars["chapterSlug"],
			Title:         r.PostFormValue("title"),
			Body:          r.PostFormValue("body"),
			AuthorComment: r.PostFormValue("authorComment"),
			IsPublished:   r.PostFormValue("isPublished"),
		}

		chapterResponse, err := ch.chapterService.EditChapter(r.Context(), request)
		if err != nil {
			ch.handleErr(err).JSON(w)
			return
		}

		ch.response.SetCode(http.StatusOK).SetMessage("OK").SetData(chapterResponse).JSON(w)
	})
}

func (ch *chapterHandler) DeleteChapter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)
		chapterID := vars["chapterId"]

		chapterResponse, err := ch.chapterService.RemoveChapter(r.Context(), user.Id, chapterID)
		if err != nil {
			ch.handleErr(err).JSON(w)
			return
		}

		ch.response.SetCode(http.StatusOK).SetMessage("OK").SetData(chapterResponse).JSON(w)
	})
}

func (ch *chapterHandler) GetChapter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		storySlug := vars["storySlug"]
		user := r.Context().Value(jwtpkg.CtxKeyUserInformation).(*jwtpkg.CustomClaims)
		chapterSlug := vars["chapterSlug"]

		chapterResponse, err := ch.chapterService.GetChapterByStorySlugAndChapterSlug(r.Context(), user.Id, storySlug, chapterSlug)
		if err != nil {
			ch.handleErr(err).JSON(w)
			return
		}

		ch.response.SetCode(http.StatusOK).SetMessage("OK").SetData(chapterResponse).JSON(w)
	})
}

func (ch *chapterHandler) GetChapters() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		storyID := vars["storyId"]

		chaptersResponse, err := ch.chapterService.GetAllChapterByStoryID(r.Context(), storyID)
		if err != nil {
			ch.handleErr(err).JSON(w)
			return
		}

		ch.response.SetCode(http.StatusOK).SetMessage("OK").SetData(chaptersResponse).JSON(w)
	})
}

func (ch *chapterHandler) handleErr(err error) *response.Response {
	switch {
	case errors.As(err, &validation.Errors{}):
		return ch.response.Error(err).SetCode(http.StatusBadRequest)
	case errors.Is(err, exception.ErrChapterNotFound):
		return ch.response.Error(err).SetCode(http.StatusNotFound)
	case errors.Is(err, storyexception.ErrStoryNotFound):
		return ch.response.Error(err).SetCode(http.StatusNotFound)
	case errors.Is(err, exception.ErrCannotLikeYourOwnChapter):
		return ch.response.Error(err).SetCode(http.StatusBadRequest)
	}

	log.Println("[ERROR]", err)
	return ch.response.Error(err).SetCode(http.StatusInternalServerError).SetMessage("internal server error")
}

func NewHandler(chapterService chapter.ChapterService) ChapterHandler {
	return &chapterHandler{
		chapterService: chapterService,
		response:       new(response.Response),
	}
}
