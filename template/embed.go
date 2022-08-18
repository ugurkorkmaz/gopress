package template

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var dist embed.FS

func Dist() http.FileSystem {
	distSub, err := fs.Sub(dist, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(distSub)
}
