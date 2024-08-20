package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"logic"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(bytes)
	tours := logic.ToursFromJson(content)

	for _, tour := range tours {
		fmt.Println(tour.Name, " ", tour.Price)
	}
}
