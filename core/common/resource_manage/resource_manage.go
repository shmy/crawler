package resource_manage

type ResourceManage interface {
	GetOne()
	FreeOne()
	Has() bool
}
