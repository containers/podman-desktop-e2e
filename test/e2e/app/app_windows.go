package app

import (
	"github.com/adrianriobo/goax/pkg/os/windows/powershell"
)

func cleanup() error {
	ps := powershell.New()
	removePodmanDesktopHome := "Remove-Item \"$env:USERPROFILE\\.podman-desktop\" -Recurse -Force -Erroraction silentlycontinue"
	ps.Execute(removePodmanDesktopHome)
	// if err != nil {
	// 	return err
	// }
	removePodmanDesktopShare := "Remove-Item \"$env:USERPROFILE\\.local\\share\\containers\\podman-desktop\" -Recurse -Force -Erroraction silentlycontinue"
	ps.Execute(removePodmanDesktopShare)
	// TODO delete this path
	// C:\Users\crcqe\AppData\Roaming\Podman Desktop
	return nil
}
