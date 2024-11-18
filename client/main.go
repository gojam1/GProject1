package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
	"github.com/gojam1/GoDemo/server"
	"github.com/gojam1/GoDemo/client"
)

func main() {
	response, err := http.Get("http://localhost")
	if err != nil{
		return
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
		}
	
}
func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
func ReadFile(filename string) ([]byte, error){
	return os.ReadFile((filename))
}