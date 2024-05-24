package api

import (
	"github.com/gin-gonic/gin"
	"go-movie-api/repository"
)

type API struct {
	actorRepo    repository.ActorRepository
	directorRepo repository.DirectorRepository
	movieRepo    repository.MovieRepository
	gin          *gin.Engine
}

func NewAPI(actorRepo repository.ActorRepository, directorRepo repository.DirectorRepository, movieRepo repository.MovieRepository) API {
	api := API{
		actorRepo:    actorRepo,
		directorRepo: directorRepo,
		movieRepo:    movieRepo,
		gin:          gin.Default(),
	}

	api.setupRouter()

	return api
}

func (api *API) setupRouter() {
	api.gin.GET("/health", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	actor := api.gin.Group("/actors")
	actor.GET("/", api.GetAllActor)
	actor.GET("/:id", api.GetActorById)
	actor.POST("/", api.CreateActor)

	director := api.gin.Group("/directors")
	director.GET("/", api.GetAllDirector)
	director.GET("/:id", api.GetDirectorById)
	director.POST("/", api.CreateDirector)

	movie := api.gin.Group("/movies")
	movie.GET("/", api.GetAllMovie)
	movie.GET("/:id", api.GetMovieById)
	movie.POST("/", api.CreateMovie)
	movie.GET("/:id/actors", api.GetMovieActors)
	movie.GET("/:id/director", api.GetMovieDirectors)
	movie.POST("/:id/actors/:actor_id", api.AddActorToMovie)
}

func (api *API) Start() {
	api.gin.Run(":8080")
}
