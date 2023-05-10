package erpnext

type EmployeeIDList struct {
	Data []EmployeeID `json:"data"`
}

type EmployeeID struct {
	EmployeeID string `json:"name"`
}

type EmployeeDetailsList struct {
	Data EmployeeDetails `json:"data"`
}
type EmployeeDetails struct {
	ID            string `json:"employee"`
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	LastName      string `json:"last_name"`
	JoiningDate   string `json:"date_of_joining"`
	RelievingDate string `json:"relieving_date"`
	Status        string `json:"status"`
}
