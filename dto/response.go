package dto

const (
	STATUS_SUCCESS = "success"
	STATUS_ERROR   = "error"
)

type Response struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ResponseMeta struct {
	AfterCursor  string `json:"after_cursor"`
	BeforeCursor string `json:"before_cursor"`
}
