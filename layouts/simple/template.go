package simple

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

type FileTemplate struct {
	Name     string
	Literal  string
	Values   interface{}
}

func (p *FileTemplate) WriteFile() error {
	t := template.Must(template.New(p.Name).Parse(p.Literal))
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, p.Values); err != nil {
		return err
	}
	err := ioutil.WriteFile(p.Name, tpl.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
