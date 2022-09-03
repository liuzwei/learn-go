package hash

import "testing"

func TestMd5Encrypt(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{
			"hello",
			"5d41402abc4b2a76b9719d911017c592",
		},
	}

	for _, test := range tt {
		result := Md5Encrypt(test.in)
		if result != test.out {
			t.Errorf("Md5Encrypt error, want=%v, got=%v", test.out, result)
		}
	}
}

func TestSha256(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{
			"hello",
			"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}

	for _, test := range tt {
		result := Sha256(test.in)
		if result != test.out {
			t.Errorf("Md5Encrypt error, want=%v, got=%v", test.out, result)
		}
	}
}

func TestSha1(t *testing.T) {
	tt := []struct {
		in  string
		out string
	}{
		{
			"hello",
			"aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
	}

	for _, test := range tt {
		result := Sha1(test.in)
		if result != test.out {
			t.Errorf("Md5Encrypt error, want=%v, got=%v", test.out, result)
		}
	}
}
