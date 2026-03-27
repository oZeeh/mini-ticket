// Package users is the user feature for the mini-ticket
package users

import (
	"backend/users/interfaces"
	"backend/users/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service interfaces.Service
}

func New(s interfaces.Service) *Controller {
	return &Controller{service: s}
}

// Create godoc
// @Summary Create a user
// @Description Creates a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User payload"
// @Success 201 {object} models.CreateUserResponse
// @Router /users [post]
func (c *Controller) Create(ctx *gin.Context) {
	var req models.CreateUserRequest
	println("comecei")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := c.service.Create(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id.Hex(),
	})
}

func (c *Controller) RegisterRoutes(r *gin.Engine) {

	group := r.Group("/users")

	group.POST("", c.Create)
}
