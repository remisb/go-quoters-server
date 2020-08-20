# go-quoters-server

Project is based on spring boot rest api server from ... and converted to golang.
- Original spring repository: https://github.com/spring-guides/quoters
- Tutorial: [Building a RESTful quotation service with Spring](https://spring.io/blog/2014/08/21/building-a-restful-quotation-service-with-spring)

###### Default config settings:

```yaml
SrvConfig {
    Host: "localhost",
    Port: 8888,
    Log: "./quoters.log",
}

DbConfig {
    Host:     "localhost",
    Port:     "3306",
    User:     "root",
    Password: "",
    Name:     "quoter",
}
```

### API endpoints

```raml
title: Quoter API
baseUri: http://localhost:8888/
mediaType: application/json
types:
  Quote:
    type: object
    properties:
      Id:    string
      Quote: string
/api:
  /quote:
    type: object
    get:
      description: Fetch a random quote
      responses:
        200:
          body:
            application/json:
              type: Quote[]
    /random:
      description: Fetch a random quote
      type: collection
      get:
        responses:
          200:
            body:
              application/json:
                type: Quote
    /{id}:
      description: Fetch quote id
      type: object
      get:
        responses:
          200:
            body:
              application/json:
                type: Quote
```
