package shadowlink

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed build/*
var build embed.FS

func Build() http.FileSystem {
	fs, err := fs.Sub(build, "build")
	if err != nil {
		panic(err)
	}

	return http.FS(fs)
}
