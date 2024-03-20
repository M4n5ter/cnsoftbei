package middleware

import (
	"io"
	"io/fs"
	"text/template"

	"github.com/m4n5ter/cnsoftbei/cmd/core"
)

type Middleware struct {
	t    *template.Template
	Name string
}

func New(fsys fs.FS, name string) *Middleware {
	m := &Middleware{}
	t := template.New("middleware.tmpl")
	t.Funcs(core.Funcs)
	m.t = template.Must(t.ParseFS(fsys, "templates/middleware.tmpl"))

	m.Name = name
	return m
}

func (m *Middleware) Generate(w io.Writer) error {
	return m.t.Execute(w, m)
}
