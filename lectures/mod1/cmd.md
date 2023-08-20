# Basic Linux commands

How files are organized in Linux?
---
- file contains data
- directory or folder contains files or/and subdirectories or nothing
  - directory is a special file
- all files and directories are organized as a tree
- every file or directory has a unique path from the root of the tree 
  - with zero or more intermediate folders sit in between 


Special directories
---
| symbol | directory |
| --- | --- |
| / | root of all files and directories |
| . | current directory, or present working directory |
| .. | parent folder |
| - | previous folder in the browsing history |
| ~ | home folder of the current user |
| ~someone | someone's home folder |


Absolute path vs. relative path
---
- relative path begins from the current folder
- absolute path begins from the root


Move around in the tree
---
- anytime we are working inside a folder 
  - the current folder, or
  - the present working directory

| command | purpose |
| --- | --- |
| pwd | print working directory |
| cd< | go to the home folder of current user |
| cd folder | go to folder |


Commands operating on folders
---
| command | purpose |
| --- | --- |
| tree dir | show folders and files rooted at dir as a tree |
| mkdir dir1 dir2 ... | make multiple directories |
| mkdir -p dir1/dir2/dir3 | make nested directories |
| rmdir dir | remove empty directory |
| rm -rf dir | remove non-empty directory |
| cp file1 file2 ... filen dir | copy multiple files to dir |
| ○ cp -r source_dir dest_dir<br> ○ rsync -s source_dir dest_dir | copy directories recursively |


Look for files
---
| command | purpose |
| --- | --- |
| find folder -name filename | search filename in folder recursively |
| locate "*pub*" | search all filenames containing pub system-wide |
| find dir -iname "*.pdf" -exec rm {} \\; | find and remove all pdf files in dir |



Commands operating on files
---
| command | purpose |
| --- | --- |
| touch file1 file2 ... filen | create multiple empty files |
| rm file1 file2 ... filen | remove multiple files |
| cp source_file dest_file | copy a file to another |
| mv source_file dest_file | rename a file |
| ln -s linked_file link | create a symbol link |
| ls | list all regular files in the current folder excluding hidden files whose names begin with . |
| ls -a | display all files in the current folder including hidden files |
| ls -l | display a long listing |
| ls -t | list by time (most recent files first) |
| ls -S | list by size (largest files first) |
| ls -r | list with a reverse sort order |
| ls -lSr | long list with smallest files first |


Commands show file content
---
| command | purpose |
| --- | --- |
| cat file1 file2 ... filen | concatenate and display file contents |
| more file1 file2 ... filen | display the contents of multiple files page by page |
| more file1 file2 ... filen | similar to more but with more features |
| head -n file | show the first n lines |
| tail -n file | show the last n lines |


Search/sort file content
---
| command | purpose |
| --- | --- |
| grep substring file | show only lines  in file containing substring |
| grep -i substring file | case insensitive search |
| grep -v substring file | show all the lines NOT containing substring |
| grep -r substring dir | search through all the files in a directory recursively |
| sort file | sort lines in a given file |
| sort -u file | sort lines, only display duplicate ones once |


Compare files and folders
---
| command | purpose |
| --- | --- |
| diff file1 file2 | compare two files |
| diff -r dir1 dir2 | compare two folders recursively |


Find command help
---
| command | purpose |
| --- | --- |
| ls -h | short help |
| ls --help | long help |
| man ls | manual page of ls |
| info ls | info page of ls |
| help | show shell builtin commands |
| help command | help of builtin commands |


Further information
---
- [Introduction to Unix/Linux commands by Bootlin](https://bootlin.com/doc/legacy/command-line/)



# Basic Windows commands
| command | purpose |
| --- | --- |
| exit | Quits the CMD.EXE program (command interpreter) |
| help | Shows command help |
| attrib | Displays or sets file attributes |
| assoc | Associates an extension with a file type (FTYPE). |
| copy, xcopy, robocopy | Copies files and folders |
| move | Moves a file to a new location |
| del, erase | Deletes one or more files |
| ren, rename | Renames a file or directory |
| type | Prints the content of a file to the console |
| more | Displays the contents of a file or files, one screen at a time |
| comp, fc | Compares files |
| where | Outputs one or more locations of a file or a file name pattern |
| find | Searches for a string in files or input, outputting matching lines |
| findstr | Searches for regular expressions or text strings in files |
| sort | Sorts alphabetically, from A to Z or Z to A, case insensitive |
| dir | Displays a list of files and subdirectories in a directory |
| tree | Displays a tree of all subdirectories of the current directory to any level of recursion or depth |
| cd, chdir | Displays or sets the current directory |
| md, mkdir | Creates a directory |
| rd, rmdir | Removes the directory |


Further information
---
- [Windows batch programming](https://github.com/ufidon/its372/blob/master/lectures/module01/README.md)