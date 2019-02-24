package handler

import (
	"time"

	"github.com/timw255/flex-go"
)

// CustomEntity ...
type CustomEntity struct {
	flex.KinveyEntity
	IsActive *bool   `json:"isActive,omitempty"`
	Key      *string `json:"key,omitempty"`
}

// OnDeleteAll ...
func OnDeleteAll(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	count, err := objectsCollection.Remove("")
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	res := make(map[string]interface{})
	res["count"] = count

	resBytes, err := json.Marshal(res)

	return complete.SetBody(resBytes).Done()
}

// OnDeleteByID ...
func OnDeleteByID(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	count, err := objectsCollection.RemoveByID("5c4ce2291be4947c2a91d262")
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	res := make(map[string]interface{})
	res["count"] = count

	resBytes, err := json.Marshal(res)

	return complete.SetBody(resBytes).Done()
}

// OnDeleteByQuery ...
func OnDeleteByQuery(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	count, err := objectsCollection.Remove(`{"isActive":false}`)
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	res := make(map[string]interface{})
	res["count"] = count

	resBytes, err := json.Marshal(res)

	return complete.SetBody(resBytes).Done()
}

// OnGetAll ...
func OnGetAll(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	data, err := objectsCollection.Find("")
	if err != nil {
		modules.Logger.Error(err.Error())
		return complete.SetBody([]byte(`["` + err.Error() + `"]`)).Done()
	}

	result := make([]CustomEntity, 0)
	if err := json.Unmarshal(data, &result); err != nil {
		modules.Logger.Error(err.Error())
		return complete.SetBody([]byte(`["` + err.Error() + `"]`)).Done()
	}

	json, err := json.Marshal(result)
	if err != nil {
		modules.Logger.Error(err.Error())
		return complete.SetBody([]byte(`["` + err.Error() + `"]`)).Done()
	}

	return complete.SetBody(json).Done()
}

// OnGetByID ...
func OnGetByID(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	data, err := objectsCollection.FindByID("5c4ce2291be4947c2a91d262")
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	result := CustomEntity{}
	if err := json.Unmarshal(data, &result); err != nil {
		modules.Logger.Error(err.Error())
	}

	json, err := json.Marshal(result)

	return complete.SetBody(json).Done()
}

// OnGetByQuery ...
func OnGetByQuery(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	data, err := objectsCollection.Find(`{"isActive":false}`)
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	result := make([]CustomEntity, 0)
	if err := json.Unmarshal(data, &result); err != nil {
		modules.Logger.Error(err.Error())
	}

	json, err := json.Marshal(result)

	return complete.SetBody(json).Done()
}

// OnGetCount ...
func OnGetCount(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	count, _ := objectsCollection.Count("")

	res := make(map[string]interface{})
	res["count"] = count

	resBytes, _ := json.Marshal(res)

	return complete.SetBody(resBytes).Done()
}

// OnGetCountByQuery ...
func OnGetCountByQuery(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	count, _ := objectsCollection.Count(`{"isActive":true}`)

	res := make(map[string]interface{})
	res["count"] = count

	resBytes, _ := json.Marshal(res)

	return complete.SetBody(resBytes).Done()
}

// OnInsert ...
func OnInsert(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	entity := &CustomEntity{
		IsActive: flex.Bool(true),
		Key:      flex.String("some key"),
	}

	entity.KinveyEntity = modules.KinveyEntity.NewKinveyEntity("")

	entity.ACL.AddReaderRole("0f350bba-1145-e342-cb5a-223f314b650d")

	data, err := objectsCollection.Save(entity)
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	newEntity := CustomEntity{}
	if err := json.Unmarshal(data, &newEntity); err != nil {
		modules.Logger.Error(err.Error())
	}

	json, err := json.Marshal(newEntity)

	return complete.SetBody(json).Done()
}

// OnUpdate ...
func OnUpdate(context *flex.Request, complete flex.KinveyCompletionHandler, modules flex.Modules) (*flex.Task, *flex.Task) {
	dataStore := modules.DataStore.NewDataStore(true, true)
	objectsCollection := dataStore.NewCollection("objects")

	t := time.Now()

	entity := &CustomEntity{
		Key: flex.String(t.String()),
	}

	entity.ID = flex.String("5c4ce2291be4947c2a91d262")

	data, err := objectsCollection.Save(entity)
	if err != nil {
		modules.Logger.Error(err.Error())
	}

	return complete.SetBody(data).Done()
}
