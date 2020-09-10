package util

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

//
func Inertia(v interface{}) template.HTML {
	retVal, _ := json.Marshal(v)
	return template.HTML(fmt.Sprintf("<div id='app' data-page='%s'></div>", string(retVal)))
}

//
func JsonEncode(v interface{}) template.JS {
	retVal, _ := json.Marshal(v)
	return template.JS(string(retVal))
}

//
func JsonEncodeRaw(v interface{}) string {
	retVal, _ := json.Marshal(v)
	return string(retVal)
}

//
func Mix(path string, manifestPath ...string) template.HTML {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Check if another manifest path was given
	mPath := GetEnvOrDefault("INERTIA_PUBLIC_PATH", "public") + "/mix-manifest.json"
	if len(manifestPath) > 0 {
		mPath = manifestPath[0]
	}
	manifestFile, err := os.Open(mPath)
	defer manifestFile.Close()
	if err != nil {
		return template.HTML(path)
	}

	manifestData, err := ioutil.ReadAll(manifestFile)
	if err != nil {
		return template.HTML(path)
	}

	var manifest map[string]string
	if err = json.Unmarshal(manifestData, &manifest); err != nil {
		return template.HTML(path)
	}

	return template.HTML(manifest[path])
}
