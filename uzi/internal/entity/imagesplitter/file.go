package imagesplitter

type FileMeta struct {
	ContentType string
}

type File struct {
	FileMeta FileMeta
	FileBin  []byte
}
