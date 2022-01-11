package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

type Item struct {
	Id string `json:"id"`
	HashCode string
}

func (i *Item) UnmarshalJSON(data []byte) error  {
	type aliasType Item
	aliasItem := &struct{*aliasType}{aliasType:(*aliasType)(i)}
	if err := json.Unmarshal(data, &aliasItem); err != nil {
		return err
	}
	i.HashCode = fmt.Sprintf("%x", sha1.Sum(data))
	return nil
}

func main() {
	data := []byte(`{"id":"foo"}`)
	item := Item{}
	err := json.Unmarshal(data, &item)
	fmt.Println("err: ", err)
	fmt.Println("item: ", item)
}
