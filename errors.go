package docs_code

import "errors"

// ErrUnsupportedDocType is returned by Validate and Generate when the provided
// DocType is not one of the supported document types.
var ErrUnsupportedDocType = errors.New("unsupported document type")
