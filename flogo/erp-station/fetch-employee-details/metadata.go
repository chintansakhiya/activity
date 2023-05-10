package fetchemployee

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

type Input struct {
	employeeId string `md:"employeeId,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, err := coerce.ToString(values["employeeId"])
	if err != nil {
		return err
	}
	r.employeeId = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"employeeId": r.employeeId,
	}
}

type Activity struct {
}
