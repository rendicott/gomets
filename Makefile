version := 0.2
packageNameNix := gomets-linux-amd64-$(version).tar.gz

build_dir := output
build_dir_linux := output-linux

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./$(build_dir_linux)/gomets -ldflags "-X main.version=$(version)"
	@cd ./$(build_dir_linux) && tar zcf ../$(build_dir)/$(packageNameNix) . 

deps:
	go get -t ./...

build: clean deps configure build-linux

configure:
	mkdir $(build_dir)
	mkdir $(build_dir_linux)

clean:
	rm -f *.coverprofile
	rm -rf $(build_dir)
	rm -rf $(build_dir_linux)

deploy: build
	scp output-linux/gomets coppy:/tmp
