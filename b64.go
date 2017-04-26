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

func randomTable() string {
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
  return out
}

func base64Encode(src, table string)string{
  enc := base64.NewEncoding(table)
  out := enc.EncodeToString([]byte(src))
  return out
}

func base64Decode(src, table string)string{
  enc := base64.NewEncoding(table)

  out, err := enc.DecodeString(src)
  if err != nil {
    panic(err)
  }
  return string(out)
}

func b64Version(){
  fmt.Println("b64 v1.1")
}

func main() {
  recursion := flag.Int("m", 1, "decode or encode multiple times")
  random := flag.String("r", "", "use random table and output table to specific data")
  custom := flag.String("t", "", "use custom table")
  decode := flag.Bool("d", false, "decode mode")
  encode := flag.Bool("e", false, "encode mode")
  version := flag.Bool("v", false, "b64 version")

  flag.Parse()

  // check special case
  if *version {
    b64Version()
    return
  }

  // check mode
  if *decode && *encode {
    fmt.Println("You can only select either decode(-d) or encode(-e).")
    return
  } else if !*decode && !*encode {
    fmt.Println("Please either decode(-d) or encode(-e).")
    return
  }

  // check table
  table := make([]string, *recursion)
  if len(*random) != 0 && len(*custom) != 0 {
    fmt.Println("You can only select either custom table(-t) or random table(-r)")
    return
  } else if len(*random) != 0 {
    file, err := os.Create(*random)
    if err != nil {
      panic(err)
    }
    for i:=0 ; i<*recursion ; i++ {
      table[i] = randomTable()
      _, err = file.WriteString(table[i])
      if err != nil {
        panic(err)
      }
    }
  } else if len(*custom) != 0 {
    customTable, err := ioutil.ReadFile(*custom)
    if err != nil {
      panic(err)
    }
    if len(customTable)/64 == 1 {
      for i:=0;i<*recursion;i++{
        table[i] = string(customTable)
      }
    }else if len(customTable)/64 != *recursion {
      fmt.Println("Tables are not enough.")
    }else{
      for i:=0;i<*recursion;i++{
        table[i] = string(customTable[64*i:64*(i+1)])
      }
    }
  } else {
    for i:=0;i<*recursion;i++{
      table[i] = defaultTable
    }
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
    plainText, err := ioutil.ReadFile(inputFile)
    if err != nil {
      panic(err)
    }
    cipherText := string(plainText)
  if *encode {
    for i:=0;i<*recursion;i++{
      cipherText = base64Encode(cipherText, table[i])
    }
    fmt.Println(cipherText)
  }else if *decode {
    for i:=*recursion-1;i>=0;i--{
      cipherText = base64Decode(cipherText, table[i])
    }
    fmt.Println(cipherText)
  }
}
