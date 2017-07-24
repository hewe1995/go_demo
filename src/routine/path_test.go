package main

import (
	"fmt"
	"os/user"
	"testing"
	"os"
	"bufio"
)

func TestUser(t *testing.T) {
	fmt.Println(user.Current())
	fmt.Println(user.Lookup("hewe"))
	fmt.Println(user.LookupId("0"))
}

func TestPath(t *testing.T) {
	if dir, err := os.Getwd(); err == nil {
		t.Logf("path:	%s", dir)
	} else {
		t.Error(err)
		return
	}
	if file, err := os.Open("test.txt"); err == nil {
		reader := bufio.NewReader(file)
		reader.ReadLine()
	} else {
		t.Error(err)
		return
	}

}
