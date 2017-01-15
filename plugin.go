package main

import (
	"errors"
	"strings"
	"fmt"
	"io/ioutil"
)

type (
	Repo struct {
		Owner string
		Name  string
		Link  string
	}

	Build struct {
		Number   int
		Event    string
		Status   string
		Created  int64
		Started  int64
		Finished int64
		Link     string
	}

	Commit struct {
		Sha     string
		Ref     string
		Link    string
		Branch  string
		Tag     string
		Message string
		Name    string
		Email   string
	}

	Config struct {
		Templates string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
	}
)

func (p Plugin) Exec() error {
	if strings.TrimSpace(p.Config.Templates) == "" {
		return errors.New("No templates specified")
	}
	files := strings.Split(strings.TrimSpace(p.Config.Templates), ",")
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		content, err := Render(string(data), p)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(file, []byte(content), 0644)
		if err != nil {
			return err
		}
		fmt.Printf("processed %s\n", file)
	}
	return nil
}
