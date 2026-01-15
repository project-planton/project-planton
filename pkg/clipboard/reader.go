package clipboard

import (
	"github.com/pkg/errors"
	"golang.design/x/clipboard"
)

// Read returns the current text content from the system clipboard.
// Returns an error if the clipboard cannot be accessed or is empty.
func Read() ([]byte, error) {
	if err := clipboard.Init(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize clipboard access")
	}

	content := clipboard.Read(clipboard.FmtText)
	if len(content) == 0 {
		return nil, errors.New("clipboard is empty")
	}

	return content, nil
}
