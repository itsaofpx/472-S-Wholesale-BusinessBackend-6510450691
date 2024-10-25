package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/adapters/api"
	"github.com/ppwlsw/sa-project-backend/adapters/database"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"gorm.io/gorm"
)

func SetUpRouters(app *fiber.App, db *gorm.DB) {

	userRepo := database.InitiateUserPostgresRepository(db)
	userService := usecases.InitiateUserService(userRepo)
	userHandler := api.InitiateUserHandler(userService)

	productRepo := database.InitiateProductPostGresRepository(db)
	productService := usecases.InitiateProductsService(productRepo)
	productHandler := api.InitiateProductHandler(productService)

	transactionRepo := database.InitiateTransactionPostGresRepository(db)
	transactionService := usecases.InitiateTransactionService(transactionRepo)
	transactionHandler := api.InitiateTransactionHandler(transactionService)

	authService := usecases.InitiateAuthService(userRepo)
	authHandler := api.InitiateAuthHandler(authService)

	shipmentRepo := database.InitiateShipmentPostgresRepository(db)
	shipmentService := usecases.InitiateShipmentService(shipmentRepo)
	shipmentHandler := api.InitiateShipmentHandler(shipmentService)

	orderRepo := database.InitiateOrderPostgresRepository(db)
	orderService := usecases.InitiateOrderService(orderRepo)
	orderHandler := api.InitiateOrderHandler(orderService)

	packageRepo := database.InitiatePackagePostgresRepository(db)
	packageService := usecases.InitiatePackageService(packageRepo)
	packageHandler := api.InitiatePackageHandler(packageService)

	orderLineRepo := database.InitiateOrderLinePostgresRepository(db)
	orderLineService := usecases.InitiateOrderLineService(orderLineRepo)
	orderLineHandler := api.InitiateOrderLineHandler(orderLineService)

	tierListRepo := database.InitiateTierListPostgres(db)
	tierListService := usecases.InitiateTierListService(tierListRepo)
	tierListHandler := api.InitiateTierListHandler(tierListService)
	tierListHandler.TierListUsecase.InitialTierList()

	supplierRepo := database.InitiateSupplierPostgresRepository(db)
	supplierService := usecases.InitiateSupplierService(supplierRepo)
	supplierHandler := api.InitiateSupplierHandler(supplierService)

	supplierOrderListRepo := database.InitiateSupplierOrderListPostgresRepository(db)
	supplierOrderListService := usecases.InitiateSupplierOrderListService(supplierOrderListRepo)
	supplierOrderListHandler := api.InitiateSupplierOrderListHandler(supplierOrderListService)

	handlers := api.ProvideHandlers(
		userHandler, productHandler, transactionHandler,
		authHandler, shipmentHandler, orderHandler,
		packageHandler, orderLineHandler, supplierHandler,
		supplierOrderListHandler, tierListHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//User
	app.Get("/users", handlers.UserHandler.GetAllUsers)
	app.Get("/users/:id", handlers.UserHandler.GetUserByID)
	app.Put("/users/update", handlers.UserHandler.UpdateTierByUserID)

	//Tier
	app.Get("/discount/:id", handlers.TierListHandler.GetDiscountPercentByUserID)

	//Product
	app.Post("/products", handlers.ProductHandler.CreateProduct)
	app.Get("/products", handlers.ProductHandler.GetAllProducts)
	app.Get("product/:id", handlers.ProductHandler.GetProductByID)
	app.Put("products/:id", handlers.ProductHandler.UpdateProduct)

	//Transaction
	app.Post("/transaction", handlers.TransactionHandler.CreateTransaction)
	app.Post("/transactions", handlers.TransactionHandler.CreateTransaction)
	app.Get("/transactions", handlers.TransactionHandler.GetAllTransactions)
	app.Get("/transaction/:id", handlers.TransactionHandler.GetTransactionById)
	app.Get("/transaction/order/:order_id", handlers.TransactionHandler.GetTransactionByOrderId)
	app.Put("/transaction/:id", handlers.TransactionHandler.UpdateTransaction)

	//Auth
	app.Post("/register", handlers.AuthHandler.Register)
	app.Post("/login", handlers.AuthHandler.Login)

	//Shipment
	app.Post("/shipment", handlers.ShipmentHandler.CreateShipment)
	app.Post("/shipments", handlers.ShipmentHandler.CreateShipment)
	app.Get("/shipments", handlers.ShipmentHandler.GetAllShipments)
	app.Get("/shipment/:id", handlers.ShipmentHandler.GetShipmentByID)
	app.Get("/shipment/order/:order_id", handlers.ShipmentHandler.GetShipmentByOrderID)
	app.Put("/shipment/:id", handlers.ShipmentHandler.UpdateShipment)

	//Order
	app.Post("/order", handlers.OrderHandler.CreateOrder)
	app.Post("/orders", handlers.OrderHandler.CreateOrder)
	app.Get("/orders", handlers.OrderHandler.GetAllOrders)
	app.Get("/order/:id", handlers.OrderHandler.GetOrderByID)
	app.Put("/order/:id", handlers.OrderHandler.UpdateOrder)

	//Package
	app.Post("/packages", packageHandler.CreatePackage)
	app.Get("/packages/:id", packageHandler.GetPackageByID)
	app.Get("/shipment/:id/packages", packageHandler.GetAllPackagesByShipmentID)
	app.Get("/packages", packageHandler.GetAllPackages)

	//OrderLine
	app.Post("/orderLines", orderLineHandler.CreateOrderLine)
	app.Get("/orderLines/:id", orderLineHandler.GetOrderLineByID)
	app.Get("/orders/:orderID/orderLines", orderLineHandler.GetOrderLinesByOrderID)
	app.Get("orderLines", orderLineHandler.GetAllOrderLines)
	app.Put("/orderLines/:id", orderLineHandler.UpdateOrderLine)
	app.Delete("/orderLines/:id", orderLineHandler.DeleteOrderLine)

	//Supplier
	app.Post("/suppliers", supplierHandler.CreateSupplier)
	app.Put("/suppliers/:id", supplierHandler.UpdateSupplier)
	app.Get("/suppliers/:id", supplierHandler.GetSupplierByID)
	app.Get("/suppliers", supplierHandler.GetAllSuppliers)

	//SupplierOrderList
	app.Post("/supplierOrderLists", supplierOrderListHandler.CreateSupplierOrderList)
	app.Get("/supplierOrderLists/:id", supplierOrderListHandler.GetSupplierOrderListByID)
	app.Get("/suppliers/:supplierID/supplierOrderLists", supplierOrderListHandler.GetSupplierOrderListsBySupplierID)
	app.Get("supplierOrderLists", supplierOrderListHandler.GetAllSupplierOrderLists)
	app.Put("/supplierOrderLists/:id", supplierOrderListHandler.UpdateSupplierOrderList)
}
