package schemas

type Item struct {
	ID         uint    `json:"id" gorm:"primary_key"`
	Nome       string  `json:"nome"`
	Categoria  string  `json:"categoria"`
	Quantidade int64   `json:"quantidade"`
	Preco      float64 `json:"preco"`
	DataCompra string  `json:"data_compra"`
}