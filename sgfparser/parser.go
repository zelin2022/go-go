package sgfparser

import (
  "../gogame"
  // "github.com/davecgh/go-spew/spew"
  "encoding/json"
  "fmt"

  "path/filepath"
  "os"

  "io/ioutil"
  "strings"

  "strconv"
)

  // "io/ioutil"


type Info struct {
  EventName string //EV
  PlayerBlack string //PB
  PlayerWhite string //PW
  RankBlack string //RB
  RankWhite string //RW
  Komi float32 //KM
  Result string //RE
  Date string //DT
  Place string //PC
  BoardSize uint8 //SZ
  RuleSet string //RU
  RoundNumber string //RO
  Handicap string //HA

  PrePlacedMoves []gogame.Move
  Moves []gogame.Move
}
// CA GM TM C EVX SO GC

// CA character set
// GM game type
// TM time?

func (info Info)Print() {
  // spew.Dump(info)
  s, _ := json.MarshalIndent(info, "", "\t")
  fmt.Println(string(s))
}

func ParseAll(dir string, estimate uint32) []Info {
  files := GetAllFilesInDir(dir, estimate)
  infoList := make([]Info, 0, len(files))
  for _, fileName := range files {
    if fileName[len(fileName)-4:] == ".sgf" {
      println(fileName)
      infoList = append(infoList, ParseSingle(fileName))
    }
  }

  return infoList
}

func ParseSingle (fileIn string) Info {
  var infoOutput Info
  dataByte, err := ioutil.ReadFile(fileIn)
  if err != nil {
      //TODO check error
  }
  dataString := string(dataByte)
  dataString = strings.ReplaceAll(dataString, ";", "")
  dataString = strings.ReplaceAll(dataString, "\n", "")
  dataString = strings.ReplaceAll(dataString, "(", "")
  dataString = strings.ReplaceAll(dataString, ")", "")

  dataStringSplit := strings.Split(dataString, "]")

  prevKey := ""
  for _, keyValueString := range dataStringSplit {
    prevKey = infoOutput.parseKeyValue(keyValueString, prevKey)
  }

  return infoOutput
}

func (info *Info)parseKeyValue(input string, prevKey string) string{
  input = strings.TrimSpace(input)
  kvPair := strings.Split(input, "[")

  thisKeyWithMultipleValues := ""
  if len(kvPair) != 2 {
    // fmt.Println("!=2, Expect key[value  // Got " + input)
    return thisKeyWithMultipleValues
  }

  if kvPair[0] == "" {
    kvPair[0] = prevKey
  }

  /*
  EventName string //EV
  PlayerBlack string //PB
  PlayerWhite string //PW
  RankBlack string //RB
  RankWhite string //RW
  Komi float32 //KM
  Result string //RE
  Date string //DT
  Place string //PC
  BoardSize uint8 //SZ
  RuleSet string //RU
  RoundNumber string //RO
  Handicap string //HA

  PrePlacedMoves []gogame.Move //AB  AW
  Moves []gogame.Move //B W
  */
  switch kvPair[0] {
  case "B":
    if kvPair[1] == "" { // if yield
      info.Moves = append(info.Moves, gogame.Move{MoveColor:uint8(1), Yield:false})
    } else {
      info.Moves = append(info.Moves, gogame.Move{letterCoordToNumberCoord(kvPair[1]), uint8(1), false})
    }
  case "W":
    if kvPair[1] == "" { // if yield
      info.Moves = append(info.Moves, gogame.Move{MoveColor:uint8(1), Yield:false})
    } else {
      info.Moves = append(info.Moves, gogame.Move{letterCoordToNumberCoord(kvPair[1]), uint8(2), false})
    }
  case "EV":
    info.EventName = kvPair[1]
  case "PB":
    info.PlayerBlack = kvPair[1]
  case "PW":
    info.PlayerWhite = kvPair[1]
  case "BR":
    info.RankBlack = kvPair[1]
  case "WR":
    info.RankWhite = kvPair[1]
  case "KM":
    tmpFloat64, err := strconv.ParseFloat(kvPair[1], 32)
    if err != nil {
      //TODO error handle
    }
    info.Komi = float32(tmpFloat64)
  case "RE":
    info.Result = kvPair[1]
  case "DT":
    info.Date = kvPair[1]
  case "PC":
    info.Place = kvPair[1]
  case "SZ":
    tmpUint64, err := strconv.ParseUint(kvPair[1], 10, 8)
    if err != nil {

    }
    info.BoardSize = uint8(tmpUint64)
  case "RU":
    info.RuleSet = kvPair[1]
  case "RO":
    info.RoundNumber = kvPair[1]
  case "HA":
    info.Handicap = kvPair[1]
  case "FF":

  case "AB":
    info.PrePlacedMoves = append(info.PrePlacedMoves, gogame.Move{letterCoordToNumberCoord(kvPair[1]), uint8(2), false})
    thisKeyWithMultipleValues = kvPair[0]
  case "AW":
    info.PrePlacedMoves = append(info.PrePlacedMoves, gogame.Move{letterCoordToNumberCoord(kvPair[1]), uint8(2), false})
    thisKeyWithMultipleValues = kvPair[0]
  default:
    // fmt.Println("Unknown key: "+kvPair[0]+" Value: "+kvPair[1])
  }

  return thisKeyWithMultipleValues
}

func letterCoordToNumberCoord(letterCoord string) gogame.Point {
  return gogame.Point{YCoord:uint8(letterCoord[0])-uint8(97), XCoord:uint8(letterCoord[1])-uint8(97)}
}



func GetAllFilesInDir(dir string, estimate uint32) []string {
  fileList := make([]string, 0, estimate)
  filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    if err == nil  { //TODO change error check
        fileList = append(fileList, path)
    }
    return nil
  })
  return fileList
}























//
