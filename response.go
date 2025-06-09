package main

// Response represents the structure of the HTTP response containing information
// about an IP address, including its country code, organization, associated tags,
// domains, and additional data items.
type Response struct {
	IP           int        `json:"ip"`
	Country      string     `json:"country_code"`
	Organization string     `json:"org"`
	Tags         []string   `json:"tags"`
	Domains      []string   `json:"domains"`
	Data         []DataItem `json:"data"`
}

type DataItem struct {
	IP        int64    `json:"ip"`
	Port      int      `json:"port"`
	Transport string   `json:"transport"`
	Hash      int64    `json:"hash"`
	Tags      []string `json:"tags"`
	Cloud     Cloud    `json:"cloud"`
	Location  Location `json:"location"`
}

type Cloud struct {
	Region   string  `json:"region"`
	Service  *string `json:"service"`
	Provider string  `json:"provider"`
}

type Location struct {
	City        string  `json:"city"`
	RegionCode  string  `json:"region_code"`
	AreaCode    *string `json:"area_code"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
}
