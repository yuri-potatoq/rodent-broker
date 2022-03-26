package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

func echo() (string, HttpHandler) {
	return "/", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(GetKafkaConfig(KafkaEnv))
		if err != nil {
			log.Fatal(err)
		}

		writed, err := w.Write(data)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Was writed: %d bytes", writed)
	}
}

func producerRoute() (string, HttpHandler) {	
	return "/produce/{topic}", func(w http.ResponseWriter, r *http.Request) {
		var data map[string]string
		topic := mux.Vars(r)["topic"]
		r.ParseForm()

		for k, _ := range r.Form {
			data[k] = r.FormValue(k)
		}

		producer, _ := DefaultProducer(topic)
		if err := producer.ProduceAtTopic(topic, data); err != nil {
			log.Fatal(err)
		}
		
		w.WriteHeader(http.StatusOK)
	}
}

func InitServer() {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc(echo())
	v1.HandleFunc(producerRoute())
		
	srv := &http.Server{
		Handler:      r,
		Addr:         ServerAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func main() {
	Execute()
}
