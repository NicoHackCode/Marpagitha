package main

import (
	"context"
	"fmt"
	"log"
	"os"

	stdhttp "net/http"

	"github.com/NicoHackCode/Marpagitha/internal/infrastructure/database"
	"github.com/NicoHackCode/Marpagitha/internal/infrastructure/http"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting application...")

	// Cargo el archivo .env
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Fatal("Error: .env file does not exist")
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Obtengo variables de entorno
	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	databaseName := os.Getenv("DATABASE_NAME")
	mongoURI := os.Getenv("MONGO_URI")

	// Creo configuración del servidor
	config := &http.Config{
		Port:         port,
		JWTSecret:    jwtSecret,
		DatabaseName: databaseName,
		MongoURI:     mongoURI,
	}

	// Conexión a la base de datos
	client, db, err := database.NewMongoDB()
	if err != nil {
		log.Fatalf("Error initializing MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
	}()

	// Creo repositorio con la base de datos
	repo := database.NewStudentRepository(db)

	// Creo instancia del servidor
	ctx := context.Background()
	broker, err := http.NewServer(ctx, config)
	if err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}

	// Inicio el servidor y configuro rutas
	broker.Start(func(s http.Server, mux *stdhttp.ServeMux) {
		http.SetupRoutes(mux, repo) // Configuración de rutas con repositorio
	})
}
