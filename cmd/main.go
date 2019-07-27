package main

import (
	"github.com/AlyHKafoury/log-fs/filesystem"
	"github.com/AlyHKafoury/log-fs/wsclient"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"log"
    "os/user"
    "strconv"
)

func main() {
    wsclient.Connect()
    var uid, gid string
    user, err := user.Current()
    if err != nil {
        log.Println("Can't get current user default root")
        uid = "0"
        gid = "0"
    }else {
        uid = user.Uid
        gid = user.Gid
    }
    paramUID, _ := strconv.Atoi(uid)
    paramGID, _ := strconv.Atoi(gid)
	server, err := fs.Mount("test", &filesystem.Node{Name: "test"}, &fs.Options{MountOptions: fuse.MountOptions{Debug: true}, UID: uint32(paramUID), GID: uint32(paramGID)})
	if err != nil {
		log.Panic(err)
	}
	server.Wait()
}
