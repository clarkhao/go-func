package slices

// Slices is a slice type where type of item T is primitive type
// ~represent a constraint of S which must be []E
type Slices[S ~[]E, E comparable] []E

// NewSlice is a factory function to create a Slices instance
func NewSlice(s []string) Slices[[]string, string] {
	list := make(Slices[[]string, string], 0)
	for _, item := range s {
		list = append(list, item)
	}
	return list
}

// ItemEqual is a method of Slices type
// which determine if all items of list is equals to all items of s
// not considering the order of the item inside slice
// return true or false
// i.e. [1 2 3] equals [2 1 3]
func (list Slices[S, E]) ItemEqual(s Slices[S, E]) bool {
	if len(list) != len(s) {
		return false
	}
	m1 := list.ToMap()
	m2 := s.ToMap()
	if len(m1) != len(m2) {
		return false
	}
	for key, value := range m1 {
		if m2[key] != value {
			return false
		}
	}
	return true
}

// ToMap is a method of Slices which turn Slices to map
// return a new Map
func (list Slices[S, E]) ToMap() map[E]int {
	m := make(map[E]int)
	for _, item := range list {
		m[item]++
	}
	return m
}
