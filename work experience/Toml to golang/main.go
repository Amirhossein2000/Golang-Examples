package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	var favorites songs

	b, _ := ioutil.ReadFile("test.toml")

	if _, err := toml.Decode(string(b), &favorites); err != nil {
		log.Fatal(err)
	}

	fmt.Println(favorites)
}

type song struct {
	Name     string
	Duration duration
}
type songs struct {
	Song []song
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
