package main

import (
	"encoding/json"
	"fmt"
)

/*
	使用匿名结构体解析JSON数据

	JavaScript 对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。在类似的协议中，JSON 并不是唯一的一个标准协议。 XML、ASN.1 和 Google 的 Protocol Buffers 都是类似的协议，并且有各自的特色，
	基本的 JSON 类型有数字（十进制或科学记数法）、布尔值（true 或 false）、字符串，其中字符串是以双引号包含的 Unicode 字符序列，
	支持和Go语言类似的反斜杠转义特性，不过 JSON 使用的是 \Uhhhh 转义数字来表示一个 UTF-16 编码，而不是Go语言的 rune 类型。
*/

/// 定义数据结构

// 定义手机屏幕
type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}

// 定义电池
type Battery struct {
	Capacity int // 容量
}

func main() {

	/// 使用匿名结构体解析JSON数据

	// 1.生成一段json数据
	jsonData := generateJsonData()
	fmt.Println(string(jsonData)) // {"Size":5.5,"ResX":1920,"ResY":1080,"Capacity":2910,"HasTouchID":true}]

	/// 分离JSON数据

	// 2.只需要屏幕和指纹识别信息的结构和实例
	sceenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	// 反序列化到screenAndTouch
	json.Unmarshal(jsonData, &sceenAndTouch) // 通过 json.Unmarshal 反序列化 JSON 数据达成分离 JSON 数据效果
	// 输出screenAndTouch的详细结构
	fmt.Printf("%v\n", sceenAndTouch) // {{5.5 1920 1080} true}

	// 3.只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch) // {Battery:{Capacity:2910} HasTouchID:true}

}

// 生成json数据
func generateJsonData() []byte {

	// 完整数据结构
	// 定义了一个匿名结构体。这个结构体内嵌了 Screen 和 Battery 结构体，同时临时加入了 HasTouchID 字段。
	raw := &struct {
		Screen
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{
		// 为声明的匿名结构体填充屏幕数据。

		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			2910,
		},
		HasTouchID: true,
	}

	// 将数据序列化为json
	jsonData, _ := json.Marshal(raw) // 使用 json.Marshal 进行 JSON 序列化，将 raw 变量序列化为 []byte 格式的 JSON 数据。

	return jsonData
}
