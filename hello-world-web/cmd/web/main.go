package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ozan1338/study-go/pkg/config"
	"github.com/ozan1338/study-go/pkg/handlers"
	"github.com/ozan1338/study-go/pkg/render"
)

//Create Custom Error
//errors.New("whatever")

const PORT = ":8080"

var app config.AppConfig

var session *scs.SessionManager

//main is the main application func
func main() {
	

	//change this to true when in production
	app.InProdution = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProdution

	app.Session = session

	tc,err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("can't craete template cache")
	}

	app.TemplateChache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/About", handlers.Repo.About)

	_,_ = fmt.Println("Starting App on port", PORT)
	// _ = http.ListenAndServe(PORT, nil)

	serving := &http.Server {
		Addr: PORT,
		Handler: routes(&app),
	}

	err = serving.ListenAndServe()
	log.Fatal(err)
}