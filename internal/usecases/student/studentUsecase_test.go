package student

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/shashaneRanasinghe/simpleAPI/internal/models"
	"github.com/shashaneRanasinghe/simpleAPI/mocks"
	"github.com/tryfix/log"
	"testing"
)

var (
	s1 = models.Student{
		ID:        1,
		FirstName: "test1",
		LastName:  "test1",
		Year:      1,
	}
	s2 = models.Student{
		ID:        2,
		FirstName: "test2",
		LastName:  "test2",
		Year:      2,
	}
	studentList = []models.Student{s1, s2}

	returnErr = errors.New("error")
)

func TestStudentUsecase_GetAllStudents_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected []models.Student
	}

	tests := []test{
		{
			expected: studentList,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().GetAllStudents().Return(studentList, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.GetAllStudents()
		if actual[0] != test.expected[0] || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}

}

func TestStudentUsecase_GetAllStudents_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().GetAllStudents().Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.GetAllStudents()
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_GetAllStudents(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().GetAllStudents().Return(studentList, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.GetAllStudents()
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_GetStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: &s1,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().GetStudent(1).Return(&s1, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.GetStudent(1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_GetStudent_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().GetStudent(1).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.GetStudent(1)
		if test.expected != err {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_GetStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().GetStudent(1).Return(&s1, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.GetStudent(1)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_CreateStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: &s1,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().CreateStudent(&s1).Return(&s1, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.CreateStudent(&s1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_CreateStudent_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().CreateStudent(&s1).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.CreateStudent(&s1)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_CreateStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().CreateStudent(&s1).Return(&s1, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.CreateStudent(&s1)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_UpdateStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: &s2,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().UpdateStudent(&s1).Return(&s2, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.UpdateStudent(&s1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_UpdateStudent_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().UpdateStudent(&s1).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.UpdateStudent(&s1)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_UpdateStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().UpdateStudent(&s1).Return(&s2, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.UpdateStudent(&s1)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_DeleteStudent_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected *models.Student
	}

	tests := []test{
		{
			expected: &s1,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().DeleteStudent(1).Return(&s1, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.DeleteStudent(1)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_DeleteStudent_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		expected error
	}

	tests := []test{
		{
			expected: returnErr,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().DeleteStudent(1).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		_, err := student.DeleteStudent(1)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, err)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_DeleteStudent(b *testing.B) {
	ctrl := gomock.NewController(b)
	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().DeleteStudent(1).Return(&s1, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.DeleteStudent(1)
		if err != nil {
			return
		}
	}
}

func TestStudentUsecase_SearchStudents_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	data := models.StudentSearchData{
		TotalElements: 2,
		Data:          studentList,
	}

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     *models.StudentSearchData
	}

	tests := []test{
		{
			searchString: "a",
			pagination: models.Pagination{
				Page:     0,
				PageSize: 2,
			},
			sortBy: models.SortBy{
				Column:    "firstname",
				Direction: "ASC",
			},
			expected: &data,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().SearchStudent(tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(&data, nil)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.SearchStudent(test.searchString, test.pagination, test.sortBy)
		if actual != test.expected || err != nil {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func TestStudentUsecase_SearchStudents_ErrorPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     error
	}

	tests := []test{
		{
			searchString: "a",
			pagination: models.Pagination{
				Page:     0,
				PageSize: 2,
			},
			sortBy: models.SortBy{
				Column:    "firstname",
				Direction: "ASC",
			},
			expected: returnErr,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().SearchStudent(tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(nil, returnErr)

	student := NewStudent(repo)

	for _, test := range tests {
		actual, err := student.SearchStudent(test.searchString, test.pagination, test.sortBy)
		if err != test.expected {
			log.Info("Expected : %v, Got : %v ", test.expected, actual)
			t.Fail()
		}
	}
}

func BenchmarkStudentUsecase_SearchStudents(b *testing.B) {
	ctrl := gomock.NewController(b)

	data := models.StudentSearchData{
		TotalElements: 2,
		Data:          studentList,
	}

	type test struct {
		searchString string
		pagination   models.Pagination
		sortBy       models.SortBy
		expected     *models.StudentSearchData
	}

	tests := []test{
		{
			searchString: "a",
			pagination: models.Pagination{
				Page:     0,
				PageSize: 2,
			},
			sortBy: models.SortBy{
				Column:    "firstname",
				Direction: "ASC",
			},
			expected: &data,
		},
	}

	repo := mocks.NewMockStudentRepository(ctrl)
	repo.EXPECT().SearchStudent(tests[0].searchString, tests[0].pagination,
		tests[0].sortBy).Return(&data, nil).AnyTimes()

	student := NewStudent(repo)

	for i := 0; i < b.N; i++ {
		_, err := student.SearchStudent(tests[0].searchString, tests[0].pagination,
			tests[0].sortBy)
		if err != nil {
			return
		}
	}
}
