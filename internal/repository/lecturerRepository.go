package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/consts"
	"github.com/tryfix/log"
)

type LecturerRepository interface {
	GetAllLecturers() ([]models.Lecturer, error)
	GetLecturer(id int) (*models.Lecturer, error)
	CreateLecturer(lecturer *models.Lecturer) (*models.Lecturer, error)
	UpdateLecturer(lecturer *models.Lecturer) (*models.Lecturer, error)
	SearchLecturer(searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.LecturerSearchData, error)
	DeleteLecturer(id int) (*models.Lecturer, error)
}

type lecturerRepository struct {
	db *sql.DB
}

func NewLecturerRepository(db *sql.DB) *lecturerRepository {
	return &lecturerRepository{
		db: db,
	}
}

func (s *lecturerRepository) GetAllLecturers() ([]models.Lecturer, error) {

	stmt, err := s.db.Prepare("SELECT * FROM lecturers")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		log.Error(consts.DBResultsError, err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error(consts.DBRowCloseError, err)
		}
	}(rows)

	var lecturerList []models.Lecturer

	for rows.Next() {
		var st models.Lecturer

		err := rows.Scan(&st.ID, &st.FirstName, &st.LastName, &st.Year)
		if err != nil {
			log.Error(consts.DBScanRowError, err)
			return nil, err
		}

		err = rows.Err()
		if err != nil {
			log.Error(consts.DBRowsError, err)
			return nil, err
		}

		lecturerList = append(lecturerList, st)

	}

	log.Debug("getAllLecturers response : ", lecturerList)
	return lecturerList, nil
}

func (s *lecturerRepository) GetLecturer(id int) (*models.Lecturer, error) {
	var lecturer models.Lecturer

	stmt, err := s.db.Prepare("SELECT * FROM lecturers WHERE id = ?;")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &lecturer, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmt)

	err = stmt.QueryRow(id).Scan(&lecturer.ID, &lecturer.FirstName, &lecturer.LastName, &lecturer.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.Lecturer{}, errors.New(consts.LecturerNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &models.Lecturer{}, err
	}

	log.Debug("Lecturer : ", lecturer)
	return &lecturer, err
}

func (s *lecturerRepository) CreateLecturer(lecturer *models.Lecturer) (*models.Lecturer, error) {
	var st models.Lecturer

	stmt, err := s.db.Prepare("INSERT INTO lecturers (firstname,lastname,year)" +
		" VALUES (?,?,?);")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &st, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmt)

	result, err := stmt.Exec(lecturer.FirstName, lecturer.LastName, lecturer.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return &st, errors.New(consts.LecturerNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &st, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(consts.DBResultIDError, err)
	}

	lecturer.ID = int(id)

	log.Debug("Lecturer : ", *lecturer)
	return lecturer, err
}

func (s *lecturerRepository) UpdateLecturer(lecturer *models.Lecturer) (*models.Lecturer, error) {
	var st models.Lecturer

	stmt, err := s.db.Prepare("UPDATE lecturers SET firstname = ?, lastname = ?, year = ? " +
		"WHERE id = ?;")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &st, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmt)

	_, err = stmt.Exec(lecturer.FirstName, lecturer.LastName, lecturer.Year, lecturer.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &st, errors.New(consts.LecturerNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &st, err
	}

	log.Debug("Lecturer : ", *lecturer)
	return lecturer, err
}

func (s *lecturerRepository) SearchLecturer(searchString string, pagination models.Pagination,
	sortBy models.SortBy) (*models.LecturerSearchData, error) {

	query := fmt.Sprintf("SELECT *, Count(*) Over () AS TotalCount FROM "+
		"lecturers WHERE firstname LIKE '%%%s%%' || lastname LIKE '%%%s%%' ORDER BY %s %s "+
		"LIMIT %v,%v;", searchString, searchString,
		sortBy.Column, sortBy.Direction, pagination.Page, pagination.PageSize)

	stmt, err := s.db.Prepare(query)

	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		log.Error(consts.DBResultsError, err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error(consts.DBRowCloseError, err)
		}
	}(rows)

	var lecturerList []models.Lecturer
	var totalCount int
	var resp models.LecturerSearchData

	for rows.Next() {
		var st models.Lecturer

		err := rows.Scan(&st.ID, &st.FirstName, &st.LastName, &st.Year, &totalCount)
		if err != nil {
			log.Error(consts.DBScanRowError, err)
			return nil, err
		}

		err = rows.Err()
		if err != nil {
			log.Error(consts.DBRowsError, err)
			return nil, err
		}

		lecturerList = append(lecturerList, st)
	}

	resp.TotalElements = totalCount
	resp.Data = lecturerList

	log.Debug("getAllLecturers response : ", resp)
	return &resp, nil
}

func (s *lecturerRepository) DeleteLecturer(id int) (*models.Lecturer, error) {
	var lecturer models.Lecturer

	stmt, err := s.db.Prepare("DELETE FROM lecturers WHERE id = ?;")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &lecturer, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmt)

	_, err = stmt.Exec(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.Lecturer{}, errors.New(consts.LecturerNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &models.Lecturer{}, err
	}

	log.Debug("Lecturer id : ", id)
	return &lecturer, nil
}
