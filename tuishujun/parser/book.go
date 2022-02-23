package parser

import (
	"go-crawler/engine"
	"go-crawler/model"
	"regexp"
)

var titleRe = regexp.MustCompile(
	`<h1 class="text-color-high text-justify mb-1 text-xl">\s*(\S+)\s*</h1>`)
var complexRe = regexp.MustCompile(
	`<a href=".*" title=".*" class="text-color-medium">\s*(.*)・(.*)・(.*)\s*</a>`)
var updateTimeRe = regexp.MustCompile(
	`<div class="my-1">\s*更新时间:(.*)\s*</div>`)
var chapterCountRe = regexp.MustCompile(
	`<div class="my-1">\s*章节数:(.*)\s*</div>`)
var sourceRe = regexp.MustCompile(
	`<div class="mt-1">\s*来源:([^</div>]*)\s*</div>`)
var imgUrlRe = regexp.MustCompile(
	`<img src=".*" loading="lazy" data-src="(.*)" alt=".*"`)
var gradeRe = regexp.MustCompile(
	`<div class="text-4xl leading-snug" style="color: .*">(.+)</div>`)

func ParseBook(contents []byte) engine.ParseResult {
	book := model.Book{}
	book.Title = extractString(contents, titleRe)
	book.Source = extractString(contents, sourceRe)
	book.UpdateTime = extractString(contents, updateTimeRe)
	book.ChapterCount = extractString(contents, chapterCountRe)
	book.CoverImgUrl = extractString(contents, imgUrlRe)
	book.Grade = extractString(contents, gradeRe)

	match := complexRe.FindSubmatch(contents)
	if len(match) >= 4 {
		book.Author = string(match[1])
		book.WordCount = string(match[2])
		book.State = string(match[3])
	}

	return engine.ParseResult{
		Items: []interface{}{book},
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
