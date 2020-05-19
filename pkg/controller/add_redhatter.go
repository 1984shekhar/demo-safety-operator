package controller

import (
	"github.com/akoserwal/demo-safety-operator.git/pkg/controller/redhatter"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, redhatter.Add)
}
