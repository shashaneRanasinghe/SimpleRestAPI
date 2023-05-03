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
	"io/ioutil"
	"net/http"
	"strconv"
)

var student st.StudentUsecase

func StudentRoutes(r *mux.Router, db *sql.DB) {
	studentRepo := repository.NewStudentRepository(db)
	student = st.NewStudent(studentRepo)

	r.HandleFunc("/", getAllStudents).Methods("GET")
	r.HandleFunc("/getStudent/{id}", getStudent).Methods("GET")
	r.HandleFunc("/", createStudent).Methods("POST")
	r.HandleFunc("/", updateStudent).Methods("PUT")
	r.HandleFunc("/{id}", deleteStudent).Methods("DELETE")
	r.HandleFunc("/search", searchStudents).Methods("GET")

}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentListResponse

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	students, err := student.GetAllStudents()
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

func getStudent(w http.ResponseWriter, r *http.Request) {
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

	student, err := student.GetStudent(id)
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

func createStudent(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentResponse
	var newStudent models.Student

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := ioutil.ReadAll(r.Body)
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

	student1, err := student.CreateStudent(&newStudent)
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

func updateStudent(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentResponse
	var updatedStudent models.Student

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := ioutil.ReadAll(r.Body)
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

	student1, err := student.UpdateStudent(&updatedStudent)
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

func deleteStudent(w http.ResponseWriter, r *http.Request) {
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

	student, err := student.DeleteStudent(id)
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

func searchStudents(w http.ResponseWriter, r *http.Request) {
	var respModel models.StudentSearchResponse
	var reqBody models.StudentSearchRequest

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := ioutil.ReadAll(r.Body)
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

	students, err := student.SearchStudent(reqBody.SearchString, reqBody.Pagination, reqBody.SortBy)
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
