// Package users is the user feature for the mini-ticket
package users

import (
	customErrors "backend/custom_errors"
	"backend/middlewares"
	"backend/users/enums"
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

func (c *Controller) RegisterPublicRoutes(r *gin.Engine) {
	group := r.Group("/users")
	group.POST("", c.Create)
}

func (c *Controller) RegisterPrivateRoutes(r *gin.RouterGroup) {
	group := r.Group("/users")
	group.GET("/:id", c.FindByID)
	group.PUT("/:id", c.Update)
	group.DELETE("/:id", middlewares.RequireRole(enums.Admin), c.Delete)
}

// Create godoc
// @Summary		Cria um novo usuário
// @Description	Cria um novo usuário no sistema
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body		models.CreateUserRequest	true	"Payload do usuário"
// @Success		201		{object}	models.CreateUserResponse
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/users [post]
func (c *Controller) Create(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(customErrors.ErrBadRequest)
		return
	}

	id, err := c.service.Create(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
}

// FindByID godoc
// @Summary		Busca usuário por ID
// @Tags			users
// @Produce		json
// @Security		BearerAuth
// @Param			id	path		string	true	"ID do usuário"
// @Success		200	{object}	models.User
// @Failure		400	{object}	map[string]string
// @Failure		403	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/users/{id} [get]
func (c *Controller) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.service.Find(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	if user == nil {
		ctx.Error(customErrors.ErrNotFound)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary		Deleta um usuário
// @Tags			users
// @Produce		json
// @Security		BearerAuth
// @Param			id	path	string	true	"ID do usuário"
// @Success		204
// @Failure		400	{object}	map[string]string
// @Failure		403	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/users/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// Update godoc
// @Summary		Atualiza um usuário
// @Tags			users
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			id		path		string						true	"ID do usuário"
// @Param			user	body		models.CreateUserRequest	true	"Payload do usuário"
// @Success		200		{object}	models.User
// @Failure		400		{object}	map[string]string
// @Failure		403		{object}	map[string]string
// @Failure		404		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/users/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(customErrors.ErrBadRequest)
		return
	}

	user, err := c.service.Find(ctx.Request.Context(), id)
	if err != nil {
		ctx.Error(err)
		return
	}

	if user == nil {
		ctx.Error(customErrors.ErrNotFound)
		return
	}

	newUser, err := c.service.Update(ctx.Request.Context(), &req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}
