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
