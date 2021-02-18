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

<p xmlns:cc="http://creativecommons.org/ns#" xmlns:dct="http://purl.org/dc/terms/"><a property="dct:title" rel="cc:attributionURL" href="https://github.com/marekkowalczyk/filenames">filenames</a> by <a rel="cc:attributionURL dct:creator" property="cc:attributionName" href="https://github.com/marekkowalczyk">Marek Kowalczyk</a> is licensed under <a href="http://creativecommons.org/licenses/by-sa/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">CC BY-SA 4.0<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1"><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/sa.svg?ref=chooser-v1"></a></p>
