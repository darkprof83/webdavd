WEBDAVD=webdavd
WDCMD=cmd/$(WEBDAVD)/$(WEBDAVD)

WDHASH=wdhash
HASHCMD=cmd/$(WDHASH)/$(WDHASH)

.PHONY: all
all: $(WDCMD) $(HASHCMD)

$(WDCMD): cmd/$(WEBDAVD)/*.go internal/logger/*.go internal/config/*.go
	go build -C cmd/$(WEBDAVD)

$(HASHCMD): cmd/$(WDHASH)/*.go
	go build -C cmd/$(WDHASH)
