# GO 002 Fyne-02 INKKAPLUM SH

## 前文

GO 002 Fyne-02

本教程是上一个的[Go_Fyne 教程](https://www.bilibili.com/video/BV1u142187Ps)的延续。

没有太多变化的部分, 相比而言, 文档因为内容更少所以对应地, 写得更详细了。

而桌面背景和Code背景也没有变化, 依然是 Harumakigohan(はるまきごはん)的ゼロトーキング和再会, Up 暂时没有换的打算, 但是下一期视频肯定会换。Up 主要背景主要选择会集中在「ボカロ曲」范畴, 也常常会有城市风景、J-Pop 艺人(如 Backnumber, YOASOBI, 優里)。

一般而言, 选择集中在はるまきごはん、Chinozo、ナブナ(n-buna)、ナユタン星人、ピノキオピー、DECO*27、まふまふ、ヨアソビ(YOASOBI)、バックナンバー(Backnumber)、ずっと真夜中でいいのに(ZTMY)、ヨルシカ、優里、Imase、Vaundy、Ado、「25 時、ナイトコードで。」、Official 髭男 Dism、藤井風、Uru、ヒトリエ(Hitorie), 假如你也很喜欢这些 P 主/艺人/组合/乐队等等(比如说是三夜人/四夜人), 欢迎和 Up 主分享你喜欢的曲子和一些有意思的事情, Up 的视频除了分享技术和观点, 也会给大家推荐一些歌曲。

## Up 的话(前言)

本教程是对于上一个教程中间最后提出的一个整活项目的对应专项说明, 同时, 对于Go语言更加复杂的内容进行一个深入的分析。帮助各位快乐地开发有趣的 GUI 应用程序, 并且更好地学习Go语言。而对应地, 灵感就来自于 B 站也有的一些开机检测器, 关机检测器等整活项目, 而单独拿出来作为一个快速教程, 毫无疑问是有理由的, 我们会让这个程序再复杂一点。比如说在检测过程中有一定的波折, 如中间又弹出一个窗口, 窗口上有说明, 如跳广告窗口(当然我们不可能实现这个广告), 正常而言, 可以模拟最常见的情况, 比如说关注获得密码, 然后输入密码的情况——也就是输入密码进一步检测等等功能, 这里就是关注InkkaPlum频道并且三连(你关注、三连了嘛(doge))。而重要的是, 好处很多, 还可以讲一下 Go 断言、常见小技巧等 Go 核心内容以及通过概述强化各位读源码的能力, 遇到问题先看源码, 看不懂再问 GPT 是一个很好的方法。

此外, 我们将充分地利用 Fyne 的主题文件, 也就是`Theme.go`文件, 我们将让我们的程序更加符合高端检测器的实际含义。这也是之前视频没有提及到的。

录制视频不仅是传递知识和解法, 也是对于 Up 自身的一个挑战和知识的再加强, 所以如果有任何不懂的地方, 尤其是视频没有讲清楚的地方, 欢迎私信问 Up 主, Up 主会尽力给出一个可用的解法。

当然, 欢迎关注 Up 的 B 站和知乎, 同时多多三连, 当然也可以充电支持 Up 主。大家的支持将给 Up 更多的动力, Up 也会努力给大家带来更好的视频。

同时, 所有课件和代码都在 GitHub 上开源地分享, 如果感到有帮助, 欢迎给一个 Star。

## 说明

本教程中全部文字版教程和代码为 B 站: [InkkaPlum 频道](https://space.bilibili.com/290859233)和知乎: [Inkka Plum](https://www.zhihu.com/people/instead-opt)的相关教程所用, 仅供学习。

不得二次用于任何机构/个人再次录制 Go / Fyne 或其它任何语言, 框架, 架构, 工具等等教程中。

此外, 下文有任何问题或者需要改进的点, 请联系 UP 主。

## 迫在眉睫: 解决热重载问题(Hot Reload)

活用 Air 工具, 此工具的开发者当初是因为`Gin`缺乏实时重载的功能而开发, 那么也顺便解决了 Fyne 学习者的问题。

参考 Github 开源项目地址的 Readme.md

<https://github.com/air-verse/air/blob/master/README.md>

配置非常简单:

第一步:

```bash
go install github.com/air-verse/air@latest
```

第二步:

官方解释: You can initialize the .air.toml configuration file to the current directory with the default settings running the following command.

```bash
air init
```

会有一个 `.air.toml`的文件, 留意即可。

然后直接 air 即可。

在`ターミナル`内, 出现此内容即成功!

```bash
  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.52.3, built with Go go1.22.5
```

接下来解析`theme.go`文件

## theme.go 文件解析

之前的教程没有提及到, 这一次我们需要提及一下。

继续讲一些 Go 的基础内容:

## 前置内容: Go 语言的类型断言

Go 语言的类型断言(Type Assertion)是一种用于将接口类型变量转换为具体类型的机制。

### 空接口进阶

在 Go 语言中, `interface{}`是一个空接口，可以接收任何类型的值。

函数参数例子:

```go
package main
import "fmt"
func showAny(mytest interface{}){
 fmt.Println(mytest)   //直接打印输出传入的内容, 如114
}

func main() {
 showAny("Gopher Gogogo")
 showAny(114)
 showAny(114.514)
}
```

输出:

```bash
Gopher Gogogo
114
114.514
```

### 类型断言正式分析

它的语法形式为：

```go
value, ok := x.(T)
```

下面是一个简单的结合活用示例：

```go
package main

import "fmt"

func main() {
    var i interface{} = "inkkaplum"

    s, ok := i.(string)
    if ok {
        fmt.Println("断言成功，值为:", s)
    } else {
        fmt.Println("断言失败")
    }
}
```

### 另一个例子：失败的类型断言

```go
package main

import "fmt"

func main() {
    var i interface{} = 123

    s, ok := i.(string)
    if ok {
        fmt.Println("断言成功, 值为:", s)
    } else {
        fmt.Println("断言失败")
    }
}
```

零值参考:

```go
package main

import "fmt"

func main() {
    var i interface{} = "123"

    s, ok := i.(int)
    if ok {
        fmt.Println("断言成功, 值为", s)
    } else {
        fmt.Println("断言失败, 值为", s)
    }
}
//断言失败, 值为0 即可理解。
```

### 类型断言的另一种用法：不检查成功与否

有时候, 如果你确定断言一定会成功, 可以省略 `ok` 值：

公式:

```go
value := x.(T)
```

```go
package main

import "fmt"

func main() {
    var i interface{} = "inkkaplum"
    s := i.(string)
    fmt.Println(s)
}
```

但是, 必须要留意的是, 这种情况下, 如果断言失败, 就会触发 panic。因此, 除非非常确定, 否则建议使用带 `ok` 检查的类型断言。

#### 对于 Panic 的解释

案例代码:

```go
package main

import "fmt"

func main() {
    var i interface{} = 123

    // 直接进行类型断言，没有使用 `ok` 检查
    s := i.(string)

    fmt.Println(s)
}
```

输出结果:

```bash
panic: interface conversion: interface {} is int, not string
```

类型断言失败时, 如果没有使用 ok 检查, 会直接导致 panic(如上), 程序会中止执行, 所以建议使用ok检查, 当然接下来涉及到对于`panic`的处理方法。

复杂的注意点:

`Goroutine`相关内容请移步至上一期的 Go Fyne 视频, <https://www.bilibili.com/video/BV1u142187Ps>

`panic`能在当前`Goroutine`中递归执行调用方(Caller)的`defer`

`defer`: expression in defer must be function call

请思考下面的例子

```go
package main

import "fmt"

func main() {
 defer func() {
 defer func() {
 fmt.Println("test1")
 }()
 fmt.Println("test")
 }()
 defer func(){
 fmt.Println("good")
 }()

 panic("panic01")
}
```

```go
func main() {
 defer fmt.Println("main")
 go func() {
 defer fmt.Println("goroutine")
 panic("")
 }()

 time.Sleep(1 * time.Second)
}
```

输出:

```bash
goroutine
panic:

goroutine ...
```

结论: `panic`只会触发当前`Goroutine`的`defer`。

请思考下一个例子:

```go
func main() {
 defer func() {
    defer func() {
    panic("panic03")
    }()
 panic("panic02")
 }()

 panic("panic01")
}
```

输出:

```bash
panic: panic01
 panic: panic02
 panic: panic03
```

最后, 请思考下面的综合代码:

```go
package main

import (
    "fmt"
)

func causePanic() {
    panic("something went wrong!")
}

func withRecovery() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from", r)
        }

        fmt.Println("Exiting withRecovery")
    }()
    fmt.Println("Entering withRecovery")
    causePanic()

    fmt.Println("This line will not be executed")
}

func withoutRecovery() {
    defer fmt.Println("Exiting withoutRecovery")
    fmt.Println("Entering withoutRecovery")
    causePanic() 

    fmt.Println("This line will not be executed either")
}

func main() {
    fmt.Println("Starting main")

    withRecovery()

    fmt.Println("Trying withoutRecovery (this will crash the program)")
    withoutRecovery() 

    fmt.Println("This line will not be executed due to the crash")
}
```

输出

```bash
Starting main
Entering withRecovery
Recovered from something went wrong!
Exiting withRecovery
Trying withoutRecovery (this will crash the program)
Entering withoutRecovery
Exiting withoutRecovery
panic: something went wrong!

goroutine 1 [running]:
...(堆栈跟踪(Stack trace)信息)
```

顺序问题:

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("捕获到panic:", r);
        }
    }()

    var i interface{} = 123

    // 直接进行类型断言，没有使用 `ok` 检查
    s := i.(string)

    fmt.Println(s)
}
```

反例(无法捕获)

```go
package main

