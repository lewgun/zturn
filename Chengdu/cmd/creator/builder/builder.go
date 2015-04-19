//package builder do the builder work.
package builder

import (
	"bytes"
	"io"
	"os"
	"fmt"
	// "bytes"
	"errors"

	"encoding/xml"
	"io/ioutil"
	"path/filepath"

	"github.com/lewgun/zturn/Chengdu/cmd/creator/config"
)

var (
	errIllegalParams = errors.New("illegal parameters")
)

const (
	//the templateDir project's base directory
	templateDir        = "./template"
	jniDir             = "jni"
	modulePlaceholder  = "MODULE_PLACEHOLDER"
	packagePlaceholder = "PACKAGE_PLACEHOLDER"
)

const (
	elementManifest    = "manifest"
	elementApplication = "application"
	attrPackage        = "package"
	attrLabel   = "label"
	elementActivity    = "activity"
	elementMetaData    = "meta-data"
	attrValue   = "value"
	namespace = "android"
)

//Builder build a skeleton project.
type Builder struct {
	conf config.Config
}

func New(c *config.Config) *Builder {

	b := &Builder{
		conf: *c,
	}

	return b

}

func mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (b *Builder) makeAllDirs(paths []string) error {
	if paths == nil {
		return errIllegalParams
	}

	var err error

	// errLog := &bytes.Buffer{}

	for _, path := range paths {
		path = filepath.Join(b.conf.Loc, b.conf.Project, path)
		if err = mkdir(path); err != nil {
			// fmt.Fprintf(errLog, "%v\n", err)
			return err
		}
	}

	/*
	   errStr := errLog.String()
	   if errStr == "" {
	       return nil
	   }

	   return fmt.Errorf("%s", errStr)
	*/

	return nil

}

func (b *Builder) copyAllFiles(paths []string) error {
	if paths == nil {
		return errIllegalParams
	}

	// var err error
	// errLog := &bytes.Buffer{}

	for _, path := range paths {
		dst := filepath.Join(b.conf.Loc, b.conf.Project, path)
		src := filepath.Join(templateDir, path)
		if err := b.copy(dst, src); err != nil {
			return err
			// fmt.Fprintf(errLog, "%v\n", err)

		}

	}
	/*
	   errStr := errLog.String()
	   if errStr == "" {
	       return nil
	   }

	   return fmt.Errorf("%s", errStr)
	*/

	return nil

}
func (b *Builder) copy(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer d.Close()

	// do the actual work
	_, err = io.Copy(d, s)
	return err

}

func replace(path string, old, new []byte, n int) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	data = bytes.Replace(data, old, new, n)
	return ioutil.WriteFile(path, data, 0644)

}

func (b *Builder) updateElement(e *xml.StartElement) *xml.StartElement{

	println(e.Name.Local)
	switch e.Name.Local {

	case elementManifest:
		for _, v := range e.Attr {
			println("M: ", v.Name.Local)
			if v.Name.Local == attrPackage {
				v.Value = b.conf.Pkg
				break
			}
		}

	case elementApplication:
		for _, v := range e.Attr {
			if v.Name.Local == attrLabel && v.Name.Space == namespace {
				v.Value = b.conf.Project
				break
			}
		}
	case elementActivity:
		for _, v := range e.Attr {
			if v.Name.Local == attrLabel && v.Name.Space == namespace {
				v.Value = b.conf.Project
				break
			}
		}
	case elementMetaData:
		for _, v := range e.Attr {
			if v.Name.Local == attrLabel  && v.Name.Space == namespace{
				v.Value = b.conf.Project
				break
			}
		}
	default:
		//placeholder
	}

    return e
}

func (b *Builder) AndroidManifestXML(path string) error {

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	d := xml.NewDecoder(f)

	buf := &bytes.Buffer{}
	e := xml.NewEncoder(buf)

	for {
		token, err := d.Token()
		if err != nil {
			break
		}

		switch token.(type) {
		case xml.CharData: //skip the escape text
			continue

		case xml.StartElement:
			t := token.(xml.StartElement)
			token = b.updateElement(&t)

		default:
			//placeholder

		} // end of switch

		if err = e.EncodeToken(token); err != nil {
			return err
		}

	} // end of for

	if err != nil && err != io.EOF {
		return err
	}

	e.Flush()
	println(buf.Bytes())

	return ioutil.WriteFile(path+".temp", buf.Bytes(), os.ModePerm)

}


func (b *Builder) AndroidMK(path string) error {
	return replace(path, []byte(modulePlaceholder), []byte(b.conf.Project), -1)
}

func (b *Builder) allBash(path string) error {
	return replace(path, []byte(packagePlaceholder), []byte(b.conf.Pkg), -1)
}

func (b *Builder) allBat(path string) error {
	return replace(path, []byte(packagePlaceholder), []byte(b.conf.Pkg), -1)
}

func (b *Builder) makeBash(path string) error {
	return replace(path, []byte(modulePlaceholder), []byte(b.conf.Pkg), -1)
}

func (b *Builder) makeBat(path string) error {
	return replace(path, []byte(modulePlaceholder), []byte(b.conf.Pkg), -1)
}

func (b *Builder) configAll() error {

	path := filepath.Join(b.conf.Loc, b.conf.Project)

	//Android.mk
	b.AndroidMK(filepath.Join(path, jniDir, "Android.mk"))

	//all.bash
	b.allBash(filepath.Join(path, "all.bash"))

	//all.bat
	b.allBat(filepath.Join(path, "all.bat"))

	//make.bash
	b.makeBash(filepath.Join(path, "make.bash"))

	//make.bat
	b.makeBat(filepath.Join(path, "make.bat"))

	//AndroidManifest.xml
	err := b.AndroidManifestXML(filepath.Join(path, "AndroidManifest2.xml"))
	if err != nil {
		fmt.Println(err)
	}
	return nil

}

//Build build a skeleton project.
func (b *Builder) Build() error {
	di := newDirInfo(templateDir)
	di.analysis()

	var (
		err error
	)

	if err = b.makeAllDirs(di.dirs); err != nil {
		return err
	}

	if err = b.copyAllFiles(di.files); err != nil {
		return err
	}

	return b.configAll()

}

func (b *Builder) Clean() error {
	return os.RemoveAll(filepath.Join(b.conf.Loc, b.conf.Project))
}
