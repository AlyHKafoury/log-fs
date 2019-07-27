package filesystem

import (
	"context"
	"github.com/hanwen/go-fuse/fs"
	"github.com/hanwen/go-fuse/fuse"
	"syscall"
    "fmt"
    "path/filepath"
)

type Node struct {
	fs.Inode
    Name string
}

var _ = (fs.InodeEmbedder)((*Node)(nil))
var _ = (fs.NodeLookuper)((*Node)(nil))
var _ = (fs.NodeReaddirer)((*Node)(nil))
var _ = (fs.NodeWriter)((*Node)(nil))

func (n *Node) Lookup(ctx context.Context, name string, out *fuse.EntryOut) (*fs.Inode, syscall.Errno) {
	node := Node{Name: name}
    if filepath.Ext(name) == "" {
        return n.NewInode(ctx, &node, fs.StableAttr{Mode: syscall.S_IFDIR}), 0
    }
	return n.NewInode(ctx, &node, fs.StableAttr{Mode: syscall.S_IFREG}), 0
}

func (n *Node) Readdir(ctx context.Context) (fs.DirStream, syscall.Errno) {
	r := make([]fuse.DirEntry, 0)
	return fs.NewListDirStream(r), 0
}

func (n *Node) Open(ctx context.Context, flags uint32) (fh fs.FileHandle, fuseFlags uint32, errno syscall.Errno) {
    return nil, 755 ,0
}

func (n *Node) Write(ctx context.Context, f fs.FileHandle, data []byte, off int64) (written uint32, errno syscall.Errno) {
    fmt.Println(string(data))
    fmt.Println(n.Path(n.Root()))
    return uint32(len(data)), 0
}
func (n *Node) Setattr(ctx context.Context, f fs.FileHandle, in *fuse.SetAttrIn, out *fuse.AttrOut) syscall.Errno {
    return 0
}
