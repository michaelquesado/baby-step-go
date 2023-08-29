package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pgx "github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar as variaveis de ambiente")
	}
	con, err := pgx.Connect(context.Background(), os.Getenv("DATABASE"))
	if err != nil {
		fmt.Println("Erro ao tentar estabelecer conexão com a base de dados", err)
		os.Exit(1)
	}

	defer con.Close(context.Background())

	var today time.Time
	er := con.QueryRow(context.Background(), "select now() as today").Scan(&today)
	if er != nil {
		fmt.Println(er)
	}

	fmt.Println("today is", today)
}
