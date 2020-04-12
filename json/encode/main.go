package main

import (
	"bytes"
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
func (t *Test) MarshalJSON() ([]byte, error) {
	return []byte("{\"name\":\"" + t.Name + "\",\"age\":" + strconv.FormatInt(int64(t.Age), 10) + "}"), nil
}

//正确的做法
/*func (t Test) MarshalJSON() ([]byte, error) {
	return []byte("{\"name\":\"" + t.Name + "\",\"age\":" + strconv.FormatInt(int64(t.Age), 10) + "}"), nil
}*/

type Data struct {
	Test
	Number  int
	Message string
}

func main() {

	inheritByJsonMarshal()
	inheritByMarshalText()
	//HTMLEscape
	//replaces "<", ">", "&", U+2028, and U+2029 are escaped
	//to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029".
	//没有搞明白 U+2028（行分隔符） 和 U+2029（段落分隔符） 怎么显示在字符串中

	//使用omitempty可以将nil指针，nil接口，空数组，空slice，空map或者string, false进行忽略
	//proto文件转pb.go文件中 大量使用了omitempty
	t2 := struct {
		Message string `json:",omitempty"`
		Age     int
		Number  int `json:"number,omitempty"`
	}{
		Message: "<>&U +2028U +2029",
		Age:     30,
	}
	r2, _ := json.Marshal(t2)
	fmt.Println(string(r2)) //{"Message":"\u003c\u003e\u0026U +2028U +2029","Age":30}

	//对比 t2
	//1. 未设置Number，但是没有加omitempty，所以会在序列化的时候会增加 "number": 0
	//2. json标识中增加了`json:",omitempty"`, 在有值的时候会将Message作为key，没有值的时候为空
	//3. 由于Age的tag是`json:"-"`，则即使Age有值，也会序列化的时候忽略掉Age
	//4. 由于使用了`json:"-"` 来使得序列化忽略字段，那么真的需要“-” 作为键值的时候怎么办呢？ 使用 `json:"-,"` 来使得key为"-"
	//		同时在第一版本的proto go的API中，大量也使用了`json:"-"`
	//5. 由于需要和 JS 打交道 也可以使得`json:",string"` 将浮点，整数或者 boolean 在序列化的时候以 string 的格式输出
	t3 := struct {
		Message string  `json:",omitempty"`
		Age     int     `json:"-"`
		Number  int     `json:"number"`
		Index   int     `json:"-,"`
		Score   float64 `json:",string"`
	}{
		Age:   30,
		Index: 710,
		Score: 98.99,
	}
	r3, _ := json.Marshal(t3)
	fmt.Println(string(r3)) //{"number":0,"-":710,"Score":"98.99"}

	//不转义 "<",">","&" 的处理方法
	SetEscapeHTMLFalse()

	//json 对于 slice 和处理
	SliceProcess()

	//可见性的处理
	visibilityProcess()

	//关于 Map 的 json 的 Marshal
	// key 必须满足以下条件
	//1 字符串类型 那么 byte[] 呢？Invalid map key type: the comparison operators == and != must be fully defined for key type
	//2 整型
	//3 实现了encoding.TextMarshaler方法

	/*m := make(map[[2]byte]string)

	m[[2]byte{'a','b'}] = "ab"
	m[[2]byte{'A','B'}] = "AB"

	fmt.Printf("%+v\n", m)

	mapResult, err := json.Marshal(m)
	if err != nil {
		panic(err) //panic: json: unsupported type: map[[2]uint8]string
	}*/

	//fmt.Println(string(mapResult))

	tryStructAsMapKey()
	integerAsMapKey()
	boolAsMapKey()
}

type sType struct {
	Name string
	Age  int
}

func (s sType) MarshalText() (text []byte, err error) {
	return []byte(s.Name), nil
}

func tryStructAsMapKey() {
	m := make(map[sType]string, 0)

	m[sType{"luojie", 30}] = "luojie"
	m[sType{"LuoJie", 31}] = "LuoJie"

	r, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(r))
}

func integerAsMapKey() {
	m := make(map[int]string)
	m[0] = "ling"
	m[1] = "yi"
	m[2] = "er"

	r, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(r))
}

func boolAsMapKey() {
	m := make(map[bool]string)
	m[false] = "ling"
	m[true] = "yi"

	r, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(r))
}

func inheritByJsonMarshal() {
	t := Data{
		Test: Test{
			"LuoJie", 30,
		},
		Number: 1539905,
	}

	//由于 MarshalJSON 为指针接收者 直接序列化 不会有效果（此处有比较大的疑惑）
	r1, _ := json.Marshal(t) //{"Name":"LuoJie","Age":30,"Number":1539905,"Message":""}

	//使用了Test的MarshalJSON方法 Number并不会被序列化进来
	fmt.Println(string(r1)) //{"name":"LuoJie","age":30}

	r2, _ := json.Marshal(&t)
	fmt.Println(string(r2))

	t1 := &Data{} //&Data2{}
	err := json.Unmarshal(r1, t1)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	//无论是 Data2 或者 Data 作为t1 的类型 输出均为一样的
	fmt.Println(t1) // &{{LuoJie 30} 0 }
}

