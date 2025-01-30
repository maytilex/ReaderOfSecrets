package DirectoryManager

import "os"

func NewDirectoryManager(directoryWithSecrets string) Directory {
	directory := Directory{
		name:  directoryWithSecrets,
		entry: nil,
		ok:    true,
		err:   nil,
		msg:   "",
	}

	directory.checkName()

	directory.checkDirectory()

	return directory
}

func (d *Directory) GetDirName() string {
	return d.name
}

func (d *Directory) GetEntry() []os.DirEntry {
	return d.entry
}

func (d *Directory) GetStatus() bool {
	return d.ok
}

func (d *Directory) GetError() error {
	return d.err
}

func (d *Directory) GetMessage() string {
	return d.msg
}

func (d *Directory) switchStatus() {
	d.ok = false
}

func (d *Directory) checkName() {
	if d.name == "" {
		d.name = "/run/secrets/"
	} else if d.name[len(d.name)-1] != '/' {
		d.name += "/"
	}
}

func (d *Directory) checkDirectory() {
	d.entry, d.err = os.ReadDir(d.name)
	if d.err != nil {
		d.switchStatus()
		d.msg = `Error reading files from directory. Check if directory with secrets exists.`
	} else if len(d.entry) == 0 {
		d.switchStatus()
		d.msg = `Directory with secrets is EMPTY`
	}
}
