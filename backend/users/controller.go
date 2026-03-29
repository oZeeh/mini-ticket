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

func NewController(s interfaces.Service) *Controller {
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

// GetUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (c *Controller) FindByID(ctx *gin.Context) {

	id := ctx.Param("id")

	user, err := c.service.Find(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Param id path string true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	err := c.service.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// UpdateUser godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.CreateUserRequest true "User payload"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {

	id := ctx.Param("id")

	var req models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.Find(ctx.Request.Context(), id)
	if err != nil || user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	newUser, err := c.service.Update(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}

func (c *Controller) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/users")

	group.POST("", c.Create)
	group.GET("/:id", c.FindByID)
	group.PUT("/:id", c.Update)
	group.DELETE("/:id", c.Delete)
}
