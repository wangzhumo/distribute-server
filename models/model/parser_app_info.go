package model

import "image"

type ParserAppInfo struct {
	Name     string
	BundleId string
	Version  string
	Build    string
	Icon     image.Image
	Size     int64
}

type AndroidManifest struct {
	Package     string `xml:"package,attr"`
	VersionName string `xml:"versionName,attr"`
	VersionCode string `xml:"versionCode,attr"`
}
