package Core

import "git.docus.tech/kdl12138/DocusServer/Template"

func NewServer(Setting Template.Server) {
	for i := 0; i < len(Setting.Nodes.Node); i++ {
		if Setting.Nodes.Node[i].Status {
			continue
		} else {
			Setting.Nodes.Node[i].Status = true

		}
	}
}
