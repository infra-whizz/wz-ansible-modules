module main

go 1.14

replace github.com/infra-whizz/wzlib => ../../../../../wzlib

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/infra-whizz/wz-ansible-modules v0.0.0-20200720173508-2b1bb614beff // indirect
	github.com/infra-whizz/wzlib v0.0.0-20200724114653-1b20fd7a54aa
	github.com/infra-whizz/wzmodlib v0.0.0-20200720151532-7bd7a478413f
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/sys v0.0.0-20200728102440-3e129f6d46b1 // indirect
)
