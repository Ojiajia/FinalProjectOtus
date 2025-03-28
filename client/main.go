package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(
		"http://localhost:8080/get_data",
	)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	/*
		// transfer text

		s := string(body)
		fmt.Println(s)
	*/

	os.WriteFile("the-best.jpg", body, 0666)
	fmt.Println("image saved")
}
