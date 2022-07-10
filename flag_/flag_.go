package flag_

import (
	"flag"
	"log"
)

var (
	FlagStr = flag.String("string", "default", "this is string flag!!!")
)

func Main() {
	flag.Parse()
	log.Println(*FlagStr)
	log.Println(flag.Args())
}
