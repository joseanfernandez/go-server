# GO

For this example, I have used [The Movie Database API](https://developers.themoviedb.org/3/getting-started/introduction).

You must register for an API key.

### Instructions

####  Download the code
```console
git clone https://github.com/joseanfernandez/go-server.git
```

### Create your own config package
```console
cd go-server
mkdir config
cd config
touch config.go
```

#### config.go
```console
package config

var ApiKey =  "YOUR_API_KEY"
```

####  Build
```console
cd go-server
go build
```

#### How to use it
```console
./go-server
```

[Try me!](http://localhost:8080/films?key=now_playing)

