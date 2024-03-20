package service

import (
	"io"
	"io/fs"
	"text/template"

	"github.com/m4n5ter/cnsoftbei/cmd/core"
)

type Service struct {
	t    *template.Template
	Name string
}

func New(fsys fs.FS, name string) *Service {
	m := &Service{}
	t := template.New("service.tmpl")
	t.Funcs(core.Funcs)
	m.t = template.Must(t.ParseFS(fsys, "templates/service.tmpl"))

	m.Name = name
	return m
}

func (m *Service) Generate(w io.Writer) error {
	return m.t.Execute(w, m)
}
