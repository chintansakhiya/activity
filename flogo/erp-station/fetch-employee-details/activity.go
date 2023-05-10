package fetchemployee

import (
	"fmt"
	"log"

	"github.com/project-flogo/core/activity"
)

func init() {
	err := activity.Register(&Activity{})
	if err != nil {
		log.Fatal(err)
	}
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{})

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}
	err = ctx.GetInputObject(input)

	if err != nil {
		return false, err
	}
	fmt.Println(input)
	// cfg := config.GetConfig()

	// db, err := database.PostgresDBConnection(cfg)
	// if err != nil {
	// 	return false, err
	// }
	// var id erpnext.EmployeeID
	// err = json.Unmarshal([]byte(input.employeeId), &id)
	// if err != nil {
	// 	return false, err
	// }
	// err = erpnext.GetDetails(db, id.EmployeeID)
	// if err != nil {
	// 	return false, err
	// }
	return true, nil
}
