
# Restaurant order system

A barebones Go app, which can easily be deployed to Heroku.
A server side application with endpoints made to manage and take restaurant orders

This application supports the [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) article - check it out.
There will be DEV article on this system soon to help for better understanding and replication 😉 

## Running Locally

kindly make sure to have installed [Go](http://golang.org/doc/install) version 1.12 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) (that is if you plan on hosting it on heroku).

```sh
$ git clone https://github.com/Joojo7/GO-Restarunt.git
$ cd GO-Restarunt
$ go build -o bin/GO-Restarunt -v 
$ nodemon --exec heroku local  --signal SIGTERM;
```

The application should be available and running on [localhost:8000](http://localhost:8000/).


## Documentation
A swagger endpoint testing zone will be created for api testing.

If you want to have more information about using Go on Heroku, take a look at these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
