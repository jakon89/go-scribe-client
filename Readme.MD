### GO! Scribe client

It's pretty simple GO client for Scribe useful for testing :)


Installation:

```go get github.com/jakon89/go-scribe-client```

```go install github.com/jakon89/go-scribe-client```


Usage:


```$GOPATH/bin/go-scribe-client```

After startup, you can add messages in format "category::message" to batch:

```my_category::{"message":"my first message"} <ENTER>```

Available commands:
* SEND - will send all messages in batch
* CLEAR - will clear local messages
* EXIT - will exit


Example:

```
jakon89$ ./scribe-client-go -addr localhost:9090
Connected to localhost:9090 
new_file_event::{"message":"new file created!"}
LogEntry({Category:new_file_event Message:{"message":"new file created!"}}) added to batch 
SEND
1 messages sent 
EXIT
jakon89$ 
```