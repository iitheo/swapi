package filmsusecase

import (
	"github.com/iitheo/theobusha/pkg/app/config/httpresponses"
	"github.com/iitheo/theobusha/pkg/app/domains/filmservice"
	"github.com/iitheo/theobusha/pkg/app/models/filmsmodel"
)

type FilmsUsecase interface {
	GetAllFilms() httpresponses.HttpResponse
	GetCommentsByFilm(filmTitle string) httpresponses.HttpResponse
	CreateComment(comment filmsmodel.CreateCommentVM) httpresponses.HttpResponse
}

type filmsUseCase struct {
	filmsRepo filmservice.FilmsService
}

func NewFilmsUseCase(filmsService filmservice.FilmsService) FilmsUsecase {
	return filmsUseCase{filmsRepo: filmsService}
}

func (f filmsUseCase) GetAllFilms() httpresponses.HttpResponse {
	return f.filmsRepo.GetAllFilms()
}

func (f filmsUseCase) GetCommentsByFilm(filmTitle string) httpresponses.HttpResponse {
	return f.filmsRepo.GetCommentsByFilm(filmTitle)
}

func (f filmsUseCase) CreateComment(comment filmsmodel.CreateCommentVM) httpresponses.HttpResponse {
	return f.filmsRepo.CreateComment(comment)
}
