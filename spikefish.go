package main

import (
  "io/ioutil"
  "exp/ssh"
)

func main() {
  config := new(ssh.ServerConfig)
  pemBytes, err := ioutil.ReadFile("id_rsa")
  if err != nil {
    panic("Failed to load private key")
  }
  err = config.SetRSAPrivateKey(pemBytes)
  if err != nil {
    panic("Failed to parse private key")
  }
  listener, _ := ssh.Listen("tcp", "0.0.0.0:2022", config)
  sConn, err := listener.Accept()
  if err != nil {
    panic("failed to accept incoming connection")
  }
  err = sConn.Handshake(conn)
  if err != nil {
    panic("failed to handshake")
  }

}
