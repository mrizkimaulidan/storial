package category

import "net/http"

type CategoryHandler interface {
	GetAllCategory() http.Handler
}
