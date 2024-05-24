package main

import (
	"database/sql"
	"fmt"
	"go-movie-api/api"
	"go-movie-api/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
}

func Connect(credential Credential) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", credential.Host, credential.Username, credential.Password, credential.DatabaseName, credential.Port)

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func ConnectGorm(credential Credential) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", credential.Host, credential.Username, credential.Password, credential.DatabaseName, credential.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func SQLExecute(db *sql.DB) error {
	//create table actor
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS actors (id SERIAL PRIMARY KEY, name VARCHAR(255), nationality VARCHAR(255), age INT)")
	if err != nil {
		return err
	}

	//create table director
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS directors(id SERIAL PRIMARY KEY, name VARCHAR(255), nationality VARCHAR(255), age INT)")
	if err != nil {
		return err
	}

	// create table movie
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS movies(id SERIAL PRIMARY KEY, title VARCHAR(255), release_year INT, genre VARCHAR(255), duration INT, director_id INT)")
	if err != nil {
		return err
	}

	// create table movie_actor
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS movie_actors(id SERIAL PRIMARY KEY, movie_id INT, actor_id INT)")
	if err != nil {
		return err
	}

	return nil
}

func Reset(db *sql.DB, table string) error {
	_, err := db.Exec("TRUNCATE " + table)
	if err != nil {
		return err
	}

	_, err = db.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	dbCredential := Credential{
		Host:         "localhost",
		Username:     "admin",
		Password:     "admin123!",
		DatabaseName: "movie_db",
		Port:         5432,
	}

	dbConn, err := Connect(dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	gormDbConn, err := ConnectGorm(dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	err = SQLExecute(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	actorRepo := repository.NewActorRepo(dbConn)
	directorRepo := repository.NewDirectorRepo(gormDbConn)
	movieRepo := repository.NewMovieRepo(gormDbConn)

	mainAPI := api.NewAPI(actorRepo, directorRepo, movieRepo)
	mainAPI.Start()
}
