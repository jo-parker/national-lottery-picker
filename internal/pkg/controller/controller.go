package controller

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-kit/kit/endpoint"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service"
)

type Endpoints struct {
	PostEnterDraws endpoint.Endpoint
}

type LotteryController struct {
	Service service.Lottery
}

func MakeEndpoints(logger log.Logger, middlewares []endpoint.Middleware) Endpoints {
	ctrl := &LotteryController{
		&service.LotteryImpl{},
	}

	return Endpoints{
		PostEnterDraws: wrapEndpoint(ctrl.makePostEnterDraws(logger), middlewares),
	}
}

func (ctrl *LotteryController) makePostEnterDraws(logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*model.EnterDrawsRequest)
		if !ok {
			log.Printf("invalid request")
			return nil, errors.New("invalid request")
		}

		var responseErr error
		message := "success"

		if errs := ctrl.enterDraws(req); errs != nil {
			for _, err := range errs {
				message += fmt.Sprintf("%v \n", err)
			}
			responseErr = errors.New("errors entering draws")
		}

		return &model.EnterDrawsResponse{
			Message: message,
		}, responseErr
	}
}

func (ctrl *LotteryController) enterDraws(req *model.EnterDrawsRequest) []error {
	var errors []error

	for _, d := range req.Draws {
		if d.NumTickets > 4 {
			errors = append(errors, fmt.Errorf("maximum number of 4 tickets exceeded in one order: %d", d.NumTickets))
			continue
		}

		if err := ctrl.Service.EnterDraw(d, req.Credentials); err != nil {
			errors = append(errors, fmt.Errorf("entering draw failed: %s for %s", err, req.Credentials.Username))
			continue
		}
	}
	return errors
}

func wrapEndpoint(e endpoint.Endpoint, middlewares []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middlewares {
		e = m(e)
	}
	return e
}
