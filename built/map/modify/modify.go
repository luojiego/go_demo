package main

import "fmt"

type Person struct {
	Name string
}

type Dic map[int]int

func NewDic() *Dic {
	return &Dic{}
}

func (p Dic) Find(key int) int {
	return 0
}

func (p Dic) Add(key int, val int) {
	if p == nil {
		p = make(map[int]int)
	}
	p[key] = val
}

func (p Dic) Remove(key int) {
	if _, ok := p[key]; ok {
		p[key] = 0
	}
}

func main() {
	var dic Dic
	dic = map[int]int{}
	dic.Add(1, 1)
	dic.Add(2, 2)
	dic.Add(3, 3)
	fmt.Printf("after add: %+v\n", dic)
	dic.Remove(3)
	fmt.Printf("after remove: %+v\n", dic)

	return
	m := make(map[string]*Person, 0)

	m["一"] = &Person{
		Name: "1",
	}
	fmt.Println(m)

	v, ok := m["一"]
	if ok {
		// v = &Person{
		// 	Name: "1111",
		// }
		v.Name = "1111"
	}
	_ = v

	/*if _, ok := m["一"]; ok {
		m["一"] = &Person{
			Name: "1111",
		}
	}*/

	for k, v := range m {
		fmt.Println(k, v.Name)
	}
}
