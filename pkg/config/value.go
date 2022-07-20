package config

import (
	"bytes"
	"os"
	"text/template"
)

var glbEnvs map[string]string

func GetRenderedConfFromFile(path string) (out []byte,err error)  {
	var b []byte
	b,err = os.ReadFile(path)
	if err != nil {
		return
	}
	out,err = RenderContent(b)
	return

}


func RenderContent(in []byte) (out []byte,err error){
	tmpl,errRet := template.New("frp").Parse(string(in))
	if errRet != nil{
		return
	}
	buffer := bytes.NewBufferString("")
	v := GetValues()
	err = tmpl.Execute(buffer,v)
	if err != nil{
		return
	}
	out = buffer.Bytes()
	return
}

type Values struct {
	Envs map[string]string //Environment vars
}


func GetValues() *Values{
	return &Values{
		Envs: glbEnvs,
	}
}
