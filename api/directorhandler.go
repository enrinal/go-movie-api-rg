package api

import (
	"github.com/gin-gonic/gin"
	"go-movie-api/model"
	"strconv"
)

func (api *API) GetAllDirector(c *gin.Context) {
	directors, err := api.directorRepo.FetchAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, directors)
}

func (api *API) GetDirectorById(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	director, err := api.directorRepo.FetchByID(intID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, director)
}

func (api *API) CreateDirector(c *gin.Context) {
	var director model.Director
	err := c.ShouldBindJSON(&director)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = api.directorRepo.Store(&director)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Director created successfully"})
}
