package response

import (
	"errors"
	"net/http"

	"github.com/lapitskyss/go_elasticsearch/internal/pkg/render"
)

func SendError(w http.ResponseWriter, err error) {
	var badRequestError *BadRequestError

	if errors.As(err, &badRequestError) {
		render.BadRequestError(w, err)
		return
	}

	var notFoundError *NotFoundError
	if errors.As(err, &notFoundError) {
		render.NotFoundError(w)
		return
	}

	render.InternalServerError(w)
}
