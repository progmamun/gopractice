<P>A typical file is just a pointer to a place on the hard disk called an inode. A hard link creates a new pointer to the same place. A file will only be deleted from disk after all links are removed. Hard links only work on the same file system. A hard link is what you might consider a 'normal' link.

A symbolic link, or soft link, is a little different, it does not point directly to a place on the disk. Symlinks only reference other files by name. They can point to files on different filesystems. Not all systems support symlinks.</p>
