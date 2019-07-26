package provareadist

import (
	"home/appconf/filehandle"
	"home/appconf/urlparams"
	"strings"
)

var realfilename = "conflist.ini"

func Caclarea(prefilepath string, params urlparams.ParamsInfo) (bool, string) {
	areaExist := true
	filepath := ""
	prov, iscontainsProv, issetprov := caclProv(prefilepath, params)
	if iscontainsProv {
		filepath = prov + "/"
		city, iscontainsCity, issetcity := caclCity(prefilepath, filepath, params)
		if iscontainsCity {
			filepath = filepath + city + "/"
			dist, iscontainsDist, issetdist := caclDist(prefilepath, filepath, params)
			if iscontainsDist {
				filepath = filepath + dist + "/"
			} else if issetdist == true {
				areaExist = false
			}
		} else if issetcity == true {
			areaExist = false
		}
	} else if issetprov == true {
		areaExist = false
	}
	return areaExist, filepath
}

// 计算省
func caclProv(prefilepath string, params urlparams.ParamsInfo) (prov string, iscontainsProv bool, isset bool) {
	if len(params.Prov) > 0 && params.Prov[0] != "" {
		_, dirs, _ := filehandle.GetFilesAndDirs(prefilepath)
		for _, dirname := range dirs {
			iscontainsProv = strings.Contains(params.Prov[0], dirname)
			if iscontainsProv {
				prov = dirname
				break
			}
		}
		isset = true
	} else {
		isset = false
	}
	return prov, iscontainsProv, isset
}

// 计算市
func caclCity(prefilepath string, filepath string, params urlparams.ParamsInfo) (city string, iscontainsCity bool, isset bool) {
	if len(params.Area) > 0 && params.Area[0] != "" {
		_, dirs, _ := filehandle.GetFilesAndDirs(prefilepath + filepath)
		for _, dirname := range dirs {
			iscontainsCity = strings.Contains(params.Area[0], dirname)
			if iscontainsCity {
				city = dirname
				break
			}
		}
		isset = true
	} else {
		isset = false
	}
	return city, iscontainsCity, isset
}

// 计算区
func caclDist(prefilepath string, filepath string, params urlparams.ParamsInfo) (dist string, iscontainsDist bool, isset bool) {
	if len(params.Dist) > 0 && params.Dist[0] != "" {
		_, dirs, _ := filehandle.GetFilesAndDirs(prefilepath + filepath)
		for _, dirname := range dirs {
			iscontainsDist = strings.Contains(params.Dist[0], dirname)
			if iscontainsDist {
				dist = dirname
				break
			}
		}
		isset = true
	} else {
		isset = false
	}
	return dist, iscontainsDist, isset
}

func GetFilename(prefilepath string, path string, isGetConf bool) (bool, string) {
	var (
		zfile       = ""
		filename    = ""
		defaultfile = prefilepath + realfilename
		pathlist    = make(map[int]string)
		fileExist   = true
	)

	if path != "" {
		pathlist1 := strings.Split(path, "/")
		if len(pathlist1) > 0 {
			for k, v := range pathlist1 {
				pathlist[k] = v
			}
		}
	}

	if len(pathlist) > 0 {
		dist, okdist := pathlist[2]
		city, okcity := pathlist[1]
		prov, okprov := pathlist[0]
		if okdist == true && dist != "" {
			tmpzfile := prefilepath + path + realfilename
			isExist, _ := filehandle.PathExists(tmpzfile)
			if isExist == true {
				zfile = tmpzfile
			} else if isGetConf == true {
				fileExist = false
			}
		}
		if zfile == "" && okcity == true && city != "" {
			tmpzfile := prefilepath + prov + "/" + city + "/" + realfilename
			isExist, _ := filehandle.PathExists(tmpzfile)
			if isExist == true {
				zfile = tmpzfile
			} else if isGetConf == true && fileExist == true {
				fileExist = false
			}
		}
		if zfile == "" && okprov == true && prov != "" {
			tmpzfile := prefilepath + prov + "/" + realfilename
			isExist, _ := filehandle.PathExists(tmpzfile)
			if isExist == true {
				zfile = tmpzfile
			} else if isGetConf == true && fileExist == true {
				fileExist = false
			}
		}
	}

	if zfile == "" {
		filename = defaultfile
	} else {
		filename = zfile
	}
	return fileExist, filename
}
