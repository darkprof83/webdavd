WEBDAVD=webdavd
WDCMD=cmd/$(WEBDAVD)/$(WEBDAVD)

WDHASH=wdhash
HASHCMD=cmd/$(WDHASH)/$(WDHASH)

rwildcard=$(wildcard $1$2) $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2))

.PHONY: all
all: $(WDCMD) $(HASHCMD)

$(WDCMD): cmd/$(WEBDAVD)/*.go $(call rwildcard,internal/,*.go)
	go build -C cmd/$(WEBDAVD)

$(HASHCMD): cmd/$(WDHASH)/*.go $(call rwildcard,internal/,*.go)
	go build -C cmd/$(WDHASH)

.PHONY: install
install:
	install -Dm755 -t "$(DESTDIR)/usr/bin/" $(WDCMD)
	install -Dm755 -t "$(DESTDIR)/usr/bin/" $(HASHCMD)

.PHONY: uninstall
uninstall:
	rm -f "$(DESTDIR)/usr/bin/$(WEBDAVD)"
	rm -f "$(DESTDIR)/usr/bin/$(WDHASH)"

.PHONY: clean
clean:
	rm -f $(WDCMD)
	rm -f $(HASHCMD)
