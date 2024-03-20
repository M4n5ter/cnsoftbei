package api

import (
	"io"
	"io/fs"
	"text/template"

	"github.com/m4n5ter/cnsoftbei/cmd/core"
)

type API struct {
	t    *template.Template
	Name string
}

func New(fsys fs.FS, name string) *API {
	m := &API{}
	t := template.New("api.tmpl")
	t.Funcs(core.Funcs)
	m.t = template.Must(t.ParseFS(fsys, "templates/api.tmpl"))

	m.Name = name
	return m
}

func (m *API) Generate(w io.Writer) error {
	return m.t.Execute(w, m)
}
