package dto

type CreateEmployeeRequest struct {
	UserId   int64              `json:"userId"`
	Employee CreateEmployeeData `json:"employee"`
}

type CreateEmployeeData struct {
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Surname   string `json:"surName"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
	BirthDate string `json:"birthDate"`
}
