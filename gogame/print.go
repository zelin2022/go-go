package gogame

import (
  "strconv"
  "fmt"
)

func (g *Game)PrintBoard() {
  line1 := "  "
  for i := uint8(0); i < g.BoardLength; i++ {
    line1 += "  "
    if i < 10 {
      line1 += "0"
    }
    line1 += strconv.Itoa(int(i))
  }
  fmt.Println(line1)

  for j := uint8(0); j < g.BoardLength; j++ {
    line := ""
    if j < 10 {
      line += "0"
    }
    line += strconv.Itoa(int(j))
    for k := uint8(0); k < g.BoardLength; k++ {
      if k != 0 {
        line += "---"
      } else {
        line += "   "
      }

      if g.GoBoard[j][k] == 0 {
        line += "+"
      } else if g.GoBoard[j][k] == 1 {
        line += "X"
      } else if g.GoBoard[j][k] == 2 {
        line += "O"
      } else {
        // TODO error
      }
    }
    fmt.Println(line)

    if j != g.BoardLength - 1 {
      line2 := "  "
      for l := uint8(0); l < g.BoardLength; l++ {
        line2 += "   |"
      }
      fmt.Println(line2)
    }


  }

}
