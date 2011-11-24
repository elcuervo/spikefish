package main

import (
  "io/ioutil"
  . "exp/ssh"
)

func main() {
  config := new(ServerConfig)
  pemBytes, err := ioutil.ReadFile("fixtures/id_rsa_test")
  if err != nil {
    panic("Failed to load private key" + err.Error())
  }
  err = config.SetRSAPrivateKey(pemBytes)
  if err != nil {
    panic("Failed to parse private key" + err.Error())
  }
  listener, _ := Listen("tcp", "0.0.0.0:2022", config)
  conn, err := listener.Accept()
  if err != nil {
    panic("failed to accept incoming connection")
  }

  err = conn.Handshake()
  if err != nil {
    panic("failed to handshake" + err.Error())
  }

}
