package repository

import (
	"go-mongodb/db"
	"go-mongodb/model"
	"go-mongodb/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IStudentRepository interface {
	GetAll() ([]model.Student, error)
	GetOneByUsername(name string) (*model.Student, error)
	CreateOne(student model.Student) (*model.Student, error)
	GetWithPage(skip int64, limit int64) ([]model.Student, error)
}

type StudentRepository struct {
	repo *mongo.Collection
}

func NewStudentRepository(resource *db.Resource) IStudentRepository {
	studentCollection := resource.Db.Collection("students")
	studentRepository := &StudentRepository{repo: studentCollection}
	return studentRepository
}

func (s *StudentRepository) GetAll() ([]model.Student, error) {
	students := []model.Student{}
	ctx, cancel := utils.InitContext()
	defer cancel()

	cursor, err := s.repo.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var student model.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
func (s *StudentRepository) GetOneByUsername(name string) (*model.Student, error) {
	student := model.Student{}
	ctx, cancel := utils.InitContext()
	defer cancel()

	filterName := bson.D{
		{
			"$and", bson.A{
				bson.D{
					{"name", name},
				},
			},
		},
	}

	err := s.repo.FindOne(ctx, filterName).Decode(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *StudentRepository) CreateOne(student model.Student) (*model.Student, error) {
	ctx, cancel := utils.InitContext()
	defer cancel()
	_, err := s.repo.InsertOne(ctx, student)
	// student.Id = (*newId).InsertedID

	if err != nil {
		return nil, err
	}

	return &student, nil
}

// func (s *StudentRepository) GetWithPage(skip int64, limit int64) ([]*model.Student, error) {
func (s *StudentRepository) GetWithPage(skip int64, limit int64) ([]model.Student, error) {
	students := []model.Student{}
	ctx, cancel := utils.InitContext()
	defer cancel()

	// opts := options.FindOptions{
	// 	Skip: &skip,
	// 	Limit: &limit,
	// }

	optionFind := options.Find()
	optionFind.SetLimit(limit)
	optionFind.SetSkip(skip)

	// cursor, err := s.repo.Find(ctx, bson.D{}, limitz, skipz)
	cursor, err := s.repo.Find(ctx, bson.D{}, optionFind)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var student model.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
