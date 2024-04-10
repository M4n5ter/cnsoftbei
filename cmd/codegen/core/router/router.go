package router

import (
	"html/template"
	"io"
	"io/fs"

	"github.com/m4n5ter/cnsoftbei/cmd/core"
)

type Router struct {
	t    *template.Template
	Name string
}

func New(fsys fs.FS, name string) *Router {
	m := &Router{}
	t := template.New("router.tmpl")
	t.Funcs(core.Funcs)
	m.t = template.Must(t.ParseFS(fsys, "templates/router.tmpl"))

	m.Name = name
	return m
}

func (m *Router) Generate(w io.Writer) error {
	return m.t.Execute(w, m)
}
