package mapzenjs

import (
	"bufio"
	"bytes"
	"golang.org/x/net/html"
	"io"
	_ "log"
	"net/http"
	"net/http/httptest"
	"strconv"
)

type APIKeyHandler struct {
	handler http.Handler
	api_key string
}

func MapzenAPIKeyHandler(handler http.Handler, api_key string) (http.Handler, error) {

	h := APIKeyHandler{
		handler: handler,
		api_key: api_key,
	}

	return h, nil
}

func (h APIKeyHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
     
	rec := httptest.NewRecorder()
	h.handler.ServeHTTP(rec, req)

	for k, v := range rec.Header() {
		rsp.Header()[k] = v
	}

	rsp.Header().Set("X-We-Modified-This", "Yup")
	rsp.WriteHeader(200)

	body := rec.Body.Bytes()
	reader := bytes.NewReader(body)
	doc, err := html.Parse(reader)

	if err != nil {
		http.Error(rsp, err.Error(), http.StatusInternalServerError)
		return
	}

	var f func(node *html.Node, writer io.Writer)

	f = func(n *html.Node, w io.Writer) {
		
		if n.Type == html.ElementNode && n.Data == "body" {

			api_key_ns := ""
			api_key_key := "data-mapzen-api-key"
			api_key_value := h.api_key

			api_key_attr := html.Attribute{api_key_ns, api_key_key, api_key_value}
			n.Attr = append(n.Attr, api_key_attr)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, w)
		}
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	f(doc, wr)

	err = html.Render(wr, doc)

	if err != nil {
		http.Error(rsp, err.Error(), http.StatusInternalServerError)
		return
	}

	wr.Flush()

	data := buf.Bytes()	
	clen := len(data)

	req.Header.Set("Content-Length", strconv.Itoa(clen))
	rsp.Write(data)
}
