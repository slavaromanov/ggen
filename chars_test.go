package passgen

import (
  "reflect"
  "testing"
)

func neq(a, b []byte) bool {
  return !reflect.DeepEqual(a, b)
}

func TestNewAlpha(t *testing.T) {
  l := "NewAlpha: "
  switch {
  case neq(newAlphabet(true, false, false), lower):
    t.Error(l + "lower")
  case neq(newAlphabet(false, true, false), upper):
    t.Error(l + "upper")
  case neq(newAlphabet(false, false, true), puncts):
    t.Error(l + "puncts")
  case neq(newAlphabet(true, true, false), append(lower, upper...)):
    t.Error(l + "lower | upper")
  case neq(newAlphabet(true, false, true), append(lower, puncts...)):
    t.Error(l + "lower | puncts")
  case neq(newAlphabet(false, true, true), append(upper, puncts...)):
    t.Error(l + "upper | puncts")
  case neq(newAlphabet(true, true, true), append(lower, append(upper, puncts...)...)):
    t.Error(l + "all")
  }
}

func TestGetAlpha(t *testing.T) {
  if neq(GetAlpha(true, true, true, string(newAlphabet(true, true, false))), puncts) {
    t.Error("GetAlpha!")
  }
}
