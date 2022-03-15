package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ozan1338/study-go/pkg/config"
	"github.com/ozan1338/study-go/pkg/models"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

//New Template creaet new template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

//Render Template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateChache
		
	}else{
		tc, _ = CreateTemplateCache()
	}
	//get the template cache from app config


	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("couldnt get template from template chace")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
	
	// parseTemplate, _ := template.ParseFiles("./templates/"+ tmpl)
	
	// err = parseTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 
	// }
}

//Create Template Chache as map
func CreateTemplateCache() (map[string]*template.Template,error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil{
		return myCache,err
	}

	for _,page := range pages {
		name := filepath.Base(page)
		//fmt.Println("page is currently", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache,err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache,err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache,err
			}
		}

		myCache[name] = ts
	}

	return myCache,nil
}