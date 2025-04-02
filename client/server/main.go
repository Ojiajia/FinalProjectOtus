package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// staticDir := ""
	// встроенный обработчик http.FileServer,
	// для доступа к статическим файлам из
	// определенной директории по HTTP протоколу:

	//staticDir := "./"

	// чтобы все HTTP-запросы, начинающиеся с "/static/",
	// обрабатывались с  помощью http.FileServer:
	// Для инициализации обработчика http.FileServer нужно
	// использовать функцию http.FileServer() следующим образом:

	//fileServer := http.FileServer(http.Dir(staticDir))

	ip := flag.String("ip", "127.0.0.1", "IP address") //
	port := flag.String("port", "8080", "Port number")
	flag.Parse()

	println("run")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/get_data", setContent) // where setcontent is handler

	// fileServer := http.FileServer(http.Dir("./ui/static/"))

	http.ListenAndServe(*ip+":"+*port, nil)

}

/*

// transfer text

func setContent(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("error error error")
	}
	// w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Type", "text/txt")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)
	// print(string(data))
	// fmt.Println(data)

}
*/

func setContent(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("static/mig-31.jpg")
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Write(data)

}
