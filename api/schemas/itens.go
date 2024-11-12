package schemas

import (
	"time"
)

// Estrutura para a tabela Itens
type Itens struct {
	ID        string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // ID único para a tabela, gerado automaticamente
	Name      string     `gorm:"type:varchar(100);not null"`                     // Nome do item
	Categoria string     `gorm:"type:varchar(50)"`                               // Categoria do item
	CreateAt  *time.Time `gorm:"autoCreateTime"`                                 // Data de criação, gerenciada automaticamente
	UpdateAt  *time.Time `gorm:"autoUpdateTime"`                                 // Data de atualização, gerenciada automaticamente
}
