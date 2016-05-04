package goga

// The dropable interface is used to clean up GL objects.
// Use the Drop() function to drop a range of objects.
type Dropable interface {
	Drop()
}

// Drops given GL objects.
// Objects must implement the Dropable inteface.
func Drop(objects []Dropable) {
	for _, obj := range objects {
		obj.Drop()
	}
}
