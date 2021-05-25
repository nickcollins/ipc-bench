package main

import (
	"github.com/adrianleh/WTMP-client"
	"github.com/adrianleh/WTMP-middleend/types"
	"log"
)

// #cgo CFLAGS: -I..
// // ugly hack
// #cgo LDFLAGS: -L${SRCDIR}/../../build/source/common -lipc-bench-common
// #include "main.h"
import "C"

//export ServerMain
func ServerMain(args *C.Arguments) {
	size := uint64(args.size)
	count := int(args.count)

	if size%2 != 0 {
		log.Fatalf("odd size %d, it must be even", size)
	}

	typ := types.ArrayType{
		Length: size / 2,
		Typ:    types.CharType{},
	}

	if err := clientlib.Register("Bench-Server"); err != nil {
		log.Fatal(err)
	}
	if err := clientlib.AcceptType(typ); err != nil {
		log.Fatal(err)
	}

	var bench C.Benchmarks
	C.setup_benchmarks(&bench)

	datumArr := make([]uint16, size/2)
	for j := range datumArr {
		datumArr[j] = 10000
	}

	for i := 0; i < count; i++ {
		bench.single_start = C.now()

		if err := clientlib.Send(typ, "Bench-Client", datumArr); err != nil {
			log.Fatal(err)
		}

		for j := range datumArr {
			datumArr[j] = 20000
		}

		for {
			isEmpty, err := clientlib.Empty(typ)
			if err != nil {
				log.Fatal(err)
			}
			if !isEmpty {
				break
			}
		}

		datum, gErr := clientlib.Get(typ)
		if gErr != nil {
			log.Fatal(gErr)
		}
		datumArr = datum.([]uint16)

		C.benchmark(&bench)
	}

	C.evaluate(&bench, args)
}

func main() {}
