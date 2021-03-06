package gtc

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestFile(t *testing.T) {
	var dir = "./tmptest"
	if PathExist(dir) != false {
		t.Fatal("fail PathExist")
	}
	if PathNotExist(dir) != true {
		t.Fatal("fail PathNotExist")
	}
	if IsDir(dir) != false {
		t.Fatal("fail IsDir")
	}

	err := CreateDir("afhifeiwIHQNFBLqwj/b/c")
	if err == nil {
		t.Fatal("create cannot deep")
	}

	deepdir := filepath.Join("a", "b", "c")
	err = CreateAllDir(deepdir)
	if err != nil {
		t.Fatal("create all dir fail")
	}
	if !IsDir(deepdir) {
		t.Fatal("test all dir fail")
	}
	defer os.RemoveAll("a")

	err = CreateDir(dir)
	if err != nil {
		t.Fatalf("create dir fail: %s", dir)
	}
	defer os.Remove(dir)

	if runtime.GOOS == "linux" {
		fi, err := os.Stat(dir)
		if err != nil {
			t.Fatal(err)
		}
		if fi.Mode().String() != "drwxr-xr-x" {
			t.Error("createdir permission error")
		}
	}
	if PathExist(dir) != true {
		t.Fatal("after fail PathExist")
	}
	if PathNotExist(dir) != false {
		t.Fatal("after fail PathNotExist")
	}
	if IsDir(dir) != true {
		t.Fatal("after fail IsDir")
	}

	f := "main.go"
	if IsFile(f) != true {
		t.Fatal("fail IsFile")
	}
	if IsFile("/tmp/this_is_a_not_exist_file") != false {
		t.Fatal("fail IsFile No.2")
	}

	shm := "/dev/shm"
	if PathExist(shm) {
		if IsDir(shm) != true {
			t.Fatal("/dev/shm is dir")
		}
	}

	nf := "/dev/null"
	if IsFile(nf) {
		if IsCommonFile(nf) {
			t.Fatal("/dev/null is not common file")
		}
	}

	src := "go.mod"
	dst := "go.mod.bak"
	_, err = FileCopy(dst, src)
	if err != nil {
		t.Log("fail FileCopy")
		t.Fatal(err)
	}
	defer os.Remove(dst)

	if IsFile(dst) != true {
		t.Fatal("created dst, but fail IsFile")
	}
	srcText, err := FileReadByte(src)
	if err != nil {
		t.Fatal("fail FileReadByte")
	}
	dstText, err := FileReadStr(dst)
	if err != nil {
		t.Fatal("fail FileReadStr")
	}
	if string(srcText) != dstText {
		t.Fatal("fail FileCopy, src and dst are inconsistent")
	}

	dstN := "go.mod.bak.N"
	_, err = FileCopyN(dstN, src, 6)
	if err != nil {
		t.Log("fail FileCopyN")
		t.Fatal(err)
	}
	defer os.Remove(dstN)
	dstnText, err := FileReadStr(dstN)
	if err == nil {
		if dstnText != "module" {
			t.Fatal("fail FileCopyN, invalid bytes")
		}
	}
}

func TestBool(t *testing.T) {
	if IsTrue("1") != true {
		t.Fatal("1 is true")
	}
	if IsTrue("t") != true {
		t.Fatal("t is true")
	}
	if IsTrue("T") != true {
		t.Fatal("T is true")
	}
	if IsTrue("true") != true {
		t.Fatal("true is true")
	}
	if IsTrue("True") != true {
		t.Fatal("True is true")
	}
	if IsTrue("TRUE") != true {
		t.Fatal("TRUE is true")
	}
	if IsTrue("abc") == true {
		t.Fatal("abc not true")
	}
	if IsTrue("") == true {
		t.Fatal("empty not true")
	}
	if IsTrue("on") != true {
		t.Fatal("on is true")
	}

	if NotTrue("0") != true {
		t.Fatal("0 is false")
	}
	if NotTrue("f") != true {
		t.Fatal("f is false")
	}
	if NotTrue("F") != true {
		t.Fatal("F is false")
	}
	if NotTrue("false") != true {
		t.Fatal("false is false")
	}
	if NotTrue("False") != true {
		t.Fatal("False is false")
	}
	if NotTrue("FALSE") != true {
		t.Fatal("FALSE is false")
	}
	if NotTrue("abc") != true {
		t.Fatal("abc is false")
	}
	if NotTrue("") != true {
		t.Fatal("empty is false")
	}

	if IsFalse("0") != true {
		t.Fatal("0 is false")
	}
	if IsFalse("f") != true {
		t.Fatal("f is false")
	}
	if IsFalse("F") != true {
		t.Fatal("F is false")
	}
	if IsFalse("false") != true {
		t.Fatal("false is false")
	}
	if IsFalse("False") != true {
		t.Fatal("False is false")
	}
	if IsFalse("FALSE") != true {
		t.Fatal("FALSE is false")
	}
	if IsFalse("abc") != false {
		t.Fatal("abc is not false")
	}
	if IsFalse("") != false {
		t.Fatal("empty is not false")
	}
	if IsFalse("oFF") != true {
		t.Fatal("off is false")
	}
}

