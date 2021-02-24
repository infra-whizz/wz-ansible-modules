OUT := pkgpinch
DISTRO := $(shell cat /etc/os-release | grep '^NAME=' | sed -e 's/.*=//')
GCCGO :=
GO := go

SOURCES = \
	*.go

all: update-deps mainbuild compress
gcc: update-deps with-gcc compress

mainbuild: $(SOURCES)
	$(GO) build -x -o $(OUT) $(SOURCES)

with-gcc:
ifneq ($GCCGO,)
	$(GO) build -compiler $(GCCGO) -gccgoflags '-static-libgo' -o $(OUT)
else
	@echo "Cannot find gccgo"
endif

strip:
	strip $(OUT)

update-deps:
	$(GO) get -u

deps:
	$(GO) get -v ./... 2>/dev/null ||:
strip:
	strip $(OUT)

compress:
	upx -9t $(OUT)

clean:
	rm $(OUT)
	$(GO) clean -x -i

flush-cache:
	$(GO) clean -cache -modcache -i -r

.PHONY: all install clean

