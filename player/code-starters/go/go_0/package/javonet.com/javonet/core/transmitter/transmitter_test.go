//go:build performanceTests
// +build performanceTests

package transmitter

import (
	"runtime"
	"testing"

	"javonet.com/javonet/utils/binariesextractor"
	"javonet.com/javonet/utils/commandtype"
	"javonet.com/javonet/utils/connectiontype"
	"javonet.com/javonet/utils/runtimename"
	"javonet.com/javonet/utils/runtimenamehandler"
)

func TestMain(m *testing.M) {
	ActivateWithCredentials("n9B5-Km7g-Pp69-j9FE-e9A5")
	m.Run()
}

func produceInMemoryHeartBeatMessage(runtimeNumber byte) []byte {
	return []byte{runtimeNumber, 0, connectiontype.InMemory, 0, 0, 0, 0, 0, 0, runtimename.Go, commandtype.HeartBeat}
}

func sendCommand(t *testing.T, runtimeNumber byte) {
	be := binariesextractor.NewBinariesExtractor()
	runtimeName, _ := runtimenamehandler.GetName(runtimeNumber)
	be.ExtractFromBinariesZip(runtimeName)
	messageByteArray := produceInMemoryHeartBeatMessage(runtimeNumber)
	responseByteArray, err := SendCommand(messageByteArray)
	if err != nil {
		t.Fatal(t.Name() + " failed: " + err.Error())
	}
	if len(responseByteArray) != 2 {
		t.Fatal(t.Name() + " failed: wrong responseByteArray length")
	}
	if responseByteArray[0] != 49 || responseByteArray[1] != 48 {
		t.Fatal(t.Name() + " failed: wrong responseByteArray elements")
	}
}

func performance(t *testing.T, runtimeNumber byte) {
	messageByteArray := produceInMemoryHeartBeatMessage(runtimeNumber)
	sendCommand(t, runtimeNumber)
	for i := 0; i < 250000; i++ {
		SendCommand(messageByteArray)
	}
}

func TestHeartBeat_ClrRuntime(t *testing.T) {
	if runtime.GOOS == "windows" {
		sendCommand(t, runtimename.Clr)

	} else {
		t.Skip("CLR not implemented on Linux and MacOs")
	}
}

func TestHeartBeat_JvmRuntime(t *testing.T) {
	sendCommand(t, runtimename.Jvm)
}

func TestHeartBeat_NetcoreRuntime(t *testing.T) {
	sendCommand(t, runtimename.Netcore)
}

func TestHeartBeat_PerlRuntime(t *testing.T) {
	if runtime.GOOS != "darwin" {
		sendCommand(t, runtimename.Perl)
	} else {
		t.Skip("Perl not implemented on MacOs")
	}
}

func TestHeartBeat_PythonRuntime(t *testing.T) {
	if runtime.GOOS != "darwin" {
		sendCommand(t, runtimename.Python)
	} else {
		t.Skip("Python not implemented on MacOs")
	}
}

func TestHeartBeat_RubyRuntime(t *testing.T) {
	sendCommand(t, runtimename.Ruby)
}

func TestHeartBeat_NodejsRuntime(t *testing.T) {
	sendCommand(t, runtimename.Nodejs)
}

func TestPerformance_ClrRuntime(t *testing.T) {
	if runtime.GOOS == "windows" {
		performance(t, runtimename.Clr)
	} else {
		t.Skip("CLR not implemented on Linux and MacOs")
	}
}

// func TestPerformance_JvmRuntime(t *testing.T) {
// 	performance(t, runtimename.Jvm)
// }

func TestPerformance_NetcoreRuntime(t *testing.T) {
	performance(t, runtimename.Netcore)
}

func TestPerformance_PerlRuntime(t *testing.T) {
	if runtime.GOOS != "darwin" {
		performance(t, runtimename.Perl)
	} else {
		t.Skip("Perl not implemented on MacOs")
	}
}

func TestPerformance_PythonRuntime(t *testing.T) {
	if runtime.GOOS != "darwin" {
		performance(t, runtimename.Python)
	} else {
		t.Skip("Python not implemented on MacOs")
	}
}

// func TestPerformance_RubyRuntime(t *testing.T) {
// 	if runtime.GOOS == "windows" {
// 		performance(t, runtimename.Ruby)
// 	} else {
// 		t.Skip("Ruby does not work on Linux and MacOs")
// 	}
// }

func TestPerformance_NodejsRuntime(t *testing.T) {
	performance(t, runtimename.Nodejs)
}
