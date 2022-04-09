package modules_users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/quantumsheep/plouf"
	modules_users_dto "github.com/quantumsheep/plouf/example/modules/users/dto"
	"github.com/stretchr/testify/assert"
)

func TestUsersControllerCreate(t *testing.T) {
	os.Setenv("database_driver", "sqlite")
	os.Setenv("database_path", ":memory:")

	usersModule := &UsersModule{}
	worker, err := plouf.NewWorkerMock(usersModule)
	assert.NoError(t, err)

	dto, err := json.Marshal(modules_users_dto.CreateUserBodyDTO{
		Username: "John Doe",
	})
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(dto)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := worker.NewContext(req, rec)

	if assert.NoError(t, usersModule.UsersController.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestUsersControllerGetCreateUser(t *testing.T) {
	os.Setenv("database_driver", "sqlite")
	os.Setenv("database_path", ":memory:")

	usersModule := &UsersModule{}
	worker, err := plouf.NewWorkerMock(usersModule)
	assert.NoError(t, err)

	// Create user
	{
		dto, err := json.Marshal(modules_users_dto.CreateUserBodyDTO{
			Username: "John Doe",
		})
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(dto)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := worker.NewContext(req, rec)

		if assert.NoError(t, usersModule.UsersController.CreateUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	}

	// Fetch created user (ID = 1)
	{
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		c := worker.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, usersModule.UsersController.GetById(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}
