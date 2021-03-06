package back_end

import (
	"env/controllers"
	"env/models"
	"image/png"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

type PictureController struct {
	controllers.BaseController
}

// 查看（Get）功能
func (c *PictureController) Get() {
	///////////////////////////////////////////////////////
	// files := make([]string, 0)
	files := make([]models.Picture, 0)
	dirPath, _ := os.Getwd()
	dirPath = dirPath + "/static/picture"
	// 获得所有文件
	if dir, err := ioutil.ReadDir(dirPath); err == nil {
		PthSep := string(os.PathSeparator)
		suffix := strings.ToUpper(".png")
		for _, fi := range dir {
			if fi.IsDir() {
				continue
			}
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				// files = append(files, dirPath+PthSep+fi.Name())
				beego.Debug(fi)
				var file models.Picture
				file.Name = fi.Name()
				file.Path = "static/picture" + PthSep + fi.Name()
				file.Link = "http://" + beego.AppConfig.String("WEB_URL") + "/static/picture" + PthSep + fi.Name()
				pngfile, err := os.Open(file.Path)
				if err != nil {
					beego.Debug(err)
					continue
				} else {
					img, _ := png.DecodeConfig(pngfile)
					file.Height = img.Height
					file.Width = img.Width
				}
				files = append(files, file)
			}
		}
		// suffix_jpg := strings.ToUpper(".jpg")
		// for _, fi := range dir {
		// 	if fi.IsDir() {
		// 		continue
		// 	}
		// 	if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix_jpg) {
		// 		// files = append(files, dirPath+PthSep+fi.Name())
		// 		beego.Debug(fi)
		// 		var file models.Picture
		// 		file.Name = fi.Name()
		// 		file.Path = "static/picture" + PthSep + fi.Name()
		// 		file.Link = "http://" + beego.AppConfig.String("WEB_URL") + "/static/picture" + PthSep + fi.Name()
		// 		jpgfile, err := os.Open(file.Path)
		// 		if err != nil {
		// 			beego.Debug(err)
		// 			continue
		// 		} else {
		// 			img, _, _ := image.DecodeConfig(jpgfile)
		// 			file.Height = img.Height
		// 			file.Width = img.Width
		// 		}
		// 		files = append(files, file)
		// 	}
		// }
	} else {
		beego.Debug(err)
	}
	beego.Debug(files)
	c.Data["Picture"] = files
	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "picture"
}

// 增加（Add）功能
func (c *PictureController) Add() {
	// 增加上传图片功能
	filepath := "static/picture"
	_, _, err := c.GetFile("pic")
	name := c.GetString("name")
	filetype := c.GetString("filetype")
	beego.Debug(err)
	if err == nil {
		os.MkdirAll(filepath, 0777)
		if err := c.SaveToFile("pic", filepath+"/"+name+filetype); err != nil {
			beego.Debug(err)
		}
	}
	// c.TplName = "back_end/public.html"
	// c.Data["Tpl"] = "picture"
	c.Redirect("/picture", 302)
}

// 删除（Del）功能
func (c *PictureController) Del() {
	name := c.GetString("name")
	beego.Debug(name)
	err := os.Remove("static/picture/" + name)
	if err != nil {
		beego.Debug(err)
	}
	beego.Debug("已删除", name)
	c.Redirect("/picture", 302)
}

// 查找一篇文章
func (c *PictureController) Search() {
	files := make([]models.Picture, 0)
	dirPath, _ := os.Getwd()
	dirPath = dirPath + "/static/picture"
	content := c.GetString("content")
	// 获得所有文件
	if dir, err := ioutil.ReadDir(dirPath); err == nil {
		PthSep := string(os.PathSeparator)
		suffix := strings.ToUpper(".png")
		for _, fi := range dir {
			if fi.IsDir() {
				continue
			}
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				var file models.Picture
				file.Name = fi.Name()
				if findString(file.Name, content) {
					file.Path = "static/picture" + PthSep + fi.Name()
					file.Link = "http://" + beego.AppConfig.String("WEB_URL") + "/static/picture" + PthSep + fi.Name()
					pngfile, err := os.Open(file.Path)
					if err != nil {
						beego.Debug(err)
						continue
					} else {
						img, _ := png.DecodeConfig(pngfile)
						file.Height = img.Height
						file.Width = img.Width
					}
					files = append(files, file)
				}
			}
		}
	} else {
		beego.Debug(err)
	}
	beego.Debug(files)
	c.Data["Picture"] = files
	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "picture"
}

func findString(name string, input string) bool {
	beego.Debug(input)
	beego.Debug(name)
	re := regexp.MustCompile(input + "{1}")
	beego.Debug(re)
	return re.MatchString(name)
}
