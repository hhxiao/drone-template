package main

import (
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

	Job struct {
		Started int64
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
		Templates []string
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
		Job    Job
		Var    map[string]string
	}
)

func (p Plugin) Exec() error {
	for _, file := range p.Config.Templates {
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
		fmt.Printf("Processed %s\n", file)
	}
	return nil
}
