

build-java: 
	mkdir -p testdata/classes
	javac -d testdata/classes $$(find testdata -name *.java)
	jar -c -f testdata/dist/test.jar --main-class main.Main -C testdata/classes .
	cd testdata/dist && jar xf test.jar 

test: build-java
	go test ./...