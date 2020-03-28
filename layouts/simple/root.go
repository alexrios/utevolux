package simple

type RootValues struct {
	Package string
	Module string
}

const Dummy = `package {{.Package}}

import "github.com/google/uuid"

type Dummy struct {
	ID   uuid.UUID
}

type DummyService interface {
	Dummies() ([]*Dummy, error)
	CreateDummy(person Dummy) (uuid.UUID, error)
	RemoveDummy(uuid.UUID, error)
	Dummy(uuid.UUID) (Dummy, error)
}
`

const GoMod = `module {{.Module}}

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/urfave/negroni v1.0.0
)
`