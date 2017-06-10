go build fetch.go
go build findlinks.go
./fetch http://golang.org | ./findlinks
