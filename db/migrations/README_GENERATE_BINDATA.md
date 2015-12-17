#Generating bindata.go

If an existing sql migration is modified or a new one is inserted the bindata.go file has to be regenerated.

In order to do that, make sure you have installed go-bindata

```
go get -u github.com/jteeuwen/go-bindata/...
```

Once that's installed, go to the root of this repo (i.e. `safenotes-api/`) and run the following command:

```
go-bindata -ignore ".*\.(?:go|txt|md)$" -pkg migrations -o db/migrations/bindata.go db/migrations/
```

The `bindata.go` file should have been regenerated, and has to be commited with your sql changes.

If the file isn't regenerated, the app will not pick up the changes to the sql when it runs.
