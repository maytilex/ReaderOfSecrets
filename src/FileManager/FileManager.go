package FileManager

import (
	"os"
)

type FileManager struct {
	ok  bool
	err error
	msg string
}

func NewFileManager() FileManager {
	return FileManager{
		ok:  true,
		err: nil,
		msg: "",
	}
}

func (f *FileManager) GetStatus() bool {
	return f.ok
}

func (f *FileManager) GetError() error {
	return f.err
}

func (f *FileManager) GetMessage() string {
	return f.msg
}

func (f *FileManager) switchStatus() {
	f.ok = false
}

func (f *FileManager) ReadEntry(entry []os.DirEntry, dirName string, contentZero bool) map[string][]byte {
	arrBytes := map[string][]byte{}

	for _, fileEntry := range entry {
		if !fileEntry.IsDir() {

			bytes, err := os.ReadFile(dirName + fileEntry.Name())
			if err != nil {
				f.switchStatus()
				f.err = err
				f.msg = "Error reading file from secrets directory. Check if the secret exists."
				return nil
			} else if contentZero == false && len(bytes) == 0 {
				f.switchStatus()
				f.err = err
				f.msg = "The content of the secret is 0"
				return nil
			}
			arrBytes[fileEntry.Name()] = bytes
		}
	}
	return arrBytes
}