import "fmt"

func main() {
    var i interface{} = 123

    // 直接进行类型断言，没有使用 `ok` 检查
    s := i.(string)

    fmt.Println(s)

    defer func() {
        if r := recover(); r != nil {
            fmt.Println("捕获到panic:", r)
        }
    }()
}
```

反例(无法捕获)

```go
func main() {
 defer fmt.Println("inkka")
 if err := recover(); err != nil {
 fmt.Println(err)
 }

 panic("something went wrong")
}
```

### Go 还有一些有趣的点——`unsafe.Pointer`和`uintptr`

和本案例无关, 但是需要提及一下, 尤其是有时在看源码是会看到, 所以有必要说明!

案例:

```go
package main

import "fmt"

func main() {
 a := 114514
 b := a
 fmt.Println(b)
 b = &a + 1
}
```

输出:

```bash
# command-line-arguments
invalid operation: &a + 114514 (mismatched types *int and untyped int)
```

案例2:

```go
func main(){
 num := 114
 numPointer := &num

 numb := (*float32)(numPointer)
 fmt.Println(numb)
}
```

输出:

```bash
# command-line-arguments
cannot convert numPointer (variable of type *int) to type *float32
```

解决方法:

活用`unsafe.Pointer`

定义:

unsafe 包下的 Pointer,

源码

```go
type Pointer *ArbitraryType
```

案例：

```go
package main

