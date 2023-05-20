package lecturer

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/mocks"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	lecturer0 = models.Lecturer{
		ID:        0,
		FirstName: "Charles",
		LastName:  "Leclerc",
		Year:      3,
	}
	lecturer1 = models.Lecturer{
		ID:        1,
		FirstName: "Charles",
		LastName:  "Leclerc",
		Year:      3,
	}
	lecturer2 = models.Lecturer{
		ID:        2,
		FirstName: "Carlos",
		LastName:  "Sainz",
		Year:      1,
	}
	errLecturer  = &models.Lecturer{}
	lecturerList = []models.Lecturer{lecturer1, lecturer2}
	ErrResponse  = errors.New("error Getting Lecturers")

	expectedResponseError = `{"status":"Error","data":{"id":0,"firstname":"","lastname":"","year":0},"message":"Error Getting Lecturers "}`
)

func NewMockLecturerHandler_HappyPath(ctrl *gomock.Controller) *LecturerHandler {
	lecturer := mocks.NewMockLecturerUsecase(ctrl)

	data := models.LecturerSearchData{
		TotalElements: 2,
		Data:          lecturerList,
	}

	lecturer.EXPECT().GetAllLecturers().Return(lecturerList, nil)
	lecturer.EXPECT().GetLecturer(1).Return(&lecturer1, nil)
	lecturer.EXPECT().CreateLecturer(&lecturer0).Return(&lecturer1, nil)
	lecturer.EXPECT().UpdateLecturer(&lecturer1).Return(&lecturer1, nil)
	lecturer.EXPECT().DeleteLecturer(1).Return(&lecturer1, nil)
	lecturer.EXPECT().SearchLecturer("charl", models.Pagination{Page: 0, PageSize: 2},
		models.SortBy{Column: "firstname", Direction: "ASC"}).Return(&data, nil)

	return &LecturerHandler{
		lecturer: lecturer,
	}
}

