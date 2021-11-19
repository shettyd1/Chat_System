package main

import (
    "fmt"
    "net/http"
    "os"
    "encoding/json"
    "time"

     
)
type Account struct {
    Username []string  "username"
    Password string   "password"
}

type User struct{
    Username  []string  "username"
    Passsword string    "password"
    Login bool
  }
type Logged struct {

 }
const BUFFERSIZE int = 1024
var newClients_conn = make(map[net.Conn]string)
var allClients_conn = make(map[net.Conn]string)
var currentLoggedUser = make(map[net.Conn]interface{})
var allLoggedIn_conns = make(map[net.Conn]interface{})
var lostclient = make(chan net.Conn)

func main() {

godotenv.Load()
    if len(os.Args) != 2 {
        fmt.Sprintf("Usage: %s <port>\n", os.Args[0])
        os.Exit(0)
    }
    port := os.Args[1]
    if len(port) > 5 {
        fmt.Println("Invalid port value. Try again!")
        os.Exit(1)
    }
    server, err := net.Listen("tcp", ":"+port)
    if err != nil {
        fmt.Printf("Cannot listen on port '" + port + "'!\n")
        os.Exit(2)
    }
    fmt.Sprintf("Server is listening on port '%s' ...\n", port)
    newclient := make(chan net.Conn)
    go func() {
        for {
            client_conn, _ := server.Accept()
            newclient <- client_conn
            
        }
    }()
    for {
        select {
       case client_conn := <-newclient:
            allClients_conn[client_conn] = client_conn.RemoteAddr().String()
            go client_goroutine(client_conn, net.Conn)
        case client_conn := <-lostclient:
            delete(allClients_conn, client_conn)
            byemessage := fmt.Sprintf("A client '%s' DISCONNECTED!\n# of connected clients: %d\n", client_conn.RemoteAddr().String(), len(allClients_conn))
            
    }
 json.ConnectDatabase(allLoggedIn_conns, currentLoggedUser)
 
}
func client_goroutine(client_conn net.Conn) {
    welcomemessage := fmt.Sprintf("A new client '%s' connected!\n# of connected clients: %d\n", client_conn.RemoteAddr().String(), len(allClients_conn))
    fmt.Println(welcomemessage)
    go public_chat([]byte(welcomemessage))
    var buffer [BUFFERSIZE]byte

    go func() {
        for {
            byte_received, read_err := client_conn.Read(buffer[0:])
            if read_err != nil {
                fmt.Println("Error in receiving...")
                lostclient <- client_conn
                return
            }
            go public_chat(buffer[0:byte_received])
        }
    }()
}

func login(client_conn net.Conn) {
  
var user User
   err :=json.Unmarshal(data, &user)
   if err!=nil || user.Username == " " || user.Password == "" {
      fmt.Sprintf("Parsing error %s\n", err)
      return false
}

func checkLogin(client_conn net.Conn) {
  
 var logged Logged
  if (user) == allLoggedIn_conns[client_conn].Logged(login) {
  return len(map[currentLoggedUser[client_conn]])
  }
}
   
func checkAccount(client_conn net.Conn) {

 account Account {
err := json.Unmarshal(&Login, &checklogin)
if err == nil {
   return checkAccount.(Username)
}
return false
}
}


