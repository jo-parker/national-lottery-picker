package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

func GetEnterDrawsHandler(ep endpoint.Endpoint, options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		ep,
		decodePostEnterDrawsRequest,
		encodePostEnterDrawsResponse,
		options...,
	)
}

func decodePostEnterDrawsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.EnterDrawsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodePostEnterDrawsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	resp, ok := response.(*model.EnterDrawsResponse)
	if !ok {
		return errors.New("error decoding")
	}
	return json.NewEncoder(w).Encode(resp)
}
