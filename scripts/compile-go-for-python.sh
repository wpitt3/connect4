#!/bin/bash

cd ../go
go build -buildmode=c-shared -o c4mcts.so c4mcts.go tree.go treeSearch.go utils.go
