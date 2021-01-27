package dictionary

type dictionary map[string]string
type dictionaryError string
const (
	ErrNotFound = dictionaryError("No such word in the dictionary")
	ErrWordExists = dictionaryError("This word already exists in the dictionary")
)

func (de dictionaryError) Error() string {
	return string(de)
}

func (d dictionary) Search(word string) (description string, err error) {
	description, ok := d[word]

	if !ok {
		err = ErrNotFound
	}

	return description, err
}

func (d dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}

func (d dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if err != nil {
		return ErrNotFound
	}

	delete(d, word)
	return nil
}
