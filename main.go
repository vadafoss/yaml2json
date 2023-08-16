package main

import (
	"fmt"
	"log"
	"strings"

	"k8s.io/apimachinery/pkg/util/yaml"
)

type Foo struct {
	Name string `json:"name" yaml:"name"`
}

const fooStr = `
name: bar
`

func main() {
	decoder := yaml.NewYAMLOrJSONDecoder(strings.NewReader(fooStr), 4096)
	f := Foo{}
	if err := decoder.Decode(&f); err != nil {
		log.Fatalf("failed to decode foo: %v", err)
	}

	fmt.Println("foo", f)
}
