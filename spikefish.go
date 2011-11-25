package main

import (
  "fmt"
  "io/ioutil"
  "exp/ssh"
)

func passwordAuth(user, passwd string) bool {
  fmt.Printf(passwd)
  return true
}

func pubKeyAuth(user, algorithm string, pubkey []byte) bool {
  fmt.Printf("user: %s, alg: %s, pub: %x", user, algorithm, pubkey)
  return true
}

func main() {
  config := new(ssh.ServerConfig)
  config.PasswordCallback = passwordAuth
  config.PubKeyCallback = pubKeyAuth

  pemBytes, err := ioutil.ReadFile("fixtures/id_rsa_test")

  if err != nil {
    panic("Failed to load private key" + err.Error())
  }

  err = config.SetRSAPrivateKey(pemBytes)
  if err != nil {
    panic("Failed to parse private key" + err.Error())
  }

  listener, _ := ssh.Listen("tcp", "0.0.0.0:2022", config)

  fmt.Println("SpikeFish")
  fmt.Println("    listening for incoming connection on :2022")

  conn, err := listener.Accept()

  if err != nil {
    panic("failed to accept incoming connection")
  }

  err = conn.Handshake()
  if err != nil {
    panic("failed to handshake" + err.Error())
  }

}
