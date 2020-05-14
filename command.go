package audacity

import (
	"encoding/json"
	"fmt"
	"log"
)

// Command sends a command and returns the response.
// The number of arguments must be even, in the form of key and value pairs.
func Command(cmd string, args ...string) (resp Response) {
	if k := len(args); k > 0 {
		cmd += ":"
		if k%2 != 0 {
			log.Fatal("[BUG] can't process command with odd number of arguments")
		}
		for i := 0; i < k; i = +2 {
			cmd += fmt.Sprintf(` %s="%s"`, args[i], args[i+1])
		}
	}

	// Send a single command.
	sendCmd(cmd)

	// decode response
	json.NewDecoder(fromPipe).Decode(&resp)
	fmt.Printf("Command: '%s'\nResponse: %s\n\n", cmd, resp)

	return resp
}

func sendCmd(cmd string) {
	if toPipe == nil {
		if err := initPipes(); err != nil {
			log.Printf("can't run audacity command %s: %s", cmd, err)
			return
		}
	}

	toPipe.Write([]byte(cmd + eol))
}

func command_debug(cmd string) {

	// Send a single command.
	sendCmd(cmd)

	// show message
	buf := make([]byte, 15)
	fromPipe.Read(buf)
	fmt.Printf("Debugging Command: '%s'\nResponse: %s\n\n", cmd, buf)

	// // Return the command resp.
	// str := ""
	// buf = make([]byte, 1)
	// done := -1
	// for done < 1 {
	// 	// BatchCommand finished: OK
	// 	fromPipe.Read(buf)
	// 	switch buf[0] {
	// 	case '\n':
	// 		done++
	// 	default:
	// 		done = -1
	// 	}
	// 	str += string(buf)
	// 	fmt.Println("MORE:", str)
	// }
}
