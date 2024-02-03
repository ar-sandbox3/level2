# level2

This is an attempt to solve the level2 challenge.
However, there are some assumptions, and concepts (read: best practices) are skipped for brevity.
As an example: we don't do error handling yet here.
Also, the way we extract substring is too naive and assuming the input is correct.

## Running locally

This project is intentionally simple, and contains only tests.
To run the tests:

```console
go test ./...
```

> [!TIP]
> While this project contains only tests, you can import this "module" to build a "main" program that depends on it.
> ```
> go get github.com/ar-sandbox3/level2
> ```
