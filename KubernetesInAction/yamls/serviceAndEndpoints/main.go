package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"Hello World!!!")
	})

	fmt.Println("Listing on port 8093")
	if err := http.ListenAndServe(":8093", nil); err != nil {
		log.Fatal("Error while serving the application")
	}
}
