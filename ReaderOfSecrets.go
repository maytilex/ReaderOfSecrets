package ReaderOfSecrets

import (
	"github.com/maytilex/ReaderOfSecrets/src/DirectoryManager"
	"github.com/maytilex/ReaderOfSecrets/src/FileManager"
)

func NewReader(directoryWithSecrets string, immediatelyReading bool, contentZero bool) (reader *ReaderOfSecrets) {
	reader = &ReaderOfSecrets{
		flags: flags{
			immediatelyReading: immediatelyReading,
			contentZero:        contentZero,
		},
		dManager: DirectoryManager.NewDirectoryManager(directoryWithSecrets),
		fManager: FileManager.NewFileManager(),
		secrets:  map[string][]byte{},
	}

	switch false {
	case reader.flags.immediatelyReading, reader.dManager.GetStatus(), reader.fManager.GetStatus():
		return
	}

	reader.secrets = reader.fManager.ReadEntry(reader.dManager.GetEntry(), reader.dManager.GetDirName(), contentZero)

	return
}

func (r *ReaderOfSecrets) GetError() error {
	switch false {
	case r.dManager.GetStatus():
		return r.dManager.GetError()
	case r.fManager.GetStatus():
		return r.fManager.GetError()
	default:
		return nil
	}
}

func (r *ReaderOfSecrets) Secret(NameSecretAsFileName string) (secret []byte, ok bool) {
	secret, ok = r.secrets[NameSecretAsFileName]
	return
}

func (r *ReaderOfSecrets) AllSecrets() map[string][]byte {
	return r.secrets
}