func TestString(t *testing.T) {
	s := []string{"1", "", "a"}
	if StrInSlice("", s) != true {
		t.Fatal("empty string in slice")
	}
	if StrInSlice("1", s) != true {
		t.Fatal("1 in slice")
	}
	if StrInSlice("a", s) != true {
		t.Fatal("a in slice")
	}
	if StrInSlice("b", s) == true {
		t.Fatal("b not in slice")
	}

	has, index := InArraySlice("", s)
	if has != true {
		t.Fatal("empty string in InArraySlice")
	}
	if index != 1 {
		t.Fatal("empty index error")
	}

	has, _ = InArraySlice("1", s)
	if has != true {
		t.Fatal("1 in InArraySlice")
	}

	has, _ = InArraySlice("a", s)
	if has != true {
		t.Fatal("a in InArraySlice")
	}

	has, _ = InArraySlice("b", s)
	if has == true {
		t.Fatal("b not in InArraySlice")
	}

	has, _ = InArraySlice(1, s)
	if has == true {
		t.Fatal("1(number) not in InArraySlice")
	}

	has, _ = InArraySlice("a", "abc")
	if has == true {
		t.Fatal("a not in abc(string), InArraySlice not allow string")
	}

	has, _ = InArraySlice("a", map[string]int{"a": 1})
	if has == true {
		t.Fatal("a not in map, InArraySlice not allow map")
	}

	var s2 [][]string
	s2 = append(s2, s)
	s2 = append(s2, []string{"b"})
	s2 = append(s2, []string{"1"})

	has, _ = InArraySlice(s, s2)
	if has != true {
		t.Fatal("s in s2")
	}

	has, _ = InArraySlice([]string{"b"}, s2)
	if has != true {
		t.Fatal("[]string{b} in s2")
	}

	has, _ = InArraySlice([]string{"c"}, s2)
	if has == true {
		t.Fatal("[]string{c} not in s2")
	}

	has, _ = InArraySlice([]string{"1"}, s2)
	if has != true {
		t.Fatal("[]string{1} in s2")
	}

	has, _ = InArraySlice([]int{1}, s2)
	if has == true {
		t.Fatal("[]int{1} not in s2")
	}

	sf := []string{"a", "b"}
	if FindSlice(sf, "a") != 0 {
		t.Fatal("err, it should be 0 for FindSlice")
	}
	if FindSlice(sf, "c") != -1 {
		t.Fatal("err, not found, -1 for FindSlice")
	}

	ss1 := "abcd"
	ss2 := "你好，世界"
	if SubStr(ss1, 0, 2) != "ab" {
		t.Fatal("err for SubStr test 1")
	}
	if SubStr(ss2, 1, 3) != "好，" {
		t.Fatal("err for SubStr test 2")
	}
}

func TestTool(t *testing.T) {
	md5k := "hello world!"
	md5v := "fc3ff98e8c6a0d3087d515c0473f8677"
	if MD5(md5k) != md5v {
		t.Fatal("md5 fail")
	}
	md5f, err := ioutil.TempFile("", "md5-test.txt")
	if err != nil {
		t.Fatal("tempfile error")
	}
	md5f.Write([]byte(md5k))
	md5f.Close()
	md5fv, err := MD5File(md5f.Name())
	if err != nil {
		t.Fatal("md5file raise error")
	}
	if md5v != md5fv {
		t.Fatal("md5file result error")
	}
}

func BenchmarkSubStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStr("abaghjpiowpoejgre8786awef86wer8962", 3, 15)
	}
}
