package erpnext

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/doug-martin/goqu/v9"
)

func GetDetails(db *goqu.Database, id string) error {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://10.0.8.190:8090/api/resource/Employee/%s", id), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token 472e6bdd5f355a2:a5b7f758c8cfbf0")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	detail, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var employeeDetails EmployeeDetailsList
	err = json.Unmarshal(detail, &employeeDetails)
	if err != nil {
		return err
	}
	fmt.Println(employeeDetails.Data)
	sql, _, err := db.Insert("employee").Rows(
		goqu.Record{
			"id":             employeeDetails.Data.ID,
			"first_name":     employeeDetails.Data.FirstName,
			"middle_name":    employeeDetails.Data.MiddleName,
			"last_name":      employeeDetails.Data.LastName,
			"joining_date":   employeeDetails.Data.JoiningDate,
			"relieving_date": employeeDetails.Data.RelievingDate,
			"status":         employeeDetails.Data.Status,
		}).OnConflict(goqu.DoUpdate("id", goqu.Record{
		"first_name":     employeeDetails.Data.FirstName,
		"middle_name":    employeeDetails.Data.MiddleName,
		"last_name":      employeeDetails.Data.LastName,
		"joining_date":   employeeDetails.Data.JoiningDate,
		"relieving_date": employeeDetails.Data.RelievingDate,
		"status":         employeeDetails.Data.Status,
	})).ToSQL()
	fmt.Println("sql:", sql)
	if err != nil {
		return err

	}
	fmt.Println("6")

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
