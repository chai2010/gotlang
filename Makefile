# 狗头语 版权 @2019 柴树杉。

default:
	go run ./got/main.go hello.gotmpl aa bb cc 11 22 33

help:
	go run ./got/main.go -h

dev:
	go run ./got/main.go ./examples/04-prime.gotmpl

fib:
	go run ./got/main.go ./examples/05-fib.gotmpl

test:
	go run ./got/main.go ./examples/01-helloworld.gotmpl
	go run ./got/main.go ./examples/02-sum100.gotmpl
	go run ./got/main.go ./examples/03-xrange.gotmpl
	go run ./got/main.go ./examples/04-recursion.gotmpl
	go run ./got/main.go ./examples/05-fib.gotmpl

clean:
