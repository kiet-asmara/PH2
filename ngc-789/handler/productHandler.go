package handler

import (
	"gin-ex/config"
	"gin-ex/entity"
	"gin-ex/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// swagger embed files
)

// @Summary      Add Product
// @Description  Add a new product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Product
// @Failure      400  {object}  utils.APIErrorResponse
// @Failure      500  {object}  utils.APIErrorResponse
// @Router       /products [Post]
func AddProduct(c *gin.Context) {
	var product entity.Product

	err := c.BindJSON(&product)
	if err != nil {
		log.Printf("error AddProduct: %s \n", err.Error())
		utils.ErrorMessage(c, &utils.ErrBadRequest)
		return
	}

	config.DB.Create(&product)

	c.JSON(200, product)
}

func GetAllProducts(c *gin.Context) {
	products := []entity.Product{}

	err := config.DB.Model(&entity.Product{}).Preload("Stores").Find(&products).Error
	if err != nil {
		log.Panicln("error GetAllProducts: ", err.Error())
	}

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	product := entity.Product{}
	id := c.Params.ByName("id")
	idint, _ := strconv.Atoi(id)

	if err := config.DB.Preload("Stores").First(&product, "Product_id = ?", idint).Error; err != nil {
		log.Printf("error GetProductById: %v \n", err.Error())
		utils.ErrorMessage(c, &utils.ErrBadRequest)
		return
	}
	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	idInt, _ := strconv.Atoi(id)

	body := entity.ProductUpdate{}

	if err := c.BindJSON(&body); err != nil {
		log.Printf("error UpdateProduct: %s \n", err.Error())
		utils.ErrorMessage(c, &utils.ErrBadRequest)
		return
	}

	var product entity.Product

	if result := config.DB.First(&product, id); result.Error != nil {
		log.Printf("error UpdateProduct: %v \n", result.Error)
		utils.ErrorMessage(c, &utils.ErrBadRequest)
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
		log.Printf("error UpdateProduct: %v \n", err.Error())
		utils.ErrorMessage(c, &utils.ErrDataNotFound)
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted product id=" + id})
}
