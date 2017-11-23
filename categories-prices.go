package categoriesPrices

import (
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"math"
)

type CategoryRemoteApi struct {
	Paging struct {
		Total int `json:"total"`
	} `json:"paging"`
	Results []struct {
		Price      float64    `json:"price"`
	} `json:"results"`
}

func GetPrices(categoryId string)(map[string]float64){

	var prices = make(chan float64,2)
	defer close(prices)
	go getMaxPrice(categoryId, prices)
	go getMinPrice(categoryId, prices)

	count := 0
	var category = make(map[string]float64,3)
	var data []float64
	for i:= range prices{
		count++

		data = append(data, i)

		if count== 2 {

			if(data[0]<data[1]){
				category["min"],category["max"] = data[0],data[1]
			} else {
				category["min"],category["max"] = data[1],data[0]

			}

			category["suggested"] = getSuggestedPrice(category["min"], category["max"])

			break
		}
	}

	return category
}

func getData(categoryId, order string)(float64){
	apiClient := http.Client{
		Timeout: time.Second * 10,
	}

	countrySite := categoryId[:3]

	url := "https://api.mercadolibre.com/sites/"+countrySite+"/search?limit=1&category="+categoryId+"&sort="+order

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return formatPrice(0.0)
	}

	req.Header.Set("User-Agent", "price-api")

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
		return formatPrice(0.0)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return formatPrice(0.0)
	}

	categoryData := CategoryRemoteApi{}
	jsonErr := json.Unmarshal(body, &categoryData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return formatPrice(0.0)
	}


	if(categoryData.Paging.Total>0){
		return formatPrice(categoryData.Results[0].Price)
	}

	return formatPrice(0.0)
}

func getMaxPrice(categoryId string, channel chan<- float64){
	channel <- getData(categoryId, "price_desc")
}

func getMinPrice(categoryId string,channel chan<- float64){
	channel <- getData(categoryId, "price_asc")
}

func getSuggestedPrice(minPrice, maxPrice float64)(float64){
	return formatPrice((minPrice+maxPrice)/2)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func formatPrice(value float64)(float64){
	output := math.Pow(10, float64(2))
	return float64(round(value * output)) / output
}