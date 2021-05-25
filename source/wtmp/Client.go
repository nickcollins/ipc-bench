package main

import (
	"github.com/adrianleh/WTMP-client"
	"github.com/adrianleh/WTMP-middleend/types"
	"log"
)

// #cgo CFLAGS: -I..
// // ugly hack
// #cgo LDFLAGS: -L${SRCDIR}/../../build/source/common -lipc-bench-common -lm
// #include "main.h"
import "C"

//export ClientMain
func ClientMain(args *C.Arguments) {
	size := uint64(args.size)
	count := int(args.count)

	if size%2 != 0 {
		log.Fatalf("odd size %d, it must be even", size)
	}

	typ := types.ArrayType{
		Length: size / 2,
		Typ:    types.CharType{},
	}

	if err := clientlib.Register("Bench-Client"); err != nil {
		log.Fatal(err)
	}
	if err := clientlib.AcceptType(typ); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < count; i++ {
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

		// Dummy operation
		datumArr := datum.([]uint16)
		for j := range datumArr {
			datumArr[j] = 20000
		}

		if err := clientlib.Send(typ, "Bench-Server", datum); err != nil {
			log.Fatal(err)
		}
	}
}
