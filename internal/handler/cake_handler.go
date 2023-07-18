package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cake-store/internal/cake"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CakeHandler struct {
	service cake.CakeService
}

func NewCakeHandler(service cake.CakeService) *CakeHandler {
	return &CakeHandler{
		service: service,
	}
}

func (h *CakeHandler) CreateCake(c echo.Context) error {
	req := new(cake.CreateCakeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errors})
	}

	cake := &cake.CreateCakeRequest{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}

	createdCake, err := h.service.CreateCake(cake)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":       "failed to create cake",
			"description": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, createdCake)
}

func (h *CakeHandler) GetAllCakes(c echo.Context) error {
	cakes, err := h.service.GetAllCakes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve cakes",
		})
	}

	return c.JSON(http.StatusOK, cakes)
}

func (h *CakeHandler) GetCakeByID(c echo.Context) error {

	idStr := c.Param("id")
	if idStr == "" {
		log.Println("[GetCakeByID][Error] Missing id parameter")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":       "Missing id parameter",
			"description": "Please provide a valid id",
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Failed to convert id to int: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":       "Failed to convert id to int",
			"description": err.Error(),
		})
	}

	req := &cake.GetCakeByIDRequest{
		ID: id,
	}

	cake, err := h.service.GetCakeByID(req.ID)
	if err != nil {
		log.Println("Failed to retrieve cake: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":       "Failed to retrieve cake",
			"description": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, cake)
}

func (h *CakeHandler) UpdateCake(c echo.Context) error {
	idStr := c.Param("id")
	if idStr == "" {
		log.Println("[UpdateCake][Error] Missing id parameter")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":       "Missing id parameter",
			"description": "Please provide a valid id",
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Failed to convert id to int: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":       "Failed to convert id to int",
			"description": err.Error(),
		})
	}

	cakeExisting, err := h.service.GetCakeByID(id)
	if err != nil {
		log.Println("Failed to retrieve cake: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":       "Failed to retrieve cake",
			"description": err.Error(),
		})
	}

	req := new(cake.UpdateCakeRequest)
	if req.Title == "" {
		req.Title = cakeExisting.Title
	}

	if req.Description == "" {
		req.Description = cakeExisting.Description
	}

	if req.Rating == 0 {
		req.Rating = cakeExisting.Rating
	}

	if req.Image == "" {
		req.Image = cakeExisting.Image
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errors})
	}

	cakeUpdate := &cake.UpdateCakeRequest{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}

	updatedCake, err := h.service.UpdateCake(cakeUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":       "failed to update cake",
			"description": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, updatedCake)
}

func (h *CakeHandler) DeleteCake(c echo.Context) error {
	idStr := c.Param("id")

	if idStr == "" {
		log.Println("[DeleteCake][Error] Missing id parameter")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":       "Missing id parameter",
			"description": "Please provide a valid id",
		})
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Println("Failed to convert id to int: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":       "Failed to convert id to int",
			"description": err.Error(),
		})
	}

	err = h.service.DeleteCake(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":       "failed to delete cake",
			"description": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Cake deleted successfully",
	})
}
