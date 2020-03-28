package layouts

import (
	"errors"
	"fmt"
	"utevolux/layouts/simple"
)

const SimpleLayout = "simple"

type GenerateInput struct{
	Name, Package, Module string
}

func Generate(layout GenerateInput) error {
	switch layout.Name {
	case SimpleLayout:
		return simple.Generate(simple.SimpleInput{
			Package: layout.Package,
			Module:  layout.Module,
		})
	default:
		return errors.New(fmt.Sprint(layout.Name, " is an invalid layout"))
	}
}
