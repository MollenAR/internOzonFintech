package postgreSQL

import (
	"context"
	"database/sql"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	gonanoid "github.com/matoous/go-nanoid"
)

type shortUrlPsqlRepo struct {
	Db *sqlx.DB
}

const alphabet = "_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const shortUrlSize = 10

func NewPsqlRepo(db *sqlx.DB) model.ShortUrlRepository {
	return &shortUrlPsqlRepo{
		Db: db,
	}
}

func (shurlRepo *shortUrlPsqlRepo) SaveOriginalUrl(ctx context.Context, originalUrl string) (string, error) {
	var shortUrl string
	schema := `SELECT short_url FROM urls WHERE original_url = $1`
	err := shurlRepo.Db.Get(&shortUrl, schema, originalUrl)

	if err != nil && err != sql.ErrNoRows{
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	if shortUrl != "" {
		return shortUrl, nil
	}

	shortUrl, err = gonanoid.Generate(alphabet, shortUrlSize)
	if err != nil {
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	schema = `INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`

	_, err = shurlRepo.Db.Exec(schema, shortUrl, originalUrl)
	if err != nil {
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	return shortUrl, nil
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
