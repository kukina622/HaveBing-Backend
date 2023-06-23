package error

const (
	MISSING_PARAMETER      = "The request is missing a required parameter. Please provide the correct parameter and try again"
	UNAUTHORIZED           = "Unauthorized"
	RESOURCE_NOT_FOUND     = "The requested resource was not found."
	CREATE_RESOURCE_FAILED = "Failed to create the data"
	UPDATE_RESOURCE_FAILED = "Failed to update the data"
)

type ServerError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ServerError) Error() string {
	return e.Msg
}
