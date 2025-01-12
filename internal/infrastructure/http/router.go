package http

import (
	"net/http"

	"github.com/NicoHackCode/Marpagitha/internal/infrastructure/database"
	"github.com/NicoHackCode/Marpagitha/internal/infrastructure/http/handlers"
	"github.com/NicoHackCode/Marpagitha/internal/usecases"
)

func SetupRoutes(mux *http.ServeMux, studentRepo *database.StudentRepository) {
	//llamado a casos de uso
	listarEstudiantesUseCase := &usecases.ListActiveStudentsUseCase{Repo: studentRepo}

	// Configuración de rutas a los handlers
	mux.HandleFunc("/studentsactive", handlers.ListStudentsHandler(listarEstudiantesUseCase))
}
