package model

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	Genre       string `json:"genre"`
	Duration    int    `json:"duration"`
	DirectorID  int    `json:"director_id"`
}

type MovieActor struct {
	ID      int `json:"id"`
	MovieID int `json:"movie_id"`
	ActorID int `json:"actor_id"`
}

type MovieActorView struct {
	ID          int    `json:"id"`
	MovieID     int    `json:"movie_id"`
	ActorID     int    `json:"actor_id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}

type MovieDirectorView struct {
	ID          int    `json:"id"`
	MovieID     int    `json:"movie_id"`
	DirectorID  int    `json:"director_id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}
