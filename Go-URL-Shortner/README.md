# Go-Quiz-Game

This is a Simple Go(Golang) Program to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like a URL Shortener would.

For instance, if a redirect is setup for `/ggl` to `https://www.google.com`, the program looks for any incoming web requests with the path `/ggl` and redirect them.

You can also use custom YAML File to give the program a set to links to redirect.


## To Run The Program(on Linux):

### Go to the directory where the program is stored and run the program using following commands:

```go run main/main.go```

### Then go to your browser and go to

```localhost:8080```


### You can go to the these following URL redirects to access the files:

```localhost:8080/ggl``` for google.com

```localhost:8080/yt``` for youtube.com

```localhost:8080/fb``` for facebook.com

```localhost:8080/wiki``` for wikipedia.org

```localhost:8080/med``` for medium.com

```localhost:8080/9``` for 9gag.com



### Using a YAML File

You can also run the program with custom YAML File.

The default YAML File is set to `links.yaml` but the user is also able to customize the filename via command line arguments.

To customize the YAML File pass an command line argument `-yaml="<filename>.yaml"` for example if the filename is links.yaml pass the argument as: 

```go run main/main.go -yaml=links.yaml```

## To Install Go(Golang):

If your system doesn't has Go installed on it, go to the following link and download and install as per your OS:

https://golang.org/dl/
