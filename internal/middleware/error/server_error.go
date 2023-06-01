package error

type ServerError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ServerError) Error() string {
	return e.Msg
}
