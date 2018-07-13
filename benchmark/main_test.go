package main

import "testing"

func TestPrint01(t *testing.T)  {
  if print01(100) != "100" {
    t.Fatal("error")
  }
}

func TestPrint02(t *testing.T)  {
  if print02(int64(100)) != "100" {
    t.Fatal("error")
  }
}

func TestPrint03(t *testing.T)  {
  if print03(100) != "100" {
    t.Fatal("error")
  }
}
