package repository

import (
	"go-movie-api/model"
	"gorm.io/gorm"
)

type MovieRepository interface {
	FetchAll() ([]model.Movie, error)
	FetchByID(id int) (*model.Movie, error)
	Store(m *model.Movie) error
	FetchMovieWithDirector(id int) (*model.MovieDirectorView, error)
	FetchMovieWithActors(id int) ([]model.MovieActorView, error)
	StoreMovieActor(movieID int, actorID int) error
}

type movieRepoImpl struct {
	db *gorm.DB
}

func NewMovieRepo(db *gorm.DB) *movieRepoImpl {
	return &movieRepoImpl{db}
}

func (m *movieRepoImpl) FetchAll() ([]model.Movie, error) {
	var movies []model.Movie
	result := m.db.Find(&movies)
	if result.Error != nil {
		return nil, result.Error
	}

	return movies, nil
}

func (m *movieRepoImpl) FetchByID(id int) (*model.Movie, error) {
	var movie model.Movie
	result := m.db.First(&movie, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &movie, nil
}

func (m *movieRepoImpl) Store(movie *model.Movie) error {
	result := m.db.Create(movie)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *movieRepoImpl) FetchMovieWithDirector(id int) (*model.MovieDirectorView, error) {
	var movieDirector model.MovieDirectorView

	result := m.db.Table("movies").
		Select("movies.id, movies.title, directors.id as director_id, directors.name as director_name, directors.nationality as nationality, directors.age as age").
		Joins("INNER JOIN directors ON movies.director_id = directors.id").
		Where("movies.id = ?", id).
		Scan(&movieDirector)

	// using raw sql
	//result := m.db.Raw("SELECT movies.id, movies.title, directors.id as director_id, directors.name as director_name, directors.nationality, directors.age FROM movies INNER JOIN directors ON movies.director_id = directors.id WHERE movies.id = ?", id).Scan(&movieDirector)

	if result.Error != nil {
		return nil, result.Error
	}

	return &movieDirector, nil
}

func (m *movieRepoImpl) FetchMovieWithActors(id int) ([]model.MovieActorView, error) {
	var movieActor []model.MovieActorView

	result := m.db.Table("movies").
		Select("movies.id, movies.title, actors.id as actor_id, actors.name as actor_name, actors.nationality as nationality, actors.age as age").
		Joins("JOIN movie_actors ON movies.id = movie_actors.movie_id").
		Joins("JOIN actors ON movie_actors.actor_id = actors.id").
		Where("movies.id = ?", id).
		Scan(&movieActor)

	// using raw sql
	//result := m.db.Raw("SELECT movies.id, movies.title, actors.id as actor_id, actors.name as actor_name, actors.nationality, actors.age FROM movies JOIN movie_actors ON movies.id = movie_actors.movie_id JOIN actors ON movie_actors.actor_id = actors.id WHERE movies.id = ?", id).Scan(&movieActor)

	if result.Error != nil {
		return nil, result.Error
	}

	return movieActor, nil
}

func (m *movieRepoImpl) StoreMovieActor(movieID int, actorID int) error {
	movieActor := model.MovieActor{
		MovieID: movieID,
		ActorID: actorID,
	}

	result := m.db.Create(&movieActor)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
