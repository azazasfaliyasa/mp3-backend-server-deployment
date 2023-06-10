package dto

type Response struct {
	Meta MetaResponse `json:"meta"`
	Data any          `json:"data"`
}

type MetaResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
