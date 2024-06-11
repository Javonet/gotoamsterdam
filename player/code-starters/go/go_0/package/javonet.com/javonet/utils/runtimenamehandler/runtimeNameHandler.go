package runtimenamehandler

import (
	"errors"

	"javonet.com/javonet/utils/runtimename"
)

func GetName(runtimeName byte) (string, error) {
	switch runtimeName {
	case runtimename.Clr:
		return "clr", nil
	case runtimename.Go:
		return "go", nil
	case runtimename.Jvm:
		return "jvm", nil
	case runtimename.Netcore:
		return "netcore", nil
	case runtimename.Perl:
		return "perl", nil
	case runtimename.Python:
		return "python", nil
	case runtimename.Ruby:
		return "ruby", nil
	case runtimename.Nodejs:
		return "nodejs", nil
	case runtimename.Cpp:
		return "cpp", nil
	default:
		return "", errors.New("invalid runtime name")
	}
}
