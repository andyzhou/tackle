package base

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"html/template"

	"github.com/andyzhou/tinylib/util"
)

/*
 * tpl file face
 */

//face info
type TplFace struct {
	tplPath string
	sharedTplFiles []string
	tplFuncMap template.FuncMap
	util.Util
}

//construct
func NewTplFace() *TplFace {
	this := &TplFace{
		sharedTplFiles: []string{},
		tplFuncMap: map[string]any{},
	}
	return this
}

//get tpl content, STEP-3 [option]
func (f *TplFace) GetTplContent(
	tpl *template.Template,
	tplData map[string]interface{},
) (string, error) {
	//check
	if tpl == nil {
		return "", errors.New("invalid parameter")
	}

	//get tpl content
	tplBuff := bytes.NewBuffer(nil)
	err := tpl.Execute(tplBuff, tplData)
	if err != nil {
		return "", err
	}

	//must unescape html string for show
	originTpl := tplBuff.String()
	originTpl = html.UnescapeString(originTpl)
	return originTpl, nil
}

//parse tpl, STEP-2
func (f *TplFace) ParseTpl(
	mainTplFile string,
	tplPaths ...string,
) (*template.Template, error) {
	var (
		tplPath string
		err error
	)
	//check
	if mainTplFile == "" {
		return nil, errors.New("invalid parameter")
	}
	if tplPaths != nil && len(tplPaths) > 0 {
		tplPath = tplPaths[0]
	}

	//init new template obj
	tpl := template.New(mainTplFile)

	//register default tpl func
	//f.RegisterDefaultTplFunc(tpl)

	//format tpl full path
	mainTplFullPath := fmt.Sprintf("%v/%v", f.tplPath, mainTplFile)
	if tplPath != "" {
		mainTplFullPath = fmt.Sprintf("%v/%v", tplPath, mainTplFile)
	}

	//setup final tpl files
	finalTplFiles := make([]string, 0)
	finalTplFiles = append(finalTplFiles, f.sharedTplFiles...)
	finalTplFiles = append(finalTplFiles, mainTplFullPath)

	//begin parse tpl files
	tpl, err = tpl.ParseFiles(finalTplFiles...)
	return tpl, err
}

//set shared tpl files, STEP-1-1 [option]
func (f *TplFace) SetShareTplFiles(
	fileNames ...string) error {
	//check
	if fileNames == nil || len(fileNames) <= 0 {
		return nil
	}

	//reset and fill share tpl files
	f.sharedTplFiles = []string{}
	for _, fileName := range fileNames {
		tplFullPath := fmt.Sprintf("%v/%v", f.tplPath, fileName)
		f.sharedTplFiles = append(f.sharedTplFiles, tplFullPath)
	}
	return nil
}

//set tpl root path, STEP-1
func (f *TplFace) SetTplPath(
	path string) {
	f.tplPath = path
}

/////////////////
//private func
////////////////

//internal func maps
func (f *TplFace) TplFuncOfHtml(text string) template.HTML {
	return template.HTML(text)
}
