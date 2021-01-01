package shell

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"runtime"
)

func OpenInBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		return errors.Wrapf(err, "Error opening url in browser: %s", url)
	}
	return nil
}
