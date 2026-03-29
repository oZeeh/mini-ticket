package tickets

import (
	customErrors "backend/custom_errors"
	"backend/tickets/interfaces"
	"backend/tickets/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Controller struct {
	service interfaces.Service
}

func NewController(service interfaces.Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) RegisterRoutes(r *gin.Engine) {
	tickets := r.Group("/tickets")
	{
		tickets.POST("", ctrl.Create)
		tickets.GET("", ctrl.FindAll)
		tickets.GET("/:id", ctrl.FindByID)
		tickets.GET("/user/:id", ctrl.FindByUser)
		tickets.GET("/technician/:id", ctrl.FindByAssignedTechnician)
		tickets.GET("/done/:id", ctrl.FindDoneTickets)
		tickets.GET("/open/:id", ctrl.FindOpenTickets)
		tickets.PUT("/:id", ctrl.Update)
		tickets.DELETE("/:id", ctrl.Delete)
	}
}

// @Summary		Cria um novo ticket
// @Tags			tickets
// @Accept			json
// @Produce		json
// @Param			ticket	body		models.CreateTicketRequest	true	"Ticket"
// @Success		201		{object}	map[string]primitive.ObjectID
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/tickets [post]
func (ctrl *Controller) Create(c *gin.Context) {
	var req models.CreateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		// TODO: remover quando JWT estiver pronto
		req.UserID = primitive.NewObjectID()
	} else {
		req.UserID = userIDVal.(primitive.ObjectID)
	}

	id, err := ctrl.service.Create(c.Request.Context(), &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// @Summary		Lista todos os tickets
// @Tags			tickets
// @Produce		json
// @Success		200	{array}		models.TicketEntity
// @Failure		500	{object}	map[string]string
// @Security		BearerAuth
// @Router			/tickets [get]
func (ctrl *Controller) FindAll(c *gin.Context) {
	tickets, err := ctrl.service.FindAll(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// @Summary		Busca ticket por ID
// @Tags			tickets
// @Produce		json
// @Param			id	path		string	true	"Ticket ID"
// @Success		200	{object}	models.TicketEntity
// @Failure		400	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/tickets/{id} [get]
func (ctrl *Controller) FindByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	ticket, err := ctrl.service.FindByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, ticket)
}

// @Summary		Busca tickets por usuário
// @Tags			tickets
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{array}		models.TicketEntity
// @Failure		400	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/tickets/user/{id} [get]
func (ctrl *Controller) FindByUser(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	tickets, err := ctrl.service.FindByUser(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// @Summary		Busca tickets por técnico assignado
// @Tags			tickets
// @Produce		json
// @Param			id	path		string	true	"Technician ID"
// @Success		200	{array}		models.TicketEntity
// @Failure		400	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/tickets/technician/{id} [get]
func (ctrl *Controller) FindByAssignedTechnician(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	tickets, err := ctrl.service.FindByAssignedTechnitian(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// @Summary		Busca tickets concluídos por usuário
// @Tags			tickets
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{array}		models.TicketEntity
// @Failure		400	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/tickets/done/{id} [get]
func (ctrl *Controller) FindDoneTickets(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	tickets, err := ctrl.service.FindDoneTickets(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// @Summary		Busca tickets abertos por usuário
// @Tags			tickets
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{array}		models.TicketEntity
// @Failure		400	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/tickets/open/{id} [get]
func (ctrl *Controller) FindOpenTickets(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	tickets, err := ctrl.service.FindOpenTickets(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// @Summary		Atualiza um ticket
// @Tags			tickets
// @Accept			json
// @Produce		json
// @Param			id		path		string						true	"Ticket ID"
// @Param			ticket	body		models.UpdateTicketRequest	true	"Ticket"
// @Success		200		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/tickets/{id} [put]
func (ctrl *Controller) Update(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	var req models.UpdateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	req.ID = id
	userID := c.MustGet("user_id").(primitive.ObjectID)

	if err := ctrl.service.Update(c.Request.Context(), &req, userID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ticket updated"})
}

// @Summary		Deleta um ticket
// @Tags			tickets
// @Produce		json
// @Param			id	path		string	true	"Ticket ID"
// @Success		200	{object}	map[string]string
// @Failure		400	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/tickets/{id} [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.Error(customErrors.ErrBadRequest)
		return
	}

	if err := ctrl.service.Delete(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ticket deleted"})
}
