package api

type Handlers struct {
	UserHandler              *UserHandler
	ProductHandler           *ProductHandler
	TransactionHandler       *TransactionHandler
	AuthHandler              *AuthHandler
	ShipmentHandler          *ShipmentHandler
	OrderHandler             *OrderHandler
	PackageHandler           *PackageHandler
	OrderLineHandler         *OrderLineHandler
	TierListHandler          *TierListHandler
	SupplierHandler          *SupplierHandler
	SupplierOrderListHandler *SupplierOrderListHandler
	AdminHandler             *AdminHandler
}

func ProvideHandlers(UserHandler *UserHandler, ProductHandler *ProductHandler,
	TransactionHandler *TransactionHandler, AuthHandler *AuthHandler,
	ShipmentHandler *ShipmentHandler, OrderHandler *OrderHandler,
	PackageHandler *PackageHandler, OrderLineHandler *OrderLineHandler,
	SupplierHandler *SupplierHandler, SupplierOrderListHandler *SupplierOrderListHandler,
	TierListHandler *TierListHandler, AdminHandler *AdminHandler) *Handlers {
	return &Handlers{
		UserHandler:              UserHandler,
		ProductHandler:           ProductHandler,
		TransactionHandler:       TransactionHandler,
		AuthHandler:              AuthHandler,
		ShipmentHandler:          ShipmentHandler,
		OrderHandler:             OrderHandler,
		PackageHandler:           PackageHandler,
		OrderLineHandler:         OrderLineHandler,
		SupplierHandler:          SupplierHandler,
		SupplierOrderListHandler: SupplierOrderListHandler,
		TierListHandler:          TierListHandler,
		AdminHandler:             AdminHandler,
	}
}
