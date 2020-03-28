package simple

import "os"

type SimpleInput struct{
	Package, Module string
}

func Generate(input SimpleInput) error {
	values := RootValues{
		Package: input.Package,
		Module: input.Module,
	}

	err := Root(values)
	if err != nil {
		return err
	}
	err = Cmd(values)
	if err != nil {
		return err
	}
	err = Api(values)
	if err != nil {
		return err
	}
	return nil
}

//Root create domain files, in this case:
//- dummy.go
func Root(v RootValues) error {
	domain := &FileTemplate{
		Name:    "dummy.go",
		Literal: Dummy,
		Values:  v,
	}
	err := domain.WriteFile()
	if err != nil {
		return err
	}

	gomod := &FileTemplate{
		Name:    "go.mod",
		Literal: GoMod,
		Values:  v,
	}
	err = gomod.WriteFile()
	if err != nil {
		return err
	}

	return nil
}

//TODO
func Cmd(v RootValues) error {
	err := os.MkdirAll("cmd", 0755)
	if err != nil {
		return err
	}
	main := &FileTemplate{
		Name:    "cmd/main.go",
		Literal: Main,
		Values:  v,
	}
	return main.WriteFile()
}

func Api(v RootValues) error {
	err := os.MkdirAll("api", 0755)
	if err != nil {
		return err
	}
	handler := &FileTemplate{
		Name:    "api/dummy.go",
		Literal: DummyHandler,
		Values:  v,
	}
	err = handler.WriteFile()
	if err != nil {
		return err
	}

	routes := &FileTemplate{
		Name:    "api/routes.go",
		Literal: Routes,
		Values:  v,
	}
	err = routes.WriteFile()
	if err != nil {
		return err
	}
	return nil
}
