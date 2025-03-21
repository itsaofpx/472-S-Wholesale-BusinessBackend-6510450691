package api

type Handlers struct {
	UserHandler              *UserHandler
	ProductHandler           *ProductHandler
	TransactionHandler       *TransactionHandler
	AuthHandler              *AuthHandler
	OrderHandler             *OrderHandler
	OrderLineHandler         *OrderLineHandler
	TierListHandler          *TierListHandler
	SupplierHandler          *SupplierHandler
	SupplierOrderListHandler *SupplierOrderListHandler
	AdminHandler             *AdminHandler
	CreditCardHandler        *CreditCardHandler
	ChatHandler              *ChatHandler
	MessageHandler           *MessageHandler

}

func ProvideHandlers(
	UserHandler *UserHandler,
	ProductHandler *ProductHandler,
	TransactionHandler *TransactionHandler,
	AuthHandler *AuthHandler,
	OrderHandler *OrderHandler,
	OrderLineHandler *OrderLineHandler,
	CreditCardHandler *CreditCardHandler,
	SupplierHandler *SupplierHandler, SupplierOrderListHandler *SupplierOrderListHandler,
	TierListHandler *TierListHandler, AdminHandler *AdminHandler, ChatHandler *ChatHandler, MessageHandler *MessageHandler) *Handlers {
	return &Handlers{
		UserHandler:              UserHandler,
		ProductHandler:           ProductHandler,
		TransactionHandler:       TransactionHandler,
		AuthHandler:              AuthHandler,
		OrderHandler:             OrderHandler,
		OrderLineHandler:         OrderLineHandler,
		SupplierHandler:          SupplierHandler,
		SupplierOrderListHandler: SupplierOrderListHandler,
		TierListHandler:          TierListHandler,
		AdminHandler:             AdminHandler,
		CreditCardHandler:        CreditCardHandler,
		ChatHandler:              ChatHandler,
		MessageHandler:           MessageHandler,
	}
}
