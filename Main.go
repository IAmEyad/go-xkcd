package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Comic struct {

	Title string
	Alt string
	Img string
}


type DiscWebhook struct {
    Content string `json:"content"`
    Embeds  []M    `json:"embeds"`
}

type M map[string]interface{}

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
	//fmt.Printf("%s",jsonData)

	err = json.Unmarshal(jsonData, &comic)
	//fmt.Println(comic)
	post_comic(comic)

	
}

func post_comic(comic Comic) {
    
	fmt.Println(comic)
    embeds := create_embed(comic)
   // embeds = array of map[string]string but also sometimes contains map[string]map[string]string 
    
   discordData := DiscWebhook{
        Content: "Daily XKCD", 
        Embeds: embeds,
    }
    //var discordValues map[string]interface{}

    fmt.Println(discordData)

    enc := json.NewEncoder(os.Stdout)

    enc.Encode(discordData)
    jsonData, err := json.MarshalIndent(discordData, "", "    ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(jsonData))
    resp, err := http.Post("",
        "application/json",
        bytes.NewBuffer(jsonData))
    if err != nil {
        panic(err)
    }
    fmt.Println("Response status:", resp.Status)

}

func create_embed(comic Comic) []M {

   
    var mapSlice []M

    imageUrl := map[string]string{"url": comic.Img}
    image := M{"title": comic.Title, "image": imageUrl}
    //fmt.Println(image)
   
    mapSlice = append(mapSlice, image)
    testJson, _ := json.MarshalIndent(mapSlice, "","    ")

    fmt.Println(string(testJson))

    return(mapSlice)


}
