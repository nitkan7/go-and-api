package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Creating a struct which is similiar to the json response
type Response struct{
	Anime string `json:"anime"`
	Character string `json:"character"`
	Quote string `json:"quote"`
}
// 	text:=` {
//         "anime": "Naruto",
//         "character": "Itachi Uchiha",
//         "quote": "Those who forgive themselves, and are able to accept their true nature... They are the strong ones!"
//     }`


func main() {

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
 fmt.Println(responseObject.Quote)
 
}






