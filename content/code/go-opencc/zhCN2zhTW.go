package mylib

import "github.com/siongui/go-opencc"

func CN2TW(input string) string {
	c := opencc.NewConverter("zhs2zhtw_vp.ini")
	defer c.Close()
	return c.Convert(input)
}
