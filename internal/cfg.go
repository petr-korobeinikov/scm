package internal

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func ReadCfg(scmUrl string) (Cfg, error) {
	scmWorkspaceDir, _ := LookupEnvOrDefault("SCM_WORKSPACE_DIR", "~/Workspace")
	scmExpanedWorkspaceDir, err := ExpandHomeDir(scmWorkspaceDir)
	if err != nil {
		return Cfg{}, err
	}

	scmWorkspaceDirDefaultPermStr, _ := LookupEnvOrDefault("SCM_WORKSPACE_DIR_DEFAULT_PERM", "0755")
	scmWorkspaceDirDefaultPerm, err := strconv.ParseInt(scmWorkspaceDirDefaultPermStr, 8, strconv.IntSize)
	scmWorkspaceDirDefaultPermFileMode := os.FileMode(scmWorkspaceDirDefaultPerm)

	scmPathFromUrl, err := ExtractLocalPathFromScmURL(scmUrl)
	if err != nil {
		return Cfg{}, err
	}
	scmWorkingCopyPath := filepath.Join(scmExpanedWorkspaceDir, scmPathFromUrl)

	var scmPostCloneCmd ScmPostCloneCmd
	if scmPostCloneCmdStr, found := LookupEnvOrDefault("SCM_POST_CLONE_CMD", ""); found {
		tmplData := ScmPostCloneCmdTmplData{ScmWorkingCopyPath: scmWorkingCopyPath}

		scmPostCloneCmdTmpl, err := parseScmPostCloneCmdTmpl(scmPostCloneCmdStr, tmplData)
		if err != nil {
			return Cfg{}, err
		}

		cmd, args := prepareScmPostCloneCmd(scmPostCloneCmdTmpl)
		scmPostCloneCmd.Cmd = cmd
		scmPostCloneCmd.Args = args
	}

	return Cfg{
		ScmWorkspaceDirDefaultPerm: scmWorkspaceDirDefaultPermFileMode,
		ScmWorkingCopyPath:         scmWorkingCopyPath,
		ScmPostCloneCmd:            scmPostCloneCmd,
	}, nil
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
	ScmPostCloneCmd            ScmPostCloneCmd
}

type ScmPostCloneCmd struct {
	Cmd  string
	Args []string
}

type ScmPostCloneCmdTmplData struct {
	ScmWorkingCopyPath string
}
