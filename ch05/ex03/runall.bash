go build ../util/fetch.go
go build textnode.go
./fetch http://golang.org | ./textnode
