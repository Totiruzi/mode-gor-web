package randler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap {

}

// RenderTemplate using html/template
func RenderTemplate(w http.ResponseWriter, temp string) {
	// _, err := RenderAvailableTemplates(w) 
	// if err != nil {
	// 	fmt.Println("Could not render template", err)
	// }

	parseTemplate, _ := template.ParseFiles("./templates/" + temp)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Problem parsing template", http.StatusBadRequest)
		return
	}
}


func RenderAvailableTemplates(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err 
	}

	for _, page := range pages{
		name := filepath.Base(page)
		fmt.Println("Page is currently", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}