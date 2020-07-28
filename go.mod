module github.com/infra-whizz/wz-ansible-modules

go 1.14

replace github.com/infra-whizz/wzmodlib => ../wzmodlib

require (
	github.com/antonfisher/nested-logrus-formatter v1.1.0 // indirect
	github.com/infra-whizz/wzlib v0.0.0-20200709175548-7accf26d7b69
	github.com/infra-whizz/wzmodlib v0.0.0-20200720151532-7bd7a478413f
	github.com/sirupsen/logrus v1.6.0 // indirect
	golang.org/x/sys v0.0.0-20200728102440-3e129f6d46b1
)
