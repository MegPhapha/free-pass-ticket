package main

import ( 
	"fmt"
	"net/http"
	"html/template"
	"time"

)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved: " + r.URL.Path)
	t, err := template.ParseFiles("templates/index.html")
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s\n", err)
		return
	}
	

	//get today's date
date := time.Now()
expiration := date.AddDate(0, 0, 30)
	t.Execute(w, expiration.String())

}


func main() {
	http.HandleFunc("/", handler)
	http.Handle("/images/", http.StripPrefix ("/images/" ,(http.FileServer (http.Dir("images")))))
	http.Handle("/css/", http.StripPrefix ("/css/" ,(http.FileServer (http.Dir("css")))))
	http.Handle("/manuals/", http.StripPrefix ("/manuals/" ,(http.FileServer (http.Dir("manuals")))))
	fmt.Println("Listening on port 3002...")
	http.ListenAndServe(":3002", nil)

}