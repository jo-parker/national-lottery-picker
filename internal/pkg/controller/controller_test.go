package controller

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	mock_service "github.com/jpparker/national-lottery-picker/internal/pkg/service/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	testDraws = []model.Draw{
		{
			Name:       model.EuroMillions,
			NumTickets: 1,
			Day:        model.Tuesday,
			Strategy:   model.OddEven,
		},
		{
			Name:       model.Lotto,
			NumTickets: 1,
			Day:        model.Wednesday,
			Strategy:   model.OddEven,
		},
	}
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller")
}

var _ = Describe("Controller", func() {
	var (
		ctrl       *gomock.Controller
		impl       *mock_service.MockLottery
		controller LotteryController
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		impl = mock_service.NewMockLottery(ctrl)
		controller = LotteryController{Service: impl}
	})

	Describe("enterDraws", func() {
		It("should enter draws", func() {
			request := &model.EnterDrawsRequest{
				Draws: testDraws,
				Credentials: model.Credentials{
					Username: "test",
					Password: "password",
				},
			}

			impl.EXPECT().EnterDraw(testDraws[0], request.Credentials).Return(nil)
			impl.EXPECT().EnterDraw(testDraws[1], request.Credentials).Return(nil)

			errs := controller.enterDraws(request)

			Expect(errs).To(BeEmpty())
		})

		It("should enter 1 draw and return 1 err", func() {
			request := &model.EnterDrawsRequest{
				Draws: testDraws,
				Credentials: model.Credentials{
					Username: "test",
					Password: "password",
				},
			}

			impl.EXPECT().EnterDraw(testDraws[0], request.Credentials).Return(nil)
			impl.EXPECT().EnterDraw(testDraws[1], request.Credentials).Return(fmt.Errorf(""))

			errs := controller.enterDraws(request)

			Expect(len(errs)).Should(Equal(1))
		})

		It("should return err when ticket num exceeds limit", func() {
			testDraws = []model.Draw{
				{
					Name:       model.EuroMillions,
					NumTickets: 5,
					Day:        model.Tuesday,
					Strategy:   model.OddEven,
				},
			}
			request := &model.EnterDrawsRequest{
				Draws: testDraws,
				Credentials: model.Credentials{
					Username: "test",
					Password: "password",
				},
			}
			errs := controller.enterDraws(request)

			Expect(len(errs)).Should(Equal(1))
		})
	})
})
