package modules_users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/quantumsheep/plouf"
	"github.com/stretchr/testify/assert"
)

func TestUsersControllerGetById(t *testing.T) {
	usersController := &UsersController{}
	assert.NoError(t, plouf.Inject(usersController))

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, usersController.GetById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
