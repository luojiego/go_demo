package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//如何正确的给一个类实现MarshalJSON接口来让json.MarshalJSON正确调用

type Test struct {
	Name string
	Age  int
}

//stack overflow
/*func (t Test) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}*/

//不能正常被调用
/*func (t *Test) MarshalJSON() ([]byte, error) {
	return []byte("{\"name\":\"" + t.Name +"\",\"age\":" + strconv.FormatInt(int64(t.Age), 10) + "}"), nil
}*/

//正确的做法
func (t Test) MarshalJSON() ([]byte, error) {
	return []byte("{\"name\":\"" + t.Name + "\",\"age\":" + strconv.FormatInt(int64(t.Age), 10) + "}"), nil
}

type Data struct {
	Test
	Number int
}

func main() {
	t := Data{
		Test: Test{
			"luojie", 30,
		},
		Number: 15399052129,
	}

	r, _ := json.Marshal(t)

	fmt.Println(string(r))
}
