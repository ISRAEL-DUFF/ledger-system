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
