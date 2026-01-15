package clipboard

import (
	"github.com/pkg/errors"
	"github.com/zyedidia/clipboard"
)

// Read returns the current text content from the system clipboard.
// Returns an error if the clipboard cannot be accessed or is empty.
func Read() ([]byte, error) {
	// Initialize clipboard
	if err := clipboard.Initialize(); err != nil {
		return nil, errors.Wrap(err, "failed to initialize clipboard access")
	}

	// Read text content
	content, err := clipboard.ReadAll("clipboard")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read from clipboard")
	}

	if len(content) == 0 {
		return nil, errors.New("clipboard is empty")
	}

	return []byte(content), nil
}
