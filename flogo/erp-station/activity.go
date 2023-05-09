package activity

import (
	"fmt"

	"github.com/project-flogo/core/activity"
)

func init() {
	//activity.Register(&Activity{}, New) to create instances using factory method 'New'
	_ = activity.Register(&Activity{})
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

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

	fmt.Println(input)

	return true, nil
}
