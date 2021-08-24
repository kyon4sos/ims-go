package dto

type UserDto struct {
	Username string `json:"username" validate:"required,gte=6,lte=20,tname=用户名"`
	Password string `json:"password" validate:"required,gte=6,lte=20"`
}

type UserRegDto struct {
	UserDto
	NickName string `json:"nickname" binding:"required,gte=6,lte=20"`
}