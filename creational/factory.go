package creational

import "io"

// 工厂方法创建设计模式允许创建对象，而不必指定将要创建的对象的确切类型。
// 示例实现展示了如何使用不同的后端，如内存、磁盘存储。

// Store 对象类型
type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}
