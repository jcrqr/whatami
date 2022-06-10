package main

import "strings"

type filterFlag []string

func (i filterFlag) String() string {
	return strings.Join(i, ", ")
}

func (i *filterFlag) Set(value string) error {
	*i = append(*i, value)

	return nil
}
