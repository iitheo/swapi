package filmsrepo

import (
	"fmt"
	"github.com/iitheo/theobusha/pkg/app/config/dbconfig"
	"github.com/iitheo/theobusha/pkg/app/config/httpresponses"
	"github.com/iitheo/theobusha/pkg/app/models/filmsmodel"
	"github.com/iitheo/theobusha/pkg/app/services/bushaservices"
	"time"
)

type pgrepoadapter struct{}

func (pgrepoadapter) CreateComment(comment filmsmodel.CreateCommentVM) httpresponses.HttpResponse {
	var (
		resp httpresponses.HttpResponse
	)

	newComment := filmsmodel.FilmsComments{
		Comment:     comment.Comment,
		Title:       comment.Title,
		CommentTime: time.Now().UTC(),
		CommentIP:   comment.CommentIP,
	}

	conn, err := dbconfig.DBConn()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}
	defer conn.Close()

	sqlStatement := `
INSERT INTO films (comment, title, comment_time, comment_ip)
VALUES ($1, $2, $3, $4)
RETURNING id`
	err = conn.QueryRow(sqlStatement, newComment.Comment, newComment.Title, newComment.CommentTime, newComment.CommentIP).Scan(&newComment.ID)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}

	resp.Success = true
	resp.Message = fmt.Sprintf("new comment successfully added.")
	resp.Data = newComment
	return resp
}

func (pgrepoadapter) GetCommentsByFilm(filmTitle string) httpresponses.HttpResponse {
	var (
		resp           httpresponses.HttpResponse
		listOfComments filmsmodel.GetCommentsByFilmVM
	)

	conn, err := dbconfig.DBConn()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, comment, title, comment_time, comment_ip FROM films WHERE title=$1`, filmTitle)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}
	defer rows.Close()
	for rows.Next() {
		var comment filmsmodel.FilmsComments
		err = rows.Scan(&comment.ID, &comment.Comment, &comment.Title, &comment.CommentTime, &comment.CommentIP)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			resp.Error = err.Error()
			return resp
		}
		listOfComments.FilmsComments = append(listOfComments.FilmsComments, comment)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}
	listOfComments.CommentCount = len(listOfComments.FilmsComments)

	if listOfComments.CommentCount > 0 {
		resp.Message = fmt.Sprintf("%d comments successfully fetched", listOfComments.CommentCount)
	} else {
		resp.Message = fmt.Sprintf("no comments were found")
	}

	resp.Success = true
	resp.Data = listOfComments
	return resp
}

func (pgrepoadapter) GetAllFilms() httpresponses.HttpResponse {
	var (
		resp           httpresponses.HttpResponse
		listOfComments []filmsmodel.FilmsComments
		listOfFilms    []filmsmodel.FilmsCommentsVM
		swapiFilmList  filmsmodel.FilmsVM
	)
	mapOfComments := make(map[string][]filmsmodel.CommentInfo)
	swapiChan := make(chan httpresponses.HttpResponse)

	go bushaservices.GetAllSwapiFilms(swapiChan)

	conn, err := dbconfig.DBConn()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id, comment, title, comment_time, comment_ip FROM films")
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}
	defer rows.Close()
	for rows.Next() {
		var comment filmsmodel.FilmsComments
		err = rows.Scan(&comment.ID, &comment.Comment, &comment.Title, &comment.CommentTime, &comment.CommentIP)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			resp.Error = err.Error()
			return resp
		}
		listOfComments = append(listOfComments, comment)
	}

	err = rows.Err()
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		resp.Error = err.Error()
		return resp
	}

	for _, v := range listOfComments {
		commentInfo := filmsmodel.CommentInfo{
			Comment:     v.Comment,
			CommentTime: v.CommentTime,
			CommentIP:   v.CommentIP,
		}
		mapOfComments[v.Title] = append(mapOfComments[v.Title], commentInfo)
	}

	swapiResult := <-swapiChan
	if !swapiResult.Success {
		return swapiResult
	} else {
		if value, ok := swapiResult.Data.(filmsmodel.FilmsVM); ok {
			swapiFilmList = value
		} else {
			resp.Success = false
			resp.Message = fmt.Sprintf("error fetching swapi film list")
			resp.Error = fmt.Sprintf("error fetching swapi film list")
			return resp
		}
	}

	for _, v := range swapiFilmList.Results {
		singleFilmItem := filmsmodel.FilmsCommentsVM{
			Title:          v.Title,
			OpeningCrawl:   v.OpeningCrawl,
			ReleaseDate:    v.ReleaseDate,
			Characters:     v.Characters,
			URL:            v.URL,
			CommentDetails: nil,
			CommentCount:   0,
		}
		if value, ok := mapOfComments[v.Title]; ok {
			singleFilmItem.CommentDetails = value
			singleFilmItem.CommentCount = len(value)
		}
		listOfFilms = append(listOfFilms, singleFilmItem)
	}

	resp.Success = true
	resp.Message = fmt.Sprintf("user record(s) successfully fetched")
	resp.Data = listOfFilms
	return resp
}
