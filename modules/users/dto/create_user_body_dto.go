package modules_users_dto

type CreateUserBodyDTO struct {
	Username string `json:"username" validate:"required"`
}
