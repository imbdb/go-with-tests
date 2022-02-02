package maps

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound     = DictionaryErr("word not found")
	ErrUpdateFailed = DictionaryErr("update failed, word not found")
	ErrWordExist    = DictionaryErr("word already exist in dictionary")
	ErrDeleteFailed = DictionaryErr("delete failed, word not found")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {

	_, err := d.Search(word)
	switch err {
	case nil:
		return ErrWordExist
	case ErrNotFound:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrUpdateFailed
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrDeleteFailed
	default:
		return err
	}
	return nil
}
