package main

import "testing"

func Test_index(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		want      int
		predicate func(int) bool
	}{
		{"With index", []int{1, 2, 3, 4}, 0, func(i int) bool {
			if i == 1 {
				return true
			}
			return false
		}},
		{"Without index", []int{1, 2, 3, 4}, -1, func(i int) bool {
			if i == 5 {
				return true
			}
			return false
		}},
	}
	for _, tt := range tests {
		if got := index(tt.elements, tt.predicate); got != tt.want {
			t.Errorf("index() = %v, want %v", got, tt.want)
		}
	}
}
func Test_any(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		want      bool
		predicate func(int) bool
	}{
		{"Any elem is true", []int{1, 2, 3, 4}, true, func(i int) bool {
			if i == 1 {
				return true
			}
			return false
		}},

		{"Any elem is false", []int{1, 2, 3, 4}, false, func(i int) bool {
			if i == 5 {
				return true
			}
			return false
		}},
	}
	for _, tt := range tests {
		if got := any(tt.elements, tt.predicate); got != tt.want {
			t.Errorf("any() = %v, want %v", got, tt.want)
		}
	}
}

func Test_none(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		want      bool
		predicate func(int) bool
	}{
		{"None elem is false", []int{1, 2, 3, 4}, false, func(i int) bool {
			if i == 1 {
				return true
			}
			return false
		}},

		{"None elem is true", []int{1, 2, 3, 4}, true, func(i int) bool {
			if i == 5 {
				return true
			}
			return false
		}},
	}
	for _, tt := range tests {
		if got := none(tt.elements, tt.predicate); got != tt.want {
			t.Errorf("none() = %v, want %v", got, tt.want)
		}
	}
}

func Test_all(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		want      bool
		predicate func(int) bool
	}{
		{"All elem is true", []int{4,4,4,4}, true, func(i int) bool {
			return i == 4
		}},
		{"All elem is false",[]int{1,2,3,4,5,6,7,8},false, func(i int) bool {
				return i == 8
		}},
	}
	for _, tt := range tests {
		if got := all(tt.elements, tt.predicate); got != tt.want {
			t.Errorf("Test for %v all() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
func TestFind(t *testing.T) {

	if find([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i == 2
	}) != 2 {
		t.Error("...")
	}
	func() {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("panic!")
			}
		}()
		find([]int{1, 2, 3}, func(i int) bool {
			return i == 6
		})
	}()
	func() {
		defer func() {
			err := recover()
			if err == nil {
				t.Error("panic")
			}
		}()
		find([]int{1, 2, 3}, func(i int) bool {
			return i == -1
		})
	}()

}

