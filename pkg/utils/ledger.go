package utils

import (
	"crypto/rand"

	"github.com/aidarkhanov/nanoid"
)

func CustomAlphabet(alphabet string, size int) func() (string, error) {
	// generateBytesBuffer returns random bytes buffer
	generateBytesBuffer := func(step int) ([]byte, error) {
		buffer := make([]byte, step)
		if _, err := rand.Read(buffer); err != nil {
			return nil, err
		}

		return buffer, nil
	}

	// id, err := nanoid.Format(generateBytesBuffer, alphabet, size)
	// if err != nil {
	// 	panic(err)
	// }

	generatorFn := func() (string, error) {
		id, err := nanoid.Format(generateBytesBuffer, alphabet, size)

		return id, err
	}

	// fmt.Println(id)
	return generatorFn
}

func NewAccountIdGenerator(size int) func() (string, error) {
	return CustomAlphabet("0123456789", size)
}

func DeleteArrayItem[T any](index int32, array []T) ([]T, bool) {
	arrayLen := len(array)
	// newArray := make([]T, 0)

	if index < 0 || index >= int32(arrayLen) {
		return array, false
	}

	array = append(array[:index], array[index+1:]...)

	return array, true
}

func GetArrayItemIndex[T comparable](item T, arr []T) (int, bool) {
	for i, v := range arr {
		if item == v {
			return i, true
		}
	}

	return -1, false
}