import (
 "fmt"
 "unsafe"
)
func main(){
 num := 114
 numPointer := &num

 numb := (*float32)(unsafe.Pointer(numPointer))
 fmt.Println(numb)
}
```

输出, 没有任何问题!

```bash
0xc000096068
```

活用Go来了解

```go
package main

import "fmt"
import "unsafe"

func main() {
    name := "inkka"
    var age uint32 = 17
    var studentID float32 = 13

    fmt.Printf("name type and size: %T, %d\n", name, unsafe.Sizeof(name))
    fmt.Printf("age type and size: %T, %d\n", age, unsafe.Sizeof(age))
    fmt.Printf("studentID type and size: %T, %d\n", studentID, unsafe.Sizeof(studentID))
    fmt.Printf("name len: %T, %d\n", name, len(name))

}
```

分析`uintptr`

```go
// uintptr is an integer type that is large enough to hold the bit pattern of
// any pointer.
type uintptr uintptr
```

注意!

uintptr 是一个整型(这个非常重要)!

unsafe 包支持了这些函数来完成到`uintptr`的转换, `Sizeof`即为我们刚刚用到的。

```go
func Sizeof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Alignof(x ArbitraryType) uintptr
```

分析代码:

```go
package main

import (
 "fmt"
 "unsafe"
)

func main() {

 arr := []int{10, 20, 30, 40, 50}

 ptr := unsafe.Pointer(&arr[0])

 ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0]))


 *(*int)(ptr) = 25

 fmt.Println(arr) // 输出 [10, 25, 30, 40, 50]
}
```

再次理解:

请参考下面的简单案例代码:

```go
package main

import (
 "fmt"
 "unsafe"
)

