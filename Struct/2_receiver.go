package main

import (
	"fmt"
	"math"
)

/*
	方法和接收器

func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}

接收器变量：接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母，而不是 self、this 之类的命名。例如，Socket 类型的接收器变量应该命名为 s，Connector 类型的接收器变量应该命名为 c 等。
接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型。

Go 方法是作用在接收器（receiver）上的一个函数，接收器是某种类型的变量，因此方法是一种特殊类型的函数。
接收器类型可以是（几乎）任何类型，不仅仅是结构体类型，任何类型都可以有方法，
甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型
但是接收器不能是一个接口类型，因为接口是一个抽象定义，而方法却是具体实现

类型 T（或 T）上的所有方法的集合叫做类型 T（或 T）的方法集。

*/

type Point struct {
	X int
	Y int
}

type Bag struct {
	items []int
}

// Go语言为任意类型(int类型)添加方法
// 将int定义为MyInt类型
type MyInt int

// 为MyInt添加Add()方法
func (m MyInt) Add(other int) int {
	return other + int(m)
}

// 为MyInt添加IsZero()方法
func (m MyInt) IsZero() bool {
	return int(m) == 0
}

// 1) 面向过程实现方法
// 将一个物品放入背包的过程
func Insert(b *Bag, itemId int) { // 两个参数，第一个是背包指针（*Bag），第二个是物品 ID（itemid）。
	b.items = append(b.items, itemId)
}

// 2) Go语言的结构体方法
// (b*Bag) 表示接收器
// 每个方法只能有一个接收器
func (b *Bag) Insert(itemId int) {
	b.items = append(b.items, itemId)
}

func main() {

	bag := new(Bag)

	// 1) 面向过程实现方法
	// 将 *Bag 参数放在第一位，强调 Insert 会操作 *Bag 结构体，
	// 但实际使用中，并不是每个人都会习惯将操作对象放在首位，一定程度上让代码失去一些范式和描述性。
	Insert(bag, 1001)

	// 2) Go语言的结构体方法
	bag.Insert(1002)

	// 3) 理解指针类型的接收器
	// 指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的 this 或者 self。

	p := new(Property)
	p.SetValue(100)
	fmt.Println(p.Value())

	// 4) 理解非指针类型的接收器
	// 当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，
	// 在非指针接收器的方法中可以获取接收器的成员值，但修改后无效。
	// Point 属于小内存对象，在函数返回值的复制过程中可以极大地提高代码运行效率，
	// 初始化点
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	// 与另外一个点相加
	result := p1.Add(p2)
	// 输出结果
	fmt.Println(result)

	/*
		总结：
			小对象由于值复制时的速度较快，所以适合使用非指针接收器
			大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针。
	*/

	// 5) Go语言为任意类型添加方法
	var b MyInt
	fmt.Println(b.IsZero()) // true
	b = 1
	fmt.Println(b.Add(2)) // 3

	/// 示例：二维矢量模拟玩家移动
	// 实例化玩家对象，并设速度为0.5
	player := NewPlayer(0.5)
	// 让玩家移动到3,1点
	player.MoveTo(Vec2{3, 1})
	// 如果没有到达就一直循环
	for !player.IsArrived() {
		// 更新玩家位置
		player.Update()
		// 打印每次移动后的玩家位置
		fmt.Println(player.Pos())
	}

	/*
		打印：
		{0.47434163 0.15811388}
		{0.94868326 0.31622776}
		{1.4230249 0.47434163}
		{1.8973665 0.6324555}
		{2.3717082 0.7905694}
		{2.8460498 0.94868326}
	*/

}

// 定义属性结构
type Property struct {
	value int // 属性值
}

// 3) 理解指针类型的接收器
func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

// 4) 理解非指针类型的接收器
// 非指针接收器的加方法
func (p Point) Add(otherPoint Point) Point {
	// 成员值与参数相加后返回新的结构
	// 这个方法不能修改 Point 的成员 X、Y 变量，而是在计算后返回新的 Point 对象，
	return Point{p.X + otherPoint.X, p.Y + otherPoint.Y}
}

/// 示例：二维矢量模拟玩家移动

type Vec2 struct {
	X, Y float32
}

// 加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// 减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 乘
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// 距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 使用 Normalize() 方法将方向矢量变为模为 1 的单位化矢量
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec2{0, 0}
}

type Player struct {
	currPos   Vec2    // 当前位置
	targetPos Vec2    // 目标位置
	speed     float32 // 移动速度
}

// 移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 是否到达
func (p *Player) IsArrived() bool {
	// 通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经到达目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// 逻辑更新
func (p *Player) Update() {
	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		// 移动完成后，更新当前位置
		p.currPos = newPos
	}
}

// 创建新玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
