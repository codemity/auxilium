package validate

type format string

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
