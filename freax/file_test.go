package freax_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/g0rbe/gmod/freax"
)

func TestOpenFile(t *testing.T) {

	f, err := freax.OpenFile("file.go", freax.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("Error: %s\n", err)
	}
	defer f.Close()

}

func TestFileRead(t *testing.T) {

	f, err := freax.OpenFile("file.go", freax.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("Error open: %s\n", err)
	}
	defer f.Close()

	buf := make([]byte, 7)

	n, err := f.Read(buf)
	if err != nil {
		t.Fatalf("Error read: %s\n", err)
	}
	if n != 7 {
		t.Fatalf("Error read: bytes read: %d, should be 7\n", n)
	}

	t.Logf("%s\n", buf)

}

func TestFileIoReadAll(t *testing.T) {

	f, err := freax.OpenFile("file.go", freax.O_RDONLY, 0)
	if err != nil {
		t.Fatalf("Error open: %s\n", err)
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		t.Fatalf("Error read: %s\n", err)
	}

	t.Logf("Bytes read: %d\n", len(buf))

}

func TestFileWrite(t *testing.T) {

	defer os.RemoveAll("test.txt")

	f, err := freax.OpenFile("test.txt", freax.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Fatalf("Error open: %s\n", err)
	}
	defer f.Close()

	buf := make([]byte, 7)

	n, err := f.Write(buf)
	if err != nil {
		t.Fatalf("Error write: %s\n", err)
	}
	if n != 7 {
		t.Fatalf("Error write: bytes write: %d, should be 7\n", n)
	}

}

func BenchmarkFileReadWrite(b *testing.B) {

	buf1 := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	buf2 := []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}

	b.Cleanup(func() { os.RemoveAll("test.txt") })

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f, err := freax.OpenFile("test.txt", freax.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			b.Fatalf("Open: %s\n", err)
		}

		n, err := f.Write(buf1)
		if err != nil {
			b.Fatalf("Write: %s\n", err)
		}

		if n != len(buf1) {
			b.Fatalf("Write: bytes write: %d, should be %d\n", n, len(buf1))
		}

		_, err = f.Seek(0, os.SEEK_SET)
		if err != nil {
			b.Fatalf("Seek: %s\n", err)
		}

		n, err = f.Read(buf2)
		if err != nil {
			b.Fatalf("Read: %s\n", err)
		}
		if n != len(buf1) {
			b.Fatalf("Read: bytes read: %d, should be %d\n", n, len(buf1))
		}

		if !bytes.Equal(buf1, buf2) {
			b.Fatalf("Buf differs: %v / %v\n", buf1, buf2)
		}

		f.Close()
	}
}

func BenchmarkOsFileReadWrite(b *testing.B) {

	buf1 := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	buf2 := []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}

	b.Cleanup(func() { os.RemoveAll("test.txt") })

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f, err := os.OpenFile("test.txt", freax.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			b.Fatalf("Open: %s\n", err)
		}

		n, err := f.Write(buf1)
		if err != nil {
			b.Fatalf("Write: %s\n", err)
		}

		if n != len(buf1) {
			b.Fatalf("Write: bytes write: %d, should be %d\n", n, len(buf1))
		}

		_, err = f.Seek(0, os.SEEK_SET)
		if err != nil {
			b.Fatalf("Seek: %s\n", err)
		}

		n, err = f.Read(buf2)
		if err != nil {
			b.Fatalf("Read: %s\n", err)
		}
		if n != len(buf1) {
			b.Fatalf("Read: bytes read: %d, should be %d\n", n, len(buf1))
		}

		if !bytes.Equal(buf1, buf2) {
			b.Fatalf("Buf differs: %v / %v\n", buf1, buf2)
		}

		f.Close()
	}
}

// func BenchmarkOpenFile(b *testing.B) {

// 	for i := 0; i < b.N; i++ {
// 		f, _ := freax.OpenFile("file.go", freax.O_RDONLY, 0)
// 		f.Close()
// 	}
// }

// func BenchmarkOsOpenFile(b *testing.B) {

// 	for i := 0; i < b.N; i++ {
// 		f, _ := os.OpenFile("file.go", os.O_RDONLY, 0)
// 		f.Close()
// 	}
// }
