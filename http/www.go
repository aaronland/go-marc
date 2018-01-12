package http

import (
	"github.com/thisisaaronland/go-marc/assets/html"
	"html/template"
	gohttp "net/http"
)

type HTMLVars struct {
}

func WWWHandler() (gohttp.Handler, error) {

	tpl, err := html.Asset("templates/html/index.html")

	if err != nil {
		return nil, err
	}

	t := template.New("index")

	t, err = t.Parse(string(tpl))

	if err != nil {
		return nil, err
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
