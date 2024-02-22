package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func killCommand(args []string) error {
	if len(args) < 1 {
		return errors.New("kill: usage: kill pid")
	}
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	return syscall.Kill(pid, syscall.SIGKILL)
}

func psCommand() error {
	return runCommand("ps", []string{})
}

func cdCommand(flag string) error {
	err := os.Chdir(flag)
	if err != nil {
		return errors.New("bash: cd: No such file or directory")
	}
	return nil
}

func pwdCommand() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func echoCommand(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func runCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func runPipeline(commands [][]string) error {
	cmds := make([]*exec.Cmd, len(commands))

	for i, command := range commands {
		cmds[i] = exec.Command(command[0], command[1:]...)

		if i > 0 {
			r, w := io.Pipe()
			cmds[i-1].Stdout = w
			cmds[i].Stdin = r
		}
	}

	cmds[len(cmds)-1].Stdout = os.Stdout

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func start() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal()
		}
		parts := strings.Fields(string(line))
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "cd":
			cdCommand(parts[1])
		case "pwd":
			fmt.Println(pwdCommand())
		case "ps":
			err = psCommand()
		case "echo":
			echoCommand(parts[1:])
		case "kill":
			err = killCommand(parts[1:])
		default:
			if strings.Contains(string(line), "|") {
				pipelineParts := strings.Split(string(line), "|")
				commands := make([][]string, len(pipelineParts))
				for i, part := range pipelineParts {
					commands[i] = strings.Fields(part)
				}
				err = runPipeline(commands)
			} else {
				err = runCommand(parts[0], parts[1:])
			}

		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	start()
}
