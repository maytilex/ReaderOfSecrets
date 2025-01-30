package DirectoryManager

import "os"

type Directory struct {
	name  string
	entry []os.DirEntry
	ok    bool
	err   error
	msg   string
}
