package ReaderOfSecrets

import (
	"github.com/maytilex/ReaderOfSecrets/src/DirectoryManager"
	"github.com/maytilex/ReaderOfSecrets/src/FileManager"
)

type flags struct {
	immediatelyReading bool
	contentZero        bool
}

type Secrets map[string][]byte

type ReaderOfSecrets struct {
	flags
	dManager DirectoryManager.Directory
	fManager FileManager.FileManager
	secrets  Secrets
}
