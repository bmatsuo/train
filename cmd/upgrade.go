package main

import (
	"fmt"
)

const CmdBinPath = "$GOPATH/bin/train"

func upgrade() {
	bash("go get -u github.com/bmatsuo/train")
	bash("go build -o " + CmdBinPath + " github.com/bmatsuo/train/cmd")
	fmt.Println("Installed latest train command into " + CmdBinPath)
}
