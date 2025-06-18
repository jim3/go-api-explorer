package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Upload an exported Kismet file to parse
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
		fmt.Println(v.DeviceBaseName)
		fmt.Println(v.Encryption)
		fmt.Println(v.MacAddr)
	}

	fmt.Println(resp)
}

// Returns all Wi-Fi access points and their associated clients
func kismetlookup() {
	url := "http://<username>:<password>@localhost:2501/devices/views/phydot11_accesspoints/devices.json"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTTP Status is:", res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resp []KismetResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	for i, v := range resp {
		fmt.Printf("--- Processing Device Record #%d ---\n", i)
		fmt.Println("################################")
		fmt.Printf("Device Name: %s\n", v.DeviceBaseName)
		fmt.Printf("SSID MAC: %s\n", v.MacAddr)
		fmt.Printf("Device Key: %s\n", v.DeviceKey)

		// Check if deviceRecord has an AssociatedClientMap
		if v.Dot11Device != nil && v.Dot11Device.AssociatedClientMap != nil {
			fmt.Println("  Associated Clients:")
			// Inner loop
			for key, value := range v.Dot11Device.AssociatedClientMap {
				fmt.Println("    MAC:", key)
				fmt.Println("    DEVICE_KEY:", value) // This is the device key of the client
			}
		} else {
			fmt.Println("  No associated client map found for this device (might not be an AP or no clients seen).")
		}
		fmt.Println("----------------------------------------------")
	}
	fmt.Println("Kismet lookup completed successfully.")

}

// Returns all services that have been found on the given host IP.
func lookup(w http.ResponseWriter, r *http.Request) {
	APIKEY := os.Getenv("APIKEY")
	if APIKEY == "" {
		fmt.Println("Error: MY_API_KEY environment variable not set.")
		return
	}
	fmt.Printf("Successfully retrieved API Key from environment: %s\n", APIKEY)

	// Form the URL
	ipAddr := r.URL.Query().Get("ip")
	URL := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ipAddr, APIKEY)

	// Make the call
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
	kismetlookup()
	mux.Handle("/", http.FileServer(http.Dir("/home/jim3/code/github.com/jim3/go-api-explorer/")))

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.ListenAndServe()
}

func prettyPrint(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
