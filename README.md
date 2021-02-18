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