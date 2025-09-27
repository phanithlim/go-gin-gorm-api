package response

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"message"`
	Data   interface{} `json:"data,omitempty"`
}
