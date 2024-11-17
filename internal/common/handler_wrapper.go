package common

import (
	"fmt"

	"github.com/thuongtruong109/gouse/types"
	"github.com/thuongtruong109/soundlib/pkg/helpers"
)

type CommonHandler struct {
	helper helpers.Helper
	scope  string
}

func NewCommonHandler(helper helpers.Helper, scope string) *CommonHandler {
	return &CommonHandler{
		helper: helper,
		scope:  scope,
	}
}

func (ch *CommonHandler) ErrorWrapper(desc string, err error) {
	if err != nil {
		ch.helper.OutputError(desc, fmt.Sprintf("%v", err))
		return
	}
}

func (ch *CommonHandler) SuccessWrapper(desc string, body_output []string) {
	ch.helper.OutputSuccess(desc)
	bodyLen := len(body_output)
	if bodyLen > 0 {
		if bodyLen > 1 {
			helpers.TableOutput[string, string, interface{}](ch.scope, body_output, "Total: "+types.IntToString(bodyLen))
			return
		}

		helpers.TableOutput[string, string, interface{}](ch.scope, body_output, nil)
		return
	}
}

func (ch *CommonHandler) SuccessWrapper2(desc string, body_output []string, time string) {
	ch.helper.OutputSuccess(desc)
	bodyLen := len(body_output)
	if bodyLen > 0 {
		if bodyLen > 1 {
			helpers.TableOutput2[string, string, interface{}, string, string](ch.scope, body_output, "Total: "+types.IntToString(bodyLen), "\t- Status: "+desc, "- Time: "+time)
			return
		}

		helpers.TableOutput[string, string, interface{}](ch.scope, body_output, nil)
		return
	}
}
