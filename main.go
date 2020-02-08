package main

func index(elements []int, predicate func(int) bool) int {
	for ind, element := range elements {
		if predicate(element) {
			return ind
		}
	}
	return -1
}
func all(elements []int,predicate func(int)bool)bool{
	return index(elements, func(i int) bool {
		return !predicate(i)
	}) == -1
	}

func any(elements []int, predicate func(int) bool) bool {
	return index(elements,predicate) != -1
}

func none(elements []int, predicate func(int)bool)bool  {
	return index(elements,predicate) == -1
}


func find(elements []int, predicate func(int) bool) int {
	return elements[index(elements, predicate)]
}

func main() {
}


