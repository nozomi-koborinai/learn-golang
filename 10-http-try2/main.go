package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("fortune").Parse(`<html><body>{{.Name}}さんの運勢は「<b>{{.Fortune}}</b>」です</body></html>`))
}

type fortune struct {
	Name    string
	Fortune string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("p")
		var result string = "小吉"
		if name == "Gopher" {
			result = "大吉"
		}
		f := &fortune{Name: name, Fortune: result}
		err := tpl.Execute(w, f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", nil)
}
