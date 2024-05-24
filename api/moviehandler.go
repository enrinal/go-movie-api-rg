package api

import (
	"github.com/gin-gonic/gin"
	"go-movie-api/model"
	"strconv"
)

func (api *API) GetAllMovie(c *gin.Context) {
	movies, err := api.movieRepo.FetchAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, movies)
}

func (api *API) GetMovieById(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	movie, err := api.movieRepo.FetchByID(intID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, movie)
}

func (api *API) CreateMovie(c *gin.Context) {
	var movie model.Movie
	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = api.movieRepo.Store(&movie)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Movie created successfully"})
}

func (api *API) GetMovieActors(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	actors, err := api.movieRepo.FetchMovieWithActors(intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, actors)
}

func (api *API) GetMovieDirectors(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	directors, err := api.movieRepo.FetchMovieWithDirector(intID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, directors)
}

func (api *API) AddActorToMovie(context *gin.Context) {
	movieID := context.Param("id")
	actorID := context.Param("actor_id")

	intMovieID, err := strconv.Atoi(movieID)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	intActorID, err := strconv.Atoi(actorID)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = api.movieRepo.StoreMovieActor(intMovieID, intActorID)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Actor added to movie successfully"})

}
