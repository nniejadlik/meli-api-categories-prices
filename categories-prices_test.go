package categoriesPrices

import (
	"testing"
	"fmt"
	"math/rand"
)

var Categories = [8]string {"MLA1234", "MLA1235", "MLA123", "12345", "MLA109291", "MLA5725", "MLA4711", "MLA6520"}

func TestGetPrices(t *testing.T){

	for _,i:= range Categories{
		prices := GetPrices(i)

		if !(prices["max"] >= 0){
			t.Error(fmt.Sprintf("The maximum price was expected for category %s greater than zero, but instead got %v", i, prices["max"] ))
		}

		if !(prices["min"] >= 0){
			t.Error(fmt.Sprintf("The minimum price was expected for category %s greater than zero, but instead got %v", i, prices["min"] ))
		}

		if !(prices["suggested"] >= 0){
			t.Error(fmt.Sprintf("The suggested price was expected for category %s greater than zero, but instead got %v", i, prices["suggested"] ))
		}
	}
}

func TestGetData(t *testing.T) {

	order := map[string]string {"MLA1234":"price_asc", "MLA1235":"price_desc", "MLA123":"price_asc", "12345":"price_desc", "MLA109291":"price_asc", "MLA5725":"price_desc", "MLA4711":"price_asc", "MLA6520":"price_desc"}

	for _, category := range Categories {
		price := getData(category, order[category])
		if !(price >= 0){
			t.Error(fmt.Sprintf("The price was expected greater than zero, but instead got %v", price ))
		}
	}
}


func TestGetMaxPrice(t *testing.T){

	var prices = make(chan float64,2)


	for _,i:= range Categories{
		go getMaxPrice(i,prices)
	}

	count := 0
	for i:= range prices{
		count++

		if !(i >= 0){
			t.Error(fmt.Sprintf("The maximum price was expected greater than zero, but instead got %v", i ))
		}

		if count == 8{
			break
		}
	}

}


func TestGetMinPrice(t *testing.T){
	var prices = make(chan float64,2)
	categories := [8]string{"MLA1234","MLA1235","MLA123","12345","MLA109291","MLA5725","MLA4711","MLA6520"}

	for _,i:= range categories{
		go getMinPrice(i,prices)
	}

	count := 0
	for i:= range prices{
		count++
		if !(i >= 0){
			t.Error(fmt.Sprintf("The minimum price was expected greater than zero, but instead got %v", i ))
		}

		if count == 8{
			break
		}
	}

}


func TestGetsuggestedPrice(t *testing.T){

	for i:=0; i<10; i++{
		var x,y float64 = rand.Float64(), rand.Float64()
		price := getSuggestedPrice(x,y)

		if !(price >= 0){
			t.Error(fmt.Sprintf("The suggested price was expected greater than zero, but instead got %v", price ))
		}
	}
}


func TestFormatPrice(t *testing.T){
	prices := [8]float64{0.0,1.5,6,98.98,9090,7,98.098,76}

	for _,i:= range prices{
		price := formatPrice(i)

		if !(price >= 0){
			t.Error(fmt.Sprintf("The price was expected greater than zero, but instead got %v", price ))
		}
	}
}


func benchmarkGetPrices(cant int, b *testing.B){
	category := "MLA6520"

	b.N = cant

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				GetPrices(category)
			}
		})
}

func BenchmarkGetPrices1(b *testing.B) {
	benchmarkGetPrices(1, b)
}

func BenchmarkGetPrices5(b *testing.B) {
	benchmarkGetPrices(5, b)
}

func BenchmarkGetPrices10(b *testing.B) {
	benchmarkGetPrices(10, b)
}

func BenchmarkGetPrices25(b *testing.B) {
	benchmarkGetPrices(25, b)
}

func BenchmarkGetPrices50(b *testing.B) {
	benchmarkGetPrices(50, b)
}

func BenchmarkGetPrices100(b *testing.B) {
	benchmarkGetPrices(100, b)
}