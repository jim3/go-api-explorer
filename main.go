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

// Uses an HTML <form> element to upload form data to the server
func upload2parser(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", header.Header)

	// Create a nil slice
	var resp []KismetResponse

	// Create a new decoder
	decoder := json.NewDecoder(file)

	// Decode the response body into the resp slice
	if err := decoder.Decode(&resp); err != nil {
		fmt.Println(err)
	}

	// If no error occurs, we can use the slice of items in our program
	for _, v := range resp {
		fmt.Println(v.SSID)
		fmt.Println(v.Encryption)
		fmt.Println(v.MacAddr)
	}

	fmt.Println(resp) // [{CatheadBiscuits WPA2 WPA2-PSK AES-CCMP 38:3F:B3:84:63:F8}]
}

// Uses an HTML <form> element to get the IP address and then queries the Shodan AP
func lookup(w http.ResponseWriter, r *http.Request) {
	APIKEY := os.Getenv("APIKEY")
	if APIKEY == "" {
		fmt.Println("Error: MY_API_KEY environment variable not set.")
		return
	}
	fmt.Printf("Successfully retrieved API Key from environment: %s\n", APIKEY)

	// Get the ip address value from the HTML form via url package
	ipAddr := r.URL.Query().Get("ip")

	// Use Sprintf to return a formatted URL string
	URL := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ipAddr, APIKEY)

	// Make a GET request to the `/shodan/host/{ip}` endpoint
	// The endpoint returns all services that have been found on the given host IP.
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	// Create a nil slice of resp
	var resp Response

	decoder := json.NewDecoder(res.Body)

	// Decode the response body into the resp slice
	if err := decoder.Decode(&resp); err != nil {
		log.Fatal(err)
	}

	res.Body.Close()
	prettyPrint(resp)

}

func main() {
	mux := http.NewServeMux()
	// Register the upload2parser function to handle requests to "/upload"
	mux.HandleFunc("/upload", upload2parser)

	// Register the lookup function to handle the request1 from ""
	mux.HandleFunc("/lookup", lookup)

	// Serve the HTML forms
	mux.Handle("/", http.FileServer(http.Dir("/home/<Path-To-HTML-File>")))

	// Create a new http.Server struct
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.ListenAndServe()
}
