package usecase

import (
	"context"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	gonanoid "github.com/matoous/go-nanoid"
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
		// todo error handle
		return model.SaveOriginalUrlResponse{}, err
	}

	bothUrls := model.BothUrls{
		ShortUrl:    shortUrl,
		OriginalUrl: originalUrl,
	}

	err = shUrlUsecase.shortUrlRepo.SaveOriginalUrl(ctx, bothUrls)
	if err != nil {
		// todo error handle
		return model.SaveOriginalUrlResponse{}, err
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
		// todo error handle
		return model.GetOriginalUrlResponse{}, err
	}

	response := model.GetOriginalUrlResponse{
		Status:      http.StatusOK,
		OriginalUrl: originalUrl,
	}

	return response, nil
}
