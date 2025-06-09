package main

// KismetResponse represents the structure of the JSON response from the Kismet Wireless API.
type KismetResponse struct {
	SSID       string `json:"kismet.device.base.name"`
	Encryption string `json:"kismet.device.base.crypt"`
	MacAddr    string `json:"kismet.device.base.macaddr"`
}
