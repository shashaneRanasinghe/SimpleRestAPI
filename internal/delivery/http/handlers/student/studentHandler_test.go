package student

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/mocks"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	student0 = models.Student{
		ID:        0,
		FirstName: "Charles",
		LastName:  "Leclerc",
		Year:      3,
	}
	student1 = models.Student{
		ID:        1,
		FirstName: "Charles",
		LastName:  "Leclerc",
		Year:      3,
	}
	student2 = models.Student{
		ID:        2,
		FirstName: "Carlos",
		LastName:  "Sainz",
		Year:      1,
	}
	studentList = []models.Student{student1, student2}

	//expectedResponseError = `{"status":"Error","data":{"id":0,"firstname":"","lastname":"","year":0},"message":"Error Getting Students "}`
)

func NewMockStudentHandler_HappyPath(ctrl *gomock.Controller) *StudentHandler {
	student := mocks.NewMockStudentUsecase(ctrl)

	data := models.StudentSearchData{
		TotalElements: 2,
		Data:          studentList,
	}

	student.EXPECT().GetAllStudents().Return(studentList, nil)
	student.EXPECT().GetStudent(1).Return(&student1, nil)
	student.EXPECT().CreateStudent(&student0).Return(&student1, nil)
	student.EXPECT().UpdateStudent(&student1).Return(&student1, nil)
	student.EXPECT().DeleteStudent(1).Return(&student1, nil)
	student.EXPECT().SearchStudent("charl", models.Pagination{Page: 0, PageSize: 2},
		models.SortBy{Column: "firstname", Direction: "ASC"}).Return(&data, nil)

	return &StudentHandler{
		student: student,
	}
}

func TestStudentRoutes_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mux.NewRouter()

	studentHandler := NewMockStudentHandler_HappyPath(ctrl)

	r.HandleFunc("/", studentHandler.getAllStudents).Methods("GET")
	r.HandleFunc("/getStudent/{id}", studentHandler.getStudent).Methods("GET")
	r.HandleFunc("/", studentHandler.createStudent).Methods("POST")
	r.HandleFunc("/", studentHandler.updateStudent).Methods("PUT")
	r.HandleFunc("/{id}", studentHandler.deleteStudent).Methods("DELETE")
	r.HandleFunc("/search", studentHandler.searchStudents).Methods("GET") //studentHandler := NewStudentHandler(database.NewDatabase().GetConnection())

	testCases := []struct {
		name           string
		url            string
		method         string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Get All Students",
			url:            "/",
			method:         "GET",
			requestBody:    "",
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":[{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},{"id":2,"firstname":"Carlos","lastname":"Sainz","year":1}],"message":"Student Queried Successfully"}`,
		},
		{
			name:           "Get Specific Students",
			url:            "/getStudent/1",
			method:         "GET",
			requestBody:    "",
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Student Queried Successfully"}`,
		},
		{
			name:           "Create Student",
			url:            "/",
			method:         "POST",
			requestBody:    `{"firstname":"Charles","lastname":"Leclerc","year":3}`,
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Student Created Successfully"}`,
		},
		{
			name:           "Update Student",
			url:            "/",
			method:         "PUT",
			requestBody:    `{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3}`,
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Student Updated Successfully"}`,
		},
		{
			name:           "Delete Specific Student",
			url:            "/1",
			method:         "DELETE",
			requestBody:    "",
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},"message":"Student Deleted Successfully"}`,
		},
		{
			name:           "Search Students",
			url:            "/search",
			method:         "GET",
			requestBody:    `{"searchString":"charl","sortBy": {"column":"firstname","direction":"ASC"},"pagination": {"page":0,"pageSize":2}}`,
			expectedStatus: 200,
			expectedBody:   `{"status":"Success","data":{"totalElements":2,"data":[{"id":1,"firstname":"Charles","lastname":"Leclerc","year":3},{"id":2,"firstname":"Carlos","lastname":"Sainz","year":1}]},"message":"Student Queried Successfully"}`,
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
