module github.com/infra-whizz/wz-ansible-modules

go 1.14

replace github.com/infra-whizz/wzmodlib => ../wzmodlib

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.0 // indirect
	github.com/infra-whizz/wzlib v0.0.0-20201210130450-2b56d9cf0495
	github.com/infra-whizz/wzmodlib v0.0.0-20201210130531-822420f20753
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f
)
