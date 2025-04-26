package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`// these are the aliases for all these fileds
	Price    int
	Platform string `json:"website"`

	Password string   `json:"-"`
	Tags     []string `json: "tags,omitempty"`
}


// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string 

// 	Password string  
// 	Tags     []string 
// }

func main() {
	fmt.Println("Welcome to JSON Video")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in ", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in ", "bcd123", []string{"full-satck", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in ", "hit123", nil},
	}

	//package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, " ", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
                "coursename": "MERN Bootcamp",
                "Price": 199,
                "website": "LearnCodeOnline.in ",
                "Tags": [
                        "full-satck",
                        "js"
                ]
        }`)

	var lcoCourse course 
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)


	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T \n",k, v , v)
	}


}
