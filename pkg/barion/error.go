package barion

// ErrorResponse is the general error type
type ErrorResponse struct {
	Status string
	Errors *[]barionError `json:"Errors"`
}

type barionError struct {
	ErrorCode   string
	Title       string
	Description string
	//EndPoint    string
	AuthData   string
	HappenedAt string
}
