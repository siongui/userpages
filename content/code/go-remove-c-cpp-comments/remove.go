package rmcmt

import (
	"io/ioutil"
	"regexp"
)

func RemoveCStyleComments(content []byte) []byte {
	// http://blog.ostermiller.org/find-comment
	ccmt := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)
	return ccmt.ReplaceAll(content, []byte(""))
}

func RemoveCppStyleComments(content []byte) []byte {
	cppcmt := regexp.MustCompile(`//.*`)
	return cppcmt.ReplaceAll(content, []byte(""))
}

func RemoveCAndCppComments(srcpath, dstpath string) {
	b, err := ioutil.ReadFile(srcpath)
	if err != nil {
		panic(err)
	}
	b = RemoveCppStyleComments(RemoveCStyleComments(b))
	err = ioutil.WriteFile(dstpath, b, 0644)
	if err != nil {
		panic(err)
	}
}
