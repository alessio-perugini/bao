package bao

import (
	"testing"
)

func TestExtractNIpFromString(t *testing.T) {
	ExtractNIpFromString("auth.log:Apr 26 06:26:05 vps681577 182.61.2.67 sshd[11904]: Failed password for invalid user cstrike from 182.61.2.67 port 44624 ssh2", 1)
}
