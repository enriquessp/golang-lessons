package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
)

func isValid(e error) {
	if e != nil {
		fmt.Println("Erro ao invocar API", e)
		os.Exit(1)
	}
}

type apiRequest struct {
	url string
	path string
	source string
	routes string
	pacageGroup string
}

func (a apiRequest) getURL() string {
	return a.url + a.path +"ages=30&preferences=preferences%3DpersistLog%2Clanguage%3Apt_BR%2Ccurrency%3ABRL&source=" + a.source + "&routes=" + a.routes + "&packageGroup=" + a.pacageGroup
}

func main() {

	client := &http.Client{}

	reqAPI := apiRequest {
//		url: "http://192.168.99.100:31020",
		url: "https://gwa-cvc-hom.reservafacil.tur.br",
		path: "/gwaereo/v0/flights?",
		source: "AMD",
		routes: "SAO%2CMIA%2C2018-09-09%2BMIA%2CSAO%2C2018-09-19",
		pacageGroup: "VHI",
	}
	
	req, err := http.NewRequest("GET", reqAPI.getURL(), nil )	

	isValid(err)

	req.Header.Add("Gtw-Username", "scvc")
	req.Header.Add("Gtw-Password", "scvc")
	req.Header.Add("Gtw-Branch-Id", "1000")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	isValid(err)

	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	isValid(err)
	
	flightsResponse := flightsResponse{}
	isValid( json.Unmarshal(body, &flightsResponse) )

	flightsResponse.printRecommendations()
	flightsResponse.printMeta()
}
