package usecases

import (
	"github.com/NicoHackCode/Marpagitha/internal/domain"
	"github.com/NicoHackCode/Marpagitha/internal/infrastructure/database"
)

type ListActiveStudentsUseCase struct {
	Repo *database.StudentRepository
}

func (uc *ListActiveStudentsUseCase) Execute() ([]domain.Student, error) {
	//Obtengo los estudiantes activos
	students, err := uc.Repo.GetActiveStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}
