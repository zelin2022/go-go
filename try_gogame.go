package main

import (
  "./gogame"
)

func main() {
  game := gogame.MakeGame(19)
  game.MakeMove_a(3,3,0)
}
