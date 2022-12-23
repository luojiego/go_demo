package main

//GOSSAFUNC=main go build main.go

//https://eddycjy.com/posts/go/map/2019-04-07-why-map-no-order/
func main() {
	m := map[string]int{
		"一": 1,
		"二": 2,
		"三": 3,
	}

	for k, v := range m {
		println(k, v)
	}
}