func main() {
 var x = 114

 p := &x
 up := unsafe.Pointer(p) 
 uptr := uintptr(up)   
 fmt.Printf("pointer: %p;十进制为：%d\n", p, p)
// pointer: 0xc00000a0c8; 十进制为：824633761992(Up案例)
 fmt.Printf("uintptr: %d\n", uptr) // 地址的十进制表示
// uintptr: 824633761992
}

```

当然, 我们也可以在原案例上修改, 以再一次验证:

```go
package main

import (
 "fmt"
 "unsafe"
)

func main() {

 arr := []int{10, 20, 30, 40, 50}

 ptr := unsafe.Pointer(&arr[0])

 fmt.Println(ptr)
 fmt.Println(uintptr(ptr))

 ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0]))


 *(*int)(ptr) = 25

 fmt.Println(arr) 
}
```

输出(Up的例子):

```bash
0xc0000a8030
824634409008
[10 25 30 40 50]
```

活用工具: 可以使用计算器的十六(`HEX`)转十进制(`DEC`)换算一下。

### 继续分析 Fyne

参考文档:
<https://docs.fyne.io/extend/custom-theme.html>

### 定义主题

官方文档:

```go
type myTheme struct {}
```

我们的案例:

```go
type myTheme struct {
  font fyne.Resource
}
```

## 官方文档阅读小技巧: 接口满足检查

而接下来还有一段文字:

(It is a good idea to assert that we implement an interface so that compile errors are closer to the defining type.)

```go
var _ fyne.Theme = (*myTheme)(nil)
```

翻译: 断言我们实现了某个接口是一个好主意，这样可以使编译错误更接近定义类型的地方。

`_`可以叫做空标识符。

比如

在多重赋值时，如果你不需要使用所有的值, 可以使用`_`来忽略它们。

```go
write, err := storage.Writer(cfg.CurrentFile)
```

```go
write, _ := storage.Writer(cfg.CurrentFile)
```

当使用空标识符来接收一个值时, 作为无用变量的占位符，可以接受任何类型的值，而该值将被无害地丢弃。

### 公式

```go
var _ Interface = (*Type)(nil)
```

### 实际例子

有一个接口 `fyne.Theme`，其中定义了多个方法：(位置: `theme.go`)

```go
package fyne

//...

type Theme interface {
 Color(ThemeColorName, ThemeVariant) color.Color
 Font(TextStyle) Resource
 Icon(ThemeIconName) Resource
 Size(ThemeSizeName) float32
}
```

现在我们定义了一个类型 `myTheme`，并希望它实现 `fyne.Theme` 接口, 但是我们注释掉了 Size 部分的代码：

```go
package main

import (
 "fyne.io/fyne/v2"
 "fyne.io/fyne/v2/theme"
 "image/color"
)

type myTheme struct {
 font fyne.Resource
}

func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource {
 return m.font
}

func (m *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
 return theme.DefaultTheme().Color(n, v)
}

func (m *myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
 return theme.DefaultTheme().Icon(n)
}

// func (m *myTheme) Size(n fyne.ThemeSizeName) float32 {
//  return theme.DefaultTheme().Size(n)
// }

var _ fyne.Theme = (*myTheme)(nil)
```

报错:

```bash
.\theme.go:29:20: cannot use (*myTheme)(nil) (value of type *myTheme) as fyne.Theme value in variable declaration: *myTheme does not implement fyne.Theme (missing method Size)
.\main.go:16:24: cannot use &myTheme{…} (value of type *myTheme) as fyne.Theme value in argument to a.Settings().SetTheme: *myTheme does not implement fyne.Theme (missing method Size)
```

好处: 提高代码的健壮性和可维护性。如上文的`.\theme.go:29:20`。

反例:

假如没有活用接口满足检查, 报错如下:

```bash
.\main.go:16:24: cannot use &myTheme{…} (value of type *myTheme) as fyne.Theme value in argument to a.Settings().SetTheme: *myTheme does not implement fyne.Theme (missing method Size)
```

可以看到是`.\main.go:16:24` 并不在`theme.go`文件内, 位置不够清晰, 不具有较好的健壮性和可维护性。

## main.go代码

```go
package main

import (
 "fyne.io/fyne/v2/app"
 "fyne.io/fyne/v2/container"
 "fyne.io/fyne/v2/theme"
 "fyne.io/fyne/v2/widget"
)

