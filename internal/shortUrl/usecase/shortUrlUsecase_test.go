package usecase

import (
	"context"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestSaveOriginalUrl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := &mocks.ShortUrlRepository{}
		mockRepo.On("SaveOriginalUrl",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("string")).Return("qwerUNB12_", nil)

		testUsecase := NewShortUrlUsecase(mockRepo)

		testResp, err := testUsecase.SaveOriginalUrl(context.TODO(), "ozon.ru/fen")
		assert.NoError(t, err)
		expectedResp := model.SaveOriginalUrlResponse{
			Status:   http.StatusOK,
			ShortUrl: "qwerUNB12_",
		}
		assert.Equal(t, expectedResp, testResp)
	})

	t.Run("fail", func(t *testing.T) {
		mockRepo := &mocks.ShortUrlRepository{}
		mockRepo.On("SaveOriginalUrl",
			mock.AnythingOfType("*context.emptyCtx"),
			mock.AnythingOfType("string")).Return("", errors.New("test error"))

		testUsecase := NewShortUrlUsecase(mockRepo)

		testResp, err := testUsecase.SaveOriginalUrl(context.TODO(), "ozon.ru/fen")
		if assert.Error(t, err) {
			assert.IsType(t, errors.Wrap(errors.New("test error"), ""), err)
		}
		expectedResp := model.SaveOriginalUrlResponse{}
		assert.Equal(t, expectedResp, testResp)
	})
}
