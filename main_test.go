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
		{"All elem is true", []int{1, 2, 3, 4}, true, func(i int) bool {
			if i == 1 {
				return true
			}
			return false
		}},
	}
	for _, tt := range tests {
		if got := all(tt.elements, tt.predicate); got != tt.want {
			t.Errorf("Test for %v all() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_find(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		want      int
		predicate func(int) bool
	}{
		{"Found elem", []int{1, 2, 3, 4}, 2, func(i int) bool {
			if i == 2 {
				return true
			}
			return false
		}},
	}

	for _, tt := range tests {
		if got := find(tt.elements, tt.predicate); got != tt.want {
			t.Errorf("all() = %v, want %v", got, tt.want)
		}
	}
}

