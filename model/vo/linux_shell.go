package vo

import (
	"github.com/airdb/scf-airdb/model/po"
	"gorm.io/gorm"
)

type LinuxShell struct {
	ID uint `json:"id"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Command string
	Tags    string
	Comment string
	Ref     string
}

func FromPoLinuxShell(shell *po.TabLinuxShell) *LinuxShell {
	return &LinuxShell{
		ID:      shell.ID,
		Command: shell.Command,
		Tags:    shell.Tags,
		Comment: shell.Comment,
		Ref:     shell.Ref,
	}
}

func ToPoLinuxShell(shell *LinuxShell) *po.TabLinuxShell {
	return &po.TabLinuxShell{
		Model:   gorm.Model{},
		Command: shell.Command,
		Tags:    shell.Tags,
		Comment: shell.Comment,
		Ref:     shell.Ref,
	}
}

type CreateLinuxShellReq struct {
	Command string `json:"command"`
	Comment string `json:"comment"`
	Ref     string `json:"ref"`
	Tags    string `json:"tags"`
}

type CreateLinuxShellResp struct {
}

type ListLinuxShellReq struct {
	PageNo   int `form:"page_no"`
	PageSize int `form:"page_size"`
}

type ListLinuxShellResp struct {
	Shells []*LinuxShell `json:"shells"`
}

func ListLinuxShell(req *ListLinuxShellReq) *ListLinuxShellResp {
	var resp ListLinuxShellResp

	shells := po.ListLinuxShell(req.PageNo, req.PageSize)

	for _, shell := range shells {
		_shell := FromPoLinuxShell(shell)

		resp.Shells = append(resp.Shells, _shell)
	}

	return &resp
}

func CreateLinuxShell(req *CreateLinuxShellReq) *CreateLinuxShellResp {
	shell := &LinuxShell{
		Command: req.Command,
		Tags:    req.Tags,
		Comment: req.Comment,
		Ref:     req.Ref,
	}

	po.CreateLinuxShell(ToPoLinuxShell(shell))
	return &CreateLinuxShellResp{}
}
