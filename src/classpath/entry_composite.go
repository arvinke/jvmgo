package classpath

import (
    "errors"
    "strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
    compositEntry := []Entry{}

    for _, path := range strings.Split(pathList, pathListSeparator) {
        entry := newEntry(path)
        compositEntry = append(compositEntry, entry)
    }

    return compositEntry
}

func (ce CompositeEntry) readClass(className string) ([]byte, Entry, error) {
    for _, entry := range ce {
        data, from, err := entry.readClass(className)
        if err == nil {
            return data, from, nil
        }
    }

    return nil, nil, errors.New("class not found: " + className)
}

func (ce CompositeEntry) String() string {
    strs := make([]string, len(ce))

    for i, entry := range ce {
        strs[i] = entry.String()
    }

    return strings.Join(strs, pathListSeparator)
}
