package resource_manage

type ResourceManageChan struct {
	mc chan int
}

func NewResourceManageChan(num int) *ResourceManageChan {
	mc := make(chan int, num)
	return &ResourceManageChan{
		mc: mc,
	}
}

func (r *ResourceManageChan) GetOne() {
	r.mc <- 1
}

func (r *ResourceManageChan) FreeOne() {
	<-r.mc
}

func (r *ResourceManageChan) Has() bool {
	return len(r.mc) != 0
}
