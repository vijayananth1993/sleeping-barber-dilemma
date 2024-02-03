package barber_shop

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"sleeping-barber-dilemma/customer"
	mock_customer "sleeping-barber-dilemma/customer/mocks"
	"testing"
)

type BarberShopTestSuite struct {
	suite.Suite
	mockCustomer *mock_customer.MockCustomer
	barberShop   barberShop
}

func (suite *BarberShopTestSuite) SetupTest() {
	controller := gomock.NewController(suite.T())
	suite.mockCustomer = mock_customer.NewMockCustomer(controller)
	suite.barberShop = barberShop{
		waitingRoom:         make(chan customer.Customer, 3),
		numberOfBarbers:     3,
		barbers:             make([]Barber, 3),
		barberReturningHome: make(chan bool, 3),
	}
}

func TestBarberShopTestSuite(t *testing.T) {
	suite.Run(t, new(BarberShopTestSuite))
}

func (suite *BarberShopTestSuite) TestIsShopClose_WhenShopNotClosed_ShouldReturnFalse() {
	suite.Equal(false, suite.barberShop.IsShopClose())
}

func (suite *BarberShopTestSuite) TestIsShopClose_WhenShopClosed_ShouldReturnTrue() {
	suite.barberShop.isShopClosed = 1
	suite.Equal(true, suite.barberShop.IsShopClose())
}

func (suite *BarberShopTestSuite) TestWaitTillAllBarberReturnsHome_ShouldConsumeEventsFromBarberReturningHomeChannel() {
	suite.barberShop.barberReturningHome <- true
	suite.barberShop.barberReturningHome <- true
	suite.barberShop.barberReturningHome <- true

	suite.barberShop.WaitTillAllBarberReturnsHome()

	suite.Len(suite.barberShop.barberReturningHome, 0)
}

func (suite *BarberShopTestSuite) TestCustomerVisit_WhenWaitingRoomIsFull_ShouldReturnWithoutHaircut() {
	customer1 := customer.NewCustomer(1)
	customer2 := customer.NewCustomer(2)
	suite.barberShop.waitingRoom <- customer1
	suite.barberShop.waitingRoom <- customer1
	suite.barberShop.waitingRoom <- customer1

	suite.barberShop.CustomerVisit(customer2)

	close(suite.barberShop.waitingRoom)
	for c := range suite.barberShop.waitingRoom {
		suite.NotEqual(customer2, c)
	}
}

func (suite *BarberShopTestSuite) TestCustomerVisit_WhenWaitingRoomIsNotFull_ShouldReturnWithHaircut() {

	suite.mockCustomer.EXPECT().GetCustomerId().Return(1).Times(1)
	suite.mockCustomer.EXPECT().WaitForHaircutToBeCompleted().Times(1)

	suite.barberShop.CustomerVisit(suite.mockCustomer)

	close(suite.barberShop.waitingRoom)
	suite.Len(suite.barberShop.waitingRoom, 1)
	suite.Equal(suite.mockCustomer, <-suite.barberShop.waitingRoom)
}

func (suite *BarberShopTestSuite) TestBarberReturnsToHome_ShouldSendEventInBarberReturningHomeChannel() {

	suite.barberShop.BarberReturnsToHome()

	suite.Len(suite.barberShop.barberReturningHome, 1)
	suite.Equal(true, <-suite.barberShop.barberReturningHome)
}
