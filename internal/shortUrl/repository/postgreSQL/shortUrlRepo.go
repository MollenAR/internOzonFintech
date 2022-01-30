package postgreSQL

import (
	"context"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
)

type shortUrlPsqlRepo struct {
}

func NewPsqlRepo() model.ShortUrlRepository {
	return &shortUrlPsqlRepo{}
}

func (shurlRepo *shortUrlPsqlRepo) SaveOriginalUrl(ctx context.Context, bothUrls model.BothUrls) error {
	return nil
}

func (shurlRepo *shortUrlPsqlRepo) GetOriginalUrl(crx context.Context, shortUrl string) (string, error) {
	return "", nil
}
