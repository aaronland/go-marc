package http

import (
	"fmt"
	"html/template"
	gohttp "net/http"
)

type HTMLVars struct {
}

// MARC034Handler returns a `net/http.Handler` instance for converting MARC 034 fields.
func MARC034Handler(t *template.Template) (gohttp.Handler, error) {

	t = t.Lookup("marc_034")

	if t == nil {
		return nil, fmt.Errorf("Can't find marc_034 template")
	}

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		vars := HTMLVars{}
		err := t.Execute(rsp, vars)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
