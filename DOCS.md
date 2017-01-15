---
date: 2016-01-01T00:00:00+00:00
title: Template
author: hhxiao
tags: [ coding ]
repo: github.com/hhxiao/drone-template
image: hhxiao/drone-template
---

The Template plugin injects building environment variables into template files. The below pipeline configuration demonstrates simple usage:

```
pipeline:
  template:
    image: hhxiao/drone-template
    templates: src/main/run.java,src/main/version.java
```

# Parameter Reference

**templates**
: Template files to inject building environment variables

# Template Reference

**repo.owner**
: repository owner

**repo.name**
: repository name

**repo.link**
: repository link

**build.status**
: build status type enumeration, either `success` or `failure`

**build.event**
: build event type enumeration, one of `push`, `pull_request`, `tag`, `deployment`

**build.number**
: build number

**build.commit**
: git sha for current commit

**build.message**
: git message for current commit

**build.branch**
: git branch for current commit

**build.tag**
: git tag for current commit

**build.ref**
: git ref for current commit

**build.author**
: git author for current commit

**build.email**
: git author email for current commit

**build.link**
: link the the build results in drone

**build.created**
: unix timestamp for build creation

**build.started**
: unix timestamp for build started

# Template Function Reference

**uppercasefirst**
: converts the first letter of a string to uppercase

**uppercase**
: converts a string to uppercase

**lowercase**
: converts a string to lowercase. Example `{{lowercase build.author}}`

**datetime**
: converts a unix timestamp to a date time string. Example `{{datetime build.started}}`

**success**
: returns true if the build is successful

**failure**
: returns true if the build is failed

**truncate**
: returns a truncated string to n characters. Example `{{truncate build.sha 8}}`

**urlencode**
: returns a url encoded string

**since**
: returns a duration string between now and the given timestamp. Example `{{since build.started}}`
