package controller

import (
	"context"
	"net/http/httptest"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ctx                = context.Background()
	mockResponseWriter = httptest.NewRecorder()
)

var _ = Describe("Transport", func() {
	Describe("encodePostEnterDrawsResponse", func() {
		It("should encode response", func() {
			response := &model.EnterDrawsResponse{
				Message: "test response",
			}

			err := encodePostEnterDrawsResponse(ctx, mockResponseWriter, response)

			Expect(mockResponseWriter.Body.String()).To(ContainSubstring("test response"))
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return error if response not of type EnterDrawsResponse", func() {
			response := "test response"

			err := encodePostEnterDrawsResponse(ctx, mockResponseWriter, response)
			Expect(err).To(HaveOccurred())
		})
	})
})
