package main

// KismetResponse represents the structure of the JSON response from the Kismet Wireless API.
type KismetResponse struct {
	Dot11Device    *Dot11Device `json:"dot11.device"`
	DeviceBaseName string       `json:"kismet.device.base.name"`
	Encryption     string       `json:"kismet.device.base.crypt"`
	MacAddr        string       `json:"kismet.device.base.macaddr"`
	DeviceKey      string       `json:"kismet.device.base.key"`
}

// dot11.device.advertised_ssid_map is an ARRAY of objects.
type Dot11Device struct {
	AdvertisedSSIDMap   []AdvertisedSSID  `json:"dot11.device.advertised_ssid_map"`
	AssociatedClientMap map[string]string `json:"dot11.device.associated_client_map"`
}

// AdvertisedSSID represents a single object within the "advertised_ssid_map" array.
type AdvertisedSSID struct {
	// These are direct string fields within each advertised SSID object
	SSID      string `json:"dot11.advertisedssid.ssid"`
	SSIDCrypt string `json:"dot11.advertisedssid.crypt_string"`
}
