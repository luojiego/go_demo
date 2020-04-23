package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"repeated/proto/testpb"
)

func main() {
	t := testpb.Test{}
	t.List = append(t.List, false)
	t.List = append(t.List, false)
	t.List = append(t.List, false, false, false)
	r, err := proto.Marshal(&t)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(r), r)

	t1 := &testpb.Test{}
	err = proto.Unmarshal(r, t1)
	fmt.Printf("%+v\n", t1)
}