type Test2 struct {
	Name string
	Age  int
}

func (t Test2) MarshalText() ([]byte, error) {
	return []byte("{\"name\":\"" + t.Name + "\",\"age\":" + strconv.FormatInt(int64(t.Age), 10) + "}"), nil
}

type Data2 struct {
	Test2
	Number  int
	Message string
}

func inheritByMarshalText() {
	t := Data2{
		Test2: Test2{
			"LuoJie", 30,
		},
		Number: 1539905,
	}

	r1, _ := json.Marshal(t)

	//使用了 Test2 的 MarshalText 方法 Number 并不会被序列化进来
	//和 Marshal 不同的是 结果中的 " 都增加的转义符
	fmt.Println(string(r1)) //"{\"name\":\"LuoJie\",\"age\":30}"

	r2, _ := json.Marshal(&t)
	fmt.Println(string(r2)) //"{\"name\":\"LuoJie\",\"age\":30}"

	t1 := &Data2{} //&Data1{}
	err := json.Unmarshal(r1, t1)
	if err != nil {
		//无论使用Data1 还是 Data2 作为t1的类型 均在Unmarshal的时候出错了
		fmt.Println("err: ", err) //err:  json: cannot unmarshal string into Go value of type main.Data2
		return
	}

	//TODO 此处需要研究一下如果decode
}

//因为直接调用json.MarshalJSON "<",">","&"会被转义，当有需求是不能被转义的时候
//需要用encode.SetEscapeHTML(false)来进行encode
func SetEscapeHTMLFalse() {
	type Test struct {
		Content string
	}
	t := new(Test)
	t.Content = "http://www.baidu.com?id=123&test=1"
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(t)
	fmt.Println(bf.String()) //{"Content":"http://www.baidu.com?id=123&test=1"}
	//参考：https://hacpai.com/article/1524558037151
}

//Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string
//and a nil slice encodes as the null JSON value.
//S3 中的数据 '1','2','3' 经过base64编码之后的结果即为 MTIz
//当 slice 为nil时，序列化的结果为null
//当 slice 为空时，序列化的结果为[]
func SliceProcess() {
	type Test struct {
		S1         []int
		S2         []string
		S3         []byte
		NilSlice   []int
		EmptySlice []int
	}

	t := Test{
		S1:         []int{1, 2, 3},
		S2:         []string{"1", "2", "3"},
		S3:         []byte{'1', '2', '3'},
		EmptySlice: make([]int, 0),
	}

	r, _ := json.Marshal(t)
	fmt.Println(string(r)) // {"S1":[1,2,3],"S2":["1","2","3"],"S3":"MTIz","NilSlice":null,"EmptySlice":[]}

}

func visibilityProcess() {
	type Test1 struct {
		Name string
		Age  int
	}
	type Test2 struct {
		NickName string
		Age      int
	}
	type Data1 struct {
		Test1
		Test2
	}

	d1 := Data1{
		Test1{"LuoJie", 30},
		Test2{"Roger", 30},
	}

	r1, _ := json.Marshal(d1)
	//两个 Age 都被丢弃了
	fmt.Println(string(r1)) //{"Name":"LuoJie","NickName":"Roger"}

	type Test3 struct {
		Name string
		Age  int `json:"age"`
	}
	type Data2 struct {
		Test3
		Test2
	}
	d2 := Data2{
		Test3{"LuoJie", 30},
		Test2{"Roger", 30},
	}

	r2, _ := json.Marshal(d2)
	//由于 Test3 使用了 json tag, Test3 和 Test2 中的 Age 便不再冲突
	fmt.Println(string(r2))

	type Test4 struct {
		NickName string
		Age      int `json:"age"`
	}
	type Data3 struct {
		Test3
		Test4
	}
	d3 := Data3{
		Test3{"LuoJie", 30},
		Test4{"Roger", 30},
	}

	r3, _ := json.Marshal(d3)
	//由于 Test3 和 Test4 均使用了相同 json tag, Test3 和 Test4 中的Age 依然冲突
	fmt.Println(string(r3)) //{"Name":"LuoJie","NickName":"Roger"}

	type Data4 struct {
		Test3
		Test4
		Age int `json:"age"`
	}
	d4 := Data4{
		Test3{"LuoJie", 30},
		Test4{"Roger", 30},
		500,
	}

	r4, _ := json.Marshal(d4)
	//由于 Test3 和 Test4 和 Data4 中的Age 冲突，但是在该情况下 Data4 中的 Age 是可以正常序列化的
	fmt.Println(string(r4)) //{"Name":"LuoJie","NickName":"Roger","age":500}
}
