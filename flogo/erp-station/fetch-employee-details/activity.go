package fetchemployee

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chintansakhiya/activity/flogo/erp-station/fetch-employee-details/config"
	"github.com/chintansakhiya/activity/flogo/erp-station/fetch-employee-details/database"
	"github.com/chintansakhiya/activity/flogo/erp-station/fetch-employee-details/pkg/erpnext"
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// New optional factory method, should be used if one activity instance per configuration is desired
// func New(ctx activity.InitContext) (activity.Activity, error) {

// 	s := &Settings{}
// 	err := metadata.MapToStruct(ctx.Settings(), s, true)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ctx.Logger().Debugf("Setting: %s", s.ASetting)

// 	act := &Activity{} //add aSetting to instance

// 	return act, nil
// }

// Activity is an sample Activity that can be used as a base to create a custom activity

// Metadata returns the activity's metadata
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
	fmt.Println("fron activity chintan", input)

	// output := &Output{AnOutput: input.EmployeeId}
	// err = ctx.SetOutputObject(output)
	// if err != nil {
	// 	return true, err
	// }
	cfg := config.GetConfig()

	db, _ := database.PostgresDBConnection(cfg)
	if err != nil {
		return false, err
	}
	var id erpnext.EmployeeID
	err = json.Unmarshal([]byte(input.EmployeeId), &id)
	if err != nil {
		return false, err
	}
	log.Fatal(id)
	err = erpnext.GetDetails(db, id.EmployeeID)
	if err != nil {
		return false, err
	}

	return true, nil
}
