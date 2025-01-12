package database

import (
	"context"
	"time"

	"github.com/NicoHackCode/Marpagitha/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository struct {
	collection *mongo.Collection
}

func NewStudentRepository(db *mongo.Database) *StudentRepository {
	return &StudentRepository{
		collection: db.Collection("Estudiantes"),
	}
}

func (r *StudentRepository) GetActiveStudents() ([]domain.Student, error) {
	// Creo contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Creo el filtro
	filter := bson.M{"activo": true}

	// Realizo la consulta
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Mapeo los resultados a un slice de domain.Student
	var students []domain.Student
	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}
