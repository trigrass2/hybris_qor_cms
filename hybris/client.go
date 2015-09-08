package hybris

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/theplant/hybris_qor_cms/db"
)

const (
	BaseUlr = "http://localhost:9001/ws410/rest/"
)

func GetProduct(id string) (product *db.Product) {

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", BaseUlr+"catalogs/apparelProductCatalog/catalogversions/Staged/products/"+id, nil)

	req.Header.Add("Authorization", `Basic YWRtaW46bmltZGE=`)
	req.Header.Add("Accept", `application/json`)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}
	if resp.StatusCode >= 300 {
		err = errors.New("request url: " + req.URL.String() + "status: " + resp.Status)
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	if e := dec.Decode(&product); e != nil {
		fmt.Println("Decode Failure : ", err)
	}
	return
}

func GetPrice(uri string) (price *db.Price) {

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", uri, nil)

	req.Header.Add("Authorization", `Basic YWRtaW46bmltZGE=`)
	req.Header.Add("Accept", `application/json`)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}
	if resp.StatusCode >= 300 {
		err = errors.New("request url: " + req.URL.String() + "status: " + resp.Status)
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	if e := dec.Decode(&price); e != nil {
		fmt.Println("Decode Failure : ", err)
	}
	return
}
