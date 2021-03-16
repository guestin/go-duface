package duface

import (
	_ "embed"
)

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/guestin/mob/merrors"
	"github.com/guestin/mob/mio"
	"io"
	"os"
	"testing"
)

//go:embed test_api_key.txt
var kTestApiKey string

//go:embed test_app_secret_key.txt
var kTestAppSecretKey string

var _testCli = NewClient(context.TODO(), kTestApiKey, kTestAppSecretKey)

func Test_ClientImpl_FaceDetect(t *testing.T) {
	f, err := os.Open("me.jpeg")
	merrors.AssertError(err, "open failed")
	defer mio.CloseIgnoreErr(f)
	faceBytes, err := io.ReadAll(f)
	merrors.AssertError(err, "read file failed")
	imgBase64 := base64.StdEncoding.EncodeToString(faceBytes)
	detectResult, err := _testCli.FaceDetect(
		&ImageData{
			Data: imgBase64,
			Type: BASE64,
		}, nil)
	merrors.AssertError(err, "detect failed")
	fmt.Println(*detectResult)
}

func Test_ClientImpl_ListLibraries(t *testing.T) {
	libs, err := _testCli.ListLibraries(0, 100)
	merrors.AssertError(err, "list libs failed")
	fmt.Println(libs)
}
