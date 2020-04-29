# libp2p-mpi/core

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

## How does it work?

### _Interfaces_

All of the interfaces use by the main [mpi](./mpi.go) interface are defined in the [type_definition.go](./type_definition.go) file.

Using `mpi.SetInitFunctions()` you can set the init functions for all other interfaces.

#### ExtHost

[ExtHost](./host.go) is an extended go-libp2p host interface that implements functions to manage peerstores for each interpreter.

#### Store

The [Store](./ipfs.go) interface is an ipfs interface to store interpreters.

#### Remote

The [Remote](./remote.go) interface implements the connection between two peers and peer reseting.

#### Interface

The [interface](./interface.go) interface implements the interactions between a `SlaveComm` interface and a local interpreter.

#### SlaveComm

The [SlaveComm](./slaveComm.go) interface handles the interactions between the `Remotes` and a local `Interface`.

#### MasterComm

The [MasterComm](./masterComm.go) interface is a wrap-around of the `SlaveComm` interface.

#### standardFunctionsCloser

The [standardInterface](./standardInterface.go) interface is used in all other classes to handle the functions in the `standardFunctionsCloser` interface.

### _Peer reset_

The peer reset algorithm of libp2p-mpi is also defined in the [type_definition.go](./type_definition.go) file in the `ResetReader` function:

```go
func ResetReader(received int, sent []interface{}, sendToRemote func(interface{}), pushToComm func(string)) (readFromRemote func(string)) {
  offset := received

  for _, msg := range sent {
    sendToRemote(msg)
  }

  return func(msg string) {
    if offset > 0 {
      offset--
      return
    }

    pushToComm(msg)
  }
}
```

This function takes as argument:
 - the number of messages already received (`received`) to know how many messages to discard from the remote,
 - the list of all sent messages (`sent`) to re-send them,
 - the function that handles sending messages to the remote (`sendToRemote`),
 - and the function that pushes messages back to the comm (`readFromRemote`).

It returns a function that handles new messages from the remote (`readFromRemote`).

Initializing messages that have to be send at remote reset are append before the list of sent messages (`sent`) and passed to the `ResetReader` function.
