package barber_shop

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	mock_barber_shop "sleeping-barber-dilemma/barber_shop/mocks"
	"sleeping-barber-dilemma/customer"
	mock_customer "sleeping-barber-dilemma/customer/mocks"
	"testing"
)

type BarberTestSuite struct {
	suite.Suite
	mockBarberShop  *mock_barber_shop.MockBarberShop
	mockCustomer    *mock_customer.MockCustomer
	mockWaitingRoom chan customer.Customer
	barber          barber
}

func (suite *BarberTestSuite) SetupTest() {
	controller := gomock.NewController(suite.T())
	suite.mockBarberShop = mock_barber_shop.NewMockBarberShop(controller)
	suite.mockCustomer = mock_customer.NewMockCustomer(controller)
	suite.mockWaitingRoom = make(chan customer.Customer)
	suite.barber = barber{id: 1, barberShop: suite.mockBarberShop}
}

func TestBarberTestSuite(t *testing.T) {
	suite.Run(t, new(BarberTestSuite))
}

func (suite *BarberTestSuite) TestWork_WhenShopIsClosed_ShouldHaircutAllWaitingCustomerAndReturnHome() {
	go func() {
		suite.mockWaitingRoom <- suite.mockCustomer
		close(suite.mockWaitingRoom)
	}()

	suite.mockBarberShop.EXPECT().IsShopClose().Return(true).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)
	suite.mockCustomer.EXPECT().GetCustomerId().Return(1).Times(2)
	suite.mockCustomer.EXPECT().HaircutCompleted().Times(1)
	suite.mockBarberShop.EXPECT().BarberReturnsToHome().Times(1)

	suite.barber.Work()

	suite.Len(suite.mockWaitingRoom, 0)
}

func (suite *BarberTestSuite) TestWork_WhenShopIsOpen_ShouldHaircutCustomersTillShopCloses() {
	go func() {
		suite.mockWaitingRoom <- suite.mockCustomer
		close(suite.mockWaitingRoom)
	}()

	suite.mockBarberShop.EXPECT().IsShopClose().Return(false).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)
	suite.mockCustomer.EXPECT().GetCustomerId().Return(1).Times(2)
	suite.mockCustomer.EXPECT().HaircutCompleted().Times(1)
	suite.mockBarberShop.EXPECT().IsShopClose().Return(true).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)
	suite.mockBarberShop.EXPECT().BarberReturnsToHome().Times(1)

	suite.barber.Work()

	suite.Len(suite.mockWaitingRoom, 0)
}

func (suite *BarberTestSuite) TestWork_WhenBarberIsSleepingAndCustomerArrives_ShouldAwake() {

	suite.mockBarberShop.EXPECT().IsShopClose().Return(false).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)

	suite.mockBarberShop.EXPECT().IsShopClose().DoAndReturn(func() bool {
		go func() {
			suite.mockWaitingRoom <- suite.mockCustomer
		}()
		return false
	}).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)
	suite.mockCustomer.EXPECT().GetCustomerId().Return(1).Times(3)
	suite.mockCustomer.EXPECT().HaircutCompleted().Times(1)
	suite.mockBarberShop.EXPECT().IsShopClose().DoAndReturn(func() bool {
		go func() {
			close(suite.mockWaitingRoom)
		}()
		return true
	}).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)
	suite.mockBarberShop.EXPECT().BarberReturnsToHome().Times(1)

	suite.barber.Work()

	suite.Len(suite.mockWaitingRoom, 0)
}

func (suite *BarberTestSuite) TestWork_WhenBarberIsSleepingAndShopClosedAndNoCustomersWaiting_ShouldReturnHome() {

	suite.mockBarberShop.EXPECT().IsShopClose().Return(false).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)

	suite.mockBarberShop.EXPECT().IsShopClose().DoAndReturn(func() bool {
		close(suite.mockWaitingRoom)
		return true
	}).Times(1)
	suite.mockBarberShop.EXPECT().IsShopClose().Return(true).Times(1)
	suite.mockBarberShop.EXPECT().GetWaitingRoom().Return(suite.mockWaitingRoom).Times(1)
	suite.mockBarberShop.EXPECT().BarberReturnsToHome().Times(1)

	suite.barber.Work()

	suite.Len(suite.mockWaitingRoom, 0)
}
