package audacity

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	winio "github.com/Microsoft/go-winio"
)

var (
	toPipe   io.WriteCloser
	fromPipe io.ReadCloser
	eol      string
)

func initPipes() (err error) {

	errFormat := "dialing %s pipe failed: %s\nPossible solution: Ensure Audacity is running with mod-script-pipe, see: https://manual.audacityteam.org/man/scripting.html#Enable_mod-script-pipe."

	for tries := 0; tries < 2; tries++ {

		if tries == 1 {
			// TODO start audacity via command
			// "C:\Program Files (x86)\Audacity\audacity.exe"
		}

		if runtime.GOOS == "windows" {
			eol = "\r\n"

			t := time.Minute

			// TODO winio still necessary?
			if fromPipe, err = winio.DialPipe("\\\\.\\pipe\\FromSrvPipe", &t); err != nil {
				if tries == 0 {
					continue
				}
				return fmt.Errorf(errFormat, err, "sending")
			}

			if toPipe, err = winio.DialPipe("\\\\.\\pipe\\ToSrvPipe", &t); err != nil {
				return fmt.Errorf(errFormat, err, "receiving")
			}

		} else {
			toName := "/tmp/audacity_script_pipe.to." + string(os.Getuid())
			fromName := "/tmp/audacity_script_pipe.from." + string(os.Getuid())
			eol = "\n"

			if toPipe, err = os.Open(toName); err != nil {
				if tries == 0 {
					continue
				}
				return fmt.Errorf(errFormat, err, "sending")
			}

			if fromPipe, err = os.Open(fromName); err != nil {
				return fmt.Errorf(errFormat, err, "receiving")
			}
		}
	}
	return nil
}

// ClosePipes closes the pipes to Audacity. Make sure you call this function when you are done.
func ClosePipes() {
	if toPipe != nil {
		toPipe.Close()
		toPipe = nil
	}
	if fromPipe != nil {
		fromPipe.Close()
		fromPipe = nil
	}
}
