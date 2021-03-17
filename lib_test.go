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

func Test_LibraryImpl_DeleteUser(t *testing.T) {
	libs, err := _testCli.NewLibrary("test", false)
	merrors.AssertError(err, "new lib failed")
	err = libs.DeleteUser("su")
	merrors.AssertError(err, "delete user failed")
}

func Test_LibraryImpl_ListUsers(t *testing.T) {
	libs, err := _testCli.NewLibrary("test", false)
	merrors.AssertError(err, "new lib failed")
	ulist, err := libs.ListUsers(0, 100)
	merrors.AssertError(err, "delete user failed")
	dump(t, ulist)
}

func Test_LibraryImpl_ListUserFaces(t *testing.T) {
	libs, err := _testCli.NewLibrary("test", false)
	merrors.AssertError(err, "new lib failed")
	ll, err := libs.ListUserFaces("su")
	merrors.AssertError(err, "list user face failed")
	dump(t, ll)
}

func Test_LibraryImpl_DeleteFace(t *testing.T) {
	libs, err := _testCli.NewLibrary("test", false)
	merrors.AssertError(err, "new lib failed")
	err = libs.DeleteFace("su", "${which face token}")
	merrors.AssertError(err, "delete user failed")
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
