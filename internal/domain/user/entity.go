package user

type User struct {
	UserID   int    `json:"user_id" binding:"required"`
	UserName string `json:"user_name" binding:"required"`
	Pass     string `json:"pass" binding:"required"`
}
