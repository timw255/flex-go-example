package handler

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/timw255/flex-go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type endpointMessage struct {
	Message string `json:"message"`
}

// KinveyEndpointHandler ...
func KinveyEndpointHandler(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	endpointRunner := modules.EndpointRunner.NewEndpointRunner(true)
	testEndpoint := endpointRunner.NewEndpoint("test")

	requestMessage := endpointMessage{
		Message: "ping!",
	}

	requestData, err := json.Marshal(requestMessage)
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	responseData, err := testEndpoint.Execute(requestData)
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	return complete.SetBody(responseData).Done()
}

// LocalEndpointHandler ...
func LocalEndpointHandler(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	message := endpointMessage{
		Message: "Hello from SomeHandler (Go)!",
	}

	json, err := json.Marshal(message)
	if err != nil {
		modules.Logger.Error(err.Error())
		return complete.SetBody([]byte(`["There wan an error :("]`)).Done()
	}

	return complete.SetBody(json).Done()
}
