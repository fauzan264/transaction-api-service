package user

type RegisterUserinput struct {
	Name			string `json:"name" binding:"required"`
	NIK				string `json:"nik" binding:"required"`
	PhoneNumber		string `json:"phone_number" binding:"required"`
}