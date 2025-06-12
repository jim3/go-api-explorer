package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func prettyPrint(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

func upload2parser(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", header.Header)
	var resp []KismetResponse
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&resp); err != nil {
		fmt.Println(err)
	}

	for _, v := range resp {
		fmt.Println(v.SSID)
		fmt.Println(v.Encryption)
		fmt.Println(v.MacAddr)
	}

	fmt.Println(resp)
}

func lookup(w http.ResponseWriter, r *http.Request) {
	APIKEY := os.Getenv("APIKEY")
	if APIKEY == "" {
		fmt.Println("Error: MY_API_KEY environment variable not set.")
		return
	}
	fmt.Printf("Successfully retrieved API Key from environment: %s\n", APIKEY)

	ipAddr := r.URL.Query().Get("ip")
	URL := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ipAddr, APIKEY)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	var resp Response
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resp); err != nil {
		log.Fatal(err)
	}

	res.Body.Close()
	prettyPrint(resp)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", upload2parser)
	mux.HandleFunc("/lookup", lookup)
	mux.Handle("/", http.FileServer(http.Dir("/home/<Path-To-HTML-File>")))

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
