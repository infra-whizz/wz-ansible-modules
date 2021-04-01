module github.com/infra-whizz/wz-ansible-modules

go 1.14

//replace github.com/infra-whizz/wzmodlib => ../wzmodlib

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/infra-whizz/wzlib v0.0.0-20210306212611-2af49aea1704
	github.com/infra-whizz/wzmodlib v0.0.0-20210308141613-815923370ebc
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/sys v0.0.0-20210331175145-43e1dd70ce54
)
