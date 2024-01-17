package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"regexp"

	"github.com/gomarkdown/markdown"
)

const (
	mdDir        = "posts"
	templateName = "post.html"
)

var (
	titleRe      = regexp.MustCompile(`title : "([A-Za-z0-9 ]+)"`)
	dateRe       = regexp.MustCompile(`date: (\d{4}-\d{2}-\d{2}T[\d:\+]+)`)
	tagsRe       = regexp.MustCompile(`tags: \[([\w", ]+)\]`)
	categoriesRe = regexp.MustCompile(`tags: \[([\w", ]+)\]`)
)

func main() {
	md, err := getMarkdown(mdDir)
	if err != nil {
		slog.Error("failed to get markdown content", slog.Any("err", err))
		os.Exit(1)
	}

	template, err := getTemplate(mdDir + templateName)
	if err != nil {
		slog.Error("failed to get tempalte", slog.Any("err", err))
		os.Exit(1)
	}

	for title, post := range md {
		rendered, err := templatePost(template, title, post)
		if err != nil {
			slog.Error("failed to render md", slog.Any("err", err))
			os.Exit(1)
		}

		fd, err := os.Create(fmt.Sprintf("%s-rendered.html", title))
		if err != nil {
			slog.Error("failed to create output file", slog.Any("err", err))
			os.Exit(1)
		}

		written, err := fd.Write(rendered)
		if err != nil {
			slog.Error("failed to write to output file", slog.Any("err", err))
		}

		slog.Info("finised writing render", slog.Any("title", title), slog.Any("bytes_written", written))
	}
}

type Post struct {
	Body  string
	Title string
}

func templatePost(t []byte, title string, post []byte) ([]byte, error) {
	var p Post
	p.Body = string(markdown.ToHTML(post, nil, nil))
	p.Title = title

	temp, err := template.New("post").Parse(string(t))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err := temp.Execute(&b, p); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func getTemplate(templatePath string) ([]byte, error) {
	t, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func getMarkdown(basePath string) (map[string][]byte, error) {
	files, err := os.ReadDir(basePath)
	if err != nil {
		slog.Error("failed to read md dir", slog.Any("err", err))
		return nil, err
	}

	out := make(map[string][]byte)
	for _, file := range files {
		slog.Info("f", slog.Any("filename", file.Name()))
		fileContent, err := os.ReadFile(file.Name())
		if err != nil {
			return nil, err
		}

		out[file.Name()] = fileContent
	}

	return out, nil
}
