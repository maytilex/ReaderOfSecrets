package DirectoryManager_test

import (
	"github.com/maytilex/ReaderOfSecrets/src/DirectoryManager"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestDirectoryManager(t *testing.T) {
	checkCorrectPath(t)
	checkFailedPath(t)
	checkRandomPath(t)
}

func checkCorrectPath(t *testing.T) {
	dManager := DirectoryManager.NewDirectoryManager("")
	assert.Equal(t, dManager.GetDirName(), "/run/secrets/") // "" empty path changes on "/run/secrets/"

	dManager = DirectoryManager.NewDirectoryManager("/run/secrets")
	assert.Equal(t, dManager.GetDirName(), "/run/secrets/") // add "/"

	dManager = DirectoryManager.NewDirectoryManager("/")
	assert.Equal(t, dManager.GetDirName(), "/") // "/" is directory, nothing needs to be done
}

func checkRandomPath(t *testing.T) {
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

	dManger := DirectoryManager.NewDirectoryManager(dirPath)

	for _, entry := range dManger.GetEntry() {
		assert.Equal(t, entry.Name(), randStr)
	}

	err = os.Remove(filePath)
	assert.Nil(t, err) // Failed to Remove files in dir

	err = os.Remove(dirPath)
	assert.Nil(t, err) // Failed remove directory
}

func checkFailedPath(t *testing.T) {

}
