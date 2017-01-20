package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var version string // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "template plugin"
	app.Usage = "template plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:   "templates",
			Usage:  "template files",
			EnvVar: "PLUGIN_TEMPLATES,PLUGIN_TEMPLATE",
		},

		//
		// repo args
		//

		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "repo.link",
			Usage:  "repository link",
			EnvVar: "DRONE_REPO_LINK",
		},

		//
		// build args
		//

		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.IntFlag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.IntFlag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.IntFlag{
			Name:   "build.finished",
			Usage:  "build finished",
			EnvVar: "DRONE_BUILD_FINISHED",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},

		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},

		//
		// commit args
		//

		cli.StringFlag{
			Name:   "build.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "build.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "build.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Value:  "",
			Usage:  "git commit tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.message",
			Usage:  "git commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "build.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
	}
	//for _, e := range os.Environ() {
	//	pair := strings.Split(e, "=")
	//	fmt.Println(pair[0] + ":" + os.Getenv(pair[0]))
	//}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
			Link:  c.String("repo.link"),
		},
		Build: Build{
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Created:  int64(c.Int("build.created")),
			Started:  int64(c.Int("build.started")),
			Finished: int64(c.Int("build.finished")),
			Link:     c.String("build.link"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Commit: Commit{
			Sha:     c.String("build.sha"),
			Ref:     c.String("build.ref"),
			Branch:  c.String("build.branch"),
			Tag:     c.String("build.tag"),
			Message: c.String("build.message"),
			Name:    c.String("build.author"),
			Email:   c.String("build.email"),
		},
		Config: Config{
			Templates: c.StringSlice("templates"),
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
