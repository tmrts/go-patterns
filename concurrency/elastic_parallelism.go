package elastic_parallelism

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

var wg sync.WaitGroup

// walkFiles starts a goroutine to walk the directory tree at root and send the
// path of each regular file on the string channel.  It sends the result of the
// walk on the error channel.  If done is closed, walkFiles abandons its work.
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	errc := make(chan error, 1)
	paths := make(chan string)
	go func() { // HL
		// Close the paths channel after Walk returns.
		defer close(paths) // HL
		// No select needed for this send, since errc is buffered.
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error { // HL
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path: // HL
			case <-done: // HL
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

// A result is the product of reading and summing a file using MD5.
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.  In that case,
// MD5All does not wait for inflight read operations to complete.
func MD5All(root string) (map[string][md5.Size]byte, error) {
	// MD5All closes the done channel when it returns; it may do so before
	// receiving all the values from c and errc.
	done := make(chan struct{})
	adjust := make(chan struct{}, 1024)
	defer close(done)

	paths, errc := walkFiles(done, root)

	// Start a fixed number of goroutines to read and digest files.
	c := make(chan result, 1024*10) // HLc
	expand(2, adjust, done, paths, c)
	expand(2, adjust, done, paths, c)
	shink(4, adjust)
	expand(20, adjust, done, paths, c)

	go func() {
		wg.Wait()
		close(c) // HLc
	}()
	// End of pipeline. OMIT

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	// Check whether the Walk failed.
	if err := <-errc; err != nil { // HLerrc
		return nil, err
	}
	return m, nil
}

func shink(num int64, adjust chan struct{}) {
	var i int64
	for i = 0; i < num; i++ {
		adjust <- struct{}{}
	}
}

// digester reads path names from paths and sends digests of the corresponding
// files on c until either paths or adjust is closed or done is closed.
func expand(num int64, adjust <-chan struct{}, done <-chan struct{}, paths <-chan string, c chan<- result) {
	var i int64
	for i = 0; i < num; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			for {
				select {
				case path := <-paths:
					if path == "" {
						//close(paths) may happens very early
						return
					}
					data, err := ioutil.ReadFile(path)
					c <- result{path, md5.Sum(data), err}
				case <-done:
					return
				case <-adjust:
					return
				}
			}
		}()
	}
}

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
