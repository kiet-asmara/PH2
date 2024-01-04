package handlers

import (
	"net/http"
	"ngc-11/config"
	"ngc-11/helpers"
	"ngc-11/model"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	products := []model.Product{}

	err := config.DB.Model(&model.Product{}).Find(&products).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, products)
}

func BuyProduct(c echo.Context) error {
	// get product_id & quantity
	var trInput model.TransactionInput
	err := c.Bind(&trInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// get user id
	id, err := helpers.GetUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// get user info
	var user model.User
	err = config.DB.First(&user, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// get product info
	var product model.Product
	err = config.DB.First(&product, trInput.Product_id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}
	if product.Stock < trInput.Quantity {
		return echo.NewHTTPError(http.StatusBadRequest, "Insufficient stock")
	}
	if user.Deposit_amount < product.Price*float32(trInput.Quantity) {
		return echo.NewHTTPError(http.StatusBadRequest, "Insufficient funds")
	}

	// transaction
	tx := config.DB.Begin()

	// reduce deposit
	err = tx.Model(&user).Update("deposit_amount", user.Deposit_amount-product.Price*float32(trInput.Quantity)).Error
	if err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update deposit")
	}

	// reduce stock
	err = tx.Model(&product).Update("stock", product.Stock-trInput.Quantity).Error
	if err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update stock")
	}

	// record transaction
	tr := model.Transaction{
		User_id:      user.User_id,
		Product_id:   trInput.Product_id,
		Quantity:     trInput.Quantity,
		Total_amount: product.Price * float32(trInput.Quantity),
	}

	err = tx.Create(&tr).Error
	if err != nil {
		tx.Rollback()
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to record transaction")
	}

	tx.Commit()

	return c.JSON(http.StatusCreated, map[string]any{
		"message":     "transaction success",
		"transaction": tr,
	})
}
