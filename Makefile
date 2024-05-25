

.PHONY: build-java
build-java: 
	mkdir -p testdata/classes
	javac -d testdata/classes $$(find testdata -name *.java)
	mkdir -p testdata/dist
	jar -c -f testdata/dist/test.jar --main-class main.Main -C testdata/classes .
	cd testdata/dist && jar xf test.jar 

.PHONY: test
test: build-java
	go test ./...