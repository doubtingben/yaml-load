package main

import (
	"fmt"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var data = `
octets:
  - production:
      &production-firstoctet
      firstoctet: 10
  - protected:
      &protected-secondoctet
      secondoctet: 20
  - docker:
      &docker-subnet
      thirdoctet: 100
hostgroups:
  - name: dockerL0X
    instances: 3
    startAt: 161
    octets:
      first:
        *production-firstoctet
      second:
        *protected-secondoctet
      third:
        *docker-subnet
`

type T struct {
	HostGroups []struct {
		Name      string
		Instances int
		FirstIP   int
	}
	Subnets []struct {
		Name   string
		Subnet string
	}
}

func main() {

	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
