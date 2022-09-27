package dropbox

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

var config = dropbox.Config{
	Token: os.Getenv("DROPBOX_TOKEN"),
}

func Upload(f_name string, content io.Reader) error {
	// client := files.New(config)
	// arg := files.NewCommitInfo("/" + f_name)
	// if _, err := client.Upload(arg, content); err != nil {
	// 	return err
	// }
	return nil
}

func Download(f_name string) ([]byte, error) {
	client := files.New(config)
	arg := files.NewDownloadArg("/" + f_name)
	_, content, err := client.Download(arg)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
