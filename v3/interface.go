package main

import (
  "fmt"
  "os/exec"
  "bufio"
  "context"
  "errors"
)

type Interface interface {
  Start()
  SetMessageHandler(func([]byte, []byte) error)
  SetReadHandler(func([]byte) ([]byte, error))
}

var (
  HeaderNotUnderstood = errors.New("Header not understood")
  CommandNotUnderstood = errors.New("Command not understood")
  //NotMatserComm = errors.New("Not the MasterComm")
  NotEnoughFields = errors.New("Not enough field")
  EmptyString = errors.New("Received an empty string")

  nilMessageHandler = func([]byte, []byte) error {return nil}
  nilReadHandler = func([]byte) ([]byte, error) {return []byte{}, nil}

  InterfaceSendHeader = "Send"
  InterfaceReadHeader = "Read"
)

func NewInterface(ctx context.Context, file string, n []int, i []int) (Interface, error) {
  cmdArgs := append([]string{file + "/run.py", fmt.Sprint(n), fmt.Sprint(i)})
  inter := StdInterface {
    Ctx: ctx,
    Idx: i,
    N: n,
    Cmd: exec.CommandContext(ctx, "python3", cmdArgs...),
    MessageHandler: &nilMessageHandler,
    ReadHandler: &nilReadHandler,
  }

  return &inter, nil
}

type StdInterface struct {
  Ctx context.Context
  MessageHandler *func([]byte, []byte) error
  ReadHandler *func([]byte) ([]byte, error)
  Idx []int
  N []int
  Cmd *exec.Cmd
}

func (s *StdInterface)Start() {
  defer recover()

  stdin, err := s.Cmd.StdinPipe()
	if err != nil {
    return
	}

  stdout, err := s.Cmd.StdoutPipe()
	if err != nil {
    return
	}

  err = s.Cmd.Start()
  if err != nil {
    return
	}

  scanner := bufio.NewScanner(stdout)
  go func(){
    defer recover()

    for scanner.Scan() {
      switch scanner.Text() {
      default:
        return

      case InterfaceReadHeader:
        if scanner.Scan() {
          address := scanner.Text()
          msg, err:= (*s.ReadHandler)([]byte(address))
          if err != nil {
            return
          }

          _, err = fmt.Fprintln(stdin, msg)
          if err != nil {
            return
          }
        }

      case InterfaceSendHeader:
        var address, value string
        if scanner.Scan() {
          address = scanner.Text()
        }

        if scanner.Scan() {
          value = scanner.Text()
          err = (*s.MessageHandler)([]byte(address), []byte(value))
          if err != nil {
            return
          }
        }
      }
    }
  }()
}

func (s *StdInterface)SetMessageHandler(handler func([]byte, []byte) error) {
  s.MessageHandler = &handler
}

func (s *StdInterface)SetReadHandler(handler func([]byte) ([]byte, error)) {
  s.ReadHandler = &handler
}
