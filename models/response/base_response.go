package response

type APIBaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
