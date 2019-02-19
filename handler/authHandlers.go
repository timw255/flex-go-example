package handler

import (
	"github.com/timw255/flex-go"
)

// AuthHandler ...
func AuthHandler(context *flex.Request, complete flex.AuthCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	return complete.SetToken("41jknt32ntl34h234bthj3b24t").OK().Done()
}
