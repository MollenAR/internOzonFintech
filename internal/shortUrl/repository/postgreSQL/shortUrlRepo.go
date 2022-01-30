package postgreSQL

import (
	"context"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type shortUrlPsqlRepo struct {
	Db *sqlx.DB
}

func NewPsqlRepo(db *sqlx.DB) model.ShortUrlRepository {
	return &shortUrlPsqlRepo{
		Db: db,
	}
}

func (shurlRepo *shortUrlPsqlRepo) SaveOriginalUrl(ctx context.Context, bothUrls model.BothUrls) error {
	schema := `INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`

	_, err := shurlRepo.Db.Exec(schema, bothUrls.ShortUrl, bothUrls.OriginalUrl)
	if err != nil {
		return errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	return nil
}

func (shurlRepo *shortUrlPsqlRepo) GetOriginalUrl(crx context.Context, shortUrl string) (string, error) {
	schema := `SELECT original_url FROM urls WHERE short_url = $1`
	var originalUrl string

	err := shurlRepo.Db.Get(&originalUrl, schema, shortUrl)
	if err != nil {
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	return originalUrl, nil
}
