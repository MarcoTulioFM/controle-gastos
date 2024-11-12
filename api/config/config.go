package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/MarcoTulioFM/controle-gastos/api/schemas" // Importa o pacote onde estão as estruturas do banco de dados
)

var DB *gorm.DB

// Configura e inicia a conexão com o PostgreSQL usando GORM
func Config() {

	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load("/home/zezepopo/repositorios/controle-gastos/api/config/.env"); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Parâmetros de conexão com o PostgreSQL
	dsn := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s TimeZone=America/Sao_Paulo",
    os.Getenv("DATABASE_HOST"),   // Aqui, substitua por um host compatível com IPv4, se necessário
    os.Getenv("DATABASE_USER"),
    os.Getenv("DATABASE_PASSWORD"),
    os.Getenv("DATABASE_NAME"),
)

	// Configuração do logger do GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,   // Tempo limite para consultas lentas
			LogLevel:      logger.Silent, // LogLevel pode ser ajustado (Silent, Error, Warn, Info)
			Colorful:      true,
		},
	)

	// Conexão com o banco usando GORM
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")

	// Executa o AutoMigrate para as tabelas especificadas
	if err := DB.AutoMigrate(&schemas.Itens{}, &schemas.Walmart{}); err != nil {
		log.Fatalf("Erro ao realizar a migração automática: %v", err)
	}
}
