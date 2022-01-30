package usecase

import (
	"context"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/pkg/errors"
	"net/http"
)

const alphabet = "-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const shortUrlSize = 10

type shortUrlUsecase struct {
	shortUrlRepo model.ShortUrlRepository
}

func NewShortUrlUsecase(shurlRepo model.ShortUrlRepository) model.ShortUrlUsecase {
	return &shortUrlUsecase{
		shortUrlRepo: shurlRepo,
	}
}

func (shUrlUsecase *shortUrlUsecase) SaveOriginalUrl(ctx context.Context, originalUrl string) (model.SaveOriginalUrlResponse, error) {
	shortUrl, err := gonanoid.Generate(alphabet, shortUrlSize)
	if err != nil {
		return model.SaveOriginalUrlResponse{}, errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	bothUrls := model.BothUrls{
		ShortUrl:    shortUrl,
		OriginalUrl: originalUrl,
	}

	err = shUrlUsecase.shortUrlRepo.SaveOriginalUrl(ctx, bothUrls)
	if err != nil {
		return model.SaveOriginalUrlResponse{}, errors.Wrap(err, "")
	}

	response := model.SaveOriginalUrlResponse{
		Status:   http.StatusOK,
		ShortUrl: shortUrl,
	}

	return response, nil
}

func (shUrlUsecase *shortUrlUsecase) GetOriginalUrl(ctx context.Context, shortUrl string) (model.GetOriginalUrlResponse, error) {
	originalUrl, err := shUrlUsecase.shortUrlRepo.GetOriginalUrl(ctx, shortUrl)
	if err != nil {
		return model.GetOriginalUrlResponse{}, errors.Wrap(err, "")
	}

	response := model.GetOriginalUrlResponse{
		Status:      http.StatusOK,
		OriginalUrl: originalUrl,
	}

	return response, nil
}
