package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't Update non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search word in Dictionary
func (mydict Dictionary) Search(word string) (string, error) {
	value, exists := mydict[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the Dictionary
func (mydict Dictionary) Add(word, definition string) error {
	_, err := mydict.Search(word)
	switch err {
	case errNotFound:
		mydict[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

// Update a Word
func (mydict Dictionary) Update(word, definition string) error {
	_, err := mydict.Search(word)
	switch err {
	case nil:
		mydict[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a Word
func (mydict Dictionary) Delete(word string) {
	delete(mydict, word)
}
