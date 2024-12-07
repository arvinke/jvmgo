package classpath

import (
    "archive/zip"
    "errors"
    "io"
    "path/filepath"
)

type ZipEntry struct {

    absPath string
}

func newZipEntry(path string) *ZipEntry {
    absPath, err := filepath.Abs(path)
    if err != nil {
        panic(err)
    }

    return &ZipEntry{absPath}
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
    r, err := zip.OpenReader(e.absPath)
    if err != nil {
        return nil, nil, err
    }

    defer r.Close()
    for _, f := range r.File {
        if f.Name == className {
            rc, err := f.Open()
            if err != nil {
                return nil, nil, err
            }

            defer r.Close()
            data, err := io.ReadAll(rc)
            if err != nil {
                return nil, nil, err
            }

            return data, e, nil
        }
    }

    return nil, nil, errors.New("class not found: " + className)
}

func (e *ZipEntry) String() string {
    return e.absPath
}
