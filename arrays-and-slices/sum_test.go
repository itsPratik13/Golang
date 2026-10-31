package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Arrays used", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d but wanted %d, given %v", got, want, numbers)
		}
	})
	t.Run("Slices used",func(t *testing.T) {
		slice:=[]int{1,2,3}
		got:=Sum(slice);
		want:=6
		if got!=want{
			t.Errorf("got %d but want %d , given %v ",got,want,slice)
		}

	})
}
