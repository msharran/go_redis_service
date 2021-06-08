package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	createKeyJSON     = `{"key":"test-key","value":"test-value"}`
	getSuccessJSON    = `{"key":"test-key","value":"test-value"}` + "\n"
	searchSuccessJSON = `[{"key":"test-key","value":"test-value"}]` + "\n"
)

func TestSetKey(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/set", strings.NewReader(createKeyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.SetKey(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/get/:key")
	c.SetParamNames("key")
	c.SetParamValues("test-key")

	// Assertions
	if assert.NoError(t, h.GetKey(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, getSuccessJSON, rec.Body.String())
	}
}

func TestSearchByPrefix(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/search/search?prefix=test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.Search(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, searchSuccessJSON, rec.Body.String())
	}
}

func TestSearchBySuffix(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/search/search?suffix=key", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.Search(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, searchSuccessJSON, rec.Body.String())
	}
}
