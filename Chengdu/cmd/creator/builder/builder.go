//package builder do the builder work.
package builder

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	attrLabel          = "label"
	elementActivity    = "activity"
	elementMetaData    = "meta-data"
	attrValue          = "value"
	namespace          = "android"
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

func (b *Builder) updateElement(t xml.Token) {

	e := t.(xml.StartElement)
	switch e.Name.Local {

	case elementManifest:
		for i := range e.Attr {
			if e.Attr[i].Name.Local == attrPackage {
				e.Attr[i].Value = b.conf.Pkg
				break
			}

		}

	case elementApplication:
		for i := range e.Attr {
			if e.Attr[i].Name.Local == attrLabel {
				e.Attr[i].Value = b.conf.Project
				break
			}

		}
	case elementActivity:
		for i := range e.Attr {
			if e.Attr[i].Name.Local == attrLabel {
				e.Attr[i].Value = b.conf.Project
				break
			}

		}
	case elementMetaData:
		for i := range e.Attr {
			if e.Attr[i].Name.Local == attrValue {
				e.Attr[i].Value = b.conf.Project
				break
			}

		}
	default:
		//placeholder
	}

	//t = e

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
	e.Indent(" ", "  ")

	for {
		token, err := d.Token()
		if err != nil {
			break
		}

		switch token.(type) {
		case xml.CharData: //skip the escape text
			//t := token.(xml.CharData)
			//fmt.Println(string(t))
			continue

		case xml.StartElement:
			//	t := token.(xml.StartElement)
			//	fmt.Println("START: ", t.Name.Local )
			//	token = *(b.updateElement(&t))
			b.updateElement(token)
		case xml.EndElement:

			//	t := token.(xml.EndElement)
			//	fmt.Println("END: ", t.Name.Local )

		default:
			//placeholder

		} // end of switch

		//fmt.Println("T: ", token )
		if err = e.EncodeToken(token); err != nil {
			fmt.Println("ERR: ", err, "TOKEN: ", token)
			return err
		}

	} // end of for

	if err != nil && err != io.EOF {
		return err
	}

	e.Flush()

	return ioutil.WriteFile(path, buf.Bytes(), os.ModePerm)

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
	//b.AndroidMK(filepath.Join(path, jniDir, "Android.mk"))

	//all.bash
	//b.allBash(filepath.Join(path, "all.bash"))

	//all.bat
	//b.allBat(filepath.Join(path, "all.bat"))

	//make.bash
	//b.makeBash(filepath.Join(path, "make.bash"))

	//make.bat
	//b.makeBat(filepath.Join(path, "make.bat"))

	//AndroidManifest.xml
	err := b.AndroidManifestXML(filepath.Join(path, "AndroidManifest.xml"))
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
