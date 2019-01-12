package main

import "testing"

func TestAdd(t *testing.T) {
	expected := 5
  actual := myAdd(3, 2)
  if (expected != actual) {
    t.Error("Nope")
  }else {
	  t.Log("第一个测试通过了") //记录一些你期望记录的信息
  }
}
