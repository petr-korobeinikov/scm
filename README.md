# scm 💪

## It's time 🕒 to clone ⤵️ interesting 🧐 repo faster ⏩ and clearer 🧹

`scm` is a tool that aims to keep your workspace to be strongly organized.

![Usage example](demo.svg)

Watch on [asciinema.org](https://asciinema.org/a/387451).

There is an example of how your workspace directory structure would look like:

```shell
> tree -L 2 ~/Workspace/
├── github.com
│   ├── VictoriaMetrics
│   ├── fluent
│   ├── github
│   ├── golang
│   ├── goreleaser
│   ├── hashicorp
│   ├── micro
│   ├── prometheus
│   ├── timescale
│   ├── topolvm
│   └── wrouesnel
├── hg.nginx.org
│   └── nginx
└── private-project-storage.tld
    └── private-project-team
```

Just type:

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
