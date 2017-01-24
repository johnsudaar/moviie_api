package models

type Movie struct {
	ID          int64   `json:"id" bson:"movie_id"`
	Title       string  `json:"title" bson:"title"`
	Backdrop    string  `json:"backdrop" bson:"backdrop"`
	Poster      string  `json:"poster" bson:"poster"`
	Overview    string  `json:"overview" bson:"overview"`
	ReleaseDate string  `json:"release_date" bson:"release_date"`
	VoteAverage float64 `json:"vote_average" bson:"vote_average"`
	Seen        bool    `json:"seen" bson:"seen"`
}

/*
private Long id;
private String title;
private String backdrop;
private String poster;
private String overview;
private String releaseDate;
private double voteAverage;
private boolean seen = false;
*/
