package main

// KismetResponse represents the structure of the JSON response from the Kismet Wireless API.
type KismetResponse struct {
	Dot11Device    *Dot11Device `json:"dot11.device"`
	DeviceBaseName string       `json:"kismet.device.base.name"`
	Encryption     string       `json:"kismet.device.base.crypt"`
	MacAddr        string       `json:"kismet.device.base.macaddr"`
	DeviceKey      string       `json:"kismet.device.base.key"`
}

type Dot11Device struct {
	// dot11.device.advertised_ssid_map is an ARRAY of objects,
	// so it's a slice (Go's array) of another struct type.
	AdvertisedSSIDMap []AdvertisedSSID `json:"dot11.device.advertised_ssid_map"`
	// ... potentially other fields from "dot11.device.*"
	// CORRECTED:
	// dot11.device.associated_client_map is an OBJECT where keys are dynamic (MACs)
	// and values are strings (Kismet device keys).
	AssociatedClientMap map[string]string `json:"dot11.device.associated_client_map"`
}

// AdvertisedSSID represents a single object within the "advertised_ssid_map" array.
// This struct is also defined OUTSIDE of Dot11Device, then used as a type inside it.
type AdvertisedSSID struct {
	// These are direct string fields within each advertised SSID object
	SSID      string `json:"dot11.advertisedssid.ssid"`
	SSIDCrypt string `json:"dot11.advertisedssid.crypt_string"`
	// ... potentially other fields from "dot11.advertisedssid.*"
}
