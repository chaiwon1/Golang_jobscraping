package mydict

import "errors"

// Dictionary type
type Sajeon map[string]string

// Errors
var (
	errNotFound   = errors.New("No such word")
	errAlready    = errors.New("The word already exists")
	errCantUpdate = errors.New("Can't Update")
)

// Search method
func (s Sajeon) Search(word string) (string, error) {
	definition, exists := s[word]
	if exists {
		return definition, nil
	}
	return "", errNotFound
}

// Add method(using if)
func (s Sajeon) Add(word, definition string) error {
	_, err := s.Search(word)
	if err == errNotFound {
		s[word] = definition
	} else if err == nil {
		return errAlready
	}
	return nil
}

// Add method(using case)
func (s Sajeon) Add(word, definition string) error {
	_, err := s.Search(word)
	switch err {
	case errNotFound:
		s[word] = definition
	case nil:
		return errAlready
	}
	return nil
}

// Update method
func (s Sajeon) Update(word, newdef string) error {
	_, err := s.Search(word)
	if err == nil {
		s[word] = newdef
	} else if err == errNotFound {
		return errCantUpdate
	}
	return nil
}

// Delete method
func (s Sajeon) Delete(word string) error {
	delete(s, word)
	return nil
}
