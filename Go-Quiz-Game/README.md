# Go-Quiz-Game

This is a Simple Go(Golang) Program to read in a quiz provided via a CSV file and will then give the quiz to a user also keeping track of how many questions they get right and how many they get incorrect.
The quiz also has a timer. The quiz stops as soon as the time limit has exceeded.

At the end of the quiz the program outputs the total number of questions correct and how many questions there were there in total. Questions given invalid answers or unanswered are considered incorrect.

The default CSV file is set to `problems.csv` but the user is able to customize the filename via a flag.
The default time limit is set to `30 seconds` but the user is able to customize via a flag.



##To Run The Program(on Linux):

Go to the directory where the program is stored and run the program using following commands:

```go run main.go```


The default CSV file is set to `problems.csv` but the user is also able to customize the filename via command line arguments.
To customize the CSV file pass an command line argument `-csv="<filename>"` for example if the filename is abc.csv pass the argument as: 

```go run main.go -csv="abc.csv"```


The default time limit is set to `30 seconds`, but the user is also able to customize the time limit via command line arguments.
To customize the time limit pass an command line argument `-limit=<seconds>` for example if the time limit is 5 seconds pass the argument as:

```go run main.go -csv="abc.csv"```



##To Install Go(Golang):

If your system doesn't has Go installed on it, go to the following link and download and install as per your OS:

https://golang.org/dl/
