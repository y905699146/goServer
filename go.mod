module goServer

go 1.12

replace (
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190626221950-04f50cda93cb
	golang.org/x/text => github.com/golang/text latest
	golang.org/x/image => github.com/golang/image latest
)

require (
	github.com/golang/protobuf v1.3.1
	github.com/sirupsen/logrus v1.4.2
)
