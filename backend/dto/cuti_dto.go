package dto

type CreateCutiRequest struct {
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Informations string `json:"information"`
}

type CreateCutiResponse struct {
	UserName       string `json:"users_name"`
	SubmissionDate string `json:"submission_date"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Information    string `json:"information"`
	StatusCuti     string `json:"status_cuti"`
}
