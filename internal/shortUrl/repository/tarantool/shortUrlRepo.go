package tarantool

import (
	"context"
	"fmt"
	"github.com/MollenAR/internOzonFintech/internal/shortUrl/model"
	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/tarantool/go-tarantool"
)

type shortUrlTarantoolRepo struct {
	Db *tarantool.Connection
}

const alphabet = "_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const shortUrlSize = 10

func NewTarantoolRepo(tarantoolConn *tarantool.Connection) model.ShortUrlRepository {
	return &shortUrlTarantoolRepo{
		Db: tarantoolConn,
	}
}

func (shurlRepo *shortUrlTarantoolRepo) SaveOriginalUrl(ctx context.Context, originalUrl string) (string, error) {
	resp, err := shurlRepo.Db.Select("urls", "secondary", 0, 1, tarantool.IterEq, []interface{}{originalUrl})
	if err != nil {
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	if len(resp.Data) != 0 {
		return resp.Tuples()[0][0].(string), nil
	}

	shortUrl, err := gonanoid.Generate(alphabet, shortUrlSize)
	if err != nil {
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	resp, err = shurlRepo.Db.Insert("urls", []interface{}{shortUrl, originalUrl})
	if err != nil {
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}
	fmt.Println(resp.Data)

	return shortUrl, nil
}

func (shurlRepo *shortUrlTarantoolRepo) GetOriginalUrl(crx context.Context, shortUrl string) (string, error) {
	resp, err := shurlRepo.Db.Select("urls", "primary", 0, 1, tarantool.IterEq, []interface{}{shortUrl})
	if err != nil {
		return "", errorTypes.ErrTryAgainLater{
			Reason: err.Error(),
		}
	}

	if len(resp.Data) == 0 {
		return "", errorTypes.ErrWrongOriginalUrl{
			Reason: "no original Url fund in tarantool bd"}
	}

	return resp.Tuples()[0][1].(string), nil
}
