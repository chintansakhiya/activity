package fetchemployee

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/chintansakhiya/activity/flogo/erp-station/fetch-employee-details/database"
	"github.com/chintansakhiya/activity/flogo/erp-station/fetch-employee-details/pkg/erpnext"
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ctx.Logger().Debugf("Input: %s", input.EmployeeId)
	fmt.Println("fron activity chintan", input.EmployeeId)

	// output := &Output{AnOutput: input.EmployeeId}
	// err = ctx.SetOutputObject(output)
	// if err != nil {
	// 	return true, err
	// }
	// cfg := config.GetConfig()

	db, err := database.PostgresDBConnection()
	if err != nil {
		return false, err
	}
	var id erpnext.EmployeeID
	temp := strings.Replace(input.EmployeeId, `":""}`, "", -5)
	temp = strings.Replace(temp, `{"`, "", -2)
	temp = strings.Replace(temp, `\`, "", -1)

	err = json.Unmarshal([]byte(temp), &id)
	if err != nil {
		fmt.Println(err)
	}

	// err = json.Unmarshal([]byte(input.EmployeeId), &id)
	// if err != nil {
	// 	return false, err
	// }
	if err != nil {
		log.Fatal(err)

	}

	err = erpnext.GetDetails(db, id.EmployeeID)
	if err != nil {
		return false, err
	}

	return true, nil
}
