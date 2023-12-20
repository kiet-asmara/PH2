package handler

import (
	"gin-ex/config"
	"gin-ex/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product entity.Product

	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed bind", "error": err.Error()})
		c.Abort()
		return
	}

	config.DB.Create(&product)

	c.JSON(200, product)
}

func GetAllProducts(c *gin.Context) {
	products := []entity.Product{}

	err := config.DB.Model(&entity.Product{}).Preload("Stores").Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed query", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	product := entity.Product{}
	id := c.Params.ByName("id")
	idint, _ := strconv.Atoi(id)

	if err := config.DB.Preload("Stores").First(&product, "Product_id = ?", idint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed query",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, _ := strconv.Atoi(id)

	body := entity.ProductUpdate{}

	if err := c.BindJSON(&body); err != nil {
		c.Abort()
		return
	}

	var product entity.Product

	if result := config.DB.First(&product, id); result.Error != nil {
		c.Abort()
		return
	}

	product = entity.Product{
		Product_id:  uint(idInt),
		Name:        body.Name,
		Description: body.Description,
		Image_url:   body.Image_url,
		Price:       body.Price,
	}

	config.DB.Save(&product)

	c.JSON(http.StatusOK, &product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, _ := strconv.Atoi(id)
	var product entity.Product

	if err := config.DB.First(&product, idInt).Error; err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted product id=" + id})
}
