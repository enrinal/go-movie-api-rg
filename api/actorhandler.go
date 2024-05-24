package api

import (
	"github.com/gin-gonic/gin"
	"go-movie-api/model"
	"strconv"
)

func (api *API) GetAllActor(c *gin.Context) {
	actors, err := api.actorRepo.FetchAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, actors)
}

func (api *API) GetActorById(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	actor, err := api.actorRepo.FetchByID(intID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, actor)
}

func (api *API) CreateActor(c *gin.Context) {
	var actor model.Actor
	err := c.ShouldBindJSON(&actor)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = api.actorRepo.Store(&actor)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Actor created successfully"})
}
