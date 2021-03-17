package duface

import (
	"github.com/guestin/mob/merrors"
	"testing"
)

func Test_LibraryImpl_RegisterFace(t *testing.T) {
	libs, err := _testCli.NewLibrary("test", false)
	merrors.AssertError(err, "new lib failed")
	imgData := readImage("me.jpeg")
	r, err := libs.RegisterFace("su", imgData, nil)
	merrors.AssertError(err, "register facefailed")
	dump(t, r)
}

func Test_LibraryImpl_Search(t *testing.T) {
	libs, err := _testCli.NewLibrary("test", false)
	merrors.AssertError(err, "new lib failed")
	imgData := readImage("me.jpeg")
	r, err := libs.Search(imgData, nil)
	merrors.AssertError(err, "search failed")
	dump(t, r)
}

func Test_LibraryImpl_MultiSearch(t *testing.T) {
	libs, err := _testCli.NewLibrary("test", false)
	merrors.AssertError(err, "new lib failed")
	imgData := readImage("me.jpeg")
	r, err := libs.MultiSearch(imgData, nil)
	merrors.AssertError(err, "search failed")
	dump(t, r)
}
