package models

import (
	"errors"
	"log"
	"strconv"

	radix "github.com/armon/go-radix"
)

// Data .
type Data struct {
	ID  int    `json:"id"`
	Foo string `json:"foo"`
}

var r = radix.New()

// Create .
func Create(d Data) (data Data, err error) {
	if d.Foo == "" {
		err = errors.New("empty value")
		log.Println("ERR: ", err)
		return data, nil
	}
	existingCount := r.Len()
	d.ID = existingCount + 1
	r.Insert(strconv.Itoa(d.ID), d.Foo)
	return d, nil
}

// Get .
func Get(id int) (data Data, err error) {
	d, exists := r.Get(strconv.Itoa(id))
	if exists == false {
		err = errors.New("Not found")
		log.Println("ERR : ", err)
		return data, nil
	}
	data.ID = id
	data.Foo = d.(string)
	return data, nil
}
