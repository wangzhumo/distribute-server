package models

type BundlePlatformType int
type BundleFileExtension string

const (
	BundleFileExtensionAndroid BundleFileExtension = ".apk"
	BundleFileExtensionIOS     BundleFileExtension = ".ipa"
)

const (
	BundlePlatformTypeAndroid BundlePlatformType = 1 + iota
	BundlePlatformTypeIOS
)

func (platformType BundlePlatformType) Extention() BundleFileExtension {
	var ext BundleFileExtension
	if platformType == BundlePlatformTypeAndroid {
		ext = BundleFileExtensionAndroid
	} else if platformType == BundlePlatformTypeIOS {
		ext = BundleFileExtensionIOS
	}
	return ext
}

func (ext BundleFileExtension) IsValid() bool {
	if ext == BundleFileExtensionAndroid {
		return true
	} else if ext == BundleFileExtensionIOS {
		return true
	}
	return false
}

func (ext BundleFileExtension) IsValidImage() bool {
	if ext == ".png" {
		return true
	} else if ext == ".jpeg" {
		return true
	} else if ext == ".PNG" {
		return true
	} else if ext == ".JPEG" {
		return true
	}
	return false
}

func (ext BundleFileExtension) IsValidApk() bool {
	return ext == ".apk"
}

func (ext BundleFileExtension) PlatformType() BundlePlatformType {
	var platformType BundlePlatformType
	if ext == BundleFileExtensionAndroid {
		platformType = BundlePlatformTypeAndroid
	} else if ext == BundleFileExtensionIOS {
		platformType = BundlePlatformTypeIOS
	}
	return platformType
}

func (platformType BundlePlatformType) String() string {
	var out string
	if platformType == BundlePlatformTypeAndroid {
		out = "android"
	} else if platformType == BundlePlatformTypeIOS {
		out = "ios"
	}
	return out
}
