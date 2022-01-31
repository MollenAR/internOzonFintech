package postgreSQL

import (
	"context"
	"database/sql"
	"github.com/MollenAR/internOzonFintech/internal/tools/errorTypes"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"regexp"
	"testing"
)

func TestGetOriginalUrl(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatal(err)
	}

	testRepo := NewPsqlRepo(db)

	t.Run("succes", func(t *testing.T) {
		rows := sqlxmock.NewRows([]string{"original_url"}).AddRow("ozon.ru/fen")

		query := `SELECT original_url FROM urls WHERE short_url = $1`

		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("9csUIp8_nM").WillReturnRows(rows)

		testResponse, err := testRepo.GetOriginalUrl(context.TODO(), "9csUIp8_nM")
		assert.NoError(t, err)
		assert.Equal(t, "ozon.ru/fen", testResponse)
	})

	t.Run("fail", func(t *testing.T) {
		query := `SELECT original_url FROM urls WHERE short_url = $1`

		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("9csUIp8_nM").WillReturnError(errors.New("test error"))

		testResponse, err := testRepo.GetOriginalUrl(context.TODO(), "9csUIp8_nM")
		if assert.Error(t, err) {
			assert.Equal(t, errorTypes.ErrTryAgainLater{Reason: "test error"}, err)
		}
		assert.Equal(t, "", testResponse)
	})
}

func TestSaveOriginalUrl(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatal(err)
	}

	testRepo := NewPsqlRepo(db)

	t.Run("success, original url already exits", func(t *testing.T) {
		rows := sqlxmock.NewRows([]string{"short_url"}).AddRow("9csUIp8_nM")

		query := `SELECT short_url FROM urls WHERE original_url = $1`

		mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("ozon.ru/fen").WillReturnRows(rows)

		testResponse, err := testRepo.SaveOriginalUrl(context.TODO(), "ozon.ru/fen")
		assert.NoError(t, err)
		assert.Equal(t, "9csUIp8_nM", testResponse)
	})

	t.Run("success, original url doesn't exits", func(t *testing.T) {
		firstQuery := `SELECT short_url FROM urls WHERE original_url = $1`

		mock.ExpectQuery(regexp.QuoteMeta(firstQuery)).WithArgs("ozon.ru/fen").WillReturnError(sql.ErrNoRows)

		secondQuery := `INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`

		mock.ExpectExec(regexp.QuoteMeta(secondQuery)).WithArgs(sqlxmock.AnyArg(), sqlxmock.AnyArg()).WillReturnResult(sqlxmock.NewResult(1, 1))

		testResponse, err := testRepo.SaveOriginalUrl(context.TODO(), "ozon.ru/fen")
		assert.NoError(t, err)
		assert.Len(t, testResponse, 10)
	})

}
