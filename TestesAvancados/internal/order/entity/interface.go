package entity

// Interface  para não ficar dependente da função
// assim caso eu precise mudar de base de dados preciso apenas
// implementar a interface
type OrderRepositoryInterface interface {
	Save(order *Order) error
}
