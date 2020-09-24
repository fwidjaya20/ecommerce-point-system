package models

type UserPointResponse struct {
	TotalPoint float64 `json:"total_point"`
}

type UserPointInquiryResponse struct {
	RequestId string `json:"request_id"`
}