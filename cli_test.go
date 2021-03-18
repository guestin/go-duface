package duface

import (
	_ "embed"
	"encoding/json"
)

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/guestin/mob/merrors"
	"os"
	"testing"
)

//go:embed test_api_key.txt
var kTestApiKey string

//go:embed test_app_secret_key.txt
var kTestAppSecretKey string

var _testCli = NewClient(context.TODO(), kTestApiKey, kTestAppSecretKey)

func Test_ClientImpl_FaceDetect(t *testing.T) {
	imgData := readImage("me.jpeg")
	r, err := _testCli.FaceDetect(
		imgData, nil)
	merrors.AssertError(err, "detect failed")
	dump(t, r)
}

func Test_ClientImpl_ListLibraries(t *testing.T) {
	libs, err := _testCli.ListLibraries(0, 100)
	merrors.AssertError(err, "list libs failed")
	dump(t, libs)
}

// test tools

func readImage(fileName string) *ImageData {
	imgBytes, err := os.ReadFile(fileName)
	merrors.AssertError(err, "read file failed")
	imgBase64 := base64.StdEncoding.EncodeToString(imgBytes)
	return &ImageData{
		Data: imgBase64,
		Type: BASE64,
	}
}

func dump(t testing.TB, v interface{}) {
	vbyts, err := json.MarshalIndent(v, "", "  ")
	merrors.AssertError(err, "marshal to json failed")
	fmt.Printf("test case: [%s]\n%s\n", t.Name(), vbyts)
}
