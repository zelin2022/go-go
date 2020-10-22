package gogame

import (
  "strconv"
  "fmt"
)

type Point struct {
  YCoord uint8
  XCoord uint8
}

type Move struct {
  MovePoint Point
  MoveColor uint8
  Yield bool
}

type Game struct {
  BoardLength uint8
  GoBoard [][]uint8

  History []Move
}

// ----------------------------------------------------------------------
func MakeGame(len uint8) Game {
  board_slice := make([][]uint8, len, len)
  for i := range board_slice {
    board_slice[i] = make([]uint8, len, len)
  }
  history_slice := make([]Move, len*len*2, len*len*2)
  return Game{BoardLength: len, GoBoard: board_slice, History: history_slice}
}

// Point to string
func (p Point)toString() string {
  return "(" + strconv.Itoa(int(p.YCoord)) + "," + strconv.Itoa(int(p.XCoord)) + ")"
}

// ----------------------------------------------------------------------

// check if a point is on(within) a board of certain length
func (p *Point)checkXYValid(BoardLength uint8) bool {
  return p.XCoord < BoardLength && p.YCoord < BoardLength
}

// ----------------------------------------------------------------------

// get up/down/left/right neighbors, provided they are valid on board
func (p *Point)getNeighbors(BoardLength uint8) []Point {
  var neighbors []Point

  possibleNeighbor := *p
  possibleNeighbor.YCoord -= 1
  if possibleNeighbor.checkXYValid(BoardLength) {
    neighbors = append(neighbors, possibleNeighbor)
  }
  possibleNeighbor = *p
  possibleNeighbor.XCoord+=1
  if possibleNeighbor.checkXYValid(BoardLength) {
    neighbors = append(neighbors, possibleNeighbor)
  }
  possibleNeighbor = *p
  possibleNeighbor.YCoord+=1
  if possibleNeighbor.checkXYValid(BoardLength) {
    neighbors = append(neighbors, possibleNeighbor)
  }
  possibleNeighbor = *p
  possibleNeighbor.XCoord-=1
  if possibleNeighbor.checkXYValid(BoardLength) {
    neighbors = append(neighbors, possibleNeighbor)
  }

  return neighbors
}

// ----------------------------------------------------------------------

// Find if a point is contained in a slice
// In other language I'd use SliceofObject.contains, but in Go it's Object.containedIn
func (p *Point)containedIn(slice []Point) bool {
  for _, a := range slice {
    if a.XCoord == p.XCoord && a.YCoord == p.YCoord {
        return true
    }
  }
  return false
}

// ----------------------------------------------------------------------

// alias for MakeMove
func (g *Game)MakeMove_a( y uint8, x uint8, color uint8) {
  g.MakeMove(Move{Point{y,x}, color, false})
}

// ----------------------------------------------------------------------

// Make a move!
func (g *Game)MakeMove(move Move)  {
  if move.Yield { //do nothing if player yield the move
    return
  }
  // check x and y validity
  if !move.MovePoint.checkXYValid(g.BoardLength) {
    // TODO missing error handling
  }
  // check for invalid moves, aka suicide moves, or ko
  if g.checkKo(move) {
    // TODO handle error, or tell usr invalid move
    fmt.Println("INVALID MOVE(Ko); ", move.MovePoint.toString(), " ", move.MoveColor)
    return
  }

  // set the stone
  g.SetStoneColorOnBoard(move.MovePoint, move.MoveColor)

  // check kill
  hasKill := g.checkKill(move)

  if !hasKill {
    if g.checkSuicide(move) {
      // remove the stone
      g.SetStoneColorOnBoard(move.MovePoint, uint8(0))
      // TODO handle error, or tell usr invalid move
      fmt.Println("INVALID MOVE(Suicide): ", move.MovePoint.toString(), " ", move.MoveColor)
      return
    }
  }




  // check win
  g.check_win_CHN()

}

// ----------------------------------------------------------------------

func (g *Game)checkKo(move Move) bool {
  // check if this move is equal to last last move, (last move of the same player)
  if len(g.History) < 2 {
    return false
  }
  return move == g.History[len(g.History) - 2]
}

// ----------------------------------------------------------------------

func (g *Game)checkSuicide(move Move) bool {
  // TODO if it counter-captures, then it's not suicide
  return g.checkAreaCaptured(move, false)
}

// ----------------------------------------------------------------------

func (g *Game)GetStoneColorOnBoard(point Point) uint8 {
  return g.GoBoard[point.YCoord][point.XCoord]
}

// ----------------------------------------------------------------------

func (g *Game)SetStoneColorOnBoard(point Point, MoveColor uint8) {
  if g.GoBoard[point.YCoord][point.XCoord] != 0 {
    // TODO missing error handing, point not empty
    return
  }
  g.GoBoard[point.YCoord][point.XCoord] = MoveColor
}

// ----------------------------------------------------------------------

func (g *Game)ClearStoneColorOnBoard(point Point) {
  g.GoBoard[point.YCoord][point.XCoord] = uint8(0)
}






























//
