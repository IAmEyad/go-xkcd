package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	//"os"
)

type Comic struct {

	Title string
	Alt string
	Img string
}


type DiscWebhook struct {
    Content string `json:"content"`
    Embeds []Embed  `json:"embeds"`
}

type Embed struct {

    Title string `json:"title"`
    Image Image `json:"image"`
}

type Image struct {
    URL string `json:"url"`
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
	//fmt.Printf("%s",jsonData)

	err = json.Unmarshal(jsonData, &comic)
	//fmt.Println(comic)
	post_comic(comic)

	
}

func post_comic(comic Comic) {
    
	fmt.Println(comic)
    
    image := Image{
        URL: comic.Img,
    
    }


    // Create a slice embed because discord requires an array of embed objects
    embed := []Embed{
        {
        Title: comic.Title,
        Image: image,
    },
}

    // embeds = array of map[string]string but also sometimes contains map[string]map[string]string 
    
   discordData := DiscWebhook{
        Content: "Daily XKCD", 
        Embeds: embed,
    }
    //var discordValues map[string]interface{}

    fmt.Println(discordData)

    //enc := json.NewEncoder(os.Stdout)

    //enc.Encode(discordData)
    jsonData, err := json.MarshalIndent(discordData, "", "    ")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(jsonData))
    resp, err := http.Post("webhook_here",
        bytes.NewBuffer(jsonData))
    if err != nil {
        panic(err)
    }
    fmt.Println("Response status:", resp.Status)

}
