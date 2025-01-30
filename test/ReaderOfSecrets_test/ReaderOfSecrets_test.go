package ReaderOfSecrets_test

import (
	"github.com/maytilex/ReaderOfSecrets"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestReaderOfSecrets(t *testing.T) {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	randStr := b.String() // Random characters for dirname & filename

	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	dirPath := pwd + "/" + randStr + "/"

	err = os.Mkdir(dirPath, 0777)
	assert.Nil(t, err) // Failed to create directory for test
	if err != nil {
		t.Error("asddsa")
	}

	filePath := dirPath + randStr

	create, err := os.Create(filePath)
	assert.Nil(t, err) // Failed to create file in new directory

	_, err = create.Write([]byte(randStr))
	assert.Nil(t, err) // Failed to write in new file

	t.Run("", func(t *testing.T) {
		reader := ReaderOfSecrets.NewReader(dirPath, true, true)
		secret, ok := reader.Secret(randStr)
		assert.Equal(t, ok, true)
		assert.Equal(t, string(secret), randStr)
	})

	err = os.Remove(filePath)
	assert.Nil(t, err) // Failed to Remove files in dir

	err = os.Remove(dirPath)
	assert.Nil(t, err) // Failed remove directory
}
