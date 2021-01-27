# scm 💪

## It's time 🕒 to clone ⤵️ interesting 🧐 repo faster ⏩ and clearer 🧹

`scm` is a tool that aims to keep your workspace to be strongly organized.

[![Usage example](demo.svg)](https://asciinema.org/a/387451)

Hi, folks!

I wrote a simple tool to obtain a working copy of any git repo.

My main purpose is to keep my workspace directory as clean as possible and strongly organized.
Usually, your projects folder looks like this:

```shell
project-1
project-2
project-3
...
```

But I prefer a much more strong outline, like this:

```shell
> tree -L 2 ~/Workspace/
├── github.com
│   ├── VictoriaMetrics
│   ├── fluent
│   ├── github
│   ├── golang
...
├── hg.nginx.org
│   └── nginx
└── private-project-storage.tld
    └── private-project-team
```

Copy interesting repo URL and paste it into your terminal:

```shell
scm https://github.com/pkorobeinikov/scm
```

It will clone `https://github.com/pkorobeinikov/scm` into `~/Workspace/github.com/pkorobeinikov/scm`.

It's also possible to clone `hg`-repo. So command:

```shell
scm hg http://hg.nginx.org/nginx
```

will clone `scm hg http://hg.nginx.org/nginx/` into `~/Workspace/hg.nginx.org/nginx`.

## Configuration

Put this into your `.rc`-file:

```shell
export SCM_WORKSPACE_DIR="~/Projects"         # defaults to ~/Workspace
export SCM_WORKSPACE_DIR_DEFAULT_PERM="0755"  # defaults to 0755

export SCM_POST_CLONE_CMD="idea {{.ScmWorkingCopyPath}}"  # defaults to ""
```

## Building from source

```shell
go build -o ~/Bin/scm main.go
```

## Running tests

```shell
go test -cover -v ./internal
```

---

Thank you for your attention! 🤝

Any feedback will be highly appreciated. 😊
