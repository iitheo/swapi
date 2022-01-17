package filmscontroller

import (
	"fmt"
	"github.com/iitheo/theobusha/pkg/api/libs/httplibraries"
	"github.com/iitheo/theobusha/pkg/app/config/httpresponses"
	"github.com/iitheo/theobusha/pkg/app/domains/filmservice"
	"github.com/iitheo/theobusha/pkg/app/models/filmsmodel"
	"github.com/iitheo/theobusha/pkg/app/services/bushaservices"
	"github.com/iitheo/theobusha/pkg/app/services/helperservices"
	"github.com/iitheo/theobusha/pkg/app/usecases/filmsusecase"
	"github.com/iitheo/theobusha/pkg/repositories/filmsrepo"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var (
	filmsService = filmsrepo.NewFilmsRepo("pgdb").(filmservice.FilmsService)
	filmsUseCase = filmsusecase.NewFilmsUseCase(filmsService)
)

func GetCommentsByFilm(res http.ResponseWriter, req *http.Request) {
	c := httplibraries.C{Req: req, Res: res}
	var (
		resp httpresponses.HttpResponse
	)
	filmTitle := c.Params("filmTitle")
	if filmTitle == "" {
		resp.Success = false
		resp.Message = fmt.Sprintf("film title is required")
		resp.Error = fmt.Sprintf("film title is required")
		httplibraries.Response400(res, resp)
		return
	}

	resp = filmsUseCase.GetCommentsByFilm(filmTitle)

	httplibraries.Response(res, resp)
}

func GetCharactersByFilm(res http.ResponseWriter, req *http.Request) {
	c := httplibraries.C{Req: req, Res: res}
	var (
		resp               httpresponses.HttpResponse
		swapiFilmList      filmsmodel.FilmsVM
		swapiCharacterList filmsmodel.People
		charactersPerMovie filmsmodel.PeopleItemVM
		actualFilm         filmsmodel.FilmItem
		characterCount     int     = 0
		totalHeightCm      float64 = 0.0
		sortBy                     = req.URL.Query().Get("sortby")
		filterBy                   = req.URL.Query().Get("filterby")
		sortType                   = req.URL.Query().Get("sorttype")
	)

	mapOfFilms := make(map[string]filmsmodel.FilmItem)
	mapOfCharacters := make(map[string]filmsmodel.PeopleItem)

	swapiFilmsChan := make(chan httpresponses.HttpResponse)
	swapiCharactersChan := make(chan httpresponses.HttpResponse)

	filmTitle := c.Params("filmTitle")
	if filmTitle == "" {
		resp.Success = false
		resp.Message = fmt.Sprintf("film title is required")
		resp.Error = fmt.Sprintf("film title is required")
		httplibraries.Response400(res, resp)
		return
	}

	go bushaservices.GetAllSwapiFilms(swapiFilmsChan)
	go bushaservices.GetAllSwapiCharacters(swapiCharactersChan)

	swapiFilmsResult := <-swapiFilmsChan
	if !swapiFilmsResult.Success {

	} else {
		if value, ok := swapiFilmsResult.Data.(filmsmodel.FilmsVM); ok {
			swapiFilmList = value

		} else {
			resp.Success = false
			resp.Message = fmt.Sprintf("error fetching swapi film list")
			resp.Error = fmt.Sprintf("error fetching swapi film list")
		}
	}

	swapiCharactersResult := <-swapiCharactersChan
	if !swapiCharactersResult.Success {

	} else {
		if value, ok := swapiCharactersResult.Data.(filmsmodel.People); ok {
			swapiCharacterList = value

		} else {
			resp.Success = false
			resp.Message = fmt.Sprintf("error fetching swapi film list")
			resp.Error = fmt.Sprintf("error fetching swapi film list")

		}
	}

	for _, v := range swapiFilmList.Results {
		mapOfFilms[v.Title] = v
	}

	for _, v := range swapiCharacterList.Results {
		mapOfCharacters[v.URL] = v
	}

	if value, ok := mapOfFilms[filmTitle]; ok {
		actualFilm = value
	} else {
		resp.Success = false
		resp.Message = fmt.Sprintf("Movie %s could not be found", filmTitle)
		resp.Error = fmt.Sprintf("Movie %s could not be found", filmTitle)
		httplibraries.Response404(res, resp)
		return
	}

	for _, v := range actualFilm.Characters {
		if value, ok := mapOfCharacters[v]; ok {
			charactersPerMovie.ListOfCharacters = append(charactersPerMovie.ListOfCharacters, value)
		}
	}

	resp.Success = true

	if len(charactersPerMovie.ListOfCharacters) > 0 {
		resp.Message = fmt.Sprintf("%d characters were found", characterCount)

		switch filterBy {
		case "male":
			var listOfMaleCharacters []filmsmodel.PeopleItem
			for _, v := range charactersPerMovie.ListOfCharacters {
				if strings.TrimSpace(strings.ToLower(v.Gender)) == "male" {
					listOfMaleCharacters = append(listOfMaleCharacters, v)
				}
				charactersPerMovie.ListOfCharacters = listOfMaleCharacters
			}
		case "female":
			var listOfFemaleCharacters []filmsmodel.PeopleItem
			for _, v := range charactersPerMovie.ListOfCharacters {
				if strings.TrimSpace(strings.ToLower(v.Gender)) == "female" {
					listOfFemaleCharacters = append(listOfFemaleCharacters, v)
				}
				charactersPerMovie.ListOfCharacters = listOfFemaleCharacters
			}
		case "n/a":
			var listOfNA []filmsmodel.PeopleItem
			for _, v := range charactersPerMovie.ListOfCharacters {
				if strings.TrimSpace(strings.ToLower(v.Gender)) == "n/a" {
					listOfNA = append(listOfNA, v)
				}
				charactersPerMovie.ListOfCharacters = listOfNA
			}
		default:
			//do nothing
		}

		switch sortBy {
		case "name":
			if sortType == "asc" {
				sort.Slice(charactersPerMovie.ListOfCharacters, func(i, j int) bool {
					return charactersPerMovie.ListOfCharacters[i].Name < charactersPerMovie.ListOfCharacters[j].Name
				})
			} else if sortType == "desc" {
				sort.Slice(charactersPerMovie.ListOfCharacters, func(i, j int) bool {
					return charactersPerMovie.ListOfCharacters[i].Name > charactersPerMovie.ListOfCharacters[j].Name
				})
			}

		case "gender":
			if sortType == "asc" {
				sort.Slice(charactersPerMovie.ListOfCharacters, func(i, j int) bool {
					return charactersPerMovie.ListOfCharacters[i].Gender < charactersPerMovie.ListOfCharacters[j].Gender
				})
			} else if sortType == "desc" {
				sort.Slice(charactersPerMovie.ListOfCharacters, func(i, j int) bool {
					return charactersPerMovie.ListOfCharacters[i].Gender > charactersPerMovie.ListOfCharacters[j].Gender
				})
			}

		case "height":
			if sortType == "asc" {
				sort.Slice(charactersPerMovie.ListOfCharacters, func(i, j int) bool {
					return charactersPerMovie.ListOfCharacters[i].HeightInt < charactersPerMovie.ListOfCharacters[j].HeightInt
				})
			} else if sortType == "desc" {
				sort.Slice(charactersPerMovie.ListOfCharacters, func(i, j int) bool {
					return charactersPerMovie.ListOfCharacters[i].HeightInt > charactersPerMovie.ListOfCharacters[j].HeightInt
				})
			}

		default:
			//do nothing
		}
		charactersPerMovie.CharactersCount = len(charactersPerMovie.ListOfCharacters)

		for _, v := range charactersPerMovie.ListOfCharacters {
			if convertedNumber, err := strconv.ParseFloat(v.Height, 64); err == nil {
				totalHeightCm += convertedNumber
				v.HeightInt = int(convertedNumber)
			}
			characterCount++
		}
		charactersPerMovie.CharactersCount = characterCount

		totalHeightFt := totalHeightCm * 0.0328084
		totalHeightFtToInt := int(totalHeightFt)
		totalHeightFtToInches := (totalHeightFt - float64(totalHeightFtToInt)) * 12

		charactersPerMovie.TotalHeightCM = fmt.Sprintf("%.2fcm", totalHeightCm)
		charactersPerMovie.TotalHeightFtIn = fmt.Sprintf("%dft and %.2finches", totalHeightFtToInt, totalHeightFtToInches)

	} else {
		resp.Message = fmt.Sprintf("no characters were found")
	}
	resp.Data = charactersPerMovie

	httplibraries.Response(res, resp)
}

func GetAllFilms(res http.ResponseWriter, req *http.Request) {

	resp := filmsUseCase.GetAllFilms()

	httplibraries.Response(res, resp)
}

func CreateComment(res http.ResponseWriter, req *http.Request) {
	c := httplibraries.C{Req: req, Res: res}
	var (
		newComment filmsmodel.CreateCommentVM
		resp       httpresponses.HttpResponse
	)
	err := c.BindJSON(&newComment)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		httplibraries.Response400(res, resp)
		return
	}
	if (newComment.Comment == "") || (newComment.Title == "") {
		resp.Success = false
		resp.Message = fmt.Sprintf("both comment and film name are required")
		resp.Error = fmt.Sprintf("both comment and film name are required")
		httplibraries.Response400(res, resp)
		return
	}

	if len(newComment.Comment) > 500 {
		resp.Success = false
		resp.Message = fmt.Sprintf("Maximum allowed comment length is 500 characters. You entered %d characters", len(newComment.Comment))
		resp.Error = fmt.Sprintf("Maximum allowed comment length is 500 characters. You entered %d characters", len(newComment.Comment))
		httplibraries.Response400(res, resp)
		return
	}

	newComment.CommentIP, _ = helperservices.GetIP(req)

	resp = filmsUseCase.CreateComment(newComment)
	if !resp.Success {
		httplibraries.Response500(res, resp)
		return
	}

	httplibraries.Response(res, resp)
}
