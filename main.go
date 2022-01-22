package cmdstream

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

type CommandStream struct {
	Messager Messager
	Cmd      *exec.Cmd
}

func NewCommandStream(command *exec.Cmd, messager Messager) *CommandStream {
	return &CommandStream{
		Cmd:      command,
		Messager: messager,
	}
}

func (cs *CommandStream) RunCommand() {

	stdout, err := cs.Cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cs.Cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cs.Cmd.Start(); err != nil {
		log.Fatal(err)
		return
	}

	s := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for s.Scan() {
		cs.Messager.SendMessage([]byte(s.Text()))
	}

	if err := cs.Cmd.Wait(); err != nil {
		log.Fatal(err)
	}

}
