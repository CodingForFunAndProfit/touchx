# touchx

Creates/touches a file in folder. If folder/s do not exist they will be created first.

While learning Go I created many repositories and always had to create a folderstructure. For example:
mkdir cmd
mkdir server
touch cmd/server/main.go

I could reduce it to:
mkdir -p cmd/server
touch cmd/server/main.go

or even in one line:
mkdir -p cmd/server && touch cmd/server/main.go

To save a bit time and learn some more with go I wrote this program, so I can do the above with one line, like:

touchx cmd/server/main.go

cases to test:

1. file in current dir does not exist.
2. file in current dir does exist.
3. path/file, file does not exist in folder path
4. path/file, path does not exist

better:

1. file does exist (./file, path/file, path/subpath/file) and every path -> exit
2. else: file or path does not exist
    1. path(s) exists but not the file -> create file (subpath could not exist!)
    2. path(s) does not exist (and so does the file) -> create path(s) && file
3. file exists as directory
