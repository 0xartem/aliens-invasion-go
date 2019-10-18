package alieninvasion

type CityProvider interface {
    GetRandomCity() *City
    DestroyCity(name string) error
}
