//
// Warning:
//   This file is generated by go compiler, don't change it!!!
//   Build on: Deepin GNU/Linux 20.8 \n \l
//
package typex

import "fmt"

type Version struct {
	Version     string
	ReleaseTime string
}

func (v Version) String() string {
	return fmt.Sprintf("{\"releaseTime\":\"%s\",\"version\":\"%s\"}", v.ReleaseTime, v.Version)
}

var DefaultVersion = Version{
	Version:   `v0.4.4`,
	ReleaseTime: "2023-04-17 09:41:25",
}

