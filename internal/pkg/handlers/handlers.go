package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service"
)

type LotteryHandler struct {
	Service service.Lottery
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func EnterDraws(r events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse,
	error,
) {
	var req *model.EnterDrawsRequest
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return nil, errors.New("invalid draw data")
	}

	handler := &LotteryHandler{
		&service.LotteryImpl{},
	}

	errs := handler.Service.EnterDraws(req.Draws, req.Credentials)
	if errs != nil {
		var errStrs []string
		for _, err := range errs {
			errStrs = append(errStrs, err.Error())
		}

		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(strings.Join(errStrs, "; ")),
		})
	}

	return apiResponse(http.StatusCreated, "success")
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, "method not allowed")
}
