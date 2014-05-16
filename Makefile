APPNAME:=deansify

default: clean
	script/release

clean:
	rm -f ./$(APPNAME)*gz
	rm -f ./$(APPNAME)*zip
	rm -f ./deansify
