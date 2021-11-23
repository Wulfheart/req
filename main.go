package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// file, _ := os.Open("./req.req")
	// defer file.Close()
	//
	// f2, _ := os.ReadFile("./req.req")
	//
	// fmt.Println(string(f2))
	// req, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(f2)))
	//
	//
	// req.RequestURI = ""
	// if err != nil {
	// 	panic(err)
	// }
	// req.Close = false
	//
	//
	//
	// client := http.DefaultClient
	//
	// res, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer res.Body.Close()
	// s,_ := io.ReadAll(res.Body)
	//
	// fmt.Println(string(s))
	//
	// _ = req
	// _ = res

	s, _ := os.ReadFile("./test1.env")
	var val, _ = godotenv.Unmarshal(string(s))
	val["Val2"] = "zeze"
	val["000"] = "GJHG"
	res, _ := godotenv.Marshal(val)
	fmt.Println(res)
}
