// auth/controller.go
package auth

import (
	"backend/auth/interfaces"
	"backend/auth/models"
	customErrors "backend/custom_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service interfaces.Service
}

func NewController(service interfaces.Service) *Controller {
	return &Controller{service: service}
}

// Login godoc
// @Summary		Login
// @Description	Autentica o usuário e retorna o token JWT
// @Tags			auth
// @Accept			json
// @Produce		json
// @Param			login	body		models.LoginRequest		true	"Credenciais"
// @Success		200		{object}	models.LoginResponse
// @Failure		400		{object}	map[string]string
// @Failure		403		{object}	map[string]string
// @Failure		404		{object}	map[string]string
// @Router			/auth/login [post]
func (ctrl *Controller) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	resp, err := ctrl.service.Login(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (ctrl *Controller) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/auth")
	group.POST("/login", ctrl.Login)
}
