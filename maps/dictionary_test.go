package maps

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDictionaryErr(t *testing.T) {
	Convey("Should generate error", t, func() {
		customErr := DictionaryErr("custom error")
		So(customErr, ShouldBeError, errors.New("custom error"))
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	Convey("Should return definition of known word", t, func() {
		got, err := dictionary.Search("test")
		expectedDefinition := "this is just a test"
		So(err, ShouldBeNil)
		So(got, ShouldEqual, expectedDefinition)
	})

	Convey("Should return error for unknown word", t, func() {
		got, err := dictionary.Search("notInDict")
		So(got, ShouldEqual, "")
		So(err, ShouldBeError, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	Convey("Should add new word to dictionary", t, func() {
		dictionary := Dictionary{}
		addErr := dictionary.Add("42", "meaning of life, universe and everything")
		So(addErr, ShouldEqual, nil)
		got, err := dictionary.Search("42")
		So(err, ShouldEqual, nil)
		So(got, ShouldEqual, "meaning of life, universe and everything")
	})

	Convey("Should return error if word exists", t, func() {
		word := "42"
		definition := "meaning of life, universe and everything"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		So(err, ShouldEqual, ErrWordExist)
	})
}

func TestUpdate(t *testing.T) {
	Convey("Should update definition of word", t, func() {
		word := "42"
		definition := "meaning of life, universe and everything"
		dictionary := Dictionary{word: definition}
		newDefinition := "the meaning of life, universe and everything"
		updateErr := dictionary.Update(word, newDefinition)
		So(updateErr, ShouldEqual, nil)
		got, err := dictionary.Search(word)
		So(err, ShouldEqual, nil)
		So(got, ShouldEqual, newDefinition)
	})

	Convey("Should return error if word doesn't exist", t, func() {
		dictionary := Dictionary{}
		updateErr := dictionary.Update("43", "")
		So(updateErr, ShouldEqual, ErrUpdateFailed)
	})
}

func TestDelete(t *testing.T) {
	Convey("Should delete given word", t, func() {
		word := "42"
		definition := "meaning of life, universe and everything"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		So(err, ShouldBeNil)
		_, searchErr := dictionary.Search(word)
		So(searchErr, ShouldEqual, ErrNotFound)
	})

	Convey("Should return error if the word doesn't exist", t, func() {
		dictionary := Dictionary{}
		err := dictionary.Delete("43")
		So(err, ShouldEqual, ErrDeleteFailed)
	})
}
