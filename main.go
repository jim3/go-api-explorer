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

// ----------------------------------------------

func kismetlookup() {
	// Endpoint returns all Wi-Fi access points
	// GET /devices/views/phydot11_accesspoints/devices.json
	url := "http://jim3:earth500@localhost:2501/devices/views/phydot11_accesspoints/devices.json"

	// Make the call
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTTP Status is:", res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Note, it's crucial that you create a slice here, otherwise the json has nowhere to "go" :)
	var resp []KismetResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	// ----------------------------------------------
	// Outer loop: Iterate through each KismetResponse object in the slice
	for i, deviceRecord := range resp {
		fmt.Printf("--- Processing Device Record #%d ---\n", i)
		// You can prettyPrint individual deviceRecord if your prettyPrint function accepts it
		// prettyPrint(deviceRecord) // If prettyPrint is defined for single KismetResponse

		fmt.Printf("Device Name: %s\n", deviceRecord.DeviceBaseName)
		fmt.Printf("Device MAC: %s\n", deviceRecord.MacAddr)
		fmt.Printf("Device Key: %s\n", deviceRecord.DeviceKey)

		// Now, check if this specific deviceRecord has Dot11Device data
		// and if it has an AssociatedClientMap
		if deviceRecord.Dot11Device != nil && deviceRecord.Dot11Device.AssociatedClientMap != nil {
			fmt.Println("  Associated Clients:")
			// Inner loop: Iterate through the AssociatedClientMap of the *current* deviceRecord
			for clientMAC, clientDeviceKey := range deviceRecord.Dot11Device.AssociatedClientMap {
				fmt.Println("    MAC:", clientMAC, "Device Key:", clientDeviceKey)
				fmt.Println("    DEVICE_KEY →", clientDeviceKey) // This is the device key of the client
			}
		} else {
			fmt.Println("  No associated client map found for this device (might not be an AP or no clients seen).")
		}
		fmt.Println("----------------------------------------------")
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("Kismet lookup completed successfully.")
	fmt.Println("----------------------------------------------")

}

// ----------------------------------------------

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

// ----------------------------------------------

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

// Helper Function
func prettyPrint(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

// ----------------------------------------------
