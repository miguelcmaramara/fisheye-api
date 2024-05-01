package types

type User struct {
	id       string // TODO: change this to the type of the key
	name     string
	email    string
	password string
}

type Group struct {
	creator *User
	id      string // TODO: change this to the type of the key
	name    string
	// TODO: add options to each group such as length of time of hangout
}

type GroupMember struct {
	member *User
	id     string
}

type FileBatch struct {
	uploader *User
	group    *Group
	id       string
}

type File struct {
	batch    *FileBatch
	id       string // TODO: change this to the type of the key
	fileType FileType
}
