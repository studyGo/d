package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	var s = flag.String("s", "start", "command")
	flag.Parse()

	switch *s {
	case "start":
		var cmd *exec.Cmd
		if len(os.Args) == 1 {
			os.Exit(0)
		} else if len(os.Args) == 2 {
			cmd = exec.Command(os.Args[1])
		} else {
			cmd = exec.Command(os.Args[1], os.Args[2:]...)
		}

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		f, err := os.OpenFile("./lock.pid", os.O_CREATE|os.O_WRONLY, 0660)
		if err != nil {
			log.Println("open pid file \"./lock.pid\" failed  Permission denied")
			os.Exit(0)
		}
		f.WriteString(strconv.Itoa(cmd.Process.Pid))
		os.Exit(0)
	case "stop":
		f, err := os.OpenFile("./lock.pid", os.O_RDONLY, 0660)
		if err != nil {
			log.Println("server is not start")
			os.Exit(0)
		}

		data := make([]byte, 6)
		len, err := f.Read(data)

		if err != nil {
			log.Fatal(err)
		}

		pid, _ := strconv.Atoi(string(data[:len]))
		p, err := os.FindProcess(pid)
		if err != nil {
			log.Fatal(err)
		}
		p.Kill()
		os.Exit(0)
	case "reload":
	case "watch":
	}

}
