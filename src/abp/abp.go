package main

import (
  "fmt"
)



// Fork type
type Message struct {
  content string      //  message content
}

// Memory
type Memory struct {
  // (public) channels
  Write0 chan int // channel used to change memory status
  Write1 chan int // channel used to change memory status
  Read0 chan int // channel used to read memory status
  Read1 chan int // channel used to read memory status
}

func (self *Memory) Run () {
  var state = 0

  for {
    if state == 0 {
      select {
        case <- self.Write1:
          state = 1
        case self.Read0 <- 0:
      }
    } else {
      select {
        case <- self.Write0:
          state = 0
        case self.Read1 <- 1:
      }
    }
  }
}

// Sender
func sender(send0 chan *Message, send1 chan *Message, mem *Memory) {
  var msg Message
  for {
    select {
      case <- mem.Read0:
        msg.content = "Hello"
        select {
          case send0 <- &msg:
          case <- mem.Read1:
        }
      case <- mem.Read1:
        msg.content = "Tim"
        select {
          case send1 <- &msg:
          case <- mem.Read0:
        }
    }
  }
}

// Receiver
func receiver(receive0 chan *Message, receive1 chan *Message,
              ack0 chan bool, ack1 chan bool) {
  var msg *Message
  var expectFrom = 0

  for {
    if expectFrom == 0 {
      msg = <- receive0
      fmt.Println(msg.content)
      ack0 <- true
      expectFrom = 1
    } else {
      msg = <- receive1
      fmt.Println(msg.content)
      ack1 <- true
      expectFrom = 0
    }
  }
}

// AckReceiver
func acknack(ack0 chan bool, ack1 chan bool,
             mem *Memory) {
  var seq = 0

  for {
    if seq == 0 {
      select {
        case ok := <- ack0:
          if ok {
            mem.Write1 <- 1
            _,_ = fmt.Println("OK0");
            seq = 1
          }
        case <- ack1:
      }
    } else {
      select {
        case ok := <- ack1:
          if ok {
            mem.Write0 <- 0
            _,_ = fmt.Println("OK1");
            seq = 0
          }
        case <- ack0:
      }
    }
  }
}

// main
func main () {
  fmt.Printf("Hello World\n")
  for i := 0; i < 5; i++ {
    var v int
    fmt.Printf("%d ", v)
    v = 5
  }
  fmt.Printf("\n")

  var mem = Memory {make(chan int), make(chan int),
                    make(chan int), make(chan int)}
  go mem.Run()

  var msgChan0 = make(chan *Message)
  var msgChan1 = make(chan *Message)
  var ackChan0 = make(chan bool)
  var ackChan1 = make(chan bool)

  go sender(msgChan0, msgChan1, &mem)
  go receiver(msgChan0, msgChan1, ackChan0, ackChan1)
  acknack(ackChan0, ackChan1, &mem)
}