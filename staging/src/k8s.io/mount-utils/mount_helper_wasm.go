//go:build js

package mount

import (
	"fmt"
	"strings"
	"sync"
)

func IsCorruptedMnt(err error) bool {
	return false
}

// MountInfo represents a single line in /proc/<pid>/mountinfo.
type MountInfo struct {
	// Unique ID for the mount (maybe reused after umount).
	ID int
	// The ID of the parent mount (or of self for the root of this mount namespace's mount tree).
	ParentID int
	// Major indicates one half of the device ID which identifies the device class
	// (parsed from `st_dev` for files on this filesystem).
	Major int
	// Minor indicates one half of the device ID which identifies a specific
	// instance of device (parsed from `st_dev` for files on this filesystem).
	Minor int
	// The pathname of the directory in the filesystem which forms the root of this mount.
	Root string
	// Mount source, filesystem-specific information. e.g. device, tmpfs name.
	Source string
	// Mount point, the pathname of the mount point.
	MountPoint string
	// Optional fieds, zero or more fields of the form "tag[:value]".
	OptionalFields []string
	// The filesystem type in the form "type[.subtype]".
	FsType string
	// Per-mount options.
	MountOptions []string
	// Per-superblock options.
	SuperOptions []string
}

// ParseMountInfo parses /proc/xxx/mountinfo.
func ParseMountInfo(filename string) ([]MountInfo, error) {
	return nil, fmt.Errorf("not implemented")
}

// splitMountOptions parses comma-separated list of mount options into an array.
// It respects double quotes - commas in them are not considered as the option separator.
func splitMountOptions(s string) []string {
	return []string{}
}

// isMountPointMatch returns true if the path in mp is the same as dir.
// Handles case where mountpoint dir has been renamed due to stale NFS mount.
func isMountPointMatch(mp MountPoint, dir string) bool {
	return strings.TrimSuffix(mp.Path, "\\040(deleted)") == dir
}

// PathExists returns true if the specified path exists.
// TODO: clean this up to use pkg/util/file/FileExists
func PathExists(path string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

// These variables are used solely by kernelHasMountinfoBug.
var (
	hasMountinfoBug       bool
	checkMountinfoBugOnce sync.Once
)

// kernelHasMountinfoBug checks if the kernel bug that can lead to incomplete
// mountinfo being read is fixed. It does so by checking the kernel version.
//
// The bug was fixed by the kernel commit 9f6c61f96f2d97 (since Linux 5.8).
// Alas, there is no better way to check if the bug is fixed other than to
// rely on the kernel version returned by uname.
func kernelHasMountinfoBug() bool {
	return false
}

func readMountInfo(path string) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
