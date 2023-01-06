package impls

import "net/http"

func Something(r http.ResponseWriter) {}

type MyWriter struct{}

func (w MyWriter) Header() http.Header        { panic("") }
func (w MyWriter) Write([]byte) (int, error)  { panic("") }
func (w MyWriter) WriteHeader(statusCode int) { panic("") }

func Another() {
	Something(MyWriter{})
}
