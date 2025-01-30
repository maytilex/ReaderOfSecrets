package FileManager_test

import (
	"github.com/maytilex/ReaderOfSecrets/src/FileManager"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestFileManager(t *testing.T) {
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

	filePath := dirPath + randStr

	create, err := os.Create(filePath)
	assert.Nil(t, err) // Failed to create file in new directory

	_, err = create.Write([]byte(randStr))
	assert.Nil(t, err) // Failed to write in new file

	dirEntry, err := os.ReadDir(dirPath)
	assert.Nil(t, err)

	realSecret := map[string][]byte{randStr: []byte(randStr)}

	fManager := FileManager.NewFileManager()

	readSecret := fManager.ReadEntry(dirEntry, dirPath, false)
	assert.Equal(t, realSecret, readSecret)

	err = os.Remove(filePath)
	assert.Nil(t, err) // Failed to Remove files in dir

	err = os.Remove(dirPath)
	assert.Nil(t, err) // Failed remove directory
}
