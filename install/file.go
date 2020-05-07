package install

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func PathExists(path, fileName string) bool {
	filePath := path + fileName
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Download(FileName string, FilePath string) {
	var url = FileUrl + FileName
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	f, err := os.Create(FilePath + FileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, res.Body)
}

func FileTar(Filename, FilePath, FileNewName string) {
	shellcmd := "tar xf " + Filename +".tar.gz -C " + FilePath
	cmd := exec.Command("/bin/bash", "-c", shellcmd)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err, "")
	}

	shellcmdmv := "mv " + Filename + " /usr/local/" + FileNewName
	cmdmv := exec.Command("/bin/bash", "-c", shellcmdmv)
	_, err = cmdmv.Output()
	if err != nil {
		fmt.Println(err)
	}

}