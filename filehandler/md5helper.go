package filehandler

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const filechunk = 8192

func GetFileHash(r *bufio.Reader) string {
	h := md5.New()

	if _, err := io.Copy(h, r); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetHash(b *[]byte) string {
	return fmt.Sprintf("%x", md5.Sum(*b))
}

func GetBigFileHash(f *os.File) string {
	info, _ := f.Stat()
	fileSize := info.Size()

	blocks := uint64(math.Ceil(float64(fileSize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(fileSize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		f.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}