func main() {
 a := app.New()
 a.Settings().SetTheme(&myTheme{}) // 设置自定义主题

 w := a.NewWindow("Custom Theme with Icon Override")
 w.SetContent(container.NewVBox(
 widget.NewLabel("Hello, Fyne!"),
 widget.NewButtonWithIcon("Home!!!", theme.HomeIcon(), func() {}),
 ))

 w.ShowAndRun()
}
```

## 具体配置

接下来我们思考具体配置问题:

### 官方文档: 自定义颜色

```go
func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
 if name == theme.ColorNameBackground {
  if variant == theme.VariantLight {
   return color.White
  }
  return color.Black
 }

 return theme.DefaultTheme().Color(name, variant)
}
```

解释一下!

#### 代码概览说明

```go
func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color
```

#### 1. 判断背景颜色

```go
if name == theme.ColorNameBackground {
    if variant == theme.VariantLight {
        return color.White
    }
    return color.Black
}
```

#### 2. 默认颜色处理

```go
return theme.DefaultTheme().Color(name, variant)
```

默认情况处理：使用`default`关键字

import 部分注意:

```go
"image/color"
```

```go
func (m *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
 switch n {
 case theme.ColorNameBackground:
 return color.RGBA{190, 233, 255, 1}
 case theme.ColorNameButton:
 return color.RGBA{0, 122, 255, 255}
 case theme.ColorNameDisabledButton:
  return color.RGBA{142, 142, 147, 255}
 case theme.ColorNameHover:
  return color.RGBA{230, 230, 230, 255}
 case theme.ColorNameFocus:
  return color.RGBA{255, 165, 0, 255}
 case theme.ColorNameShadow:
  return color.RGBA{0, 0, 0, 50}
 default:
  return theme.DefaultTheme().Color(n, v)
 }
}
```

默认处理不能乱写, 如果按照

```go
default:
return color.White 
```

结果将会有很大问题。

### 官方文档: 自定义应用内图标

前置命令和概念: Bundling resources 捆绑资源

Fyne 文档官方介绍翻译, 我们只需要抓关键词而已:

<https://docs.fyne.io/extend/bundle.html>

Up 总结: Fyne Bundle 工具作用:

1. 将外部资源(如图像)打包成 Go 代码, 使其成为应用程序的一部分。
2. 生成的 Go 文件包含资源的字节数据, 这样资源文件就可以与应用程序一起分发, 不需要单独的资源文件, 具体而言, 不需要如`.jpeg`的文件。

Fyne官方文档代码:

```bash
fyne bundle -o bundled.go image.png
```

```go
var resourceImagePng = &fyne.StaticResource{
 StaticName: "image.png",
 StaticContent: []byte{
 //包含资源的字节数据,
 }}
```

默认命名为 `"resource<Name>.<Ext>"`。

文档案例:

```go
img := canvas.NewImageFromResource(resourceImagePng)
```

接下来我们要利用此命令为按钮添加图标。

然后请按照官方文档, 输入此命令:

```bash
fyne bundle -o bundled.go icon.jpeg
```

然后你会看到有一个`bundled.go`文件, 请查看内部内容:

```go
// auto-generated
// Code generated by '$ fyne bundle'. DO NOT EDIT.

package main

import "fyne.io/fyne/v2"

var resourceIconJpeg = &fyne.StaticResource{
 StaticName: "icon.jpeg",
 StaticContent: []byte(
 ...
}
```

然后在`theme.go`文件中, 进行如下修改:

```go
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
 if name == theme.IconNameHome {
 return resourceIconJpeg
 }
 return theme.DefaultTheme().Icon(name)
}
```

图标库(Icon library)。

如: <https://remixicon.com/> 等等。

我们接下来只需要使用此函数:

```go
widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {})
```

第一、三个参数不必解释, 第二个参数理解如果觉得有些困难, 看一下源码即可:

```go
func HomeIcon() fyne.Resource {
 return safeIconLookup(IconNameHome) 
}
```

```go
func safeIconLookup(n fyne.ThemeIconName) fyne.Resource {
 icon := Current().Icon(n)
 if icon == nil {
 fyne.LogError("Loaded theme returned nil icon", nil)
 return fallbackIcon
 }
 return icon
}
```

### 官方文档: 字体和指定元素的大小

这里官方文档写得很简单, 我们对于字体进行一个说明。

```go
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
    return theme.DefaultTheme().Font(style)
}
```

我们的案例：

结构体:

```go
type myTheme struct {
  font fyne.Resource
}
```

```go
func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource {
 return m.font
}
```

进阶参考:

```go
type myTheme struct {
    regular   fyne.Resource
    bold      fyne.Resource
    italic    fyne.Resource
    monospace fyne.Resource
}
```

```go
func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource {
    if s.Monospace {
        return m.monospace
    }
    if s.Bold {
        return m.bold
    }
    if s.Italic {
        return m.italic
    }
    return m.regular
}
```

Size就简述一下了

```go

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
    return theme.DefaultTheme().Size(name)
}
```

设置主题:

```go
app.Settings().SetTheme(&myTheme{})
```

我们的案例(本案例不做参考):

```go
customFont := fyne.NewStaticResource("NotoSansHans-Regular.ttf", loadFont("NotoSansHans-Regular.ttf"))
```

`util.go`

```go
package main

