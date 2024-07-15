package _0_simplifyPath

import "testing"

func TestSimplifyPath(t *testing.T) {

	path1 := "/home/"
	path2 := "/../"
	path3 := "/home//foo/"
	path4 := "/a/./b/../../c/"

	path5 := "//./a//../b"

	println(simplifyPath(path1))
	println(simplifyPath(path2))
	println(simplifyPath(path3))
	println(simplifyPath(path4))
	println(simplifyPath(path5))

}
