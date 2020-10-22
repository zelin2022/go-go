package sgfparser

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "os"
  "log"
  "bytes"
  "../gogame"
)

func captureOutput(f func()) string {
    var buf bytes.Buffer
    log.SetOutput(&buf)
    f()
    log.SetOutput(os.Stderr)
    return buf.String()
}

func Test_letterCoordToNumberCoord (t *testing.T) {
  output := letterCoordToNumberCoord("aa")
  exOutput := gogame.Point{0,0}
  assert.Equal(t, exOutput, output)
}

func Test_parseKeyValue (t *testing.T) {
  
}
