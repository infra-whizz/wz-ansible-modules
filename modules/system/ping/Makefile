OUT := ping
DISTRO := $(shell cat /etc/os-release | grep '^NAME=' | sed -e 's/.*=//')
GCCGO :=
GO :=
ifeq ($(DISTRO),"Ubuntu")
	GCCGO += gccgo
	GO += go
else ifeq (($DISTRO),"openSUSE Leap")
	GCCGO += gccgo-8
	GO += go-8
endif

SOURCES = \
	*.go

all: update-deps mainbuild strip
gcc: update-deps with-gcc strip

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

compress:
	upx -9 $(OUT)
	upx -t $(OUT)

clean:
	rm $(OUT)
	$(GO) clean -x -i

flush-cache:
	$(GO) clean -cache -modcache -i -r

.PHONY: all install clean

