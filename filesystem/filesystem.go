package filesystem

import (
	"context"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"syscall"
)

type Node struct {
	fs.Inode
}

var _ = (fs.InodeEmbedder)((*Node)(nil))
var _ = (fs.NodeLookuper)((*Node)(nil))
var _ = (fs.NodeReaddirer)((*Node)(nil))

func (n *Node) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	node := Node{}
	return n.NewInode(ctx, &node, fs.StableAttr{}), 0
}

func (n *Node) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
	r := make([]fuse.DirEntry, 0, 1)
	d := fuse.DirEntry{
		Name: "Test file",
		Ino:  1000,
		Mode: fuse.S_IFREG,
	}
	r = append(r, d)
	return fs.NewListDirStream(r), 0
}
