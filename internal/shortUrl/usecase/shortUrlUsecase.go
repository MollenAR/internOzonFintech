package usecase

import (
	"context"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/pkg/errors"
	"net/http"
)

type shortUrlUsecase struct {
	shortUrlRepo model.ShortUrlRepository
}

func NewShortUrlUsecase(shurlRepo model.ShortUrlRepository) model.ShortUrlUsecase {
	return &shortUrlUsecase{
		shortUrlRepo: shurlRepo,
	}
}

func (shUrlUsecase *shortUrlUsecase) SaveOriginalUrl(ctx context.Context, originalUrl string) (model.SaveOriginalUrlResponse, error) {
	shortUrl, err := shUrlUsecase.shortUrlRepo.SaveOriginalUrl(ctx, originalUrl)
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
