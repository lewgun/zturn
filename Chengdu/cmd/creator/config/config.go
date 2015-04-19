//Package config read the config from command line
package config


import (
    "os"
    "fmt"
    "strings"
    "path/filepath"
)


//config define the config parameters.
type Config struct {
    Loc string
    Pkg string
    Project string
}

//New new a config.
func New() *Config {
    return &Config{}
}


func (f *Config) String() string {
    return fmt.Sprintf("Location: %s Package: %s Project: %s",
    f.Loc,
    f.Pkg,
    f.Project)
}

func (f *Config) location() {
    f.Loc = ""

    for {
        fmt.Println("Please input the Location where you want put your project and press ENTER.")
        fmt.Scanln(&f.Loc)
        if f.Loc !="" {
            break
        }
    }

}

func (f *Config) pkgName() {
    f.Pkg = ""

    for {
        fmt.Println("Please input the Package Name and press ENTER.")
        fmt.Scanln(&f.Pkg)
        if f.Pkg !="" {
            break
        }
    }

}

func (f *Config) defaultApp() {
    idx := strings.LastIndex(f.Pkg, ".")
    if idx == -1 {
        f.Project = f.Pkg

    } else {
        f.Project = f.Pkg[idx+1:]
    }

}
func (f *Config) appName(){
    for {

        fmt.Println("Please input the Project Name and press ENTER. if default just press ENTER.")
        fmt.Scanln(&f.Project)

        if f.Project == "" {
            f.defaultApp()
        }

        path := filepath.Join(f.Loc, f.Project)
        if !f.isExist( path ) {
            break
        }
        fmt.Printf("The Project (%s) is already existed. Please choose another name.\n", f.Project)
    }

}

func (f *Config) isExist(path string ) bool {
    _, err := os.Stat(path)

    return err == nil || os.IsExist(err)

}


//Config do the actual operations.
func (f *Config) Config() error {
    f.location()
    f.pkgName()
    f.appName()
    return nil

}
