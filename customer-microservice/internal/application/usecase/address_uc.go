package usecase

import "github.com/trng-tr/customer-microservice/internal/application/ports/out"

//AddressServiceImpl implement port d'entrée exposé à l'extreieur
type AddressServiceImpl struct {
	// DI pour envoyer à l'extérieur
	AddressService out.AddressService
}
