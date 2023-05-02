package student

import (
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/internal/repository"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/consts"
	"github.com/tryfix/log"
)

type studentUsecase struct {
	studentRepo repository.StudentRepositoryInterface
}

func NewStudent(studentRepo repository.StudentRepositoryInterface) StudentInterface {
	return &studentUsecase{
		studentRepo: studentRepo,
	}
}

func (s studentUsecase) GetAllStudents() ([]models.Student, error) {
	studentList, err := s.studentRepo.GetAllStudents()
	if err != nil {
		log.Debug("Error getting Students ", err)
		return nil, err
	}
	return studentList, nil
}

func (s studentUsecase) GetStudent(id int) (*models.Student, error) {
	student, err := s.studentRepo.GetStudent(id)
	if err != nil {
		log.Debug("Error getting Students ", err)
		return &models.Student{}, err
	}
	return student, nil
}

func (s studentUsecase) CreateStudent(student *models.Student) (*models.Student, error) {

	st, err := s.studentRepo.CreateStudent(student)
	if err != nil {
		log.Debug(consts.GetStudentsError, err)
		return &models.Student{}, err
	}
	return st, nil
}

func (s studentUsecase) UpdateStudent(student *models.Student) (*models.Student, error) {
	st, err := s.studentRepo.UpdateStudent(student)
	if err != nil {
		log.Debug(consts.GetStudentsError, err)
		return &models.Student{}, err
	}
	return st, nil
}

func (s studentUsecase) SearchStudent(searchString string, pagination models.Pagination,
	sortBy models.SortBy) (*models.StudentSearchData, error) {
	studentList, err := s.studentRepo.SearchStudent(searchString, pagination, sortBy)
	if err != nil {
		log.Debug("Error getting Students ", err)
		return nil, err
	}
	return studentList, nil
}

func (s studentUsecase) DeleteStudent(id int) (*models.Student, error) {
	student, err := s.studentRepo.DeleteStudent(id)
	if err != nil {
		log.Debug(consts.StudentDeleteError, err)
		return &models.Student{}, err
	}
	return student, nil
}
