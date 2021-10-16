package configuration

import (
	"flag"
)

type Options struct {
	Logger bool `json:"logger"`
	Town string `json:"town"`
	Units string `json:"units"`
}

func GetFlags() (o *Options) {
	l := flag.Bool("l", false, "-l logger is true, none - logger is false")
	t := flag.String("t", "moscow", "choose town")
	u := flag.String("u", "metric", "standard, metric and imperial")
	flag.Parse()

	o = &Options{Logger: *l, Town: *t, Units: *u}
	return
}
