package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/consts"
	"github.com/tryfix/log"
)

type StudentRepositoryInterface interface {
	GetAllStudents() ([]models.Student, error)
	GetStudent(id int) (*models.Student, error)
	CreateStudent(student *models.Student) (*models.Student, error)
	UpdateStudent(student *models.Student) (*models.Student, error)
	SearchStudent(searchString string, pagination models.Pagination,
		sortBy models.SortBy) (*models.StudentSearchData, error)
	DeleteStudent(id int) (*models.Student, error)
}

type studentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *studentRepository {
	return &studentRepository{
		db: db,
	}
}

func (s *studentRepository) GetAllStudents() ([]models.Student, error) {

	stmnt, err := s.db.Prepare("SELECT * FROM students")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return nil, err
	}
	defer func(stmnt *sql.Stmt) {
		err := stmnt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmnt)

	rows, err := stmnt.Query()
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

	var studentList []models.Student

	for rows.Next() {
		var st models.Student

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

		studentList = append(studentList, st)

	}

	log.Debug("getAllStudents response : ", studentList)
	return studentList, nil
}

func (s *studentRepository) GetStudent(id int) (*models.Student, error) {
	var student models.Student

	stmnt, err := s.db.Prepare("SELECT * FROM students WHERE id = ?;")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &student, err
	}
	defer func(stmnt *sql.Stmt) {
		err := stmnt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmnt)

	err = stmnt.QueryRow(id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.Student{}, errors.New(consts.StudentNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &models.Student{}, err
	}

	log.Debug("Student : ", student)
	return &student, err
}

func (s *studentRepository) CreateStudent(student *models.Student) (*models.Student, error) {
	var st models.Student

	stmnt, err := s.db.Prepare("INSERT INTO students (firstname,lastname,year)" +
		" VALUES (?,?,?);")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &st, err
	}
	defer func(stmnt *sql.Stmt) {
		err := stmnt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmnt)

	result, err := stmnt.Exec(student.FirstName, student.LastName, student.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return &st, errors.New(consts.StudentNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &st, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(consts.DBResultIDError, err)
	}

	student.ID = int(id)

	log.Debug("Student : ", *student)
	return student, err
}

func (s *studentRepository) UpdateStudent(student *models.Student) (*models.Student, error) {
	var st models.Student

	stmnt, err := s.db.Prepare("UPDATE students SET firstname = ?, lastname = ?, year = ? " +
		"WHERE id = ?;")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &st, err
	}
	defer func(stmnt *sql.Stmt) {
		err := stmnt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmnt)

	_, err = stmnt.Exec(student.FirstName, student.LastName, student.Year, student.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &st, errors.New(consts.StudentNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &st, err
	}

	log.Debug("Student : ", *student)
	return student, err
}

func (s *studentRepository) SearchStudent(searchString string, pagination models.Pagination,
	sortBy models.SortBy) (*models.StudentSearchData, error) {

	query := fmt.Sprintf("SELECT *, Count(*) Over () AS TotalCount FROM "+
		"students WHERE firstname LIKE '%%%s%%' || lastname LIKE '%%%s%%' ORDER BY %s %s "+
		"LIMIT %v,%v;", searchString, searchString,
		sortBy.Column, sortBy.Direction, pagination.Page, pagination.PageSize)

	stmnt, err := s.db.Prepare(query)

	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return nil, err
	}
	defer func(stmnt *sql.Stmt) {
		err := stmnt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmnt)

	rows, err := stmnt.Query()
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

	var studentList []models.Student
	var totalCount int
	var resp models.StudentSearchData

	for rows.Next() {
		var st models.Student

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

		studentList = append(studentList, st)
	}

	resp.TotalElements = totalCount
	resp.Data = studentList

	log.Debug("getAllStudents response : ", resp)
	return &resp, nil
}

func (s *studentRepository) DeleteStudent(id int) (*models.Student, error) {
	var student models.Student

	stmnt, err := s.db.Prepare("DELETE FROM students WHERE id = ?;")
	if err != nil {
		log.Error(consts.QueryPrepareError, err)
		return &student, err
	}
	defer func(stmnt *sql.Stmt) {
		err := stmnt.Close()
		if err != nil {
			log.Error(consts.DBStatementCloseError, err)
		}
	}(stmnt)

	_, err = stmnt.Exec(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.Student{}, errors.New(consts.StudentNotFound)
		}
		log.Error(consts.DBResultsError, err)
		return &models.Student{}, err
	}

	log.Debug("Student id : ", id)
	return &student, nil
}
