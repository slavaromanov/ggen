package main

import (
	"flag"
	"github.com/slavaromanov/passgen"
)

var (
	lower, upper, puncts bool
	ex                   string
	length               int
)

func init() {
	flag.BoolVar(&lower, "lower", true, "use lowercase")
	flag.BoolVar(&upper, "upper", true, "use uppercase")
	flag.BoolVar(&puncts, "puncts", false, "use punctuation characters")
	flag.StringVar(&ex, "exclude", "", "sequence of excluded characters")
	flag.IntVar(&length, "len", 8, "define a length of generated password")
}

func main() {
	flag.Parse()
	if ex == "" {
		ex = string('\x00')
	}
	println(passgen.RandomString(length, passgen.GetAlpha(lower, upper, puncts, ex)))
}
