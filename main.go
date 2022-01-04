package main

/*
Use golang inbuilt library
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//struct/class for json (key value pair )body - mholt.github.io/json-to-go/-
type Response struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func main() {

	// Get request
	resp, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		fmt.Println("No response from request")
	}

	//Close Data after response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	//create new object with response class
	var result Response

	//error handling if json cant unmarshal
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// Remove function
	//TODO: Fix pretty print function
	//  fmt.Println(PrettyPrint(result))

	// Loop through the data node for the FirstName
	for _, rec := range result.Data {
		fmt.Println(rec.FirstName, rec.LastName)
	}
}

// PrettyPrint to print struct in a readable way
/*func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}*/
