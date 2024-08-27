package main

import (
	"fmt"

	"github.com/jinmukeji/plat-pkg/v4/experiments/whitelist"
)

func main() {
	p, err := whitelist.LoadPolicyFromYamlDir("../testdata/policy")
	fmt.Println(p)
	fmt.Println(err)
}
