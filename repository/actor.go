package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-movie-api/model"
)

type ActorRepository interface {
	FetchAll() ([]model.Actor, error)
	FetchByID(id int) (*model.Actor, error)
	Store(s *model.Actor) error
}

type actorRepoImpl struct {
	db *sql.DB
}

func NewActorRepo(db *sql.DB) *actorRepoImpl {
	return &actorRepoImpl{db}
}

func (s *actorRepoImpl) FetchAll() ([]model.Actor, error) {
	row, err := s.db.Query("SELECT id, name, nationality, age FROM actors")
	if err != nil {
		return nil, err

	}

	var actors []model.Actor

	for row.Next() {
		var actor model.Actor
		err = row.Scan(&actor.ID, &actor.Name, &actor.Nationality, &actor.Age)
		if err != nil {
			return nil, err
		}

		actors = append(actors, actor)
	}

	return actors, nil
}

func (s *actorRepoImpl) FetchByID(id int) (*model.Actor, error) {
	row := s.db.QueryRow("SELECT id, name, nationality, age FROM actors WHERE id = $1", id)

	var actor model.Actor

	err := row.Scan(&actor.ID, &actor.Name, &actor.Nationality, &actor.Age)
	if err != nil {
		return nil, err
	}

	return &actor, nil
}

func (s *actorRepoImpl) Store(actor *model.Actor) error {
	_, err := s.db.Exec("INSERT INTO actors (name, nationality, age) VALUES ($1, $2, $3)", actor.Name, actor.Nationality, actor.Age)
	if err != nil {
		return err
	}

	return nil
}
