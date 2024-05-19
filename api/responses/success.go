package responses

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// NewSuccessResponse creates a new success response with the given data.
func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Data: data,
	}
}
