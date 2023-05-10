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
	ID            string      `json:"employee"`
	FirstName     interface{} `json:"first_name"`
	MiddleName    interface{} `json:"middle_name"`
	LastName      interface{} `json:"last_name"`
	JoiningDate   interface{} `json:"date_of_joining"`
	RelievingDate interface{} `json:"relieving_date"`
	Status        string      `json:"status"`
}