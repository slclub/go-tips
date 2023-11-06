## Golang 的业务流工具类
为了在业务流程中少些代码，不是特别要求极致性能的。本工具适合业务流，都以已经IO操作了，这点性能劣势可以完全看不见。
我们主打就是为了工作效率，以及公用方法少些bug

--------------

### strings 字符串类的

#### String(arg any) string ｜ 转字符串

简化代码，主打一个转换字符串简单，稳健，偏低效；String 方法并没有 实现全部类型的转字符串的； 请开发时候，看看源码支持哪些类型的转化

支持类型：

    可以将数字包括小数 转成 字符串
    
    - 小数我们只实现了 float64 您需要先转float64; String(float64(小数))
    
    - error 标注类型
    
    - 实现 String() string 函数方法的 对象，包含指针，和值类型
 
    - io.Reader 接口

    - 当然更包含 字符串本身

#### StrPos(source string, str_piece string) int  ｜ 字符串位置查找

在某个字符串中 查找一个字符串片段的位置；return  -1:未找到; >=0 :返回具体位置序号

抄袭PHP 的标准函数库


#### StrBegin(str string, sep string) string ｜ 截取某个字符串片段以前的字符串

```go
    str := "hello world go tips"
    if StrBegin(str, "o") != "hell" {
        t.Error("strings StrPos expect hell")
    }
    if StrBegin(str, "tip") != "hello world go " {
        t.Error("strings StrPos expect 'hello world go '")
    }
```
    
#### StrEnd(str, sep string) string | 截取某个字符串 片段 以后的字符串

```go
    str := "hello world go tips"
    if StrEnd(str, "o") != " world go tips" {
        t.Error("strings StrPos expect: world go tips")
    }
    if StrEnd(str, "tip") != "s" {
        t.Error("strings StrPos expect:s")
    }
```

--------------

### date 

格式：2006-01-02 15:04:05

#### StrToTimestamp(str string) int64 | 字符串转时间戳

#### TimestampToStr(timestamp int64) string | 时间戳转字符串

--------------

### spinLock 无IO 锁

实现了sync.Locker 接口

```go
    mu := spinlock.New()
    mu.Lock()
    // do something thread unsafe
    defer mu.Unlock()
```

--------------

### comment functions

#### IsNil(i interface{}) bool | 正确的判断 nil 和 空接口的nil

#### ConfigWithViper(file_name_any string) *viper.Viper ｜ 快速，要求不高的使用viper

#### FUNC_NAME(f interface{}) string  ｜ 获取函数名

#### IS_FUNC(fn interface{}) bool | 是否是函数

#### IS_SLICE(s interface{}) bool ｜ 是否是 切片

#### IS_STRUCT(stu interface{}) bool | 是否是结构体

#### IfThree(false)("yes", "no") | 简短的if 方式 争取一行做完，比较适合 接口类，否则还需要写 类型转换
每次使用 也是多消耗了内存的

#### Any2Int64(v any) int64 ｜ 转int64

#### Int(v any) int ｜ 转 int

#### GetRootPath() string ｜ 获取 go 服务执行文件 对应目录

--------------

### logf.Logger

主要是定义了比较通用的日志接口，多数框架上 都是实现的这个;一般不会为日志定义太多接口函数

为了不统一 就把接口定义提取 出来， 不必 每个项目都写，还不知道是不是一样 method

- logf.Logger 定义
```go
    type Logger interface {
        Printf(string, ...any)
        Print(...any)
    }
```

- loger struct | 以fmt 简单的实现了 Logger 接口

--------------

### json 包 
接口与标准的是一样的，值得拿出来的原因是：防止数字总会专成 float64 ， 我们以 int 为主

--------------

### safe 包

名字叫 的并不贴切

#### 安全的int型Map

仅仅做了key是int 的 map ； 并发安全，用的比较多，所以单独写出来了，其他类型的没怎么重复用过；就没放进去

最重要的是有了 ```sync.Map``` 高性能的 安全map

```go
    type QuickIntMap struct {
        data   map[int]any
        splock sync.Locker
    }
```

#### Slice 自定义 系列 

- 增加常用整数类型的切片操作，查找，合并，删除，比较等
- 可以直接互相强制转换标准复杂类型，这里的
```go
    s1 := []string{"1", "2", "3"}
    ss1 := SliceString(s1) // 强制转化标准类型
    []string(ss1) // 强制转化 自定义 SliceX 类型
```
- 仅仅提供 []string, []intX, []Value; Value 是自定义类型
```go
    type Value interface {
        Value() int64
    }
```
- 可以支持 原生golang 的语法 [:], append,copy 等操作 
- 多数组简易合并，是否 unique 都可以使用不同的方法


##### 1. type SliceString []string | 字符串切片

- 可以直接强转标准类型
```go
    s1 := []string{"1", "2", "3"}
    ss1 := SliceString(s1)
```

- 支持的Methods

```go
    func (this *SliceString) Append(val string)
    func (this *SliceString) AppendUnqiue(val string) 
    func (this *SliceString) AppendArr(arrs ...[]string)
    func (this *SliceString) Del(val string) int
    func (this *SliceString) DelKey(k int) string 
    func (this *SliceString) Reset() 
    func (this *SliceString) Len() int 
    func (this *SliceString) In(val string) int 
    func (this *SliceString) Join(sep string) string 
    func (this *SliceString) Range(fn func(i int, val string) bool) 
```

- 同样支持[]string 的所有操作


##### 2. type SliceIntX []intX 系列

- 其中X包含：int,int32,int64;其他的int16 类似uint等都没去实现
- 可以与标准类型互换

- 支持类型
```go
    type SliceInt []int
    type SliceInt32 []int32
    type SliceInt64 []int64
```

- 方法

以 SliceInt 为例子
```go
    func (this *SliceInt) Append(val int)
    func (this *SliceInt) AppendUnqiue(val int) 
    func (this *SliceInt) AppendArr(arrs ...[]int)
    func (this *SliceInt) Del(val int) int
    func (this *SliceInt) DelKey(k int) int
    func (this *SliceInt) Reset()
    func (this *SliceInt) Len() int
    func (this *SliceInt) In(val int) int
    func (this *SliceInt) Join(sep string) string
    func (this *SliceInt) Range(fn func(i, val int) bool) 
```

##### 3. type SliceValue []Value ｜ 加了一个变相的可比较类型， 以实现通用化

- Value 是自定义接口，任何实现它的方法的，都可以用
```go
    type Value interface {
        Value() int64
    }
```
- Value 就无法
- 除了copy 同样支持 slice 的标准操作
- 方法
```go
    func (this *SliceValue) Append(val Value)
    func (this *SliceValue) AppendUnqiue(val Value)
    func (this *SliceValue) AppendArr(arrs ...[]Value) 
    func (this *SliceValue) Del(val Value) int 
    func (this *SliceValue) DelKey(k int) Value 
    func (this *SliceValue) Reset() 
    func (this *SliceValue) Len() int 
    func (this *SliceValue) In(val Value) int 
    func (this *SliceValue) Range(fn func(i int, val Value) bool) 
```

--------------

