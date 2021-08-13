package mydict

import "errors"

/// Dictionary
type Dictionary map[string]string

var errNotFound = errors.New("error: not found")
var errCantUpdate = errors.New("error: Can't update non-existing word")
var errWordExists = errors.New("error: word exists already")

func (d Dictionary) Search(word string) (string, error) {
	value, isExists := d[word]

	if isExists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
