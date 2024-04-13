# 基本介绍
vrule它具有以下功能：
<ol>
<li>使用gogf的校验规则</li>
<li>使用验证标记或自定义验证程序进行跨字段验证。</li>
<li>切片、数组和map，允许验证多维字段的任何或所有级别。</li>
<li>能够深入map的value和slice类型结构体以进行验证</li>
<li>提取自定义字段名称，例如，可以指定在验证时提取 JSON 名称，并在生成的 FieldError 中提供该名称</li>
<li>指定过滤掉某些字段</li>  
<li> 动态替换错误提示信息</li>  
<li>可自定义的i18n错误消息。</li>
<li> 可以将需要验证的结构体信息进行缓存，加快验证效率</li>  

</ol>

---
# 安装

#### go版本要求 >=1.18

#### 使用go get
```go
go get -u github.com/wln32/vrule
```
然后将验证程序包导入到您自己的代码中<br>
```go
import github.com/wln32/vrule
```
----
# 校验规则文档
[校验规则详细信息](https://goframe.org/pages/viewpage.action?pageId=1114367)

---
# 基本使用
vrule使用v或者valid来标记字段是否需要被校验</br>
使用GetFieldError来获取对应字段的错误提示信息</br>
#### 简单示例
```go
type Basic struct {
    Int8Ptr  *int8          `v:"required"`
    String   string         `v:"required"`
    Int      int            `valid:"required"`

}
obj:=Basic{}
err := Struct(obj).(*ValidationError)
fmt.Println(err.GetFieldError("Int8Ptr"))
fmt.Println(err.GetFieldError("String"))
fmt.Println(err.GetFieldError("Int"))

```

#### required-without关联规则示例
格式: required-without:field1,field2,... </br>
必需参数(当所给定任意字段值其中之一为空时)。当前字段必须有值</br>
```go
type Foo struct {
    Bar *Bar `p:"bar" v:"required-without:Baz"`
    Baz *Baz `p:"baz" v:"required-without:Bar"`
}
type Bar struct {
    BarKey string `p:"bar_key" v:"required"`
}
type Baz struct {
    BazKey string `p:"baz_key" v:"required"`
}
err := Struct(foo)

```


<p>

</p>

---
# 注册自定义规则的验证
```go
type CustomStruct struct {
	Name string `v:"trim-length:6"`
}
fn := func(ctx context.Context, in *ruleimpl.CustomRuleInput) error {
    val := in.Value.(string)
    trimVal := strings.TrimSpace(val)
    trimLength, err := strconv.Atoi(in.Args)
    if err != nil {
        return err
    }

    if len(trimVal) != trimLength {
        return fmt.Errorf("the length of the string after removing spaces must be %d characters", trimLength)
    }
    return nil
}
// 注册自定义规则验证
err := RegisterCustomRuleFunc(RegisterCustomRuleOption{
    RuleName: "trim-length",
    Fn:       fn,
})
```


---
# 提取自定义的字段名
可以根据字段的某些属性来过滤掉一些字段
```go
type OptionFieldName struct {
	Name string `json:"name" v:"required"`
}
valid := New()

valid.SetFieldNameFunc(func(_ reflect.Type, field reflect.StructField) string {
    // 使用字段的json tag作为错误提示时的字段名{field}
	name := field.Tag.Get("json")
    if name == "" {
        name = field.Name
    }
    return name
})

```

---

---
# 过滤掉某些字段
一般情况下我们需要过滤掉一些字段，可以使用SetFilterFieldFunc这个函数来实现<br>
返回true代表需要被过滤<br>
返回false代表需要被验证<br>
```go
valid := New()

valid.SetFilterFieldFunc(func(structType reflect.Type, field reflect.StructField) bool {
	// 如果字段带有d tag的我们就过滤掉，不验证这个字段
	tag := field.Tag.Get("d")
	if tag != "" {
		return true
	}
	return false
})
```
---

# 只获取第一个错误
vrule默认获取到全部的错误，不会在第一个错误时停下</br>
如果需要第一个错误就返回需要调用StopOnFirstError</br>
```go
valid := New()
// 设置为true代表遇到第一个错误时就立即返回
valid.StopOnFirstError(true)

obj := &OptionStopOnFirstError{}
```
---
# 替换错误提示信息
required = The {field} field is required</br>
这是vrule的required规则错误提示模板</br>
如果发生错误，vrule会将{field}替换为检验的字段名</br>
例如以下
```go
type Basic struct {
	String  string    `v:"required"`		
}
```
发生错误时，vrule会返回`The String field is required`这样的错误信息</br>
除此之外，vrule还会替换以下的模板字段
<ul>
<li>{value} 字段的实际值，替换发生在运行时</li>
<li>{max} max和between规则下替换，是一个静态的数字类型的值</li>
<li>{min} min和between规则下替换，是一个静态的数字类型的值</li>
<li>{pattern} 是一个正则表达式</li>
<li>{size} size规则下，是一个静态的数字类型的值</li>
<li>{field1} 依赖于某些字段的名字</li>
<li>{value1} 依赖于某些字段的实际值</li>
<li>{format} 格式化规则</li>
</ul>   

---
# 错误提示信息
GetFieldError 基础类型，或者只要map或者slice的元素是基本类型就可以，例如map[k]int,[]int</br>
GetStructFieldError 如果字段是struct类型</br>
GetMapFieldError  如果map的v是struct类型，例如map[k]struct或者map[k]*struct</br>
GetSliceFieldError 如果切片的元素是struct类型，例如[]struct,[]*struct</br>

#### 简单示例
```go
type Test struct {
    Pass1 string `v:"required|same:Pass2"`
    Pass2 string `v:"required|same:Pass1"`
	BasicSlice []int   `v:"required"`
}
type Example struct {
    Id   int
    Name string `v:"required"`
    Pass Test
}
obj := &Example{
    Name: "",
    Pass: Test{
        Pass1: "1",
        Pass2: "2",
    },
}
err := Struct(obj).(*ValidationError)
// 可以获取到Name字段的错误
err.GetFieldError("Name")
// 
err.GetFieldError("BasicSlice")
// 可以获取到Pass结构体字段下的Pass1错误
err.GetStructFieldError("Pass").GetFieldError("Pass1")
// 可以获取到Pass结构体字段下的Pass2错误
err.GetStructFieldError("Pass").GetFieldError("Pass2")

```

#### 复杂示例

```go
type Pass struct {
    Pass1 string `v:"required|same:Pass2"`
    Pass2 string `v:"required|same:Pass1"`
}
type User struct {
    Id     int
    Name   string `v:"required"`
    Passes []Pass
}
obj := &User{
    Name: "",
    Passes: []Pass{
        {
            Pass1: "1",
            Pass2: "2",
        },
        {
            Pass1: "3",
            Pass2: "4",
        },
    },
}
// 获取name字段的错误
err.GetFieldError("Name")
// 获取Passes切片字段的第一个的结构体的错误
index1 := err.GetSliceFieldError("Passes").GetError(0)
// 获取Passes切片字段的第二个的结构体的错误
index2 := err.GetSliceFieldError("Passes").GetError(1)
// 分别获取对应索引的结构体下字段的错误信息
index1.GetFieldError("Pass1")
index1.GetFieldError("Pass2")
// 分别获取对应索引的结构体下字段的错误信息
index2.GetFieldError("Pass1")
index2.GetFieldError("Pass2")

```

---
# i18n
````go
type Test struct {
    Args  string `v:"required" `
    Args2 int64  `v:"between:2,15" `
    Args3 int    `v:"max:60" `
}
valid := New()
valid.I18nSetLanguage("zh-CN")
obj := &TestI18nStruct{
    Args:  "",
    Args2: 60,
    Args3: 61,
}
err := valid.Struct(obj).(*ValidationError)
err.GetFieldError("Args") // Args字段不能为空
err.GetFieldError("Args2")// "Args2字段值`60`必须介于 2和15之间
err.GetFieldError("Args3")// Args3字段值`61`字段最大值应当为60
````
---

