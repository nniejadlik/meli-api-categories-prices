package categoriesPrices

import (
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

type priceType float32
type CategoryRemoteApi struct {
	Paging struct {
		Total int `json:"total"`
	} `json:"paging"`
	Results []struct {
		Price      priceType    `json:"price"`
	} `json:"results"`
}

func getPrices(categoryId string)(map[string]priceType){

	var prices = make(chan priceType,2)
	defer close(prices)
	go getMaxPrice(categoryId, prices)
	go getMinPrice(categoryId, prices)

	count := 0
	var category = make(map[string]priceType,3)
	var data []priceType
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

func getData(categoryId, order string)(priceType){
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

func getMaxPrice(categoryId string, channel chan<- priceType){
	channel <- getData(categoryId, "price_desc")
}

func getMinPrice(categoryId string,channel chan<- priceType){
	channel <- getData(categoryId, "price_asc")
}

func getSuggestedPrice(minPrice, maxPrice priceType)(priceType){
	return formatPrice((minPrice+maxPrice)/2)
}

func formatPrice(value priceType)(priceType){
	return value
}