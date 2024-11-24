package errs

import "errors"

var (
	EmptyUserGroupName = errors.New("用户组名为空")
)
