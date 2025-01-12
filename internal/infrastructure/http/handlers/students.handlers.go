package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/NicoHackCode/Marpagitha/internal/domain"
	"github.com/NicoHackCode/Marpagitha/internal/usecases"
)

// Handler para listar todos los estudiantes
func ListStudentsHandler(useCase *usecases.ListActiveStudentsUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Valido el método
		if r.Method != http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(domain.ApiResponse{
				Status:  http.StatusMethodNotAllowed,
				Error:   "Método no permitido",
				Message: "Solo se permite el método GET",
			})
			return
		}

		// Ejecuto el caso de uso
		students, err := useCase.Execute()
		// si hay un error
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(domain.ApiResponse{
				Status:  http.StatusInternalServerError,
				Error:   "Error al obtener estudiantes",
				Message: err.Error(),
			})
		}
		// si no hay error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(domain.ApiResponse{
			Status:  http.StatusOK,
			Message: "Estudiantes obtenidos correctamente",
			Data:    students,
		})
	}
}