func TestLecturerRoutes_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mux.NewRouter()

	lecturerHandler := NewMockLecturerHandler_HappyPath(ctrl)

	r.HandleFunc("/", lecturerHandler.getAllLecturers).Methods("GET")
	r.HandleFunc("/getLecturer/{id}", lecturerHandler.getLecturer).Methods("GET")
	r.HandleFunc("/", lecturerHandler.createLecturer).Methods("POST")
	r.HandleFunc("/", lecturerHandler.updateLecturer).Methods("PUT")
	r.HandleFunc("/{id}", lecturerHandler.deleteLecturer).Methods("DELETE")
	r.HandleFunc("/search", lecturerHandler.searchLecturers).Methods("GET") //lecturerHandler := NewLecturerHandler(database.NewDatabase().GetConnection())

	testCases := []struct {
		name           string
		url            string
		method         string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Get All Lecturers",
			url:            "/",
			method:         "GET",
			requestBody:    "",
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":[{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},{"id":2,"firstname":"Carlos","lastname":"Sainz","year":1}],"message":"Lecturer Queried Successfully"}`,
		},
		{
			name:           "Get Specific Lecturers",
			url:            "/getLecturer/1",
			method:         "GET",
			requestBody:    "",
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Lecturer Queried Successfully"}`,
		},
		{
			name:           "Create Lecturer",
			url:            "/",
			method:         "POST",
			requestBody:    `{"firstname":"Charles","lastname":"Leclerc","year":3}`,
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Lecturer Created Successfully"}`,
		},
		{
			name:           "Update Lecturer",
			url:            "/",
			method:         "PUT",
			requestBody:    `{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3}`,
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Lecturer Updated Successfully"}`,
		},
		{
			name:           "Delete Specific Lecturer",
			url:            "/1",
			method:         "DELETE",
			requestBody:    "",
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Lecturer Deleted Successfully"}`,
		},
		{
			name:           "Search Lecturers",
			url:            "/search",
			method:         "GET",
			requestBody:    `{"searchString":"charl","sortBy": {"column":"firstname","direction":"ASC"},"pagination": {"page":0,"pageSize":2}}`,
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"totalElements":2,"data":[{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},{"id":2,"firstname":"Carlos","lastname":"Sainz","year":1}]},"message":"Lecturer Queried Successfully"}`,
		},
	}

	for _, test := range testCases {
		req := httptest.NewRequest(test.method, test.url, strings.NewReader(test.requestBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != test.expectedStatus {
			t.Errorf("Test %s : Expected status code %d, but got %d", test.name, test.expectedStatus, w.Code)
		}

		if w.Body.String() != test.expectedBody {
			t.Errorf("Test %s : Expected response body %s, but got %s", test.name, test.expectedBody, w.Body.String())
		}
	}
}

func NewMockLecturerHandler_ErrorPath(ctrl *gomock.Controller) *LecturerHandler {
	lecturer := mocks.NewMockLecturerUsecase(ctrl)

	lecturer.EXPECT().GetAllLecturers().Return(nil, ErrResponse)
	lecturer.EXPECT().GetLecturer(1).Return(errLecturer, ErrResponse)
	lecturer.EXPECT().CreateLecturer(&lecturer0).Return(errLecturer, ErrResponse)
	lecturer.EXPECT().UpdateLecturer(&lecturer1).Return(errLecturer, ErrResponse)
	lecturer.EXPECT().DeleteLecturer(1).Return(errLecturer, ErrResponse)
	lecturer.EXPECT().SearchLecturer("charl", models.Pagination{Page: 0, PageSize: 2},
		models.SortBy{Column: "firstname", Direction: "ASC"}).Return(nil, ErrResponse)

	return &LecturerHandler{
		lecturer: lecturer,
	}
}

func TestLecturerRoutes_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mux.NewRouter()

	lecturerHandler := NewMockLecturerHandler_ErrorPath(ctrl)

	r.HandleFunc("/", lecturerHandler.getAllLecturers).Methods("GET")
	r.HandleFunc("/getLecturer/{id}", lecturerHandler.getLecturer).Methods("GET")
	r.HandleFunc("/", lecturerHandler.createLecturer).Methods("POST")
	r.HandleFunc("/", lecturerHandler.updateLecturer).Methods("PUT")
	r.HandleFunc("/{id}", lecturerHandler.deleteLecturer).Methods("DELETE")
	r.HandleFunc("/search", lecturerHandler.searchLecturers).Methods("GET") //lecturerHandler := NewLecturerHandler(database.NewDatabase().GetConnection())

	testCases := []struct {
		name           string
		url            string
		method         string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Get All Lecturers",
			url:            "/",
			method:         "GET",
			requestBody:    "",
			expectedStatus: 500,
			expectedBody:   `{"status":"Error","data":null,"message":"Error Getting Lecturers "}`,
		},
		{
			name:           "Get Specific Lecturers",
			url:            "/getLecturer/1",
			method:         "GET",
			requestBody:    "",
			expectedStatus: 500,
			expectedBody:   expectedResponseError,
		},
		{
			name:           "Create Lecturer",
			url:            "/",
			method:         "POST",
			requestBody:    `{"firstname":"Charles","lastname":"Leclerc","year":3}`,
			expectedStatus: 500,
			expectedBody:   expectedResponseError,
		},
		{
			name:           "Update Lecturer",
			url:            "/",
			method:         "PUT",
			requestBody:    `{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3}`,
			expectedStatus: 500,
			expectedBody:   expectedResponseError,
		},
		{
			name:           "Delete Specific Lecturer",
			url:            "/1",
			method:         "DELETE",
			requestBody:    "",
			expectedStatus: 500,
			expectedBody:   `{"status":"Error","data":{"id":0,"firstname":"","lastname":"","year":0},"message":"Error Deleting Lecturer"}`,
		},
		{
			name:           "Search Lecturers",
			url:            "/search",
			method:         "GET",
			requestBody:    `{"searchString":"charl","sortBy": {"column":"firstname","direction":"ASC"},"pagination": {"page":0,"pageSize":2}}`,
			expectedStatus: 500,
			expectedBody:   `{"status":"Error","data":{"totalElements":0,"data":null},"message":"Error Getting Lecturers "}`,
		},
	}

	for _, test := range testCases {
		req := httptest.NewRequest(test.method, test.url, strings.NewReader(test.requestBody))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		if w.Code != test.expectedStatus {
			t.Errorf("Test %s : Expected status code %d, but got %d", test.name, test.expectedStatus, w.Code)
		}

		if w.Body.String() != test.expectedBody {
			t.Errorf("Test %s : Expected response body %s, but got %s", test.name, test.expectedBody, w.Body.String())
		}
	}
}
