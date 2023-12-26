package pathscan

type Directory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	// parent date
	HasParent  bool   `json:"has_parent"`
	ParentPath string `json:"parent"`
	ParentID   uint   `json:"parent_id"`
	// childrens data
	HasSubDir bool    `json:"has_subdirs"`
	HasFiles  bool    `json:"has_files"`
	Files     []*File `json:"files"`
}
