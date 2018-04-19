package util

// Set A basic implementation of a set, using a map
type Set struct {
	data map[interface{}]struct{}
}

// Add Adds the value to the set, returning true if it's a new value
func (set *Set) Add(val interface{}) bool {
	if _, ok := set.data[val]; !ok {
		set.data[val] = struct{}{}
		return true
	}
	return false
}

// Contains Returns true if the value is contained in the set
func (set *Set) Contains(val interface{}) bool {
	_, ok := set.data[val]
	return ok
}

// Remove Removes the value from the set if it exists and returns true if it did
func (set *Set) Remove(val interface{}) bool {
	if val, ok := set.data[val]; ok {
		delete(set.data, val)
		return true
	}
	return false
}

// Len Returns the size of the set
func (set *Set) Len() int {
	return len(set.data)
}

// Slice Returns the values as a slice
func (set *Set) Slice() *[]interface{} {
	arr := make([]interface{}, len(set.data))
	i := 0
	for key := range set.data {
		arr[i] = key
		i++
	}
	return &arr
}

// FillSlice Fills the provided slice with the values from the set
func (set *Set) FillSlice(slice *[]interface{}) {
	if slice == nil {
		panic("Cannot fill nil.")
	}
	if len(*slice) < len(set.data) {
		panic("Slice size is less than set size.")
	}
	i := 0
	for key := range set.data {
		(*slice)[i] = key
		i++
	}
}
