package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//var vflag = flag.Bool("v", false, "show verbose progress messages")

func main()  {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	//遍历文件树
	var wg sync.WaitGroup
	fileSize := make(chan int64)

	for _, root := range roots {
		wg.Add(1)
		go walkDir(&wg, root, fileSize)
	}
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(fileSize)
	}(&wg)

	var tick <-chan time.Time

	tick = time.Tick(500 * time.Millisecond)

	var nfiles, nbytes int64
	loop:
		for  {
			select {
			case size, ok := <-fileSize:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			case <- tick:
				printDiskUsage(nfiles, nbytes)
			}
		}

	printDiskUsage(nfiles, nbytes)
}

func dirents(dir string)[]os.FileInfo  {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(wg *sync.WaitGroup, dir string, fileSizes chan<- int64)  {
	defer wg.Done()
	for _, entry := range dirents(dir){
		if entry.IsDir(){
			defer wg.Done()
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(wg, subdir, fileSizes)
		} else{
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files   %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
//
//var sema = make(chan struct{}, 20)
//
//func dirents(dir string)[] os.FileInfo  {
//	sema <- struct{}
//	defer func() { <- sema }()
//}
