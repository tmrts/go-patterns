# integer

整数用8bit、16bit、32bit、64bit表示，
有符号整数分别对应int8、int16、int32、int64，用补码存储。
无符号整数uint8、uint16、unint32、unint64，用原码存储。
int、uint是使用最广泛的数值类型，是运算效率最高的数值。在编译时确定到底是多少位。

rune是int32类型的别名，表示一个Unicode(code point).
byte是uint8类型的别名，强调是原始的8位bit。

uintptr(u-int-ptr):无符号整数，大小不确定，长用于保存指针。在unsafe package中。

