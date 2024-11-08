package common

import "testing"

func TestNodeVersion(t *testing.T) {
	tests := []struct {
		major, minor, patch uint64
		err                 bool
		v                   Version
	}{
		{0, 0, 0, false, 0},
		{0, 10, 0, false, 10000},
		{0, 10, 999, false, 10999},
		{2, 0, 0, false, 2000000},
		{3, 100, 36, false, 3100036},
		{1000, 100, 36, true, 0},
		{3, 2000, 36, true, 0},
		{3, 100, 6000, true, 0},
	}

	for _, test := range tests {
		ver, err := NewNodeVersion(test.major, test.minor, test.patch)
		if err != nil {
			if test.err {
				t.Logf("major:%d minor:%d patch:%d error:%v check", test.major, test.minor, test.patch, err)
			} else {
				t.Fatalf("major:%d minor:%d patch:%d failed:%v", test.major, test.minor, test.patch, err)
			}
		} else {
			if test.err {
				t.Fatalf("major:%d minor:%d patch:%d should failed, but: %s", test.major, test.minor, test.patch, ver)
			} else {
				if ver != test.v {
					t.Fatalf("major:%d minor:%d patch:%d should be %d but %d", test.major, test.minor, test.patch, test.v, ver)
				} else {
					if ver.Major() == test.major && ver.Minor() == test.minor && ver.Patch() == test.patch {
						t.Logf("major:%d minor:%d patch:%d -> %s check", test.major, test.minor, test.patch, ver)
					} else {
						t.Fatalf("major:%d minor:%d patch:%d %s check versions failed", test.major, test.minor, test.patch, ver)
					}
				}
			}
		}
	}
}

func TestVersionString(t *testing.T) {
	tests := []struct {
		input string
		err   bool
		v     Version
	}{
		{"7.2.6", false, 7002006},
		{"7x.2.6", true, 0},
		{"7009.2.6", true, 0},
		{"7.5.0", false, 7005000},
	}

	for _, test := range tests {
		ver, err := NewVersion(test.input)
		if err != nil {
			if test.err {
				t.Logf("input:%s error:%v check", test.input, err)
			} else {
				t.Fatalf("input:%s failed: %v", test.input, err)
			}
		} else {
			if test.err {
				t.Fatalf("input:%s should error, but didn't, got: %s", test.input, ver)
			} else {
				if ver != test.v {
					t.Fatalf("input:%s should be %d but got %d", test.input, test.v, ver)
				} else {
					t.Logf("input:%s is %d check", test.input, ver)
				}
			}
		}
	}
}
