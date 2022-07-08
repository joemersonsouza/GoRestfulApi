package controllers

import (
	"golearning/api-sample/models"
	"golearning/api-sample/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IShopController interface {
	GetItems(context *gin.Context)
	GetItem(context *gin.Context)
	AddItem(context *gin.Context)
	UpdateItem(context *gin.Context)
}

type ShopController struct {
	ShopService services.IShopService
}

func NewShopController(shopService services.IShopService) IShopController {
	controller := new(ShopController)
	controller.ShopService = shopService

	return controller
}

// GetItems
// @Summary Get all stored items
// @Description get items
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ProductRequest	"List of products"
// @Failure 204 {string} string "No content"
// @Router /items [get]
func (controller *ShopController) GetItems(context *gin.Context) {
	var items = controller.ShopService.GetItems()
	if items != nil {
		context.IndentedJSON(http.StatusOK, items)
		return
	}

	context.IndentedJSON(http.StatusNoContent, gin.H{"message": "No Content"})
}

// GetItem/1
// @Summary Get a specific item
// @Description get an item by ID
// @Accept  json
// @Produce  json
// @Param   id      path   int     true  "Item ID"
// @Success 200 {object} models.ProductRequest	"Product Item"
// @Failure 404 {string} string "Item not found"
// @Router /item/{id} [get]
func (controller *ShopController) GetItem(context *gin.Context) {
	var param = context.Param("id")

	var item, err = controller.ShopService.GetItem(param)

	if &item != nil {
		context.IndentedJSON(http.StatusOK, item)
		return
	}

	context.IndentedJSON(http.StatusNotFound, err.Error())
}

// AddItem
// @Summary Add a new item
// @Description Add a new item
// @Accept  json
// @Produce  json
// @Param   Item  body   models.ProductRequest  true  "Item body"
// @Success 201 {string} string	"ok"
// @Failure 400 {string} string "Invalid body"
// @Failure 500 {string} string "Something goes wrong"
// @Router /item [post]
func (controller *ShopController) AddItem(context *gin.Context) {
	var item models.ProductRequest

	if err := context.BindJSON(&item); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	context.IndentedJSON(http.StatusCreated, controller.ShopService.AddItem(item))
}

// UpdateItem
// @Summary Change an item
// @Description change an existing item
// @Accept  json
// @Produce  json
// @Param   id  body     string  true  "Item ID"
// @Param   Item  body   models.ProductRequest  true  "Item body"
// @Success 201 {string} string	"ok"
// @Failure 400 {string} string "Invalid body"
// @Failure 404 {string} string "Item not found"
// @Failure 500 {string} string "Something goes wrong"
// @Router /item/:{id} [put]
func (controller *ShopController) UpdateItem(context *gin.Context) {
	var item models.ProductRequest
	var param = context.Param("id")

	if err := context.BindJSON(&item); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}

	var result, err = controller.ShopService.UpdateItem(param, item)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	context.IndentedJSON(http.StatusCreated, result)
}
