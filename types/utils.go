package types

type ApiResponse struct {
	Result      int64       `json:"result"`
	Description string      `json:"description"`
	ErrorCode   interface{} `json:"errCode"`
}

func NewAPIResponse(result int64, description string, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		Result:      result,
		Description: description,
		ErrorCode:   errCode,
	}
}
