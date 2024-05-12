# TCP/IP with Golang
## TCP - Transmission Control Protocol/Internet Protocol
- reliable
- ordered
- error checked
- transport layer

### Roles of the TCP/IP
- estabilishing connections
- tracking the sequence of packets
- acknowledging of the receipt data packets
- flow control
- routing packets
- fragmentation and reassemlbing of packets
- best-effort delivery

### Sockets
> sockets are endpoints to the network communication
> they allow programs to send and receive data over the internet

- socket address: IP + port number
- server sockets: to listen to connections
- client sockets: to send to connhections
- socket communication: bidirecional communication

### Segments and packages
- the data packages are called segments
- each segment has a number to ensure the reassemble
- the receiver acknowledge the receipt of segments (may retransmit)
- the receiver buffers the data to ensure to not lost

### Buffering and Streaming
- buffer is storing the data before processing
- streaming is processing the data while it is comming, without waiting for the data be 100% available

## Concurrency in GoLang
### Goroutines
- goroutines are the threads of Go
- each goroutine will handle a connection

### Channels (search more)
### Pooling (search more)
### Timeouts (search more)




https://okanexe.medium.com/the-complete-guide-to-tcp-ip-connections-in-golang-1216dae27b5a
https://www.kelche.co/blog/go/golang-bufio/
https://go.dev/tour/concurrency/1
https://gobyexample.com/testing
https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/
https://www.linkedin.com/pulse/mastering-error-handling-golang-shiva-raman-pandey/
https://go.dev/blog/error-handling-and-go
https://medium.com/@leodahal4/handle-errors-in-go-like-a-pro-5f2ab97c660b
https://stackoverflow.com/questions/22170942/how-can-i-get-an-error-message-in-a-string-in-go