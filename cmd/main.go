package main

import (
	"github.com/AlyHKafoury/log-fs/filesystem"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
)

func main() {
	server, err := fs.Mount("test", &filesystem.Node{}, &fs.Options{MountOptions: fuse.MountOptions{Debug: true}})
	if err != nil {
		log.Panic(err)
	}
	server.Wait()
}
