# Metadata Storage
 
Each inode stores metadata about a file or directory, such as its size, permissions, ownership, timestamps (for creation, modification, and last access), and the disk block locations where the actual data is stored.

# Unique Identifiers

Every file or directory on a filesystem is associated with a unique inode number, which acts as its identifier within the filesystem. This is how the filesystem keeps track of files.

# Filesystem

So (again), inodes are data structures used by Unix-like filesystems to store essential metadata about files and directories, excluding their names or actual data content. Each inode is identified by a unique inode number, and it contains information such as ownership, permissions, and file size, as well as pointers to the disk blocks that store the file's data. The filesystem maintains a table of inodes to manage all the files and directories it contains, allowing for efficient file retrieval and management without regard to the file or directory's name.

 The concept of inodes is strictly related to the filesystem and is agnostic of higher-level system abstractions e.g. mount namespaces. There is a clear seperation of concers here as the inode structure is not concerned with how and where the filesystem is used or viewed.

 (The Mount (mnt) namespace is a feature of the Linux kernel that encapsulates the set of file system mount points visible to a set of processes. By segregating mount points into distinct namespaces, processes within one Mount namespace can have a unique perspective of the file system structure, independent of the mount points and file system views in another Mount namespace. This isolation enables different processes to interact with different file system layouts as if they were on separate systems, even though they are operating on the same physical machine.)
