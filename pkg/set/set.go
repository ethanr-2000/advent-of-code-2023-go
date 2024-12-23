package set

type Set map[string]bool

func NewSetFromSlice(slice []string) Set {
	set := make(Set)
	for _, v := range slice {
		set.Add(v)
	}
	return set
}

func (s Set) Pop() string {
	for k := range s {
		s.Delete(k)
		return k
	}
	panic("Tried to pop empty set")
}

func (s Set) Delete(value string) {
	delete(s, value)
}

func (s Set) Add(value string) {
	s[value] = true
}

// Union: Return a new Set that is the union of two Sets
func (s Set) Union(other Set) Set {
	result := make(Set)
	for v := range s {
		result.Add(v)
	}
	for v := range other {
		result.Add(v)
	}
	return result
}

// Intersection: Return a new Set that is the intersection of two Sets
func (s Set) Intersection(other Set) Set {
	result := make(Set)
	for v := range s {
		if _, exists := other[v]; exists {
			result.Add(v)
		}
	}
	return result
}

// Convert Set to Slice
func (s Set) ToSlice() []string {
	result := []string{}
	for v := range s {
		result = append(result, v)
	}
	return result
}
