package goga

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// A generic resource.
// Must be cast to appropriate type.
// The name is the file name and must be unique.
type Res interface {
	GetName() string
	SetName(string)
	GetPath() string
	SetPath(string)
	GetExt() string
	SetExt(string)
}

// Resource loader interface.
// The loader accepts files by file extension.
// and loads them if accepted.
type ResLoader interface {
	Load(string) (Res, error)
	Ext() string
}

var (
	resloader []ResLoader
	resources []Res
)

// Adds a loader.
// If a loader with the same file extension exists already, false will be returned.
func AddLoader(loader ResLoader) bool {
	ext := strings.ToLower(loader.Ext())

	for _, l := range resloader {
		if strings.ToLower(l.Ext()) == ext {
			return false
		}
	}

	resloader = append(resloader, loader)

	return true
}

// Removes a loader.
// Returns false if loader could not be found.
func RemoveLoader(loader ResLoader) bool {
	for i, l := range resloader {
		if l == loader {
			resloader = append(resloader[:i], resloader[i+1:]...)
			return true
		}
	}

	return false
}

// Removes a loader by file extension.
// Returns false if loader could not be found.
func RemoveLoaderByExt(ext string) bool {
	ext = strings.ToLower(ext)

	for i, l := range resloader {
		if strings.ToLower(l.Ext()) == ext {
			resloader = append(resloader[:i], resloader[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all loaders.
func RemoveAllLoaders() {
	resloader = make([]ResLoader, 0)
}

// Returns a loader by file extension.
// If not found, nil will be returned.
func GetLoaderByExt(ext string) ResLoader {
	ext = strings.ToLower(ext)

	for _, l := range resloader {
		if strings.ToLower(l.Ext()) == ext {
			return l
		}
	}

	return nil
}

// Loads a resource by file path.
// If no loader is present for given file, an error will be returned.
// If the loader fails to load the resource, an error will be returned.
// If the resource name exists already, an error AND the resource will be returned.
// This is allows to cleanup on failure.
func LoadRes(path string) (Res, error) {
	ext := filepath.Ext(path)

	if len(ext) > 0 {
		ext = ext[1:]
	}

	loader := GetLoaderByExt(ext)

	if loader == nil {
		return nil, errors.New("No loader available for file extension " + ext)
	}

	res, err := loader.Load(path)

	if err != nil {
		return nil, err
	}

	res.SetName(filepath.Base(path))
	res.SetPath(path)
	res.SetExt(ext)

	for _, r := range resources {
		if r.GetName() == res.GetName() {
			return res, errors.New("Resource with file name " + res.GetName() + " exists already")
		}
	}

	resources = append(resources, res)

	return res, nil
}

// Loads all files from given folder path.
// If a loader is missing or fails to load the resource, an error will be returned.
// All resources will be kept until an error occures.
func LoadResFromFolder(path string) error {
	dir, err := ioutil.ReadDir(path)

	if err != nil {
		return err
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if _, err := LoadRes(filepath.Join(path, file.Name())); err != nil {
			return err
		}
	}

	return nil
}

// Returns a resource by name or nil, if not found.
func GetResByName(name string) Res {
	for _, r := range resources {
		if r.GetName() == name {
			return r
		}
	}

	return nil
}

// Returns a resource by path or nil, if not found.
func GetResByPath(path string) Res {
	for _, r := range resources {
		if r.GetPath() == path {
			return r
		}
	}

	return nil
}

// Removes a resource by name.
// Returns false if resource could not be found.
func RemoveResByName(name string) bool {
	for i, r := range resources {
		if r.GetName() == name {
			resources = append(resources[:i], resources[i+1:]...)
			return true
		}
	}

	return false
}

// Removes a resource by path.
// Returns false if resource could not be found.
func RemoveResByPath(path string) bool {
	for i, r := range resources {
		if r.GetPath() == path {
			resources = append(resources[:i], resources[i+1:]...)
			return true
		}
	}

	return false
}

// Removes all resources.
func RemoveAllRes() {
	resources = make([]Res, 0)
}
