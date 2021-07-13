package media

import (
	"github.com/mohammadMghi/auth/base"
	"github.com/mohammadMghi/auth/download"
	"github.com/mohammadMghi/auth/upload"
)

type Config struct {
	base.Config

	Upload   *upload.Config
	Download *download.Config
}
