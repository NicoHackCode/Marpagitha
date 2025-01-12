package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Creo y configuro un cliente de MongoDB y retorno la base de datos
func NewMongoDB() (*mongo.Client, *mongo.Database, error) {
	// Leo la URI y el nombre de la base de datos desde las variables de entorno
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, nil, fmt.Errorf("no se cargó el valor de la variable de entorno MONGO_URI")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		return nil, nil, fmt.Errorf("no se cargó el valor de la variable de entorno DATABASE_NAME")
	}

	// Configuro opciones del cliente
	clientOpts := options.Client().ApplyURI(uri)

	// Creo cliente de MongoDB
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("error al crear el cliente de MongoDB: %v", err)
	}

	// Creo un contexto con timeout de 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conecto al cliente
	if err := client.Connect(ctx); err != nil {
		return nil, nil, fmt.Errorf("error al conectar a MongoDB: %v", err)
	}

	// Verifico la conexión
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("error al hacer ping a MongoDB: %v", err)
	}

	log.Println("Conexión a MongoDB exitosa")
	return client, client.Database(databaseName), nil
}
