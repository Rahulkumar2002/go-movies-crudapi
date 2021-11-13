> # Golang-Movies-CRUD-API
<h3>

## In this project  , I have  created a basic CRUD API in Golang for a movie serving site. 

Let's get started , first I will create a folder and name it as **Golang-Movies-CRUD-API** after this open **Powershell** in that folder and run a command ``` go mod init  go-movies-crud-api ``` , this command will create a **.mod** file in this directory which will store our required dependencies . Next run ```go get -u github.com/gorilla/mux``` , this will add **Gorilla** to your **go.mod** file which will be used later in the project .

Why are we using mux ?<br>
Package /mux implements a request router and dispatcher for matching incoming requests to their respective handler.Basically we are using mux to handle our routes .


We will create a **main.go** file next , packages which we will be importing is :

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)
//github.com/gorilla/mux is a Golang HTTP router
```
Usage of these packages will be covered later in this project.

Creating a simple *struct* of name **Director** which will have two variables of type string , name of first variable is *Firstname* and second one is *Lastname* .

```go
type Director struct {
	Firstname string `json : "firstname"`
	Lastname  string `json : "lastname"`
}
```

Next create a *struct* named as **Movie** which will have *Id* of type *string* it will be used to provide unique id's to our collection of movies, *Isbn* type *string* , *Title* of type *string* this will be title of our movie.Atlast we will have variable *Director* of type *Director* this will contain name of our movie director, with all of these variable we will also write there **json name** as well , this **json name** will be name of these variable in our json object .
```go
type Movie struct {
	Id       string    `json :"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json : "director"`
}
```

We will declare a *slice* named as **movies** with type **Movie** , this will help us to store our struct of different movies.
```go
var movies []Movie
```

> ## let's write our **main** function .

For this we will start with a print statement which will be printing **Your server is connected on port 8000** , I have choosed 8000 you can choose according to yourself . 
`fmt.Println("Your server is connected on port 8000")`

**fmt** package is used to print statements or values of variables on terminal or web page.

Next we will be adding 3 movies to our *moveis slice* . Like this :

```go

movies = append(movies, Movie{Id: "1", Isbn: "5000", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})

movies = append(movies, Movie{Id: "2", Isbn: "5010", Title: "Movie Two", Director: &Director{Firstname: "Rahul", Lastname: "Kumar Jha"}})

movies = append(movies, Movie{Id: "3", Isbn: "5020", Title: "Movie Three", Director: &Director{Firstname: "Rahul", Lastname: "Kumar"}})

