package main

import (
	"github.com/timw255/flex-go"
	"github.com/timw255/flex-go-example/handler"
)

func main() {
	options := flex.NewOptions("localhost", 10001, "")
	flex.NewService(options, func(err error, f flex.Flex) {
		if err != nil {
			f.Logger.Error("Error initializing the Flex SDK, exiting.")
		}

		// functions
		f.Functions.Register("local", handler.LocalEndpointHandler)
		f.Functions.Register("kinvey", handler.KinveyEndpointHandler)

		// data
		widgets := f.Data.NewServiceObject("widgets")
		widgets.OnDeleteAll(handler.OnDeleteAll)
		widgets.OnDeleteByID(handler.OnDeleteByID)
		widgets.OnDeleteByQuery(handler.OnDeleteByQuery)
		widgets.OnGetAll(handler.OnGetAll)
		widgets.OnGetByID(handler.OnGetByID)
		widgets.OnGetByQuery(handler.OnGetByQuery)
		widgets.OnGetCount(handler.OnGetCount)
		widgets.OnGetCountByQuery(handler.OnGetCountByQuery)
		widgets.OnInsert(handler.OnInsert)
		widgets.OnUpdate(handler.OnUpdate)

		// auth
		f.Auth.Register("myAuth", handler.AuthHandler)
	})
}
