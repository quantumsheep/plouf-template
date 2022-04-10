package modules_users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/quantumsheep/plouf"
	"github.com/quantumsheep/plouf/example/entities"
	dto "github.com/quantumsheep/plouf/example/modules/users/dto"
)

type UsersController struct {
	plouf.Controller

	UsersService *UsersService
}

func (c *UsersController) InitRoutes(e *echo.Echo) {
	e.POST("/users", c.CreateUser)
	e.GET("/users/:id", c.GetById)
	e.GET("/users", c.GetAll)
}

func (c *UsersController) CreateUser(ctx echo.Context) error {
	var dto dto.CreateUserBodyDTO
	if err := plouf.ValidateAndBind(ctx, &dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := c.UsersService.CreateUser(&entities.User{
		Username: dto.Username,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(201, user)
}

func (c *UsersController) GetById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := c.UsersService.FindById(id)
	if err != nil {
		return err
	}

	return ctx.JSON(200, user)
}

func (c *UsersController) GetAll(ctx echo.Context) error {
	users, err := c.UsersService.Find()
	if err != nil {
		return err
	}

	return ctx.JSON(200, users)
}
