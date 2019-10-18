package alieninvasion

// CityProvider - interface to provide cities acces in the AliensManager
type CityProvider interface {
	GetRandomCity() *City
	DestroyCity(name string) error
}
