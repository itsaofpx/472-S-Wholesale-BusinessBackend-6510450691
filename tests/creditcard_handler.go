package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
    "github.com/ppwlsw/sa-project-backend/adapters/api"
    "github.com/ppwlsw/sa-project-backend/domain/entities"
)

func TestCreateCreditCard(t *testing.T) {
    app := fiber.New()
    mockUseCase := new(MockCreditCardUseCase)
    handler := api.InitiateCreditCardHandler(mockUseCase)

    app.Post("/creditcard", handler.CreateCreditCard)

    newCard := entities.CreditCard{
        UserID:      1,
        CardNumber:  "1234567812345678",
        CardHolder:  "John Doe",
        Expiration:  "2025-12-31",
        SecurityCode: "123",
    }

    mockUseCase.On("CreateCreditCard", newCard).Return(newCard, nil)

    jsonData, _ := json.Marshal(newCard)
    req := httptest.NewRequest(http.MethodPost, "/creditcard", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
    t.Log("TestCreateCreditCard passed")

    var responseCard entities.CreditCard
    json.NewDecoder(resp.Body).Decode(&responseCard)
    assert.Equal(t, newCard, responseCard)
}

func TestGetCreditCardByUserID(t *testing.T) {
    app := fiber.New()
    mockUseCase := new(MockCreditCardUseCase)
    handler := api.InitiateCreditCardHandler(mockUseCase)

    app.Get("/creditcard/:id", handler.GetCreditCardByUserID)

    card := entities.CreditCard{
        ID:           1,
        UserID:      1,
        CardNumber:  "1234567812345678",
        CardHolder:  "John Doe",
        Expiration:  "2025-12-31",
        SecurityCode: "123",
    }

    mockUseCase.On("GetCreditCardByUserID", 1).Return(card, nil)

    req := httptest.NewRequest(http.MethodGet, "/creditcard/1", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
    t.Log("TestGetCreditCardByUserID passed")

    var responseCard entities.CreditCard
    json.NewDecoder(resp.Body).Decode(&responseCard)
    assert.Equal(t, card, responseCard)
}

func TestUpdateCreditCardByUserID(t *testing.T) {
    app := fiber.New()
    mockUseCase := new(MockCreditCardUseCase)
    handler := api.InitiateCreditCardHandler(mockUseCase)

    app.Put("/creditcard/:id", handler.UpdateCreditCardByUserID)

    updatedCard := entities.CreditCard{
        CardNumber:  "8765432187654321",
        CardHolder:  "Jane Doe",
        Expiration:  "2026-12-31",
        SecurityCode: "321",
    }

    mockUseCase.On("UpdateCreditCardByUserID", 1, updatedCard).Return(updatedCard, nil)

    jsonData, _ := json.Marshal(updatedCard)
    req := httptest.NewRequest(http.MethodPut, "/creditcard/1", bytes.NewBuffer(jsonData))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
    t.Log("TestUpdateCreditCardByUserID passed")

    var responseCard entities.CreditCard
    json.NewDecoder(resp.Body).Decode(&responseCard)
    assert.Equal(t, updatedCard, responseCard)
}

func TestDeleteCreditCardByUserID(t *testing.T) {
    app := fiber.New()
    mockUseCase := new(MockCreditCardUseCase)
    handler := api.InitiateCreditCardHandler(mockUseCase)

    app.Delete("/creditcard/:id", handler.DeleteCreditCardByUserID)

    mockUseCase.On("DeleteCreditCardByUserID", 1).Return(nil)

    req := httptest.NewRequest(http.MethodDelete, "/creditcard/1", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
    t.Log("TestDeleteCreditCardByUserID passed")

    var responseMessage map[string]string
    json.NewDecoder(resp.Body).Decode(&responseMessage)
    assert.Equal(t, "Credit card deleted successfully", responseMessage["message"])
}

func TestGetCreditCardsByUserID(t *testing.T) {
    app := fiber.New()
    mockUseCase := new(MockCreditCardUseCase)
    handler := api.InitiateCreditCardHandler(mockUseCase)

    app.Get("/creditcards/:id", handler.GetCreditCardsByUserID)

    cards := []entities.CreditCard{
        {ID: 1, UserID: 1, CardNumber: "1234567812345678", CardHolder: "John Doe", Expiration: "2025-12-31", SecurityCode: "123"},
        {ID: 2, UserID: 1, CardNumber: "8765432187654321", CardHolder: "Jane Doe", Expiration: "2026-12-31", SecurityCode: "321"},
    }

    mockUseCase.On("GetCreditCardsByUserID", 1).Return(cards, nil)

    req := httptest.NewRequest(http.MethodGet, "/creditcards/1", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
    t.Log("TestGetCreditCardsByUserID passed")

    var responseCards []entities.CreditCard
    json.NewDecoder(resp.Body).Decode(&responseCards)
    assert.Equal(t, cards, responseCards)
}

func TestDeleteByCardNumber(t *testing.T) {
    app := fiber.New()
    mockUseCase := new(MockCreditCardUseCase)
    handler := api.InitiateCreditCardHandler(mockUseCase)

    app.Delete("/creditcard/number/:card_number", handler.DeleteByCardNumber)

    mockUseCase.On("DeleteByCardNumber", "1234567812345678").Return(nil)

    req := httptest.NewRequest(http.MethodDelete, "/creditcard/number/1234567812345678", nil)
    resp, _ := app.Test(req)

    assert.Equal(t, http.StatusOK, resp.StatusCode)
    t.Log("TestDeleteByCardNumber passed")

    var responseMessage map[string]string
    json.NewDecoder(resp.Body).Decode(&responseMessage)
    assert.Equal(t, "Credit card deleted successfully", responseMessage["message"])
}
