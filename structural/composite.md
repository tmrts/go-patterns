# Composite Pattern
Composite structural pattern allows composing objects into a tree-like structure and work with the it as if it was a singular object.

## Interface
```go
package main

import "fmt"

type file struct {
    name string
}

func (f *file) search(keyword string) {
    fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
    return f.name
}
```

## Implementation
`search` function will operate applies to both files and folders. For a file, it will just look into the contents of the file; for a folder, it will go through all files of that folder to find that keyword.

```go
package main

import "fmt"

type folder struct {
    components []component
    name       string
}

func (f *folder) search(keyword string) {
    fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
    for _, composite := range f.components {
        composite.search(keyword)
    }
}

func (f *folder) add(c component) {
    f.components = append(f.components, c)
}
```

## Usage
```go
file1 := &file{name: "File1"}
file2 := &file{name: "File2"}
file3 := &file{name: "File3"}

folder1 := &folder{
	name: "Folder1",
}

folder1.add(file1)

folder2 := &folder{
	name: "Folder2",
}
folder2.add(file2)
folder2.add(file3)
folder2.add(folder1)

folder2.search("rose")
```
