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
	Title       string `json:"title"`
	ActorID     int    `json:"actor_id"`
	ActorName   string `json:"actor_name"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}

type MovieDirectorView struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	DirectorID   int    `json:"director_id"`
	DirectorName string `json:"director_name"`
	Nationality  string `json:"nationality"`
	Age          int    `json:"age"`
}
