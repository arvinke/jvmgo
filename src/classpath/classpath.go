package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
    bootClassPath Entry
    extClassPath Entry
    userClassPath Entry
}

func Parse(jreOption, cpOption string) *ClassPath {
    cp := &ClassPath{}
    cp.parseBootAndExtClasspath(jreOption)
    cp.parseUserClasspth(cpOption)

    return cp
}

func exists(jreOption string) bool {
    if _, err := os.Stat(jreOption); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }

    return true
}

func  getJreDir(jreOption string) string {
    if jreOption != "" && exists(jreOption) {
        return jreOption
    }

    if exists("./jre") {
        return ".jre"
    }

    if jh := os.Getenv("JAVA_HOME"); jh != "" {
        return jh
    }

    panic("Can not find jre folder! ")
}

func (cp *ClassPath) parseBootAndExtClasspath(jreOption string)  {
    jreDir := getJreDir(jreOption)

    // jre/lib/*
    jreLibPath := filepath.Join(jreDir, "lib", "*")
    cp.bootClassPath = newWildcardEntry(jreLibPath)

    // jre/lib/ext/*
    jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
    cp.extClassPath = newWildcardEntry(jreExtPath)
}

func (cp *ClassPath) parseUserClasspth(cpOption string) {
    if cpOption == "" {
        cpOption = "."
    }

    cp.userClassPath = newEntry(cpOption)
}

func (cp *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
    className = className + ".class"

    if data, entry, err := cp.bootClassPath.readClass(className); err == nil {
        return data, entry, err
    }

    if data, entry, err := cp.extClassPath.readClass(className); err == nil {
        return data, entry, err
    }

    return cp.userClassPath.readClass(className)
}

func (cp *ClassPath) String() string {
    return cp.userClassPath.String()
}
