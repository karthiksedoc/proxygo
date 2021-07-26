package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpclient(b *http.Request) ([]byte, error) {
	//	fmt.Println(b)
	//	fmt.Println(b.URL)
	//	fmt.Println(b.RequestURI)
	//protocol := "http://"
	request_protocol := b.URL.Scheme
	fmt.Println(request_protocol)
	request_host := b.URL.Host
	fmt.Println(request_host)
	request_path := b.URL.Path
	fmt.Println(request_path)
	complete_URL := request_protocol + "://" + request_host + request_path
	fmt.Println(complete_URL)
	response, err := http.Get(complete_URL)
	//	fmt.Println(response)
	if err != nil {

		var my_response []byte
		return my_response, err
	}
	defer response.Body.Close()
	my_response, _ := ioutil.ReadAll(response.Body)

	//	fmt.Printf("httpclient func variable my_response = %T\n", my_response)
	return my_response, nil
}

func domainsearch(b string) string {
	domains := [5]string{"neverssl.com", "www.google.com", "google.com", "facebook.com", "yahoo.com"}

	for _, name := range domains {
		if name == b {
			dom_status := "Denied"
			return dom_status
		}

	}

	dom_status := "Granted"
	return dom_status
	//	webresponse, err := httpclient(b)
	//	if err != nil {
	//		web_response := []byte(nil)
	//		return web_response, "Upstream Error"
	//	}
	//    return webresponse, ""
}

func proxyserver(w http.ResponseWriter, r *http.Request) {
	domain_name := r.Host
	status := domainsearch(domain_name)

	if status == "Denied" {
		http.Error(w, "Blocked", 403)
		return
	}

	webresponse, err := httpclient(r)
	if err != nil {
		//		web_response := []byte(nil)
		w.Write([]byte("Upstream Error"))

	}
	w.Write([]byte(webresponse))
}

func main() {
	mux := http.NewServeMux()
	rik := http.HandlerFunc(proxyserver)
	mux.Handle("/", rik)
	http.ListenAndServe(":80", mux)
}
