package main

import ( 
	"fmt"
	"net/http"
	"encoding/json"
	"io"

)

type Comic struct {

	Title string
	Alt string
	Img string
}

func main() {


	var comic Comic

	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",jsonData)

	err = json.Unmarshal(jsonData, &comic)
	//fmt.Println(comic)
	post_comic(comic)

	
}

func post_comic(comic Comic) {

	fmt.Println(comic)
	

}