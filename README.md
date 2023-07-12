# GORM custom context

code from here: https://dev.to/franciscomendes10866/build-a-crud-graphql-api-with-gqlgen-and-gorm-36de

## Install

```sh
go get -u github.com/juunini/gorm-custom-context
```

## Usage

### With fiber

```go
import (
  // ...
  customContext "github.com/juunini/gorm-custom-context"
  // ...
)

// ...

app.All(
  "/query",
  adaptor.HTTPHandler(
    customContext.CreateContext(&customContext.CustomContext{
      Database: database,
    }, srv),
  ),
)
```

### in resolver

```go
context := customContext.GetContext(ctx)
author := &model.Author{
  ID:   input.ID,
  Name: input.Name,
}

context.Database.Create(author)
```
