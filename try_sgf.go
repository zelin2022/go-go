package main
//
//
//
import (
  // "./gogame"
  "./sgfparser"
  "fmt"
  "os"
  "log"
  "bytes"
)
//
//
func captureOutput(f func()) string {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    f()
    log.SetOutput(os.Stderr)
    return buf.String()
}

func main() {
  // dataList := sgfparser.ParseAll("/home/zelin/workspace/some_wq/DATA_SET", uint32(100000))
  // for _, dataOne := range dataList {
  //   dataOne.Print()
  // }
  // println(len(dataList))
  emptyInfo := sgfparser.Info{}
  output := captureOutput(func(){
    emptyInfo.Print()
  })
  fmt.Println(output)
  output = "hi"
  fmt.Println(output)
}
