package indexhandler

import (
	"fmt"
	"github.com/evt/blockchain-api/internal/e"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
)

// TestGetIndex tests indexhandler.Get
func TestGetIndex(t *testing.T) {
	var testIndexID int64 = 1

	log.SetOutput(ioutil.Discard)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	indexService := NewMockIndexService(ctrl)
	indexHandler := New(indexService)

	app := fiber.New()
	app.Get("/indexes/:id", indexHandler.Get)

	type index struct {
		Name              string
		EthPriceInWei     int64
		UsdPriceInCents   int64
		UsdCapitalization int64
		PercentageChange  int64
	}

	tests := []struct {
		name   string
		expect func()
		assert func([]byte)
	}{
		{
			name: "success",
			expect: func() {
				indexService.EXPECT().GetIndex(gomock.Any(), testIndexID).Return(index{
					Name:              "DeFi Index (3)",
					EthPriceInWei:     350000000000000000,
					UsdPriceInCents:   8500,
					UsdCapitalization: 270000000,
					PercentageChange:  25,
				}, nil)
			},
			assert: func(content []byte) {
				assert.Equal(t, []byte("{\"Name\":\"DeFi Index (3)\",\"EthPriceInWei\":350000000000000000,\"UsdPriceInCents\":8500,\"UsdCapitalization\":270000000,\"PercentageChange\":25}"), content)
			},
		},
		{
			name: "error",
			expect: func() {
				indexService.EXPECT().GetIndex(gomock.Any(), testIndexID).Return(nil, e.NewHttpError(http.StatusInternalServerError, "test error"))
			},
			assert: func(content []byte) {
				assert.Equal(t, []byte(`{"error":"test error"}`), content)
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.expect != nil {
				tc.expect()
			}

			req := httptest.NewRequest("GET", fmt.Sprintf("/indexes/%d", testIndexID), nil)

			res, err := app.Test(req, 1)
			if err != nil {
				t.Error(err)
			}
			defer res.Body.Close()

			content, err := io.ReadAll(res.Body)
			if err != nil {
				t.Error(err)
			}

			if tc.assert != nil {
				tc.assert(content)
			}
		})
	}
}
