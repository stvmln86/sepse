# Sepse

**Sepse** is Stephen's Elemental Pub/Sub Engine, a minimalist TCP pubsub client/server application written in Go 1.25 by Stephen Malone. 

```fish
# run server on ADDR 
$ sepse host ADDR 

# list all names on ADDR 
$ sepse list ADDR 

# receive messages from ADDR 
$ sepse recv ADDR NAME

# send message to ADDR 
$ sepse send ADDR NAME DATA 
```
