package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DataGroup struct {
	Name        string `yaml: "name"`
	Description string `yaml: "description"`
	Index       string `yaml: "index"`
	Sourcetype  string `yaml: "sourcetype"`
	Source      string `yaml: "source"`
	Path        string `yaml: "path"`
}

func main() {
	buf, err := ioutil.ReadFile("./sample.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	dgs := make([]DataGroup, 10)
	err = yaml.UnmarshalStrict(buf, &dgs)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dg := range dgs {
		fmt.Printf(
			"add oneshot -index %s -sourcetype %s -source %s %s\n",
			dg.Index, dg.Sourcetype, dg.Source, dg.Path)
	}
}
