package main

import (
	"fmt"
	"reflect"
	"repeated/proto/testpb"
)

type Registry struct {
	// methods 保存Struct所拥有的方法信息
	// key: Struct名称.Method名称，例如：computer.add
	methods map[string]reflect.Type
}

func (x *Registry) print() {
	for k, v := range x.methods {
		fmt.Println(k, v)
	}
}

// 注册Struct类型的方法
func (x *Registry) RegisterMethods(key string, pv interface{}) {
	if x.methods == nil{
		x.methods = make(map[string]reflect.Type)
	}

	// pt := pv.Type()
	// v := pv.Elem()
	// t := v.Type()

	// typeName := t.Name()
	// fmt.Println(typeName)

	x.methods[key] = reflect.TypeOf(pv).Elem()
}

func (x *Registry) Get(key string) (interface{}, error)  {
	if v, ok := x.methods[key]; ok {
		return reflect.New(v).Interface(), nil
	}
	return nil, fmt.Errorf("not found key: %s", key)
}

func init()  {
	reg := Registry{}
	reg.RegisterMethods("M1", &testpb.M1{})
	reg.print()
}

func main() {
	// t := testpb.Test{}
	// t.List = append(t.List, false)
	// t.List = append(t.List, false)
	// t.List = append(t.List, false, false, false)
	// r, err := proto.Marshal(&t)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(len(r), r)

	// t1 := &testpb.Test{}
	// err = proto.Unmarshal(r, t1)
	// fmt.Printf("%+v\n", t1)
}
