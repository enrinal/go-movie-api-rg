package api

import (
	"github.com/gin-gonic/gin"
	"go-movie-api/repository"
)

type API struct {
	actorRepo    repository.ActorRepository
	directorRepo repository.DirectorRepository
	gin          *gin.Engine
}

func NewAPI(actorRepo repository.ActorRepository, directorRepo repository.DirectorRepository) API {
	api := API{
		actorRepo:    actorRepo,
		directorRepo: directorRepo,
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
}

func (api *API) Start() {
	api.gin.Run(":8080")
}
