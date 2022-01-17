package filmservice

import (
	"github.com/iitheo/theobusha/pkg/app/config/httpresponses"
	"github.com/iitheo/theobusha/pkg/app/models/filmsmodel"
)

type FilmsService interface {
	GetAllFilms() httpresponses.HttpResponse
	CreateComment(comment filmsmodel.CreateCommentVM) httpresponses.HttpResponse
	GetCommentsByFilm(movieUrl string) httpresponses.HttpResponse
}
