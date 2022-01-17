package httproutes

import (
	"github.com/gorilla/mux"
	"github.com/iitheo/theobusha/pkg/api/controllers/v1/filmscontroller"
	"github.com/iitheo/theobusha/pkg/api/middleware"
	"github.com/urfave/negroni"
)

func Router() *negroni.Negroni {
	route := mux.NewRouter()

	n := negroni.Classic()
	n.Use(middleware.Cors())
	n.UseHandler(route)

	//BASE ROUTE
	//route.HandleFunc("/v1", homeHandler)

	//*****************
	// FILMS ROUTES
	//*****************
	filmsRoute := route.PathPrefix("/v1/films").Subrouter()
	filmsRoute.HandleFunc("/getallfilms", filmscontroller.GetAllFilms).Methods("GET")
	filmsRoute.HandleFunc("/createcomment", filmscontroller.CreateComment).Methods("POST")
	filmsRoute.HandleFunc("/getcommentsbyfilm/{filmTitle}", filmscontroller.GetCommentsByFilm).Methods("GET")
	filmsRoute.HandleFunc("/getcharactersbyfilm/{filmTitle}", filmscontroller.GetCharactersByFilm).Methods("GET")

	return n
}
