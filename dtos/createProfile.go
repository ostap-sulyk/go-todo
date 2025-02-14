package dtos

type CreateProfileRequestDTO struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type CreateProfileResponseDTO struct {
	URI string
}
