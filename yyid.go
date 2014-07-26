package yyid

import (
  "crypto/rand"
  "fmt"
)

func New() string {
  randomBytes := make([]byte, 16)
  rand.Read(randomBytes)
  return fmt.Sprintf("%x-%x-%x-%x-%x",
    randomBytes[0:4],
    randomBytes[4:6],
    randomBytes[6:8],
    randomBytes[8:10],
    randomBytes[10:16],
  )
}

