package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type Response struct{
	Anime string `json:"anime"`
	Character string `json:"character"`
	Quote string `json:"quote"`
}

func main() {

	text:=` {
        "anime": "Naruto",
        "character": "Itachi Uchiha",
        "quote": "Those who forgive themselves, and are able to accept their true nature... They are the strong ones!"
    }`

	textBytes:=[]byte(text)

	people1:=Response{}
	err:=json.Unmarshal(textBytes,&people1)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Quote)
 fmt.Println("Calling API...")
client := &http.Client{}
 req, err := http.NewRequest("GET", "https://animechan.vercel.app/api/random", nil)
 if err != nil {
  fmt.Print(err.Error())
 }
 req.Header.Add("Accept", "application/json")
 req.Header.Add("Content-Type", "application/json")
 resp, err := client.Do(req)
 if err != nil {
  fmt.Print(err.Error())
 }
defer resp.Body.Close()
 bodyBytes, err := ioutil.ReadAll(resp.Body)
 if err != nil {
  fmt.Print(err.Error())
 }
var responseObject Response
 json.Unmarshal(bodyBytes, &responseObject)
 fmt.Printf("API Response as struct %+v\n", responseObject.Quote)
 
}






