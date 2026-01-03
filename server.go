package main

import (
	//"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

type PageData struct {
	HostName string
	ClientIP string
}

func main() {
	tmpl := template.Must(template.ParseFiles("static/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		data := PageData{
			HostName: hostname,
			ClientIP: ip,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
		}

		// fmt.Fprintf(w, "Testing Go http server: OK")
	})

	log.Println("Server started on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))

	/*fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(":8090", nil))
	*/
}
