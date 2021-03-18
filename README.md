## 百度人脸识别SDK

> 注意:这不是百度官方的开发包

### Install

```go
go get github.com/guestin/go-duface
```

### Quick Start

#### detect a image

```go
func Test_ClientImpl_FaceDetect(t *testing.T) {
  imgData := readImage("me.jpeg")
  r, err := _testCli.FaceDetect(
    imgData, nil)
  merrors.AssertError(err, "detect failed")
  dump(t, r)
}
```

#### list libraries

```go
func Test_ClientImpl_ListLibraries(t *testing.T) {
libs, err := _testCli.ListLibraries(0, 100)
merrors.AssertError(err, "list libs failed")
dump(t, libs)
}
```

#### create library

```go
func Test_LibraryImpl_Create(t *testing.T) {
libs, err := _testCli.NewLibrary("test")
merrors.AssertError(err, "new lib failed")
err = libs.Create()
merrors.AssertError(err, "create library failed")
}
```

#### register face

```go
func Test_LibraryImpl_RegisterFace(t *testing.T) {
libs, err := _testCli.NewLibrary("test", false)
merrors.AssertError(err, "new lib failed")
imgData := readImage("me.jpeg")
r, err := libs.RegisterFace("su", imgData, nil)
merrors.AssertError(err, "register facefailed")
  dump(t, r)
}
```

#### list user faces

```go
func Test_LibraryImpl_ListUserFaces(t *testing.T) {
  libs, err := _testCli.NewLibrary("test", false)
  merrors.AssertError(err, "new lib failed")
  ll, err := libs.ListUserFaces("su")
  merrors.AssertError(err, "list user face failed")
  dump(t, ll)
}
```

#### list users

```go
func Test_LibraryImpl_ListUsers(t *testing.T) {
  libs, err := _testCli.NewLibrary("test", false)
  merrors.AssertError(err, "new lib failed")
  ulist, err := libs.ListUsers(0, 100)
  merrors.AssertError(err, "delete user failed")
  dump(t, ulist)
}
```

#### delete a user

```go
func Test_LibraryImpl_DeleteUser(t *testing.T) {
  libs, err := _testCli.NewLibrary("test", false)
  merrors.AssertError(err, "new lib failed")
  err = libs.DeleteUser("su")
  merrors.AssertError(err, "delete user failed")
}
```

#### delete user's face

```go
func Test_LibraryImpl_DeleteFace(t *testing.T) {
  libs, err := _testCli.NewLibrary("test", false)
  merrors.AssertError(err, "new lib failed")
  err = libs.DeleteFace("su", "${which face token}")
  merrors.AssertError(err, "delete user failed")
}
```

#### 1:N search

```go

func Test_LibraryImpl_Search(t *testing.T) {
  libs, err := _testCli.NewLibrary("test", false)
  merrors.AssertError(err, "new lib failed")
  imgData := readImage("me.jpeg")
  r, err := libs.Search(imgData, nil)
  merrors.AssertError(err, "search failed")
  dump(t, r)
}
```

#### M:N search

```go
func Test_LibraryImpl_MultiSearch(t *testing.T) {
libs, err := _testCli.NewLibrary("test", false)
merrors.AssertError(err, "new lib failed")
imgData := readImage("me.jpeg")
r, err := libs.MultiSearch(imgData, nil)
merrors.AssertError(err, "search failed")
dump(t, r)
}
```

#### drop library

```go
func Test_LibraryImpl_Drop(t *testing.T) {
libs, err := _testCli.NewLibrary("test")
merrors.AssertError(err, "new lib failed")
err = libs.Drop()
merrors.AssertError(err, "drop library failed")
}
```

### About US

[南京客町网络科技有限公司](https://www.guestin.cn)

### License

MIT