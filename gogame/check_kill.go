package gogame

import (
  "fmt"
)
// ----------------------------------------------------------------------

func (g *Game)checkKill(lastMove Move) bool {

  hasKill := false
  neighbors := lastMove.MovePoint.getNeighbors(g.BoardLength)

  for _, neighbor := range neighbors {
    switch neighborColor := g.GetStoneColorOnBoard(neighbor); neighborColor {
    case 0:
      // nothing
    case lastMove.MoveColor:
      // nothing
    default:
      hasKill = hasKill || g.checkAreaCaptured(Move{neighbor, neighborColor, false}, true)
    }
  }

  return hasKill
}

// ----------------------------------------------------------------------

func (g *Game)checkAreaCaptured(move Move, removeStone bool) bool {
  // TODO check if this is a valild stone MoveColor, not empty slot
  checked := make ([]Point, 0, g.BoardLength*g.BoardLength)
  toCheck := make ([]Point, 1, g.BoardLength*g.BoardLength)
  toCheck[0] = move.MovePoint
  for len(toCheck) > 0 {
    // get neighbors of the last element of toCheck
    neighbors := toCheck[len(toCheck)-1].getNeighbors(g.BoardLength)
    // now that we have the neighbors, we swap the element from toCheck to checked
    checked = append(checked, toCheck[len(toCheck)-1])
    toCheck = toCheck[:len(toCheck)-1]

    for _, neighbor := range neighbors {
      switch neighborColor := g.GetStoneColorOnBoard(neighbor); neighborColor {
      case 0:
        // this area is alive
        return false
      case move.MoveColor:
        // add to toCheck if not exist in toCheck and checked
        if !neighbor.containedIn(checked) && !neighbor.containedIn(toCheck) {
          toCheck = append(toCheck, neighbor)
        }
      default:
        // opposite MoveColor do nothing
      }
    }
  }

  // if we reach here ... checked are captured
  if removeStone {
    fmt.Print("captured: ")
    for _, deadStone := range checked {
      fmt.Print(deadStone.toString(), " ")
      g.ClearStoneColorOnBoard(deadStone)
    }
    fmt.Print("\n")
  }

  return true

}

// ----------------------------------------------------------------------
