package http

import (
	"fmt"
	"github.com/aaronland/go-http-rewrite"
	"github.com/aaronland/go-marc/static"
	"io/fs"
	gohttp "net/http"
	"path/filepath"
	"strings"
)

// StaticFileSystem returns a `net/http.FileSystem` instance for the static asset filesystem in `go-marc/static`.
func StaticFileSystem() (gohttp.FileSystem, error) {
	http_fs := gohttp.FS(static.FS)
	return http_fs, nil
}

// StaticAssetsHandler returns a `net/http.Handle` instance for serving files in StaticFileSystem.
func StaticAssetsHandler() (gohttp.Handler, error) {

	http_fs, err := StaticFileSystem()

	if err != nil {
		return nil, err
	}

	return gohttp.FileServer(http_fs), nil
}

// StaticAssetsHandlerWithPrefix returns a `net/http.Handle` instance for serving files in StaticFileSystem from prefix.
func StaticAssetsHandlerWithPrefix(prefix string) (gohttp.Handler, error) {

	fs_handler, err := StaticAssetsHandler()

	if err != nil {
		return nil, err
	}

	prefix = strings.TrimRight(prefix, "/")

	if prefix == "" {
		return fs_handler, nil
	}

	rewrite_func := func(req *gohttp.Request) (*gohttp.Request, error) {
		req.URL.Path = strings.Replace(req.URL.Path, prefix, "", 1)
		return req, nil
	}

	rewrite_handler := rewrite.RewriteRequestHandler(fs_handler, rewrite_func)
	return rewrite_handler, nil
}

func AppendStaticAssetHandlers(mux *gohttp.ServeMux) error {
	return AppendStaticAssetHandlersWithPrefix(mux, "")
}

func AppendStaticAssetHandlersWithPrefix(mux *gohttp.ServeMux, prefix string) error {

	asset_handler, err := StaticAssetsHandlerWithPrefix(prefix)

	if err != nil {
		return nil
	}

	walk_func := func(path string, info fs.DirEntry, err error) error {

		if path == "." {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if prefix != "" {
			path = appendPrefix(prefix, path)
		}

		if !strings.HasPrefix(path, "/") {
			path = fmt.Sprintf("/%s", path)
		}

		// log.Println("APPEND", path)

		mux.Handle(path, asset_handler)
		return nil
	}

	return fs.WalkDir(static.FS, ".", walk_func)
}

func appendPrefix(prefix string, path string) string {

	prefix = strings.TrimRight(prefix, "/")

	if prefix != "" {
		path = strings.TrimLeft(path, "/")
		path = filepath.Join(prefix, path)
	}

	return path
}
