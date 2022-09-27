package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"valo-tips/internal/image"
	"valo-tips/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/olahol/go-imageupload"
	"gorm.io/gorm"
)

func Status() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "api is working")
	}
}

func GetTip() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id, err := strconv.ParseInt(c.Param("id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		tx := c.Get("Tx").(*gorm.DB)

		t := new(model.Tip)
		if err := t.Get(tx, uint(id)); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, t)
	}
}

func PostTip() echo.HandlerFunc {
	host := os.Getenv("AZURE_STORAGE_URL")
	containerName := os.Getenv("AZURE_STORAGE_CONTAINER_NAME")
	return func(c echo.Context) (err error) {
		// upload stand image
		img, err := imageupload.Process(c.Request(), "stand_image")
		if img.ContentType != "image/jpeg" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid image type")
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid image.")
		}
		f_name_stand, err := image.SaveImage(img)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to upload image.")
		}

		// upload aim image
		img, err = imageupload.Process(c.Request(), "aim_image")
		if img.ContentType != "image/jpeg" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid image type")
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid image.")
		}
		f_name_aim, err := image.SaveImage(img)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to upload image.")
		}

		// parse side_id
		side_id, err := strconv.ParseInt(c.FormValue("side_id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid side_id.")
		}

		map_uuid := c.FormValue("map_uuid")
		agent_uuid := c.FormValue("agent_uuid")
		ability_slot := c.FormValue("ability_slot")
		title := c.FormValue("title")
		description := c.FormValue("description")

		// create tip
		t := &model.Tip{
			MapUUID:      map_uuid,
			SideID:       uint(side_id),
			AgentUUID:    agent_uuid,
			AbilitySlot:  ability_slot,
			Title:        title,
			Description:  description,
			StandImgPath: fmt.Sprintf("%s/%s/%s", host, containerName, f_name_stand),
			AimImgPath:   fmt.Sprintf("%s/%s/%s", host, containerName, f_name_aim),
		}

		tx := c.Get("Tx").(*gorm.DB)

		if err := t.Create(tx); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to create tip.")
		}
		return c.JSON(http.StatusOK, t)
	}
}

func GetTips() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		cond := new(model.Tip)
		map_uuid := c.QueryParam("map_uuid")
		side_id, err := strconv.ParseInt(c.QueryParam("side_id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		agent_uuid := c.QueryParam("agent_uuid")
		ability_slot := c.QueryParam("ability_slot")

		cond.MapUUID = map_uuid
		cond.SideID = uint(side_id)
		cond.AgentUUID = agent_uuid
		cond.AbilitySlot = ability_slot

		// fetch tips with conditions
		tips := new(model.Tips)
		if err := tx.Where(cond).Preload("Side").Find(tips).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, tips)
	}
}

func GetSides() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorm.DB)

		sides := new(model.Sides)
		if err := sides.GetAll(tx); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Does not exists.")
		}
		return c.JSON(http.StatusOK, sides)
	}
}
