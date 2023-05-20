package consts

// Handler Errors
const (
	JSONMarshalError      = "Json Marshalling Error "
	ResponseWriteError    = "Error Writing Response "
	RequestBodyReadError  = "Error Reading The Request Body"
	RequestBodyCloseError = "Error Closing The Request Body"
	IDError               = "Error Getting The ID"
)

// DB ERRORS
const (
	QueryPrepareError     = "Error Preparing Query "
	DBResultsError        = "Error Getting Results From The DB "
	DBRowCloseError       = "Error Closing DB Rows "
	DBScanRowError        = "Error Scanning DB Rows"
	DBRowsError           = "Error In DB Rows"
	DBResultIDError       = "Error Getting Insert ID "
	DBStatementCloseError = "Error Closing Prepared Statement"
)

const (
	StudentNotFound    = "student Not Found"
	StudentDeleteError = "Error Deleting Student"
	GetStudentsError   = "Error Getting Students "
)

const (
	LecturerNotFound    = "lecturer Not Found"
	LecturerDeleteError = "Error Deleting Lecturer"
	GetLecturersError   = "Error Getting Lecturers "
)
