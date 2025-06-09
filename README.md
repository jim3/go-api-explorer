## Go API Explorer

#### Description
An HTTP client that interacts with various API's via HTML forms, using the Go programming language. 
It is designed to be a simple and efficient way to explore APIs using HTML forms and retrieving responses.

---

#### Types/Methods/Functions/Used
- BaseURL = `https://pkg.go.dev`

- Package `url` parses URLs and implements query escaping:
  - `url.Parse` parses a raw URL string and returns a `*url.URL` struct. The URL struct contains fields such as `Scheme`, `Host`, `Path`, and `RawQuery`. [https://pkg.go.dev/net/url#Parse](https://pkg.go.dev/net/url#Parse)
  - `url.Values` is a type that represents a collection of URL query parameters. It is a map of string keys to string slices. [https://pkg.go.dev/net/url#Values](https://pkg.go.dev/net/url#Values)
  - `url.Values.Get` retrieves the first value associated with the given key. For example, if you have a URL with query parameters like `?domain=example.com&domain=test.com`, calling `Get("domain")` will
        return `"example.com"`. [https://pkg.go.dev/net/url#Values.Get](https://pkg.go.dev/net/url#Values.Get)

- Package `net/http` provides HTTP client and server implementations:
  - `http.NewRequest` creates a new HTTP request with the given method, URL, and optional body. The method is a string that specifies the HTTP method to use, such as "GET", "POST", etc. [https://pkg.go.dev/net/http#NewRequest](https://pkg.go.dev/net/http#NewRequest)
  - `http.Request` is a struct that represents an HTTP request. It contains fields such as `Method`, `URL`, `Header`, and `Body`. The `Method` field is a string that specifies the HTTP method 
        used for the request, such as "GET", "POST", etc. [https://pkg.go.dev/net/http#Request](https://pkg.go.dev/net/http#Request)
  - `/net/http#Request.Method`
  - `/net/http#Request`
  - `/net/http#Request.FormValue`

---

#### TODO
