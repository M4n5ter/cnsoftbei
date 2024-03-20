package core

import (
	"html/template"
	"strings"
	"unicode"
)

var Funcs = template.FuncMap{
	"toCamelCase": toCamelCase,
	"toLower":     strings.ToLower,
}

func toCamelCase(s string) string {
	var camelCase string
	upperNext := true

	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			camelCase += string(r)
			upperNext = false
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			if upperNext {
				camelCase += string(unicode.ToUpper(r))
			} else {
				camelCase += string(unicode.ToLower(r))
			}
			upperNext = false
		default: // 这里包括了空格、下划线等情况
			upperNext = true
		}
	}

	return camelCase
}
