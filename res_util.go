package goga

import (
	"errors"
)

// Finds and returns a Tex resource.
// If not found or when the resource is of wrong type, an error will be returned.
func GetTex(name string) (*Tex, error) {
	res := GetResByName(name)

	if res == nil {
		return nil, errors.New("Resource not found")
	}

	tex, ok := res.(*Tex)

	if !ok {
		return nil, errors.New("Resource was not of type *Tex")
	}

	return tex, nil
}

// Finds and returns a Ply resource.
// If not found or when the resource is of wrong type, an error will be returned.
func GetPly(name string) (*Ply, error) {
	res := GetResByName(name)

	if res == nil {
		return nil, errors.New("Resource not found")
	}

	ply, ok := res.(*Ply)

	if !ok {
		return nil, errors.New("Resource was not of type *Ply")
	}

	return ply, nil
}
