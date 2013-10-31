package main

import (
  "fmt"
  "os/exec"
  "bytes"
)

func receiveMulti(args ...string) {
  fmt.Println(args)
  command1 := exec.Command("ls", "-la")
  command2 := exec.Command("ls", "-l", "-a")
  command3 := exec.Command("ls", "-l -a")
  var buf1 bytes.Buffer
  var buf2 bytes.Buffer
  var buf3 bytes.Buffer
  command1.Stdout = &buf1
  command2.Stdout = &buf2
  command3.Stdout = &buf3
  _ = command1.Run()
  _ = command2.Run()
  _ = command3.Run()

  fmt.Println(
    "This is buf1\n",
    buf1.String(),
    "This is buf2\n",
    buf2.String(),
    "This is buf3\n",
    buf3.String(),
  )
}

func main() {
  args := []string{"hoge","fuga","piyo"}
  receiveMulti("hoge","fuga")
  receiveMulti(args)
}
