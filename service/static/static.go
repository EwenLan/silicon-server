package static

import (
	"net/http"
	"path"

	"github.com/EwenLan/silicon-server/configmanager"
	"github.com/EwenLan/silicon-server/slog"
)

var rootDirectory string

func getContentTypeByFilename(filename string) string {
	suffix := path.Ext(filename)
	mime := mimeType[suffix]
	slog.Debugf("get mime type = %s by filename = %s", mime, filename)
	return mime
}

func getFilePath(uri string) string {
	if uri == "/" {
		defaultPage := configmanager.GetGlobalConfig().GetDefaultPage()
		uri = defaultPage
	}

	filepath := path.Join(rootDirectory, uri)
	slog.Debugf("get file path = %s", filepath)
	return filepath
}

// SetRootDirectory
func SetRootDirectory(directory string) {
	slog.Debugf("set root directory = %s", directory)
	rootDirectory = directory
}

// ServeStatic
func ServeStatic(w http.ResponseWriter, r *http.Request) {
	slog.Debugf("static file %s %s %s from %s", r.Method, r.URL, r.Proto, r.RemoteAddr)
	loadFile := getFilePath(r.RequestURI)
	http.ServeFile(w, r, loadFile)
}
