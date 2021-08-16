package main

import (
	"fmt"
)

/*
Go语言中的 new 和 make 主要区别如下：
	make 只能用来分配及初始化类型为 slice、map、Channel 的数据。new 可以分配任意类型的数据；
	new 的主要作用是为类型申请一片内存空间，并返回指向这片内存的指针。
	new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
	new 分配的空间被清零。make 分配空间后，会进行初始化；
*/

type Student struct {
	name string
	age  string
}

func main() {

	// func new(Type) *Type

	// 【示例】使用 new 函数为变量分配内存空间。
	var sum *int
	// new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。
	sum = new(int)

	*sum = 98

	fmt.Println(*sum) // 98

	// 自定义类型也可以使用 new 函数来分配空间
	var student *Student
	student = new(Student)

	student.name = "shangkun"

	fmt.Println(student) // &{shangkun }

	// func make(t Type, size ...IntegerType) Type
	// make 函数的 t 参数必须是 chan（通道）、map（字典）、slice（切片）中的一个，并且返回值也是类型本身。

}

// 内置函数 new 会在编译期的 SSA 代码生成阶段经过 callnew 函数的处理，如果请求创建的类型大小是 0，
// 那么就会返回一个表示空指针的 zerobase 变量，在遇到其他情况时会将关键字转换成 newobject：
func callnew(t *types.Type) *Node {
	if t.NotInHeap() {
		yyerror("%v is go:notinheap; heap allocation disallowed", t)
	}
	dowidth(t)

	if t.Size() == 0 {
		z := newname(Runtimepkg.Lookup("zerobase"))
		z.SetClass(PEXTERN)
		z.Type = t
		return typecheck(nod(OADDR, z, nil), ctxExpr)
	}

	fn := syslook("newobject")
	fn = substArgTypes(fn, t)
	v := mkcall1(fn, types.NewPtr(t), nil, typename(t))
	v.SetNonNil(true)
	return v
}

// 需要提到的是，哪怕当前变量是使用 var 进行初始化，在这一阶段也可能会被转换成 newobject 的函数调用并在堆上申请内存：

func walkstmt(n *Node) *Node {
	switch n.Op {
	case ODCL:
		v := n.Left
		if v.Class() == PAUTOHEAP {
			if prealloc[v] == nil {
				prealloc[v] = callnew(v.Type)
			}
			nn := nod(OAS, v.Name.Param.Heapaddr, prealloc[v])
			nn.SetColas(true)
			nn = typecheck(nn, ctxStmt)
			return walkstmt(nn)
		}
	case ONEW:
		if n.Esc == EscNone {
			r := temp(n.Type.Elem())
			r = nod(OAS, r, nil)
			r = typecheck(r, ctxStmt)
			init.Append(r)
			r = nod(OADDR, r.Left, nil)
			r = typecheck(r, ctxExpr)
			n = r
		} else {
			n = callnew(n.Type.Elem())
		}
	}
}

// 当然这也不是绝对的，如果当前声明的变量或者参数不需要在当前作用域外生存，那么其实就不会被初始化在堆上，
// 而是会初始化在当前函数的栈中并随着函数调用的结束而被销毁。
