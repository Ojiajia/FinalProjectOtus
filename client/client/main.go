// + чтоб клиентов было много

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	// fmt.Println(text)
	//запрос на покажити все файлы
	resp, err := http.Get(
		//	"http://localhost:8080/get_data",
		text,
	)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return
	}

	// получить все файлы методом post

	// ввести нужный путь

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	text = strings.TrimSuffix(text, "\n")

	text = "http://localhost:8080/" + text

	// считать данные (.css .html .js .jpg)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Сервер вернул ошибку: %s\n", resp.Status)
		return
	}

	err = os.WriteFile("the-best.jpg", body, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}

}
