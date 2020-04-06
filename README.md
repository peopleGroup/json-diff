# json-diff-snippets
A template to implement your own json-diff

json-diff works on pub-sub pattern

There are TODOs in producer.go and consumer.go
- Implement the producer function which will generate requests and write it on a channel.
- The consumers will listen on the channel and pick up request make api calls concurrently and will write the JSON response in a file.
- After all the requests are done, the diff.js script will we invoked, which will compare the JSON responses and write the difference (added, modified & deleted) in a file.

To build and install json-diff
```sh
go install ./...
```
go install copies the built binary into your GOPATH, so you can run it.
```sh
json-diff
```
The diff script can be run standalone using node (>= 10.13.0 )
```sh
node diff.js
```