import (
 "os"
 "log"
)

func loadFont(fontPath string) []byte {
 // 加载字体文件
 fontData, err := os.ReadFile(fontPath)
 if (err != nil) {
  log.Fatalf("无法加载字体文件: %v", err)
 }
 return fontData
}
```

参考代码:

```go
package main

import (
  "log"
)

type Student struct {
  Name string
  Age  int
}

func main() {
  u := Student{
    Name: "Inkka",
    Age:  17,
  }

  log.Printf("Student name: %s, age:%d", u.Name, u.Age)
  log.Printf("Student name: %s, age:%d", u.Name, u.Age)
}
```

输出

```bash
2024/07/28 23:08:42 Student name: Inkka, age:17
2024/07/28 23:08:42 Student name: Inkka, age:17
```

请参考下面的代码:

```go
  log.Panicf("%s is sleeping", u.Name)
```

输出

```bash
panic: Inkka is sleeping
goroutine 1 [running]:
...
```

请参考下面的代码:

```go
log.Fatalf("%s is sleeping", u.Name)
```

输出

```bash
2024/07/28 23:11:51 Inkka is sleeping
exit status 1
```

defer相关

```go
type Student struct {
  Name string
  Age  int
}

func main() {
  u := Student{
    Name: "Inkka",
    Age:  17,
  }
  defer fmt.Println("good")
  log.Printf("Student name: %s, age:%d", u.Name, u.Age)
  log.Fatalf("%s is sleeping", u.Name)
} 
```

输出:

```bash
2024/07/28 23:22:42 Student name: Inkka, age:17
2024/07/28 23:22:42 Inkka is sleeping
exit status 1
```

```go
customFont := fyne.NewStaticResource("NotoSansHans.ttf", loadFont("NotoSansHans-Regular.ttf"))
```

类型为: `*fyne.StaticResource`。

加载主题:

参考官方文档, 用此代码即可

```go
app.Settings().SetTheme(&myTheme{})
```

## 修改我们的应用

原有代码:

```go
package main

import (
 "time"

 "fyne.io/fyne/v2"
 "fyne.io/fyne/v2/app"
 "fyne.io/fyne/v2/container"
 "fyne.io/fyne/v2/dialog"
 "fyne.io/fyne/v2/widget"
)

func main() {
 a := app.NewWithID("01")
 customFont := fyne.NewStaticResource("NotoSansHans.ttf", loadFont("NotoSansHans-Regular.ttf"))
 a.Settings().SetTheme(&myTheme{font: customFont})

 w := a.NewWindow("高端检测器")

 btn := widget.NewButton("检测是否开机", func() {
 progress := widget.NewProgressBarInfinite()
 progressContainer := container.NewVBox(progress)

 loadingDialog := dialog.NewCustom("正在检测...", "取消", progressContainer, w)
 loadingDialog.Show()

 go func(){time.Sleep(20 * time.Second)
 loadingDialog.Hide()
 dialog.ShowInformation("结果", "电脑是开机的", w)}()
 })

 w.SetContent(container.NewVBox(btn))
 w.Resize(fyne.Size{Width: 114, Height: 514})
 w.CenterOnScreen()
 w.ShowAndRun()
}
```

```go
package main

import (
 "time"

 "fyne.io/fyne/v2"
 "fyne.io/fyne/v2/app"
 "fyne.io/fyne/v2/container"
 "fyne.io/fyne/v2/dialog"
 "fyne.io/fyne/v2/widget"
)

