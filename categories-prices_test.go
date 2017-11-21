package categoriesPrices

import (
	"testing"
	"fmt"
	"math/rand"
)


func TestGetPrices(t *testing.T){
	fmt.Print("\ngetPrices\n")

	categories := [8]string{"MLA1234","MLA1235","MLA123","12345","MLA109291","MLA5725","MLA4711","MLA6520"}

	for _,i:= range categories{
		prices := getPrices(i)
		fmt.Println(i,prices)
	}
}


func TestGetData(t *testing.T) {
	fmt.Print("\ngetData\n")

	categories := make(map[string]string)
	categories["MLA1234"] = "price_asc"
	categories["MLA1235"] = "price_desc"
	categories["MLA123"] = "price_asc"
	categories["12345"] = "price_desc"
	categories["MLA109291"] = "price_asc"
	categories["MLA5725"] = "price_desc"
	categories["MLA4711"] = "price_asc"
	categories["MLA6520"] = "price_desc"


	for k,i := range categories {
		price := getData(k, i)
		fmt.Println(k,i,price)
	}
}


func TestGetMaxPrice(t *testing.T){
	fmt.Print("\ngetMaxPrice\n")

	var prices = make(chan priceType,2)
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
	var prices = make(chan priceType,2)
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


func TestGetSuggestedPrice(t *testing.T){
	fmt.Print("\ngetSuggestedPrice\n")

	for i:=0; i<10; i++{
		var x,y priceType = priceType(rand.Float32()), priceType(rand.Float32())
		price := getSuggestedPrice(x,y)
		fmt.Printf("Value between %v and %v is %v\n",x,y, price)
	}
}


func TestFormatPrice(t *testing.T){
	fmt.Print("\nformatPrice\n")
	prices := [8]priceType{0.0,1.5,6,98.98,9090,7,98.098,76}

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
		prices := getPrices(category)
		fmt.Printf("The prices for %s is %v\n",category,prices)
	}

	//Parallel
	fmt.Print("\nBenchmarGetPrices - Run parallel\n")

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				prices := getPrices(category)
				fmt.Printf("[Parallel] The prices for %s is %v\n",category,prices)
			}
		})
}