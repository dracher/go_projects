package helpers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	"regexp"
)

var (
	BuildNameDirList []string
)

func init() {

}

func walkFn(path string, info os.FileInfo, err error) error {

	name_pattern := regexp.MustCompile(`^\d\.\d$`)

	if info.IsDir() &&
		name_pattern.MatchString(info.Name()) &&
		err == nil {
		BuildNameDirList = append(BuildNameDirList, info.Name())
	}

	return err
}

func GetBuildNames() []string {
	r_path := getRhevhBuildRepoPath()

	BuildNameDirList = make([]string, 0)

	if FileDirExists(r_path) {
		filepath.Walk(r_path, walkFn)

		return BuildNameDirList

	} else {
		panic(fmt.Sprintf("%s is not exists", r_path))
	}
}

func getRhevhBuildRepoPath() string {
	cfg, err := beego.AppConfig.GetSection("user")
	if err != nil {
		panic(err)
	}
	return cfg["rhevh_buld_repo"]
}
