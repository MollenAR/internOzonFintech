package handler

import (
	"bytes"
	"encoding/json"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model/mocks"
	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestGetOriginalUrl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := &mocks.ShortUrlUsecase{}
		mockUsecase.On("GetOriginalUrl",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(model.GetOriginalUrlResponse{}, nil)

		testHandler := NewShortUrlHandler(mockUsecase)
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://127.0.0.1:8080/get/save123123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("shortUrl")
		c.SetParamValues("save123123")

		err := testHandler.GetOriginalUrl(c)
		assert.NoError(t, err)
	})

	t.Run("validate length fail", func(t *testing.T) {
		mockUsecase := &mocks.ShortUrlUsecase{}
		mockUsecase.On("GetOriginalUrl",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(model.GetOriginalUrlResponse{}, nil)

		testHandler := NewShortUrlHandler(mockUsecase)
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://127.0.0.1:8080/get/save123123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("shortUrl")
		c.SetParamValues("save123")

		err := testHandler.GetOriginalUrl(c)
		if assert.Error(t, err) {
			assert.Equal(t, errorTypes.ErrWrongShortUrl{Reason: ": must be in a valid format"}, err)
		}
	})

	t.Run("validate alphabet fail", func(t *testing.T) {
		mockUsecase := &mocks.ShortUrlUsecase{}
		mockUsecase.On("GetOriginalUrl",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(model.GetOriginalUrlResponse{}, nil)

		testHandler := NewShortUrlHandler(mockUsecase)
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://127.0.0.1:8080/get/save123123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("shortUrl")
		c.SetParamValues("save123-ВЯ")

		err := testHandler.GetOriginalUrl(c)
		if assert.Error(t, err) {
			assert.Equal(t, errorTypes.ErrWrongShortUrl{Reason: ": must be in a valid format"}, err)
		}
	})

	t.Run("validate empty fail", func(t *testing.T) {
		mockUsecase := &mocks.ShortUrlUsecase{}
		mockUsecase.On("GetOriginalUrl",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(model.GetOriginalUrlResponse{}, nil)

		testHandler := NewShortUrlHandler(mockUsecase)
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://127.0.0.1:8080/get/save123123", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("shortUrl")
		c.SetParamValues("")

		err := testHandler.GetOriginalUrl(c)
		if assert.Error(t, err) {
			assert.Equal(t, errorTypes.ErrWrongShortUrl{Reason: ": cannot be blank"}, err)
		}
	})
}

func TestSaveOriginalUrl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := &mocks.ShortUrlUsecase{}
		mockUsecase.On("SaveOriginalUrl",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string")).Return(model.SaveOriginalUrlResponse{}, nil)

		testHandler := NewShortUrlHandler(mockUsecase)
		e := echo.New()

		originalUrlTest := model.OriginalUrl{
			Url: "ozon.ru/fen",
		}
		body, _ := json.Marshal(originalUrlTest)

		req := httptest.NewRequest(echo.GET, "http://127.0.0.1:8080/save", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := testHandler.SaveOriginalUrl(c)
		assert.NoError(t, err)
	})

	t.Run("validate url fail", func(t *testing.T) {
		mockUsecase := &mocks.ShortUrlUsecase{}
		mockUsecase.On("SaveOriginalUrl",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("string")).Return(model.SaveOriginalUrlResponse{}, nil)

		testHandler := NewShortUrlHandler(mockUsecase)
		e := echo.New()

		originalUrlTest := model.OriginalUrl{
			Url: "ozonru/fen",
		}
		body, _ := json.Marshal(originalUrlTest)

		req := httptest.NewRequest(echo.GET, "http://127.0.0.1:8080/save", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := testHandler.SaveOriginalUrl(c)
		if assert.Error(t, err) {
			assert.Equal(t, errorTypes.ErrWrongOriginalUrl{Reason:": must be a valid URL"}, err)
		}
	})

	t.Run("validate empty url fail", func(t *testing.T) {
		mockUsecase := &mocks.ShortUrlUsecase{}
		mockUsecase.On("SaveOriginalUrl",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("string")).Return(model.SaveOriginalUrlResponse{}, nil)

		testHandler := NewShortUrlHandler(mockUsecase)
		e := echo.New()

		originalUrlTest := model.OriginalUrl{
			Url: "",
		}
		body, _ := json.Marshal(originalUrlTest)

		req := httptest.NewRequest(echo.GET, "http://127.0.0.1:8080/save", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := testHandler.SaveOriginalUrl(c)
		if assert.Error(t, err) {
			assert.Equal(t, errorTypes.ErrWrongOriginalUrl{Reason:": cannot be blank"}, err)
		}
	})
}
