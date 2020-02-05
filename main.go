package main


func index(elements []int, predicate func(int) bool) int {
	for ind, element := range elements {
		if predicate(element) {
			return ind
		}
	}
	return -1
}

func any(elements []int, predicate func(int) bool) bool {
	if index(elements,predicate) != -1{
		return true
	}
	return false
}

func none(elements []int, predicate func(int)bool)bool  {
	if index(elements,predicate) == -1{
		return true
	}
	return false
}

func all(elements []int,predicate func(int)bool)bool{
	if index(elements,predicate) == -1{
		return false
	}
	return true
}

func find(elements []int, predicate func(int) bool) int {
	return elements[index(elements, predicate)]
}

func main() {
	value := []int{1, 2, 3, 4}
	err := find(value, func(i int) bool {
		return true
	})
	nil := 0
	if !(err != nil) {
		panic("Error: did not find anything")
	}

}
