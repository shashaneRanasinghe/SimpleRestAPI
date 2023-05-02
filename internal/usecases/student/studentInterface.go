package student

import "github.com/shashaneRanasinghe/simpleAPI/internal/models"

type StudentInterface interface {
	GetAllStudents() ([]models.Student, error)
	GetStudent(id int) (*models.Student, error)
	CreateStudent(student *models.Student) (*models.Student, error)
	UpdateStudent(student *models.Student) (*models.Student, error)
	SearchStudent(searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.StudentSearchData, error)
	DeleteStudent(id int) (*models.Student, error)
}
