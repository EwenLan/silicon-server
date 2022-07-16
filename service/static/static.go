package static

import (
	"net/http"
	"path"
	"strings"

	"github.com/EwenLan/silicon-server/configmanager"
	"github.com/EwenLan/silicon-server/slog"
)

var rootDirectory string
var redirectSubpaths []string

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

func SetRedirectSubpaths(subpaths []string) {
	slog.Debugf("set redirect sub-paths %+v", subpaths)
	redirectSubpaths = subpaths
}

func checkRedirectSubpath(uri string) bool {
	parts := strings.Split(uri, "/")
	if len(parts) <= subpathIndex {
		slog.Debugf("path len = %d", len(parts))
		return false
	}
	for i := range redirectSubpaths {
		if parts[subpathIndex] == redirectSubpaths[i] {
			slog.Debugf("found sub-path = %s will be redirected", parts[subpathIndex])
			return true
		}
	}
	slog.Debugf("sub-path = %s will not be redirected", parts[subpathIndex])
	return false
}

// ServeStatic
func ServeStatic(w http.ResponseWriter, r *http.Request) {
	if checkRedirectSubpath(r.RequestURI) {
		slog.Debugf("path = %s in redirect sub-path list", r.RequestURI)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}
	slog.Debugf("static file %s %s %s from %s", r.Method, r.URL, r.Proto, r.RemoteAddr)
	loadFile := getFilePath(r.RequestURI)
	http.ServeFile(w, r, loadFile)
}
