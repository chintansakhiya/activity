package fetchemployee

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

type EmployeeID struct {
	EmployeeID EmployeeIDs `json:"name"`
}
type EmployeeIDs struct {
	Emp string `json:""`
}

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

	// db, _ := database.PostgresDBConnection(cfg)
	// if err != nil {
	// 	return false, err
	// }
	var id EmployeeID
	err = json.Unmarshal([]byte(input.EmployeeId), &id)
	if err != nil {
		return false, err
	}
	log.Fatal(id)
	// err = erpnext.GetDetails(db, id.EmployeeID)
	if err != nil {
		return false, err
	}

	return true, nil
}
