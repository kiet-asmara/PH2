package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ngc-11/config"
	"ngc-11/model"
	"os"

	"github.com/labstack/echo/v4"
)

func GetStores(c echo.Context) error {
	stores := []model.Store{}

	err := config.DB.Model(&model.Store{}).Find(&stores).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, stores)
}

func GetStoreByID(c echo.Context) error {
	// get id
	id := c.Param("id")

	var store model.Store

	// query db
	rows := config.DB.First(&store, id)
	if rows.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	// get weather from API
	body, err := GetWeather(store.Latitude, store.Longitude)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"store":   store,
		"weather": body,
	})
}

func GetWeather(lat string, lon string) (map[string]any, error) {

	// make request
	url := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?lat=%s&lon=%s", lat, lon)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return map[string]any{}, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RapidAPIKey"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("RapidAPIHost"))

	// send HTTP req w/default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return map[string]any{}, err
	}

	defer res.Body.Close()

	// read body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return map[string]any{}, err
	}

	weather := map[string]any{}

	json.Unmarshal(body, &weather)

	return weather, nil
}
