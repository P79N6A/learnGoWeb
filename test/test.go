package main

import (
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)

type Node struct {
	Alias   string   `yaml:"alias"`
	Headers []string `yaml:"headers"`
	Assert []string `yaml:"assert"`
}

func main() {
	// 待解析数据
	yamlFile, _ := ioutil.ReadFile("text.yaml")

	var nodes []Node
	//var nodes []map[string]interface{}
	_ = yaml.Unmarshal(yamlFile, &nodes)
	log.Println(nodes)
	// result
	// map[field1:小韩说课 field2:map[field3:value field4:[42 1024]]]
}
