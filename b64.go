package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
)

const (
	defaultTable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func randomTable(tableFile string) string {
	rand.Seed(time.Now().Unix())
	m := make(map[int]bool)
        out := ""
	for i := 0; i < 64; i++ {
		num := rand.Int() % 64
		if _, ok := m[num]; ok {
			i--
			continue
		} else {
			m[num] = false
			out += string(defaultTable[num])
		}
	}

        file, err := os.Create(tableFile)
        if err != nil {
          panic(err)
        }
        _, err = file.WriteString(out)
        if err != nil {
          panic(err)
        }

        return out
}

func base64Encode(src, table string){
	in, err := ioutil.ReadFile(src)
        if err != nil {
          panic(err)
        }

	enc := base64.NewEncoding(table)

	out := enc.EncodeToString(in)
	fmt.Println(out)
}

func base64Decode(src, table string){
	in, err := ioutil.ReadFile(src)
        if err != nil {
          panic(err)
        }

	enc := base64.NewEncoding(table)

	dec, err := enc.DecodeString(string(in))
	if err != nil {
	  panic(err)
        }
	fmt.Println(string(dec))
}

func main() {
	random := flag.String("r", "", "use random table and output table to specific data")
	custom := flag.String("t", "", "use custom table")
	decode := flag.Bool("d", false, "decode mode")
	encode := flag.Bool("e", false, "encode mode")

	flag.Parse()

	// check mode
	if *decode && *encode {
		fmt.Println("You can only select either decode(-d) or encode(-e).")
		return
	} else if !*decode && !*encode {
		fmt.Println("Please either decode(-d) or encode(-e).")
		return
	}

	// check table
	var table string
	if len(*random) != 0 && len(*custom) != 0 {
		fmt.Println("You can only select either custom table(-t) or random table(-r)")
		return
	} else if len(*random) != 0 {
		table = randomTable(*random)
	} else if len(*custom) != 0 {
	        customTable, err := ioutil.ReadFile(*custom)
	        if err != nil {
	          panic(err)
                }
		table = string(customTable)
	} else {
		table = defaultTable
	}

        // check input file
        args := flag.Args()
        if len(args) < 1 {
          fmt.Println("no input file")
        }else if len(args) > 1{
          fmt.Println("too many argument")
        }
        inputFile := args[0]
        
        // process
        if *encode {
          base64Encode(inputFile, table)
        }else if *decode {
          base64Decode(inputFile, table)
        }
}
