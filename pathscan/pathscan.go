package pathscan

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type PathScan struct {
	MainFilepath string       `json:"main_filepath"`
	Directories  []*Directory `json:"directories"`
}

func New(filepath string) PathScan {
	ps := PathScan{}
	ps.MainFilepath = filepath
	ps.AddMainFilepath()
	// to count the dirs
	ps.Fetch(filepath, ps.Directories[0])
	return ps
}

// Fetch the content of a directory
// the results are stored in  Directories or Files

func (ps *PathScan) Fetch(currentFilepath string, parent *Directory) ([]fs.DirEntry, error) {
	// checks if dir exists
	_, err := os.Stat(currentFilepath)
	if err != nil {
		return nil, err
	} else {
		content, err := os.ReadDir(currentFilepath)
		if err != nil {
			return nil, err
		}
		// fmt.Printf("%v", content)
		for _, dirItem := range content {
			// fmt.Printf("%v", dirItem.IsDir())
			if dirItem.IsDir() {
				if ps.IsMainFilepath(currentFilepath) {
					ps.Directories[0].HasSubDir = true
				}
				ps.AddDirectory(parent.ID, currentFilepath, dirItem)
				ps.Fetch(currentFilepath+"/"+dirItem.Name(), ps.Directories[len(ps.Directories)-1])
			} else {
				ps.AddFile(ps, currentFilepath, dirItem)
			}
		}
		return content, nil
	}
}

func (ps *PathScan) IsMainFilepath(dir string) bool {
	return ps.MainFilepath == dir
}

func (ps *PathScan) AddDirectory(parentID uint, parent string, fsDir fs.DirEntry) {
	dirName := fsDir.Name()

	// check if the current dir is the main
	hasSubdirs := ps.IsMainFilepath(dirName)
	// update the parent directory
	ps.Directories[parentID].HasSubDir = true
	path := parent + "/" + dirName

	newDir := &Directory{
		ID:         uint(len(ps.Directories)),
		Name:       dirName,
		Path:       path,
		HasParent:  true,
		ParentPath: parent,
		ParentID:   parentID,
		HasFiles:   false,
		HasSubDir:  hasSubdirs,
	}
	ps.Directories = append(ps.Directories, newDir)
}

func (dh *PathScan) AddMainFilepath() {
	newDir := &Directory{
		ID:        0,
		Name:      dh.RemoveRootFromName(dh.MainFilepath),
		Path:      dh.MainFilepath,
		HasParent: false,
		HasFiles:  false,
		HasSubDir: false,
	}
	dh.Directories = append(dh.Directories, newDir)
}

func (ps *PathScan) AddFile(dhParent *PathScan, parentPath string, fsFile fs.DirEntry) {
	fileName := fsFile.Name()
	fullpath := parentPath + "/" + fileName

	if parentPath != ps.MainFilepath {
		fullpath = parentPath + "/" + fileName
	}
	fi, err := os.Stat(fullpath)

	if err != nil {
		fmt.Println(err)
		return
	}
	newFile := &File{
		Name:     strings.TrimSuffix(fileName, filepath.Ext(fileName)),
		Path:     fileName,
		FullPath: fullpath,
		Size:     fi.Size(),
	}
	currentDirKey, _ := ps.GetCurrentDir(parentPath)
	if currentDirKey > -1 {
		ps.Directories[currentDirKey].HasFiles = true
		ps.Directories[currentDirKey].Files = append(ps.Directories[currentDirKey].Files, newFile)
	}
}

// returns the current dir key and value
func (ps *PathScan) GetCurrentDir(path string) (int, *Directory) {
	for key, val := range ps.Directories {
		if val.Path == path {
			return key, val
		}
	}
	return -1, &Directory{}
}

// remove the "./" from name
func (ps *PathScan) RemoveRootFromName(name string) string {
	return name[2:]
}

func (ps *PathScan) GetAsJSON() string {
	json, _ := json.Marshal(ps.Directories)
	return string(json)
}
