package parser

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"image"
	"io/ioutil"
	"os"

	"com.wangzhumo.distribute/models/model"
	"github.com/shogo82148/androidbinary"
	"github.com/shogo82148/androidbinary/apk"
)

func NewAppParser(name string) (*model.ParserAppInfo, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(file, stat.Size())
	if err != nil {
		return nil, err
	}

	var xmlFile *zip.File
	for _, f := range reader.File {
		switch {
		case f.Name == "AndroidManifest.xml":
			xmlFile = f
		}
	}

	info, err := parseApkFile(xmlFile)
	if err != nil {
		return nil, err
	}
	icon, label, err := parseApkIconAndLabel(name)
	info.Name = label
	info.Icon = icon
	info.Size = stat.Size()
	return info, err
}

func parseAndroidManifest(xmlFile *zip.File) (*model.AndroidManifest, error) {
	rc, err := xmlFile.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	xmlContent, err := androidbinary.NewXMLFile(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	manifest := new(model.AndroidManifest)
	decoder := xml.NewDecoder(xmlContent.Reader())
	if err := decoder.Decode(manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}

func parseApkFile(xmlFile *zip.File) (*model.ParserAppInfo, error) {
	if xmlFile == nil {
		return nil, errors.New("AndroidManifest.xml is not found")
	}

	manifest, err := parseAndroidManifest(xmlFile)
	if err != nil {
		return nil, err
	}

	info := new(model.ParserAppInfo)
	info.BundleId = manifest.Package
	info.Version = manifest.VersionName
	info.Build = manifest.VersionCode

	return info, nil
}

func parseApkIconAndLabel(name string) (image.Image, string, error) {
	pkg, err := apk.OpenFile(name)
	if err != nil {
		return nil, "", err
	}
	defer pkg.Close()

	icon, _ := pkg.Icon(&androidbinary.ResTableConfig{
		Density: 720,
	})
	if icon == nil {
		return nil, "", errors.New("Icon is not found")
	}

	label, _ := pkg.Label(nil)

	return icon, label, nil
}
