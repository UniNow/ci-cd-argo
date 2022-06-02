package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockDB = UserCollection{
		Items: []User{},
	}
	//userJSON = `{"data":[]}`
)

func TestUserIndex(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/users")
	h := &Handler{mockDB}

	// Assertions
	if assert.NoError(t, h.UserIndex(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, userJSON, rec.Body.String())
	}
}
