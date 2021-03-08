module github.com/infra-whizz/wz-ansible-modules

go 1.14

replace github.com/infra-whizz/wzmodlib => ../wzmodlib

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.0 // indirect
	github.com/infra-whizz/wzlib v0.0.0-20210306212611-2af49aea1704
	github.com/infra-whizz/wzmodlib v0.0.0-20210308141613-815923370ebc
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/magefile/mage v1.11.0 // indirect
	github.com/sirupsen/logrus v1.8.0
	golang.org/x/sys v0.0.0-20210305230114-8fe3ee5dd75b
)
