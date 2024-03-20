package main

import (
	"embed"
	_ "embed"
)

//go:embed templates/*
var f embed.FS

func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
