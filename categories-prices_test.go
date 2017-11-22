package categoriesPrices

import (
	"testing"
	"fmt"
	"math/rand"
)


func TestGetPrices(t *testing.T){
	fmt.Print("\nGetPrices\n")

	categories := map[string]map[string]PriceType{
		"MLA1234":{"max":0.0,"suggested":0.0,"min":0.0},
		"MLA1235":{"max":0.0,"suggested":0.0,"min":0.0},
		"MLA123":{"max":0.0,"suggested":0.0,"min":0.0},
		"12345":{"max":0.0,"suggested":0.0,"min":0.0},
		"MLA109291":{"max":50000,"suggested":25000.5,"min":1},
		"MLA5725":{"max":1e+10,"suggested":5e+09,"min":1},
		"MLA4711":{"max":1.1111111e+08,"suggested":5.5555556e+07,"min":1},
		"MLA6520":{"max":1e+10,"suggested":5e+09,"min":1},
		}

	for k,i:= range categories{
		prices := GetPrices(k)

		if prices["max"] != i["max"]{
			t.Error(fmt.Sprintf("Expected the max price %v for category %s  but instead got %v", i["max"], k, prices["max"] ))
		}

		if prices["min"] != i["min"]{
			t.Error(fmt.Sprintf("Expected the min price %v for category %s  but instead got %v", i["min"], k, prices["min"] ))
		}

		if prices["suggested"] != i["suggested"]{
			t.Error(fmt.Sprintf("Expected the suggested price %v for category %s  but instead got %v", i["suggested"], k, prices["suggested"] ))
		}
	}
}


func TestGetData(t *testing.T) {
	fmt.Print("\ngetData\n")

	categories := []struct {
		id string
        data struct {
        	price PriceType
			order string
		}
    }{
    	{"MLA1234", {0.0, "price_asc"}},
		{"MLA1235",{0.0, "price_desc"}},
		{"MLA123",{0.0, "price_asc"}},
		{"12345",{0.0, "price_desc"}},
		{"MLA109291",{1, "price_asc"}},
		{"MLA5725",{1e+10, "price_desc"}},
		{"MLA4711",{1, "price_asc"}},
		{"MLA6520",{1e+10, "price_desc"}},
	}


	for category := range categories {
		price := getData(category["id"], category["data"]["order"])
		if price != PriceType(category["data"]["price"]){
			t.Error(fmt.Sprintf("Expected the price %v for category %s  but instead got %v", category["data"]["price"], category["id"], price ))
		}
	}
}


func TestGetMaxPrice(t *testing.T){
	fmt.Print("\ngetMaxPrice\n")

	var prices = make(chan PriceType,2)
	categories := [8]string{"MLA1234","MLA1235","MLA123","12345","MLA109291","MLA5725","MLA4711","MLA6520"}

	for _,i:= range categories{
		go getMaxPrice(i,prices)
	}

	count := 0
	for i:= range prices{
		count++
		fmt.Println(i)

		if count == 8{
			break
		}
	}

}


func TestGetMinPrice(t *testing.T){
	fmt.Print("\ngetMinPrice\n")
	var prices = make(chan PriceType,2)
	categories := [8]string{"MLA1234","MLA1235","MLA123","12345","MLA109291","MLA5725","MLA4711","MLA6520"}

	for _,i:= range categories{
		go getMinPrice(i,prices)
	}

	count := 0
	for i:= range prices{
		count++
		fmt.Println(i)

		if count == 8{
			break
		}
	}

}


func TestGetsuggestedPrice(t *testing.T){
	fmt.Print("\ngetsuggestedPrice\n")

	for i:=0; i<10; i++{
		var x,y PriceType = PriceType(rand.Float32()), PriceType(rand.Float32())
		price := getSuggestedPrice(x,y)
		fmt.Printf("Value between %v and %v is %v\n",x,y, price)
	}
}


func TestFormatPrice(t *testing.T){
	fmt.Print("\nformatPrice\n")
	prices := [8]PriceType{0.0,1.5,6,98.98,9090,7,98.098,76}

	for _,i:= range prices{
		price := formatPrice(i)
		fmt.Printf("The price %v formated is: %v\n",i,price)
	}
}


func BenchmarkGetPrices(b *testing.B){
	category := "MLA6520"

	b.N = 500

	fmt.Print("\nBenchmarGetPrices\n")

	for i := 0; i < b.N; i++ {
		prices := GetPrices(category)
		fmt.Printf("The prices for %s is %v\n",category,prices)
	}

	//Parallel
	fmt.Print("\nBenchmarGetPrices - Run parallel\n")

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				prices := GetPrices(category)
				fmt.Printf("[Parallel] The prices for %s is %v\n",category,prices)
			}
		})
}