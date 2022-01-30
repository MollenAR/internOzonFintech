package model

import "context"

type BothUrls struct {
	ShortUrl    string
	OriginalUrl string
}

type OriginalUrl struct {
	Url string `json:"url"`
}

type SaveOriginalUrlResponse struct {
	Status   int   `json:"status"`
	ShortUrl string `json:"shortUrl"`
}

type GetOriginalUrlResponse struct {
	Status      int   `json:"status"`
	OriginalUrl string `json:"originalUrl"`
}

type ShortUrlUsecase interface {
	SaveOriginalUrl(ctx context.Context, originalUrl string) (SaveOriginalUrlResponse, error)
	GetOriginalUrl(crx context.Context, shortUrl string) (GetOriginalUrlResponse, error)
}

type ShortUrlRepository interface {
	SaveOriginalUrl(ctx context.Context, originalUrl string) (string, error)
	GetOriginalUrl(crx context.Context, shortUrl string) (string, error)
}
