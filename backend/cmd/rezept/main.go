package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sour-dough/rezept-backend/api"
	"github.com/sour-dough/rezept-backend/db"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Specify config path as first parameter")
		return
	}

	log.Println("Reading config")
	config, err := readConfig(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Opening database")
	database, err := db.InitDB(config.DBPath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	apiRunner := api.Init(database, api.APIConfig{
		Prefix:           config.APIPrefix,
		MaxUploadSize:    5 * 1024 * 1024,
		UploadPath:       config.ImagePath,
		ImageURL:         config.ImageURL,
		MaxThumbnailSize: 200,
		SignupAllowed:    config.SignupAllowed,
	})

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(config.ImagePath))))
	http.Handle("/", logMiddleware(apiRunner))

	go func() {
		err = http.ListenAndServe(config.Address, nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()
	log.Printf("Listening on address %s\n", config.Address)

	// block
	var c chan struct{}
	<-c
}
