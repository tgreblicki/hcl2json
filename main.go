package main

import "C"

import (
	"fmt"
	"log"
	"os"

	"github.com/tmccombs/hcl2json/convert"
)

func main() {

}

//export ConvertToJson
func ConvertToJson(input *C.char) *C.char {
	logger := log.New(os.Stderr, "", 0)
	var options convert.Options
	options.Simplify = false
	fmt.Println("\n Hcl input to convert -- %s" + C.GoString(input))

	converted, err := convert.Bytes([] byte(C.GoString(input)), "", options)
	if err != nil {
		logger.Fatalf("Failed to convert file: %v", err)
	}
	fmt.Printf("\n Converted json -- %s", string(converted))

	return C.CString(string(converted))
}
