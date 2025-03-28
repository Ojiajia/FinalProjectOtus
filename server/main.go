package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	ip := flag.String("ip", "127.0.0.1", "IP address")
	port := flag.String("port", "8080", "Port number")
	flag.Parse()

	println("run")

	http.HandleFunc("/get_data", setContent)

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

	data, err := os.ReadFile("mig-31.jpg")
	if err != nil {
		fmt.Println("error error error")
	}
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Write(data)

}
