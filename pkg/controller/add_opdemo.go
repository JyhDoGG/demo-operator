package controller

import (
	"github.com/JyhDoGG/demo-operator/pkg/controller/opdemo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, opdemo.Add)
}
