package schemas

import (
	"time"
)

// Estrutura para a tabela Walmart
type Walmart struct {
	ID       string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"` // ID único para a tabela, gerado automaticamente
	Name     string     `gorm:"type:varchar(100);not null"`                     // Nome do Walmart
	Address  string     `gorm:"type:varchar(255)"`                              // Endereço do Walmart
	CreateAt *time.Time `gorm:"autoCreateTime"`                                 // Data de criação, gerenciada automaticamente
	UpdateAt *time.Time `gorm:"autoUpdateTime"`                                 // Data de atualização, gerenciada automaticamente
}