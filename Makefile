include config.mk

bin/xkcd936: main.go xkcd936/xkcd936.go
	@go build -o bin/xkcd936

install: bin/xkcd936
	@mkdir -p ${PREFIX}/bin
	@install bin/xkcd936 ${PREFIX}/bin

uninstall:
	@rm -f ${PREFIX}/bin/xkcd936

clean:
	@rm -f bin/xkcd936

test:
	@cd xkcd936 && go test

