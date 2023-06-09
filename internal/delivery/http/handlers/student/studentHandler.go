package student

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/internal/repository"
	st "github.com/shashaneRanasinghe/simpleAPI/internal/usecases/student"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/consts"
	"github.com/tryfix/log"
	"io"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	student st.StudentUsecase
}

func NewStudentHandler(db *sql.DB) *StudentHandler {
	studentRepo := repository.NewStudentRepository(db)
	student := st.NewStudent(studentRepo)
	return &StudentHandler{
		student: student,
	}
}

func (handler *StudentHandler) StudentRoutes(r *mux.Router) {

	r.HandleFunc("/", handler.getAllStudents).Methods("GET")
	r.HandleFunc("/getStudent/{id}", handler.getStudent).Methods("GET")
	r.HandleFunc("/", handler.createStudent).Methods("POST")
	r.HandleFunc("/", handler.updateStudent).Methods("PUT")
	r.HandleFunc("/{id}", handler.deleteStudent).Methods("DELETE")
	r.HandleFunc("/search", handler.searchStudents).Methods("GET")

}

func (handler *StudentHandler) getAllStudents(w http.ResponseWriter, _ *http.Request) {
	var respModel models.StudentListResponse

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	students, err := handler.student.GetAllStudents()
	if err != nil {
		log.Error(consts.GetStudentsError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	respModel.Status = consts.Success
	respModel.Data = students
	respModel.Message = consts.GetStudent

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *StudentHandler) getStudent(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentResponse
	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Error(consts.IDError, err)
		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}

	student, err := handler.student.GetStudent(id)
	if err != nil {
		log.Error(consts.GetStudentsError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	respModel.Status = consts.Success
	respModel.Data = *student
	respModel.Message = consts.GetStudent

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *StudentHandler) createStudent(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentResponse
	var newStudent models.Student

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(consts.RequestBodyReadError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(consts.RequestBodyCloseError, err)
		}
	}(r.Body)

	err = json.Unmarshal(body, &newStudent)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	student1, err := handler.student.CreateStudent(&newStudent)
	if err != nil {
		log.Error(consts.GetStudentsError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	respModel.Status = consts.Success
	respModel.Data = *student1
	respModel.Message = consts.StudentCreated

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *StudentHandler) updateStudent(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentResponse
	var updatedStudent models.Student

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(consts.RequestBodyReadError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(consts.RequestBodyCloseError, err)
		}
	}(r.Body)

	err = json.Unmarshal(body, &updatedStudent)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	student1, err := handler.student.UpdateStudent(&updatedStudent)
	if err != nil {
		log.Error(consts.GetStudentsError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	respModel.Status = consts.Success
	respModel.Data = *student1
	respModel.Message = consts.StudentUpdated

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *StudentHandler) deleteStudent(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentResponse
	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Error(consts.IDError, err)
		respModel.Status = consts.Error
		respModel.Message = consts.IDError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}

	student, err := handler.student.DeleteStudent(id)
	if err != nil {
		log.Error(consts.GetStudentsError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.StudentDeleteError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	respModel.Status = consts.Success
	respModel.Data = *student
	respModel.Message = consts.StudentDeleted

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *StudentHandler) searchStudents(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentSearchResponse
	var reqBody models.StudentSearchRequest

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error(consts.RequestBodyReadError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(consts.RequestBodyCloseError, err)
		}
	}(r.Body)

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	students, err := handler.student.SearchStudent(reqBody.SearchString, reqBody.Pagination,
		reqBody.SortBy)
	if err != nil {
		log.Error(consts.GetStudentsError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetStudentsError

		w.WriteHeader(http.StatusInternalServerError)
		b, err := json.Marshal(respModel)
		if err != nil {
			log.Error(consts.JSONMarshalError, err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error(consts.ResponseWriteError, err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	respModel.Status = consts.Success
	respModel.Data = *students
	respModel.Message = consts.GetStudent

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}
