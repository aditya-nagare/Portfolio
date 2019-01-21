# String-Search-Files


This is a Go(Golang) Program that reads in multiple files in the current working directory of the program under `/files` folder and lets the user to search for the occurance of a string in the multiple files present.

The program uses Go Routines, Unbuffered Channels and Worker Pool to acheive this task concurrently.

The default string is set to `of the` but the user is able to customize it via a flag.


## To Run The Program(on Linux):


Go to the directory where the program is stored and run the program using following commands:

```go run main.go```



The default string is set to `of the` but the user is able to customize it via the search string via command line arguments.
To customize the seach string pass an command line argument `-word="<string to search>"` for example if the string is `and the` pass the argument as: 

```go run main.go -word="and the"```


## Output:


Output of the program is available in the following gist:

https://gist.github.com/aditya-nagare/baa3f53730a82be0ad74a2916e0012e5



## To Install Go(Golang):


If your system doesn't has Go installed on it, go to the following link and download and install as per your OS:

https://golang.org/dl/
