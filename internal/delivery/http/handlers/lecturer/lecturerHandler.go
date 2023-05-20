package lecturer

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/internal/repository"
	lec "github.com/shashaneRanasinghe/simpleAPI/internal/usecases/lecturer"
	"github.com/shashaneRanasinghe/simpleAPI/pkg/consts"
	"github.com/tryfix/log"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type LecturerHandler struct {
	lecturer lec.LecturerUsecase
}

func NewLecturerHandler(db *sql.DB) *LecturerHandler {
	lecturerRepo := repository.NewLecturerRepository(db)
	lecturer := lec.NewLecturer(lecturerRepo)
	return &LecturerHandler{
		lecturer: lecturer,
	}
}

func (handler *LecturerHandler) LecturerRoutes(r *mux.Router) {
	r.HandleFunc("/", handler.getAllLecturers).Methods("GET")
	r.HandleFunc("/getLecturer/{id}", handler.getLecturer).Methods("GET")
	r.HandleFunc("/", handler.createLecturer).Methods("POST")
	r.HandleFunc("/", handler.updateLecturer).Methods("PUT")
	r.HandleFunc("/{id}", handler.deleteLecturer).Methods("DELETE")
	r.HandleFunc("/search", handler.searchLecturers).Methods("GET")

}

func (handler *LecturerHandler) getAllLecturers(w http.ResponseWriter, r *http.Request) {
	var respModel models.LecturerListResponse

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	lecturers, err := handler.lecturer.GetAllLecturers()
	if err != nil {
		log.Error(consts.GetLecturersError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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
	respModel.Data = lecturers
	respModel.Message = consts.GetLecturer

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *LecturerHandler) getLecturer(w http.ResponseWriter, r *http.Request) {
	var respModel models.LecturerResponse
	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Error(consts.IDError, err)
		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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

	lecturer, err := handler.lecturer.GetLecturer(id)
	if err != nil {
		log.Error(consts.GetLecturersError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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
	respModel.Data = *lecturer
	respModel.Message = consts.GetLecturer

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *LecturerHandler) createLecturer(w http.ResponseWriter, r *http.Request) {
	var respModel models.LecturerResponse
	var newLecturer models.Lecturer

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(consts.RequestBodyReadError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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

	err = json.Unmarshal(body, &newLecturer)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	lecturer1, err := handler.lecturer.CreateLecturer(&newLecturer)
	if err != nil {
		log.Error(consts.GetLecturersError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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
	respModel.Data = *lecturer1
	respModel.Message = consts.LecturerCreated

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *LecturerHandler) updateLecturer(w http.ResponseWriter, r *http.Request) {
	var respModel models.LecturerResponse
	var updatedLecturer models.Lecturer

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(consts.RequestBodyReadError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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

	err = json.Unmarshal(body, &updatedLecturer)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	lecturer1, err := handler.lecturer.UpdateLecturer(&updatedLecturer)
	if err != nil {
		log.Error(consts.GetLecturersError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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
	respModel.Data = *lecturer1
	respModel.Message = consts.LecturerUpdated

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *LecturerHandler) deleteLecturer(w http.ResponseWriter, r *http.Request) {
	var respModel models.LecturerResponse
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

	lecturer, err := handler.lecturer.DeleteLecturer(id)
	if err != nil {
		log.Error(consts.GetLecturersError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.LecturerDeleteError

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
	respModel.Data = *lecturer
	respModel.Message = consts.LecturerDeleted

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}

func (handler *LecturerHandler) searchLecturers(w http.ResponseWriter, r *http.Request) {
	var respModel models.LecturerSearchResponse
	var reqBody models.LecturerSearchRequest

	w.Header().Set(consts.ContentType, consts.ApplicationJSON)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(consts.RequestBodyReadError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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

	lecturers, err := handler.lecturer.SearchLecturer(reqBody.SearchString, reqBody.Pagination,
		reqBody.SortBy)
	if err != nil {
		log.Error(consts.GetLecturersError, err)

		respModel.Status = consts.Error
		respModel.Message = consts.GetLecturersError

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
	respModel.Data = *lecturers
	respModel.Message = consts.GetLecturer

	b, err := json.Marshal(respModel)
	if err != nil {
		log.Error(consts.JSONMarshalError, err)
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(consts.ResponseWriteError, err)
	}
}
