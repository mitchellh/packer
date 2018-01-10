package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestChooseString(t *testing.T) {
	cases := []struct {
		Input  []string
		Output string
	}{
		{
			[]string{"", "foo", ""},
			"foo",
		},
		{
			[]string{"", "foo", "bar"},
			"foo",
		},
		{
			[]string{"", "", ""},
			"",
		},
	}

	for _, tc := range cases {
		result := ChooseString(tc.Input...)
		if result != tc.Output {
			t.Fatalf("bad: %#v", tc.Input)
		}
	}
}

func TestDownloadableURL(t *testing.T) {
	// Invalid URL: has hex code in host
	_, err := DownloadableURL("http://what%20.com")
	if err == nil {
		t.Fatal("expected err")
	}

	// Invalid: unsupported scheme
	_, err = DownloadableURL("ftp://host.com/path")
	if err == nil {
		t.Fatal("expected err")
	}

	// Valid: http
	u, err := DownloadableURL("HTTP://packer.io/path")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if u != "http://packer.io/path" {
		t.Fatalf("bad: %s", u)
	}

	// No path
	u, err = DownloadableURL("HTTP://packer.io")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if u != "http://packer.io" {
		t.Fatalf("bad: %s", u)
	}
}

func TestDownloadableURL_WindowsFiles(t *testing.T) {
	if runtime.GOOS == "windows" {
		dirCases := []struct {
			InputString string
			OutputURL   string
			ErrExpected bool
		}{ // TODO: add different directories
			{
				"C:\\Temp\\SomeDir\\myfile.txt",
				"file:///C:/Temp/SomeDir/myfile.txt",
				false,
			},
			{ // need windows drive
				"\\Temp\\SomeDir\\myfile.txt",
				"",
				true,
			},
			{ // need windows drive
				"/Temp/SomeDir/myfile.txt",
				"",
				true,
			},
			{ // UNC paths; why not?
				"\\\\?\\c:\\Temp\\SomeDir\\myfile.txt",
				"",
				true,
			},
			{
				"file:///C:\\Temp\\SomeDir\\myfile.txt",
				"file:///c:/Temp/SomeDir/myfile.txt",
				false,
			},
			{
				"file:///c:/Temp/Somedir/myfile.txt",
				"file:///c:/Temp/SomeDir/myfile.txt",
				false,
			},
		}
		// create absolute-pathed tempfile to play with
		err := os.Mkdir("C:\\Temp\\SomeDir", 0755)
		if err != nil {
			t.Fatalf("err creating test dir: %s", err)
		}
		fi, err := os.Create("C:\\Temp\\SomeDir\\myfile.txt")
		if err != nil {
			t.Fatalf("err creating test file: %s", err)
		}
		fi.Close()
		defer os.Remove("C:\\Temp\\SomeDir\\myfile.txt")
		defer os.Remove("C:\\Temp\\SomeDir")

		// Run through test cases to make sure they all parse correctly
		for _, tc := range dirCases {
			u, err := DownloadableURL(tc.InputString)
			if (err != nil) != tc.ErrExpected {
				t.Fatalf("Test Case failed: Expected err = %#v, err = %#v, input = %s",
					tc.ErrExpected, err, tc.InputString)
			}
			if u != tc.OutputURL {
				t.Fatalf("Test Case failed: Expected %s but received %s from input %s",
					tc.OutputURL, u, tc.InputString)
			}
		}
	}
}

func TestDownloadableURL_FilePaths(t *testing.T) {
	tf, err := ioutil.TempFile("", "packer")
	if err != nil {
		t.Fatalf("tempfile err: %s", err)
	}
	defer os.Remove(tf.Name())
	tf.Close()

	tfPath, err := filepath.EvalSymlinks(tf.Name())
	if err != nil {
		t.Fatalf("tempfile err: %s", err)
	}

	tfPath = filepath.Clean(tfPath)

	filePrefix := "file://"
	if runtime.GOOS == "windows" {
		filePrefix += "/"
	}

	// Relative filepath. We run this test in a func so that
	// the defers run right away.
	func() {
		wd, err := os.Getwd()
		if err != nil {
			t.Fatalf("getwd err: %s", err)
		}

		err = os.Chdir(filepath.Dir(tfPath))
		if err != nil {
			t.Fatalf("chdir err: %s", err)
		}
		defer os.Chdir(wd)

		filename := filepath.Base(tfPath)
		u, err := DownloadableURL(filename)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		expected := fmt.Sprintf("%s%s",
			filePrefix,
			strings.Replace(tfPath, `\`, `/`, -1))
		if u != expected {
			t.Fatalf("unexpected: %#v != %#v", u, expected)
		}
	}()

	// Test some cases with and without a schema prefix
	for _, prefix := range []string{"", filePrefix} {
		// Nonexistent file
		_, err = DownloadableURL(prefix + "i/dont/exist")
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		// Good file
		u, err := DownloadableURL(prefix + tfPath)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		expected := fmt.Sprintf("%s%s",
			filePrefix,
			strings.Replace(tfPath, `\`, `/`, -1))
		if u != expected {
			t.Fatalf("unexpected: %s != %s", u, expected)
		}
	}
}

func TestScrubConfig(t *testing.T) {
	type Inner struct {
		Baz string
	}
	type Local struct {
		Foo string
		Bar string
		Inner
	}
	c := Local{"foo", "bar", Inner{"bar"}}
	expect := "Config: {Foo:foo Bar:<Filtered> Inner:{Baz:<Filtered>}}"
	conf := ScrubConfig(c, c.Bar)
	if conf != expect {
		t.Fatalf("got %s, expected %s", conf, expect)
	}
}
