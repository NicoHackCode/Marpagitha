package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Port         string
	JWTSecret    string
	DatabaseName string
	MongoURI     string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	mux    *http.ServeMux
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	// voy a imprimir las variables de entorno
	fmt.Println("PORT:", config.Port)
	fmt.Println("JWT_SECRET:", config.JWTSecret)
	fmt.Println("DATABASE_NAME:", config.DatabaseName)
	fmt.Println("MONGO_URI:", config.MongoURI)

	// Validar configuración
	if config.Port == "" {
		return nil, fmt.Errorf("no se cargó el puerto")
	}

	if config.JWTSecret == "" {
		return nil, fmt.Errorf("no se cargó el secreto JWT")
	}

	if config.DatabaseName == "" {
		return nil, fmt.Errorf("no se cargó el nombre de la base de datos")
	}

	if config.MongoURI == "" {
		return nil, fmt.Errorf("no se cargó la URL de la base de datos")
	}

	// Creo la instacia del broker
	broker := &Broker{
		config: config,
		mux:    http.NewServeMux(),
	}

	return broker, nil
}

// Iniciar el servidor HTTP con las rutas configuradas
func (b *Broker) Start(binder func(s Server, mux *http.ServeMux)) {
	// Configurar rutas usando el binder
	binder(b, b.mux)
	log.Println("Servidor corriendo en el puerto", b.config.Port)
	if err := http.ListenAndServe(":"+b.config.Port, b.mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
