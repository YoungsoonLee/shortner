package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/YoungsoonLee/hexagonial/shortner"
	"github.com/vmihailenco/msgpack"
)

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func main() {
	address := fmt.Sprintf("http://localhost%s", httpPort())
	redirect := shortner.Redirect{}
	redirect.URL = "https://github.com/YoungsoonLee"

	body, err := msgpack.Marshal(&redirect)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(address, "application/x-msgpack", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	msgpack.Unmarshal(body, &redirect)
	log.Printf("%v\n", redirect)
}
