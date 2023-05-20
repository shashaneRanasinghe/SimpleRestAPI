package lecturer

import (
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/internal/repository"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/consts"
	"github.com/tryfix/log"
)

type LecturerUsecase interface {
	GetAllLecturers() ([]models.Lecturer, error)
	GetLecturer(id int) (*models.Lecturer, error)
	CreateLecturer(student *models.Lecturer) (*models.Lecturer, error)
	UpdateLecturer(student *models.Lecturer) (*models.Lecturer, error)
	SearchLecturer(searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.LecturerSearchData, error)
	DeleteLecturer(id int) (*models.Lecturer, error)
}

type lecturerUsecase struct {
	lecturerRepo repository.LecturerRepository
}

func NewLecturer(lecturerRepo repository.LecturerRepository) LecturerUsecase {
	return &lecturerUsecase{
		lecturerRepo: lecturerRepo,
	}
}

func (s lecturerUsecase) GetAllLecturers() ([]models.Lecturer, error) {
	lecturerList, err := s.lecturerRepo.GetAllLecturers()
	if err != nil {
		log.Debug(consts.GetLecturersError, err)
		return nil, err
	}
	return lecturerList, nil
}

func (s lecturerUsecase) GetLecturer(id int) (*models.Lecturer, error) {
	lecturer, err := s.lecturerRepo.GetLecturer(id)
	if err != nil {
		log.Debug(consts.GetLecturersError, err)
		return &models.Lecturer{}, err
	}
	return lecturer, nil
}

func (s lecturerUsecase) CreateLecturer(lecturer *models.Lecturer) (*models.Lecturer, error) {

	st, err := s.lecturerRepo.CreateLecturer(lecturer)
	if err != nil {
		log.Debug(consts.GetLecturersError, err)
		return &models.Lecturer{}, err
	}
	return st, nil
}

func (s lecturerUsecase) UpdateLecturer(lecturer *models.Lecturer) (*models.Lecturer, error) {
	st, err := s.lecturerRepo.UpdateLecturer(lecturer)
	if err != nil {
		log.Debug(consts.GetLecturersError, err)
		return &models.Lecturer{}, err
	}
	return st, nil
}

func (s lecturerUsecase) SearchLecturer(searchString string, pagination models.Pagination,
	sortBy models.SortBy) (*models.LecturerSearchData, error) {
	lecturerList, err := s.lecturerRepo.SearchLecturer(searchString, pagination, sortBy)
	if err != nil {
		log.Debug(consts.GetLecturersError, err)
		return nil, err
	}
	return lecturerList, nil
}

func (s lecturerUsecase) DeleteLecturer(id int) (*models.Lecturer, error) {
	lecturer, err := s.lecturerRepo.DeleteLecturer(id)
	if err != nil {
		log.Debug(consts.LecturerDeleteError, err)
		return &models.Lecturer{}, err
	}
	return lecturer, nil
}
