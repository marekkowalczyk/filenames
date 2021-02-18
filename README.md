# filenames

A `go` library to intuitively manipulate file names.

Functions:

- `FileBase` gets the file name with extension of filepath but without the directory
- `FileDirectory` gets the directory of filepath without the file name
- `FileExtension` gets the file extension of filepath without the leading dot
- `FileName` gets the file name of filepath without the directory or extension

Examples of usage:

- `FilePath:	 /Users/user/directory/foo.txt`
- `FileDirectory:	 /Users/user/directory`
- `FileBase:	 foo.txt`
- `FileName:	 foo`
- `FileExtension:	 txt`

````
// Sample program

package main

import (
	"fmt"

	"github.com/marekkowalczyk/filenames"
)

func main() {

	filepath := "/Users/user/directory/foo.txt"

	fmt.Println("FilePath:\t", filepath)

	fmt.Println("FileDirectory:\t", filenames.FileDirectory(filepath))
	fmt.Println("FileBase:\t", filenames.FileBase(filepath))
	fmt.Println("FileName:\t", filenames.FileName(filepath))
	fmt.Println("FileExtension:\t", filenames.FileExtension(filepath))

}
````