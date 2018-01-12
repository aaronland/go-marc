package http

import (
	gohttp "net/http"
)

func WWWHandler() (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		rsp.Write([]byte("PLEASE WRITE ME"))		
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
