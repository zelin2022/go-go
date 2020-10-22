package gogame



// Chinese Scoring: Territory + Stones + Komi = Score
func (g *Game)check_win_CHN() {

}

// ----------------------------------------------------------------------

// Japanese Scoring: Territory + Prisoners + Komi = Score
func (g* Game)check_win_JPN() {

}

// ----------------------------------------------------------------------

// territory is the surrounded area, with dead stones removed
func (g *Game)check_territory()(uint8, uint8) {
  var p1Territory uint8 = 0;
  var p2Territory uint8 = 0;
  checkedPoints := make ([]Point, 0, g.BoardLength*g.BoardLength)
  for  i := 0, i < g.BoardLength, i++ {
    for j := 0, j < g.BoardLength, j++ {
        if g.GoBoard[i][j] == 0 && Point{i,j}.containedIn(checkedPoints) == false {
          this.check_territory_from_1_point(Point{i,j}. checkedPoints)
        }
    }
  }
}

// returns:
// first uint8 is the color the territory belong to
// second is the size of the territory
func (g *Game)check_territory_from_1_point(thePoint Point, emptyPoints *[]Point)(uint8,uint8) {
  // get all empties in group

  isTerritory := true
  var territoryColor uint8 = 0

  var emptyPointsStartPos uint8 = len(emptyPoints)
  emptyPoints = append(emptyPoints, thePoint)
  var territorySize uint8 = 1


  for emptyPointsStartPos < len(emptyPoints) {
    neightbors := emptyPoints[emptyPointsStartPos].getNeighbors(g.BoardLength)
    for _, neighbor := range neighbors {
      switch neighborColor := g.GetStoneColorOnBoard(neighbor); neighborColor {
      case 0;
        emptyPoints = append(emptyPoints, neighbor)
        territorySize += 1
      case 1:
        if territoryColor == 1 {
          // do nothing
        } else if isTerritory == false {
          // already proven false
          // don't care
          break
        } else if territoryColor == 0{ // first stone it encounters, we are going to assume territory is this color, if proven wrong then fuck it
          territoryColor = 1
        } else {
          isTerritory = false
        }
      case 2:
        if territoryColor == 2 {
          // do nothing
        } else if isTerritory == false {
          // already proven false
          // don't care
          break
        } else if territoryColor == 0{ // first stone it encounters, we are going to assume territory is this color, if proven wrong then fuck it
          territoryColor = 2
        } else {
          isTerritory = false
        }
      default:
        // error
      }
    }
    emptyPointsStartPos += 1
  }

  if (isTerritory) {
    return territoryColor, territorySize
  }
  return 0,0
}
