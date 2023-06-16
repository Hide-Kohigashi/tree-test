package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

// Comments

func readdir(name string) ([]string, []string, error) {
    fp, err := os.Open(name)
    if err != nil {
        return nil, nil, err
    }

    list, err := fp.Readdir(-1)
    fp.Close()
    if err != nil {
        return nil, nil, err
    }

    dirs, files := []string{}, []string{}
    for _, v := range list {
        if v.IsDir() {
            dirs = append(dirs, v.Name())
        } else {
            files = append(files, v.Name())
        }
    }

    return dirs, files, nil
}

func tree(indent, path string) error {
    dirs, files, err := readdir(path)
    if err != nil {
        return err
    }

    for i, v := range files {
        s := indent + " ├─"
        if len(dirs) == 0 && i == len(files)-1 {
            s = indent + " └─"
        }

        fmt.Printf("%s %s\n", s, v)
    }

    for i, v := range dirs {
        s := indent + " ├─"
        a := " │ "
        if i == len(dirs)-1 {
            s = indent + " └─"
            a = "   "
        }

        fmt.Printf("%s %s\n", s, v)

        if err := tree(indent+a, filepath.Join(path, v)); err != nil {
            return err
        }
    }

    return nil
}

func main() {
    const Root = "."

    fmt.Printf(" %s\n", Root)

    path, err := filepath.Abs(Root)
    if err != nil {
        log.Fatal(err)
    }

    if err := tree("", path); err != nil {
        log.Fatal(err)
    }
}
