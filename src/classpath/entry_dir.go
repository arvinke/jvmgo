package classpath

import (
    "os"
    "path/filepath"
)

type DirEntry struct {

    absDir string
}

func newDirEntry(path string) *DirEntry {
    absDir, err := filepath.Abs(path)
    if err != nil {
        panic(err)
    }

    return &DirEntry{absDir}
}

func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
    fileName := filepath.Join(e.absDir, className)
    data, err := os.ReadFile(fileName)
    return data, e, err
}

func (e *DirEntry) String() string {
    return e.absDir
}
