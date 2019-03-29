package controller

import (
	"github.com/pdmack/h2oai-dai-operator/pkg/controller/driverlessai"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, driverlessai.Add)
}
