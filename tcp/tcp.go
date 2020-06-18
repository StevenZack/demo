package tcp

type Head struct {
	SourcePort, TargetPort [2]byte // 源端口，目标端口
	Sequence               [4]byte // 序号
	Acknowledgement        [4]byte // 确认号
	Flag, Window           [2]byte // 数据偏移4bit+保留6bit+标志位6bit ; 窗口16bit
	Sum, UrgPtr            [2]byte // 检验和16bit ; 紧急指针16bit
	Options                []byte  // 选项（可变长度）
	Fill                   []byte  // 填充（为了保证Head总长度是4的整数倍）
}
