package auth_struct

import auth_model "class_main_service/src/features/auth/model"

type RegisterInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginOutput struct {
	AccessToken string           `json:"access_token"`
	User        *auth_model.User `json:"user"`
}
