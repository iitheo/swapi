package filmsmodel

import "time"

type PeopleItem struct {
	Name      string        `json:"name"`
	Height    string        `json:"height"`
	HeightInt int           `json:"-"`
	Mass      string        `json:"mass"`
	HairColor string        `json:"hair_color"`
	SkinColor string        `json:"skin_color"`
	EyeColor  string        `json:"eye_color"`
	BirthYear string        `json:"birth_year"`
	Gender    string        `json:"gender"`
	Homeworld string        `json:"homeworld"`
	Films     []string      `json:"films"`
	Species   []interface{} `json:"species"`
	Vehicles  []string      `json:"vehicles"`
	Starships []string      `json:"starships"`
	Created   time.Time     `json:"created"`
	Edited    time.Time     `json:"edited"`
	URL       string        `json:"url"`
}

type People struct {
	//Count    int         `json:"count"`
	//Next     string      `json:"next"`
	//Previous interface{} `json:"previous"`
	Results []PeopleItem `json:"results"`
}

type PeopleItemVM struct {
	CharactersCount  int          `json:"characters_count"`
	TotalHeightCM    string       `json:"total_height_cm"`
	TotalHeightFtIn  string       `json:"total_height_ft_in"`
	ListOfCharacters []PeopleItem `json:"list_of_characters"`
}

type Films struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Title        string    `json:"title"`
		EpisodeID    int       `json:"episode_id"`
		OpeningCrawl string    `json:"opening_crawl"`
		Director     string    `json:"director"`
		Producer     string    `json:"producer"`
		ReleaseDate  string    `json:"release_date"`
		Characters   []string  `json:"characters"`
		Planets      []string  `json:"planets"`
		Starships    []string  `json:"starships"`
		Vehicles     []string  `json:"vehicles"`
		Species      []string  `json:"species"`
		Created      time.Time `json:"created"`
		Edited       time.Time `json:"edited"`
		URL          string    `json:"url"`
	} `json:"results"`
}

type FilmItem struct {
	Title        string   `json:"title"`
	OpeningCrawl string   `json:"opening_crawl"`
	ReleaseDate  string   `json:"release_date"`
	Characters   []string `json:"characters"`
	URL          string   `json:"url"`
}

type FilmsVM struct {
	Results []FilmItem `json:"results"`
}

type FilmsComments struct {
	ID          int       `json:"id"`
	Comment     string    `json:"comment"`
	Title       string    `json:"title"`
	CommentTime time.Time `json:"comment_time"`
	CommentIP   string    `json:"comment_ip"`
}

type GetCommentsByFilmVM struct {
	CommentCount  int             `json:"commentCount"`
	FilmsComments []FilmsComments `json:"filmsComments"`
}

type CommentInfo struct {
	Comment     string    `json:"comment"`
	CommentTime time.Time `json:"comment_time"`
	CommentIP   string    `json:"comment_ip"`
}

type FilmsCommentsVM struct {
	Title          string        `json:"title"`
	OpeningCrawl   string        `json:"opening_crawl"`
	ReleaseDate    string        `json:"release_date"`
	Characters     []string      `json:"characters"`
	URL            string        `json:"url"`
	CommentDetails []CommentInfo `json:"comment_details"`
	CommentCount   int           `json:"comment_count"`
}

type CreateCommentVM struct {
	Title     string `json:"title"`
	Comment   string `json:"comment"`
	CommentIP string
}
