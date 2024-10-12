package entity

type FileMeta struct {
	Path        string
	ContentType string
}

type File struct {
	FileMeta *FileMeta
	FileBin  []byte
}
