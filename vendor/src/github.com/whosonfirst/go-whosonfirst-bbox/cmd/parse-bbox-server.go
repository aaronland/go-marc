package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/whosonfirst/go-whosonfirst-bbox/parser"
	"log"
	"net/http"
	"os"
)

type Response struct {
	MinX float64 `json:"min_x"`
	MinY float64 `json:"min_y"`
	MaxX float64 `json:"max_x"`
	MaxY float64 `json:"max_y"`
}

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")

	flag.Parse()

	handler := func(rsp http.ResponseWriter, req *http.Request) {

		query := req.URL.Query()

		bbox := query.Get("bbox")
		scheme := query.Get("scheme")
		order := query.Get("order")

		if bbox == "" {
			http.Error(rsp, "Missing bbox parameter", http.StatusBadRequest)
			return
		}

		p, err := parser.NewParser()

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		if scheme != "" {
			p.Scheme = scheme
		}

		if order != "" {
			p.Order = order
		}

		bb, err := p.Parse(bbox)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		r := Response{
			MinX: bb.MinX(),
			MinY: bb.MinY(),
			MaxX: bb.MaxX(),
			MaxY: bb.MaxY(),
		}

		body, err := json.Marshal(r)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Access-Control-Allow-Origin", "*")
		rsp.Header().Set("Content-Type", "application/json")

		rsp.Write(body)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	err := gracehttp.Serve(&http.Server{Addr: endpoint, Handler: mux})

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
