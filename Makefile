.PHONY: build clean

build: clean
	npm update
	mkdir build_out
	go build -o build_out/entertainment .
	npm run css-prod
	cp -r assets build_out/
	rm -rf build_out/assets/build

clean:
	rm -rf build_out

