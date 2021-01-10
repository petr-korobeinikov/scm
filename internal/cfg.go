package internal

import (
	"os"
	"path/filepath"
	"strconv"
)

type Cfg struct {
	ScmWorkspaceDirDefaultPerm os.FileMode
	ScmWorkingCopyPath         string
}

func ReadCfg(scmUrl string) (Cfg, error) {
	scmWorkspaceDir := LookupEnvOrDefault("SCM_WORKSPACE_DIR", "~/Workspace")
	scmExpanedWorkspaceDir, err := ExpandHomeDir(scmWorkspaceDir)
	if err != nil {
		return Cfg{}, err
	}

	scmWorkspaceDirDefaultPermStr := LookupEnvOrDefault("SCM_WORKSPACE_DIR_DEFAULT_PERM", "0755")
	scmWorkspaceDirDefaultPerm, err := strconv.ParseInt(scmWorkspaceDirDefaultPermStr, 8, strconv.IntSize)
	scmWorkspaceDirDefaultPermFileMode := os.FileMode(scmWorkspaceDirDefaultPerm)

	scmPathFromUrl, err := ExtractLocalPathFromScmURL(scmUrl)
	if err != nil {
		return Cfg{}, err
	}
	scmWorkingCopyPath := filepath.Join(scmExpanedWorkspaceDir, scmPathFromUrl)

	return Cfg{
		ScmWorkspaceDirDefaultPerm: scmWorkspaceDirDefaultPermFileMode,
		ScmWorkingCopyPath:         scmWorkingCopyPath,
	}, nil
}
