package models

type PublishAddOrDeductPointRequest struct {
	UserId string `http_url:"user_id"`
	Point float64 `json:"point"`
}

type AddOrDeductPointRequest struct {
	Id        string
	UserId    string
	Point     float64
	PointType string
	Notes     string
}

type GetPointRequest struct {
	UserId string `http_url:"user_id"`
}