func main() {
 a := app.NewWithID("01")
 customFont := fyne.NewStaticResource("NotoSansHans.ttf", loadFont("NotoSansHans-Regular.ttf"))
 a.Settings().SetTheme(&myTheme{font: customFont})

 w := a.NewWindow("高端检测器")

 btn := widget.NewButton("检测是否开机", func() {
  showAdDialog(w)
 })

 w.SetContent(container.NewVBox(btn))
 w.Resize(fyne.Size{Width: 114, Height: 514})
 w.CenterOnScreen()
 w.ShowAndRun()
}

func showAdDialog(w fyne.Window) {
 entry := widget.NewPasswordEntry()
 form := &widget.Form{
  Items: []*widget.FormItem{
   {Text: "请输入密码继续检测", Widget: entry},
  },
  OnSubmit: func() {
   if entry.Text == "114514" {
    showLoadingDialog(w)
   } else {
    dialog.ShowInformation("错误", "密码错误, 请关注B站InkkaPlum频道三连获取正确密码", w)
   }
  },
 }

 adContent := container.NewVBox(
  widget.NewLabel("请关注B站InkkaPlum频道并三连得到密码以继续检测。"),
  form,
 )

 adDialog := dialog.NewCustom("注意", "取消", adContent, w)
 adDialog.Show()
}

func showLoadingDialog(w fyne.Window) {
 progress := widget.NewProgressBarInfinite()
 progressContainer := container.NewVBox(progress)

 loadingDialog := dialog.NewCustom("正在检测...", "取消", progressContainer, w)
 loadingDialog.Show()

 go func() {
  time.Sleep(20 * time.Second)
  loadingDialog.Hide()
  dialog.ShowInformation("结果", "电脑是开机的", w)
 }()
}

```

源码参考:

```go
type Form struct {
 BaseWidget

 Items      []*FormItem
 OnSubmit   func() `json:"-"`
 OnCancel   func() `json:"-"`
 SubmitText string
 CancelText string

 // Orientation allows a form to be vertical (a single column), horizontal (default, label then input)
 // or to adapt according to the orientation of the mobile device (adaptive).
 //
 // Since: 2.5
 Orientation Orientation

 itemGrid     *fyne.Container
 buttonBox    *fyne.Container
 cancelButton *Button
 submitButton *Button

 disabled bool

 onValidationChanged func(error)
 validationError     error
}
```

```go
type FormItem struct {
 Text   string
 Widget fyne.CanvasObject

 // Since: 2.0
 HintText string

 validationError error
 invalid         bool
 helperOutput    *canvas.Text
}
```

## 附录

可能会遇到的错误

```bash
Fyne error:  Preferences load error:
11:45:14   Cause: EOF
```

这里简单解释一下

在应用程序尝试加载用户首选项时, 遇到了一个`"EOF" (End Of File)`错误。也就是说, 在读取首选项(`Preferences.go`)文件时, 文件意外地到达了结尾, 没有读取到预期的数据, 简单而言, 表示应用程序在期望读取更多数据时已经达到了文件的末尾。但需要注意, 并不会影响运行。

## 结语

Up的Fyne第一期也算正式结束了, 未来 Up 将会更新一期使用 Fiber+Go 的教程, 而接下来就会是 Unity Node和一些观点相关、分析短文的内容了, 而 Go 更多的项目需要等 Up 以后有时间再继续拓展。

此外, 欢迎关注 Up 的 B 站频道和知乎, 并且别忘了一键三连, 当然如果愿意, 欢迎给 Up 充电支持, 您的支持是 Up 前进的动力, 将会鼓励 Up 给各位带来更好的视频。

同时, 所有课件和代码都在 GitHub 上分享, 如果感到有帮助, 欢迎给一个 Star。

扩充内容: 之后还会有一个 Fyne 教程, Up 会运用多窗口为我们的 markdown 编辑器增加一些功能, 与此同时, 还会用 Fyne 实现我们之前 Vue 案例中的蓝鼠兑换项目, 这一个教程也将在之后配合Fiber和各位见面。

Up B 站 InkkaPlum 频道

<https://space.bilibili.com/290859233>

Up 知乎

<https://www.zhihu.com/people/instead-opt>

Up 掘金

<https://juejin.cn/user/3529872175284560>

Up GitHub

<https://github.com/Slumhee>

以上 祝学习成功!

Inkka Plum
