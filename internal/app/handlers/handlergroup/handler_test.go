package handlergroup

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
)

// TestGetAll tests grouphandler.GetAll
func TestGetAll(t *testing.T) {
	const path = "/groups"

	log.SetOutput(ioutil.Discard)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	groupService := NewMockGroupService(ctrl)
	groupHandler := New(groupService)

	app := fiber.New()
	app.Get(path, groupHandler.GetAll)

	tests := []struct {
		name   string
		expect func()
		assert func([]byte)
	}{
		{
			name: "success",
			expect: func() {
				groupService.EXPECT().GetGroupIDs(gomock.Any()).Return([]int64{12, 13}, nil)
			},
			assert: func(content []byte) {
				assert.Equal(t, []byte("[12,13]"), content)
			},
		},
		{
			name: "error",
			expect: func() {
				groupService.EXPECT().GetGroupIDs(gomock.Any()).Return(nil, errors.New("test error"))
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

			req := httptest.NewRequest("GET", path, nil)

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
