package pathscan

type File struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	FullPath string `json:"fullpath"`
	Size     int64  `json:"size"`
}
