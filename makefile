include $(GOROOT)/src/Make.inc
TARG=gochan
GOFILES=\
	lib/thread.go\
	main.go
include $(GOROOT)/src/Make.cmd
