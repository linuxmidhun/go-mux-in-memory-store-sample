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
		err = errors.New("A new Foo cannot be created with empty value")
		log.Println("ERR: ", err)
		return data, err
	}
	existingCount := r.Len()
	data.ID = existingCount + 1
	data.Foo = d.Foo
	log.Println("Creating a new Foo as ", data)
	r.Insert(strconv.Itoa(data.ID), data.Foo)
	log.Println("Created new Foo as ", data)
	return data, nil
}

// Get .
func Get(id int) (data Data, err error) {
	log.Println("Looking for Foo with Id '", id, "'")
	d, exists := r.Get(strconv.Itoa(id))
	if exists == false {
		err = errors.New("Foo for id '" + strconv.Itoa(id) + "' was not found")
		log.Println("ERR : ", err)
		return data, err
	}
	data.ID = id
	data.Foo = d.(string)
	log.Println("Foo for id '", id, "' was found as ", data)
	return data, nil
}
