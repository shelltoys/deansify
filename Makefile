APPNAME:=deansify

default: clean
	script/release

clean:
	rm -f ./$(APPNAME).*
