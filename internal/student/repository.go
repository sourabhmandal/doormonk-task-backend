package student

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewStudentRepository(db *sqlx.DB) *StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}

type StudentRepository struct {
	*sqlx.DB
}

type IStudentRepository interface {
	Student(id int64) (Student, error)
	Students(limit int8, offset int8) ([]Student, error)
	CreateStudent(student Student) (Student, error)
	UpdateStudent(student Student) (Student, error)
	DeleteStudent(id int64) error
}

func (s *StudentRepository) Student(id int64) (Student, error) {
	var student Student
	if err := s.Get(&student, `SELECT * FROM students WHERE id = $1`, id); err != nil {
		return Student{}, fmt.Errorf(`error getting student from db: %w`, err)
	}
	return student, nil
}

func (s *StudentRepository) Students(limit int8, offset int8) ([]Student, error) {
	var students []Student
	if err := s.Get(&students, `SELECT * FROM students`); err != nil {
		return []Student{}, fmt.Errorf(`error getting students from db: %w`, err)
	}
	return students, nil
}

func (s *StudentRepository) CreateStudent(student Student) (Student, error) {
	var newStudent Student
	if err := s.Get(&newStudent, `INSERT INTO students VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		student.FirstName, student.LastName, student.Email, student.PhoneNumber,
		student.InstitutionName, student.PassOutYear, student.CgpiScore); err != nil {
		return Student{}, fmt.Errorf(`error creating student from db: %w`, err)
	}
	return newStudent, nil
}

func (s *StudentRepository) UpdateStudent(student Student) (Student, error) {
	var newStudent Student
	if err := s.Get(&newStudent, `UPDATE students SET first_name = $1, last_name = $2, email = $3, 
	phone_number = $4, institution_name = $5, pass_out_year = $6, cgpi_score = $7) RETURNING *`,
		student.FirstName, student.LastName, student.Email, student.PhoneNumber,
		student.InstitutionName, student.PassOutYear, student.CgpiScore); err != nil {
		return Student{}, fmt.Errorf(`error updating student from db: %w`, err)
	}
	return newStudent, nil
}

func (s *StudentRepository) DeleteStudent(id int64) error {
	if _, err := s.Exec(`DELETE FROM students WHERE id = $1`, id); err != nil {
		return fmt.Errorf(`error deleting students from db: %w`, err)
	}
	return nil
}
