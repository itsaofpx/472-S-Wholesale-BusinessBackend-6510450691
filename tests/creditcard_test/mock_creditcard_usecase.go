package tests

import (
    "github.com/stretchr/testify/mock"
    "github.com/ppwlsw/sa-project-backend/domain/entities"
)

// Mock สำหรับ CreditCardUseCase
type MockCreditCardUseCase struct {
    mock.Mock
}

func (m *MockCreditCardUseCase) CreateCreditCard(cc entities.CreditCard) (entities.CreditCard, error) {
    args := m.Called(cc)
    return args.Get(0).(entities.CreditCard), args.Error(1)
}

func (m *MockCreditCardUseCase) GetCreditCardByUserID(id int) (entities.CreditCard, error) {
    args := m.Called(id)
    return args.Get(0).(entities.CreditCard), args.Error(1)
}

func (m *MockCreditCardUseCase) UpdateCreditCardByUserID(id int, cc entities.CreditCard) (entities.CreditCard, error) {
    args := m.Called(id, cc)
    return args.Get(0).(entities.CreditCard), args.Error(1)
}

func (m *MockCreditCardUseCase) DeleteCreditCardByUserID(id int) error {
    args := m.Called(id)
    return args.Error(0)
}

func (m *MockCreditCardUseCase) GetCreditCardsByUserID(userID int) ([]entities.CreditCard, error) {
    args := m.Called(userID)
    return args.Get(0).([]entities.CreditCard), args.Error(1)
}

func (m *MockCreditCardUseCase) DeleteByCardNumber(cardNumber string) error {
    args := m.Called(cardNumber)
    return args.Error(0)
}
