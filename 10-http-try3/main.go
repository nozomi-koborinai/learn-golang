package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Result struct {
	Name   string
	Result string
}

var tpl = template.Must(template.New("fortune").Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Result}}</b>」です</body></html>"))

func main() {
	result := &Result{Name: "Gopher", Result: "大吉"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("p")
		if name != "" {
			result.Name = strings.TrimSpace(name)
		}

		resp, err := http.Get("http://example.com")
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		result.Result = strings.TrimSpace(string(body))

		if err := tpl.Execute(w, result); err != nil {
			return
		}
	})

	log.Println("starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
