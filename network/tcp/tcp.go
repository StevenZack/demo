package tcp

type Message struct {
	SourcePort, TargetPort uint16 // 源端口，目标端口
	Sequence               uint32 // 序号
	Acknowledgement        uint32 // 确认号
	Flags, Window          uint16 // 数据偏移4bit+保留6bit+标志位6bit ; 窗口16bit
	Sum, UrgPtr            uint16 // 检验和16bit ; 紧急指针16bit
	Options                []byte // 选项（可变长度）
	Fill                   []byte // 填充（为了保证Head总长度是4的整数倍）
}

// ------------ Flags 16bit------------------
// Offset 其实应该返回uint4的，因为偏移量只有半个字节
func (m *Message) Offset() uint8 {
	return uint8(m.Flags >> 12)
}

// Preserved 保留位 6bit
func (m *Message) Preserved() uint8 {
	return uint8((m.Flags >> 6) & 0b111111)
}

// URG 紧急标志位 1bit
func (m *Message) URG() bool {
	return ((m.Flags >> 5) & 1) == 1
}

// ACK 确认标志位 1bit
func (m *Message) ACK() bool {
	return ((m.Flags >> 4) & 1) == 1
}

// PSH Push标志位 1bit
func (m *Message) PSH() bool {
	return ((m.Flags >> 3) & 1) == 1
}

// RST Reset标志位 1bit
func (m *Message) RST() bool {
	return ((m.Flags >> 2) & 1) == 1
}

// SYN Sync标志位 1bit
func (m *Message) SYN() bool {
	return ((m.Flags >> 1) & 1) == 1
}

// FIN Final标志位
func (m *Message) FIN() bool {
	return (m.Flags & 1) == 1
}

