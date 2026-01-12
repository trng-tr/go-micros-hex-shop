package main

import (
	"log"
	"time"

	_ "github.com/lib/pq" //db connector
	"github.com/trng-tr/customer-microservice/internal/application/usecase"
	impl2 "github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/handlers/impl"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/routes"
	impl1 "github.com/trng-tr/customer-microservice/internal/infrastructure/out/repositories/impl"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/services"

	"database/sql"
)

func main() {
	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	var addressRepo = impl1.NewAddressRepositoryImpl(db)
	var customerRepo = impl1.NewCustomerRepositoryImpl(db)
	var addressOutPort = services.NewOutAddressServiceImpl(addressRepo)
	var customerOutPort = services.NewOutCustomerServiceImpl(customerRepo)
	var addressInPort = usecase.NewInAddressServiceImpl(addressOutPort)
	var customerInPort = usecase.NewInCustomerServiceImpl(customerOutPort, addressInPort)
	var customerHandler = impl2.NewCustomerHandlerServiceImpl(customerInPort, addressInPort)
	var addressHandler = impl2.NewAddressHandlerServiceImpl(addressInPort)
	var rr = routes.NewRoutesRegistration(addressHandler, customerHandler)
	var engine = rr.RegisterRoutes()
	engine.Run(":8081")
}

func createDatabase() (*sql.DB, error) {
	dbDns := "postgres://go-app-user:go-app-pass@localhost:5432/goappdb?sslmode=disable"
	db, err := sql.Open("postgres", dbDns)
	if err != nil {
		return nil, err
	}
	// Vérifie la connexion réelle
	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(50 * time.Minute)
	db.SetConnMaxLifetime(50. * time.Minute)
	db.SetMaxIdleConns(50)

	return db, nil
}
