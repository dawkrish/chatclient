### TCP Server
* create a listener for the specific host & port 
* accept connections on that listener
* create `reader` for that connection
  * `reader` reads the request of client (read on `conn`)
  * `reader` writes the response ot the client(`conn.write`)

### TCP Client
* connect to server (`net.Dial`)
* read user input (stdin)
* write to conn
* read server response (read on `conn`)
* write to stdout
