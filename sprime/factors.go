package main

import (
	"fmt"
	"github.com/otiai10/sprime"
	"regexp"
	"strconv"
)

type cFactors struct {
	origin int
}

var numericExp = regexp.MustCompile("([0-9]+)")

func (c *cFactors) Prepare() {
	args := getArgs()
	if len(args) < 2 {
		invalid("`factors` needs second arg like `sprime factors 12`.")
		return
	}
	if !numericExp.MatchString(args[1]) {
		invalid("`factors` arg must be number expression like `12345`.")
		return
	}
	m := numericExp.FindStringSubmatch(args[1])
	num, _ := strconv.Atoi(m[1])
	c.origin = num
}

func (c *cFactors) Perform() {
	fmt.Println(sprime.Factorize(float64(c.origin)).List())
}