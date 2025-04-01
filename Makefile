VERSION := v0.25.2

default:
	fd --exclude=Makefile --exclude=go.mod --exclude=Readme.md . './' | xargs -- rm -rf
	git clone --depth 1 --branch $(VERSION) -c advice.detachedHead=false https://github.com/evanw/esbuild
	fd --extension='.go' . 'esbuild/internal' | xargs -- sd -s 'github.com/evanw/esbuild/internal' 'github.com/matthewmueller/esbuild_internal'
	cp ./esbuild/LICENSE.md ./
	mv ./esbuild/internal/* ./
	rm -rf esbuild
	go get -u github.com/evanw/esbuild@$(VERSION)
	go mod tidy
	go test ./...
	if [ `git tag -l "$(VERSION)"` ]; then echo "tag $(VERSION) exists already"; exit 1; fi
	git add .
	git commit -m "Upgrade esbuild to $(VERSION)"
	git tag $(VERSION)
