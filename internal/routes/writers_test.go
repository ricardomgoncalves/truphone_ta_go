package routes

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/errors"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/family"
	"github.com/stretchr/testify/assert"
)

func TestWrite(t *testing.T) {
	ctl := gomock.NewController(t)
	respWriter := NewMockResponseWriter(ctl)

	t.Run("should use all methods", func(t *testing.T) {
		headers := http.Header{}
		respWriter.EXPECT().Header().Times(1).Return(headers)
		respWriter.EXPECT().WriteHeader(gomock.Eq(200)).Times(1)
		respWriter.EXPECT().Write(gomock.Any()).Times(1)

		Write(context.Background(), respWriter, 200, "a")
		assert.Equal(t, "application/json", headers.Get("Content-Type"))
	})
}

func TestWriteError(t *testing.T) {
	ctl := gomock.NewController(t)
	respWriter := NewMockResponseWriter(ctl)

	t.Run("should use all methods", func(t *testing.T) {
		headers := http.Header{}
		respWriter.EXPECT().Header().Times(1).Return(headers)
		respWriter.EXPECT().WriteHeader(gomock.Eq(500)).Times(1)
		respWriter.EXPECT().Write(gomock.Any()).Times(1)

		WriteError(context.Background(), respWriter, errors.New("test"))
		assert.Equal(t, "application/json", headers.Get("Content-Type"))
	})

	t.Run("should not write on nil error", func(t *testing.T) {
		headers := http.Header{}
		respWriter.EXPECT().Header().Times(1).Return(headers)
		respWriter.EXPECT().WriteHeader(gomock.Eq(500)).Times(1)
		respWriter.EXPECT().Write(gomock.Any()).AnyTimes()

		WriteError(context.Background(), respWriter, nil)
		assert.Equal(t, "application/json", headers.Get("Content-Type"))
	})

	t.Run("should user the right code", func(t *testing.T) {
		headers := http.Header{}
		respWriter.EXPECT().Header().Times(1).Return(headers)
		respWriter.EXPECT().WriteHeader(gomock.Eq(409)).Times(1)
		respWriter.EXPECT().Write(gomock.Any()).AnyTimes()

		WriteError(context.Background(), respWriter, family.ErrorFamilyAlreadyExists)
		assert.Equal(t, "application/json", headers.Get("Content-Type"))
	})
}
