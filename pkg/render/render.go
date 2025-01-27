package render

import (
	"net/http"
	"text/template"
)

func RenderTemplate (w http.ResponseWriter , tmpl string) {
	parseTempl , _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTempl.Execute(w , nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}