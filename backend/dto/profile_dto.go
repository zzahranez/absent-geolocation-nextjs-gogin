package dto

type ProfilePageResponse struct {
	Bio         string              `json:"bio"`
	NoHp        string              `json:"no_hp"`
	Address     string              `json:"address"`
	Avatar      *string             `json:"avatar,omitempty"`
	AccountData UserAccountResponse `json:"account"`
}

type UserAccountResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
