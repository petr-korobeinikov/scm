package internal

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func Configure(scmUrl, scmOverridePostCloneCmd string) (cfg Cfg, err error) {
	scmExpandedWorkspaceDir, err := readScmWorkspaceDir()
	if err != nil {
		return
	}

	scmWorkspaceDirDefaultPermFileMode, err := readScmWorkspaceDirDefaultPermFileMode()
	if err != nil {
		return
	}

	scmWorkingCopyPath, err := readScmWorkingCopyPath(scmExpandedWorkspaceDir, scmUrl)
	if err != nil {
		return
	}

	scmPostCloneCmd, err := readPostCloneCmd(scmWorkingCopyPath)
	if err != nil {
		return
	}

	switch {
	case scmOverridePostCloneCmd == "-":
		scmPostCloneCmd = EmptyPostCloneCmd
	}

	return Cfg{
		ScmWorkspaceDirDefaultPerm: scmWorkspaceDirDefaultPermFileMode,
		ScmWorkingCopyPath:         scmWorkingCopyPath,
		ScmPostCloneCmd:            scmPostCloneCmd,
	}, nil
}

func readScmWorkspaceDir() (string, error) {
	scmWorkspaceDir, _ := LookupEnvOrDefault("SCM_WORKSPACE_DIR", "~/Workspace")

	return ExpandHomeDir(scmWorkspaceDir)
}

func readScmWorkspaceDirDefaultPermFileMode() (os.FileMode, error) {
	scmWorkspaceDirDefaultPermStr, _ := LookupEnvOrDefault("SCM_WORKSPACE_DIR_DEFAULT_PERM", "0755")
	scmWorkspaceDirDefaultPerm, err := strconv.ParseInt(scmWorkspaceDirDefaultPermStr, 8, strconv.IntSize)
	if err != nil {
		return 0, err
	}

	return os.FileMode(scmWorkspaceDirDefaultPerm), nil
}

func readScmWorkingCopyPath(scmExpandedWorkspaceDir, scmUrl string) (string, error) {
	scmPathFromUrl, err := ExtractLocalPathFromScmURL(scmUrl)
	if err != nil {
		return "", err
	}

	return filepath.Join(scmExpandedWorkspaceDir, scmPathFromUrl), nil
}

func readPostCloneCmd(scmWorkingCopyPath string) (*scmPostCloneCmd, error) {
	if scmPostCloneCmdStr, found := LookupEnvOrDefault("SCM_POST_CLONE_CMD", ""); found {
		tmplData := ScmPostCloneCmdTmplData{ScmWorkingCopyPath: scmWorkingCopyPath}

		scmPostCloneCmdTmpl, err := parseScmPostCloneCmdTmpl(scmPostCloneCmdStr, tmplData)
		if err != nil {
			return EmptyPostCloneCmd, err
		}

		cmd, args := prepareScmPostCloneCmd(scmPostCloneCmdTmpl)

		return NewPostCloneCmd(cmd, args), nil
	}

	return EmptyPostCloneCmd, nil
}

func parseScmPostCloneCmdTmpl(unparsedTmpl string, tmplData ScmPostCloneCmdTmplData) (unparsedCmd string, err error) {
	tmpl, err := template.New("postCloneCmdTmpl").Parse(unparsedTmpl)
	if err != nil {
		return "", err
	}

	var tplBuf bytes.Buffer
	if err = tmpl.Execute(&tplBuf, tmplData); err != nil {
		return "", err
	}

	return tplBuf.String(), nil
}

func prepareScmPostCloneCmd(unpreparedCmd string) (cmd string, args []string) {
	parts := strings.Split(unpreparedCmd, " ")

	return parts[0], parts[1:]
}

type Cfg struct {
	ScmWorkspaceDirDefaultPerm os.FileMode
	ScmWorkingCopyPath         string
	ScmPostCloneCmd            *scmPostCloneCmd
}

type ScmPostCloneCmdTmplData struct {
	ScmWorkingCopyPath string
}

func (c *scmPostCloneCmd) Command() string {
	return c.cmd
}

func (c *scmPostCloneCmd) Arguments() []string {
	return c.args
}

func (c *scmPostCloneCmd) IsEmpty() bool {
	return c.cmd == ""
}

func NewPostCloneCmd(cmd string, args []string) *scmPostCloneCmd {
	return &scmPostCloneCmd{
		cmd:  cmd,
		args: args,
	}
}

var EmptyPostCloneCmd = NewPostCloneCmd("", []string{})

type scmPostCloneCmd struct {
	cmd  string
	args []string
}
