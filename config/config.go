package config

import (
	"net"
	"time"

	"github.com/AmirMirzayi/relay/utils"
)

const (
	DefaultPort      = 55555
	DefaultTimeout   = 30 * time.Second
	DefaultChunkSize = 1024 * 1024
	ProgressbarWidth = 25
)

func DefaultIP() net.IP {
	return net.IPv4zero
}

type File struct {
	Name    string   `json:"name"`
	Size    int64    `json:"size"`
	Path    string   `json:"-"`
	Parents []string `json:"parents"`
}

type Files []File

func (fs Files) TotalSize() (totalSize int64) {
	for _, f := range fs {
		totalSize += f.Size
	}
	return
}

func (fs Files) HumanReadableTotalSize() string {
	totalSize := fs.TotalSize()
	return utils.ConvertByteSizeToHumanReadable(float64(totalSize))
}
