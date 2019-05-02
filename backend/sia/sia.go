package sia

import (
	"fmt"

	"github.com/ncw/rclone/fs"
	"github.com/ncw/rclone/fs/config/configmap"
	"github.com/ncw/rclone/fs/config/configstruct"
	"github.com/ncw/rclone/fs/config/obscure"
	"github.com/ncw/rclone/fs/hash"

	httpclient "gitlab.com/NebulousLabs/Sia/node/api/client"
)

func init() {
	fs.Register(&fs.RegInfo{
		Name:        "sia",
		Description: "Sia",
		NewFs:       NewFs,
		Options: []fs.Option{
			{
				Name:     "endpoint",
				Help:     "The Sia API endpoint to use",
				Default:  "http://localhost:9980",
				Required: true,
				Advanced: true,
			}, {
				Name:     "folder",
				Help:     "folder on Sia to use. If omitted, the root folder is used.",
				Required: false,
				Default:  "",
				Advanced: true,
			}, {
				Name:       "apipassword",
				Help:       "Your api password for Sia.",
				IsPassword: true,
				Required:   true,
			},
		},
	})
}

// Options represent the configuration of the Koofr backend
type Options struct {
	Endpoint string `config:"endpoint"`
	Folder   string `config:"folder"`
	Password string `config:"apipassword"`
}
