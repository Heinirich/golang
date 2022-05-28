package main

import (
	"net/http"
)

func main()  {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello gos"))
	})

	http.HandleFunc("/welcomes",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello welcomes"))
	})
	
	err := http.ListenAndServe(":8081",nil)
	if err != nil {
		panic(err.Error())
	}
}