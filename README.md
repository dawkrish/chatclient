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

### Demo

(Crop issues due to OBS)


![lisa](https://github.com/user-attachments/assets/4e191820-dd33-4ca0-b18b-cbd8614fcc03)

