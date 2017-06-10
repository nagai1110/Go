go build fetch.go
go build textnode.go
./fetch http://golang.org | ./textnode
