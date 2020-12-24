// Copyright 2020 by 闪电侠(http://chende.ren). All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
//
// 为配合自己的博客项目，自定义 hugo 相关功能
//
package aaasdx

import (
	"github.com/gohugoio/hugo/hugofs"
	"regexp"
	"strings"
	"time"
)

// 规范 front matter 字段
// 可以将自定义 matter 格式，对全目录所有文件做一次批量转换
// 这个作为一次性使用，用完了注释
func SdxChangeFrontMatter(mp map[string]interface{})  {
	// 改变 url 的值，以时间 /年/月/日- 为 前缀路径
	tmpUrl := mp["url"]
	if tmpUrl != nil {
		urlStr := tmpUrl.(string)
		tm, _ := time.Parse(time.RFC3339, mp["date"].(string))
		// println(urlStr)
		preUrl := tm.Format("/2006/01/02150405-")
		// println(newDate)
		mp["url"] = regexp.MustCompile(`.*/`).ReplaceAllString(urlStr, preUrl)
		// println(newUrl)
		// println("")
	}
	// is tags
	// if mp["tags"] != nil {
	//
	//}
	// is categories
	// if mp["categories"] != nil {
	//
	//}
	// 如果没有 lastmod 就加上
	if mp["lastmod"] == nil && mp["date"] != nil{
		mp["lastmod"] = mp["date"]
	}
}

//meta["filename"]: D:\chende_dev\core_discovery\go_demo\hugo\asdx\content\idea\why-open-blog.md
//meta["baseDir"]: D:\chende_dev\core_discovery\go_demo\hugo\asdx
//meta["sourceRoot"]: D:\chende_dev\core_discovery\go_demo\hugo\asdx\content
//path -> idea\why-open-blog.md
//name -> why-open-blog.md

//![image-20201130225219039](img/image-20201130225219039.jpg)
//![image-20201130225130217](img/image-20201130225130217.png)


func SdxGetRelativePathPrefix(meta hugofs.FileMeta) string {
	if !strings.HasSuffix(meta["name"].(string), ".md") {
		return ""
	}
	prefix := strings.Replace(meta["path"].(string), meta["name"].(string), "", 1)
	return strings.Replace(prefix, "\\", "/", -1)
	// regexp.MustCompile(`\\`).ReplaceAllString("", preUrl)
	// regexp.MustCompile(meta["path"].(string)).ReplaceAllString(urlStr, preUrl)
}

func SdxChangeImagePath(source []byte, prefix string) []byte {
	result := regexp.MustCompile(`]\(img/image-`).ReplaceAll(source, []byte(`](/` + prefix + `img/image-`))
	return result
}