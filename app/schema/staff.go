package schema

type LogInSchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponseSchema struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

type CreateStaffSchema struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	HospitalName string `json:"hospital_name" binding:"required"`
}
