package server

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/SiberianPanda/selfcard_service/src/pkg/mail"
)

func mailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var RawAnswer = make(map[string]bool, 1)
		switch err := mail.Send(r); err {
		case nil:
			RawAnswer["success"] = true
		default:
			log.Print(err)
			RawAnswer["success"] = false
		}
		var rawDataInterface interface{} = RawAnswer
		data, _ := json.Marshal(rawDataInterface)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		// http.Redirect(w, r, "/", 301)
	default:
		http.Redirect(w, r, "/", 301)
		log.Print("redirected\n")
	}
}

func routerSet() {
	http.HandleFunc("/callback", mailHandler)
}

// Up starting http server
func Up(port string) {

	log.Println("starting server at :8080")
	routerSet()
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("Server failed with err -> " + err.Error())
	}
}
