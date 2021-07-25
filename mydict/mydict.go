package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errWordExists = errors.New("exists!")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}
	*/

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)

	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}
	*/

	switch err {
	case errNotFound:
		return errNotFound
	case nil:
		d[word] = def
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}
	*/

	switch err {
	case errNotFound:
		return errNotFound
	case nil:
		delete(d, word)
	}

	return nil
}
