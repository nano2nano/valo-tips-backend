package api

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	cloud "valo-tips/internal/cloud/dropbox"
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
	return func(c echo.Context) (err error) {
		img, err := imageupload.Process(c.Request(), "stand_img")
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}
		f_name_stand, err := image.SaveImage(img)
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}

		img, err = imageupload.Process(c.Request(), "aim_img")
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}
		f_name_aim, err := image.SaveImage(img)
		if err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}

		side_id, err := strconv.ParseInt(c.FormValue("side_id"), 0, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id.")
		}
		map_uuid := c.FormValue("map_uuid")
		agent_uuid := c.FormValue("agent_uuid")
		ability_slot := c.FormValue("ability_slot")
		title := c.FormValue("title")
		description := c.FormValue("description")
		t := &model.Tip{
			MapUUID:      map_uuid,
			SideID:       uint(side_id),
			AgentUUID:    agent_uuid,
			AbilitySlot:  ability_slot,
			Title:        title,
			Description:  description,
			StandImgPath: f_name_stand,
			AimImgPath:   f_name_aim,
		}
		tx := c.Get("Tx").(*gorm.DB)

		if err := t.Create(tx); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
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

func PostImg() echo.HandlerFunc {
	return func(c echo.Context) error {
		img, err := imageupload.Process(c.Request(), "file")
		if img.ContentType != "image/jpeg" {
			return c.String(http.StatusBadRequest, "only 'png' image")
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		thumb, err := imageupload.ThumbnailPNG(img, 896, 504)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		f_name := fmt.Sprintf("%s.jpeg", time.Now().Format("20060102150405"))
		if err := cloud.Upload(f_name, bytes.NewReader(thumb.Data)); err != nil {
			return c.JSON(http.StatusBadGateway, err)
		}

		return c.String(http.StatusOK, f_name)
	}
}

func GetImg() echo.HandlerFunc {
	return func(c echo.Context) error {
		f_name := c.Param("name")
		bs, err := cloud.Download(f_name)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		i := &imageupload.Image{
			Filename:    f_name,
			ContentType: "image/jpeg",
			Data:        bs,
			Size:        len(bs),
		}
		i.Write(c.Response().Writer)

		return nil
	}
}
