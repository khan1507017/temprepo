package v1

import "C"
import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tempApi/dto"
	"tempApi/helper"
)

type CacheControllerInf interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type CacheControllerInstance struct {
}

func CacheController() CacheControllerInf {
	return new(CacheControllerInstance)
}
func (CacheControllerInstance) Create(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}
	err = helper.ValidateInput(value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if helper.Init == true {
		c.JSON(http.StatusBadRequest, "already initialized, cannot recreate")
	} else {
		helper.Val = value
		c.JSON(http.StatusOK, "value set")
		helper.Init = true
	}
	return nil
}

func (CacheControllerInstance) Update(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}
	if helper.Init == false {
		return c.JSON(http.StatusBadRequest, "not initialized")
	}
	err = helper.ValidateInput(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	helper.Val = value
	return nil
}

func (CacheControllerInstance) Delete(c echo.Context) error {
	helper.Init = false
	return nil
}
func (CacheControllerInstance) Get(c echo.Context) error {
	if helper.Init == false {
		c.JSON(http.StatusBadRequest, "value not set")
	} else {
		c.JSON(http.StatusOK, helper.Val)
	}
	return nil
}