```

We will use **mux.NewRouter()** to create a new router . Which will be used to handle our end points . 

` r := mux.NewRouter() `

After this we will assign our end points there handler function using **HandleFunc** . We have 5 end points : 

```go
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
```
Our end points : 

```go
//This end point will be used to display our collection of movies .
"/movies"
//This end point will be used to display a particular movie with matching id to provided id at the time of hitting of end point.
"/movies/{id}"
//This end point will be used to create new movies by passing value .
"/movies"
//This end point will be used to update current movies.
"/movies/{id}"
//This end point will be used to delete movies with a particular id .
"/movies/{id}"
```
Then to handle 5 end points we also have 5 functions : 
1. **getMovies** which have method **Get** , **Get** method means it will be only taking data from the server & will not make any changes to the data on the server .This method will be fetching data of all the movies form the server .
2. **getMovie** which also have method **Get**. It will only be fetching data from the server with a id .
3. **createMovie** which have method **POST** . It will be used to add data to the server that is also explaing the use of **POST** method .
4. **updateMovie** which have method **PUT** . It will be used to update data in our server , but instead of updating the data in our movies struct we will be creating a new struct of same *id* .
5. **deleteMovie** which have method **DELETE** . It will be used to delete a struct of given *id* . 

Atlast we will be printing *starting your port at 8000* & using **log** package to log any error occuring in connecting with the server . To implement the connection with our port we will be using **ListenAndServe** method from **http** package .

```go
fmt.Print("Starting your port at 8000\n")
log.Fatal(http.ListenAndServe(":8000" , r))
```

After this we will be creating all of our functions which will be used when we will hit our end points.

> ## getMovies function 

  In this function we will be taking two arguments  ,first will be our **w http.ResponseWriter**(A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.) and second **r *http.Request**(A Request represents an HTTP request received by a server or to be sent by a client).After that we will set it's header by this `w.Header().Set("Content-type" , "application/json")` , then we will use our **json.NewEncoder()**(NewEncoder returns a new encoder that writes to w.) from our **encoding/json** package . 
  `json.NewEncoder(w).Encode(movies)`.
```go
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
```
  our **getMovies** function is completed.
   Moving on to **deleteMovie**.
  > ## deleteMovie function 
  We will be covering **deleteMovie** function before other middle functions because it's easy to understand. 

  In **deleteMovie** we will be taking same arguments like **getMovies** , setting its header to content-type  as application/json. 
  
  Gorilla mux provides a nice way to extract parameters from a request’s path, using **mux.Vars** which is a map[string]string. We will be extracting parameters from our request using **mux.Vars** and assigning it to params for future use  .
  
  `params := mux.Vars(r)` 
  
  After this we will be using for loop to range our **movies slice** and to match requested *id* with **movies slice's id's** . If *id* matches we will be deleting element of that id from our slice. To implement deleting we will be usign this syntax : `append(movies[:index} , movies[index+1:]...)` after this we will break out of the loop.Then we will encode it using our **NewCoder** function .
  Our function will look something like this.
  ```go
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
  ```
> # getMovie function : 
  This funciton will be used to fetch a movie from the server according to provided *id* .

  In **getMovie** same two arguments , then we will set header with same method . Then we will extract variable data from our request using **mux.Vars** and assign it to params .

  Then we will range over movies and match our params id with available movies *id's* , if we find a match then we will serve it .

  ```go
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}
  ```


  > # createMovie function

  I'm declaring a variable `var movie Movie` .

  This function will be used to create new movies by gettig data through response body . We will be using **NewDecoder** to decode this value `_ = json.NewDecoder(r.Body).Decode(&movie)` . 
 
 To make *id* random we will use **rand.Intn** to generate a random *id* then we will convert it to string using *strconv.Itoa* and assign it to **movie.id** . 

 Next we will be adding our **movie** variable to **movies slice** using *append* . alast encoding it .

Our code will look like this:

```go
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}
```

> # updateMovie function 

In this we will update our movies but we won't be altering there data we will be creating new movies element with same id and updating it's data and then deleting the old one .

```go
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
```
We will be assigning provided id to params and then range over movies to match if we get it then we will delete it using same method alike our **deleteMovie** function , then doing the same steps like **createMovie** . This will update our movies element which we want to update .

```go
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
```

Our program is completed let test it using **Postman** . To learn more about **Postman** checkout this [video](https://youtu.be/VywxIQ2ZXw4).

## For our **getMovies** : 

![Hitting end points for getMovies function ](./images/Screenshot%20(174).png)


  

## For our **getMovie** ;
![Getting response for id = 2](./images/Screenshot%20(175).png)


## For our **createMovie** :
 go to Body > raw > json : 
    paste this : 
 ```json
  {
        "isbn": "5070",
        "title": "Movie Six",
        "Director": {
            "Firstname": "Lex",
            "Lastname": "Lee"
        }
  }

 ```

![getting response](./images/Screenshot%20(176).png)


## For our **updateMovie** : 

This will return updated movie 

![Updated movie](./images/Screenshot%20(177).png)


## For our **deleteMovie**  :

It will return remaing element of **movies slice**

![remaing movies after deletion](./images/Screenshot%20(178).png)



Our fucntion work completely fine . Thank u for reading this . Feel free to fork and use code of this repo.
  
This project is made form youtube [video](https://youtu.be/TkbhQQS3m_o)
</h3>
