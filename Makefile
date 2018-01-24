include config.mk

xkcd936: xkcd936.c
	${CC} ${CFLAGS} -lsodium -o xkcd936 web2.c xkcd936.c

install:
	@echo Installing to to ${PREFIX}/bin ...
	@mkdir -p ${PREFIX}/bin
	@cp -f xkcd936 ${PREFIX}/bin
	@chmod 755 ${PREFIX}/bin/xkcd936

uninstall:
	@echo Uninstalling from ${PREFIX}/bin ...
	@rm -f ${PREFIX}/bin/xkcd936

clean:
	@rm -f xkcd936
