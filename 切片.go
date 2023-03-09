package main

import "fmt"

/*方式一：通过数组创建切片 array[startIndex:endIndex]*/
func method1() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	//数组从下标0开始取，一直取到下标3前面的索引
	var s = arr[2:4]
	fmt.Println(s) //[3 4]
	//切片的len = slice结束index - slice开始index
	fmt.Println(len(s)) //2
	//切片的cap = arr总元素数量-slice开始index
	fmt.Println(cap(s)) //4
	// 数组地址就是数组首元素的地址
	fmt.Printf("%p\n", &arr)    //0xc00000c540
	fmt.Printf("%p\n", &arr[0]) //0xc00000c540
	//切片地址就是数组种指定开始元素的地址，
	//arr[0:3]的开始地址是0，所以也是arr[0]的地址
	fmt.Printf("%p\n", s) //0xc00000c540
}

/*方式二：通过make函数创建 make(类型，长度，容量)*/
func method2() {
	//第一个参数：指定切片数据类型
	//第二个参数：指定切片的长度
	//第三个参数：指定切片的容量
	var s = make([]int, 3, 5)
	fmt.Println(s)      //[0 0 0]
	fmt.Println(len(s)) //3
	fmt.Println(cap(s)) //5
}

/*方式三：通过go提高的语法糖快速创建*/
//和创建数组一样，但是不指定长度
//长度和容量相等
func method3() {
	var s = []int{1, 2, 3}
	fmt.Println(s)      //[1 2 3]
	fmt.Println(len(s)) //3
	fmt.Println(cap(s)) //3
}

/*指定位置*/
func pointIndex() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	//同时指定开始位置与结束位置
	fmt.Println(arr[0:2]) //[1 2]

	//只指定结束位置
	fmt.Println(arr[:3]) //[1 2 3]

	//只指定开始位置
	fmt.Println(arr[1:]) //[2 3 4 5 6]

	//都不确定
	fmt.Println(arr[:]) //[1 2 3 4 5 6]
}

//可以通过切片名称[索引]方式操作切片
func use1() {
	var s = []int{1, 2, 6}
	s[1] = 666
	fmt.Println(s) // [1 666 6]
}

//如果通过切片名称[索引]方式操作切片, 不能越界
func use2() {
	var s = []int{1, 2, 6}
	// panic: runtime error: index out of range [3] with length 3
	s[3] = 666
}

//append(slice, value...)
func use3() {
	var s = []int{1, 2, 6}
	fmt.Println("追加数据前:", s)      // [1 2 6]
	fmt.Println("追加数据前:", len(s)) // 3
	fmt.Println("追加数据前:", cap(s)) // 3

	s = append(s, 666)
	//append函数会在切片末尾添加一个元素, 并返回一个追加数据之后的切片
	fmt.Println("追加数据后:", s)      // [1 2 6 666]
	fmt.Println("追加数据后:", len(s)) // 4
	//append函数每次给切片扩容都会按照原有切片容量*2的方式扩容
	fmt.Println("追加数据后:", cap(s)) // 6
}

//利用append函数追加数据时,如果追加之后没有超出切片的容量,那么返回原来的切片,
//如果追加之后超出了切片的容量,那么返回一个新的切片
func use4() {
	var test1 = make([]int, 3, 5)
	fmt.Printf("追加前地址: %p\n", test1) //0xc00000c540
	t := append(test1, 666)
	fmt.Println(t)
	fmt.Printf("追加后地址: %p\n", t) //0xc00000c540

	var test2 = make([]int, 3, 3)
	fmt.Printf("追加前地址: %p\n", test2) //0xc000018138
	t0 := append(test2, 666)
	fmt.Println(t0)
	fmt.Printf("追加后地址: %p\n", t0) //0xc00000c5a0
}

//两个切片的指向修改，造成引用传递
func use5() {
	var s1 = []int{1, 3, 5}
	var s2 = make([]int, 5)
	fmt.Printf("赋值前:%p\n", s1) // 0xc0420600a0
	fmt.Printf("赋值前:%p\n", s2) // 0xc042076060
	// 将s2的指向修改为s1, 此时s1和s2底层指向同一个数组
	s2 = s1
	fmt.Printf("赋值后:%p\n", s1) // 0xc0420600a0
	fmt.Printf("赋值后:%p\n", s2) // 0xc0420600a0
	fmt.Println(s1)            // [1 3 5]
	fmt.Println(s2)            // [1 3 5]
	s2[1] = 666
	fmt.Println(s1) // [1 666 5]
	fmt.Println(s2) // [1 666 5]
}

//copy(目标切片, 源切片), 会将源切片中数据拷贝到目标切片中，只造成值传递
func use6() {
	var s1 = []int{1, 3, 5}
	var s2 = make([]int, 5)
	fmt.Printf("赋值前:%p\n", s1) // 0xc0420600a0
	fmt.Printf("赋值前:%p\n", s2) // 0xc042076060
	// 将s1中的数据拷贝到s2中,, 此时s1和s2底层指向不同数组
	copy(s2, s1)
	fmt.Printf("赋值后:%p\n", s1) // 0xc0420600a0
	fmt.Printf("赋值后:%p\n", s2) // 0xc042076060
	fmt.Println(s1)            // [1 3 5]
	fmt.Println(s2)            // [1 3 5 0 0]
	s2[1] = 666
	fmt.Println(s1) // [1 3 5]
	fmt.Println(s2) // [1 666 5 0 0]
}

//copy只能把小容器装入到大容器中
func use7() {
	// 容量为3
	var s1 = []int{1, 10, 5}
	// 容量为5
	var s2 = make([]int, 5)
	fmt.Println("拷贝前:", s2) // [0 0 0 0 0]
	// s2容量足够, 会将s1所有内容拷贝到s2
	copy(s2, s1)
	fmt.Println("拷贝后:", s2) // [1 10 5 0 0]
}

//copy容量不够，就只能放多少个就放多少个
func use8() {
	// 容量为3
	var s1 = []int{1, 10, 5}
	// 容量为2
	var s2 = make([]int, 2)
	fmt.Println("拷贝前:", s2) // [0 0]
	// s2容量不够, 会将s1前2个元素拷贝到s2中
	copy(s2, s1)
	fmt.Println("拷贝后:", s2) // [1 10]
}

//可以通过切片再次生成新的切片, 两个切片底层指向同一数组
func action1() {
	arr := [5]int{1, 3, 5, 7, 9}
	s1 := arr[0:4]
	s2 := s1[0:3]
	fmt.Println(s1) // [1 3 5 7]
	fmt.Println(s2) // [1 3 5]
	// 由于底层指向同一数组, 所以修改s2会影响s1
	s2[1] = 666
	fmt.Println(s1) // [1 666 5 7]
	fmt.Println(s2) // [1 666 5]
}

//和数组不同, 切片只支持判断是否为nil, 不支持==、!=判断
func action2() {
	var arr1 [3]int = [3]int{1, 3, 5}
	var arr2 [3]int = [3]int{1, 3, 5}
	var arr3 [3]int = [3]int{2, 4, 6}
	// 首先会判断`数据类型`是否相同,如果相同会依次取出数组中`对应索引的元素`进行比较,
	// 如果所有元素都相同返回true,否则返回false
	fmt.Println(arr1 == arr2) // true
	fmt.Println(arr1 == arr3) // false

	s1 := []int{1, 3, 5}
	s2 := []int{1, 3, 5}
	//fmt.Println(s1 == s2) // 编译报错
	fmt.Println(s1 != nil) // true
	fmt.Println(s2 == nil) // false
}

//不同于数组，只声明但没有被创建的切片是不能使用的
func action3() {
	// 数组声明后就可以直接使用, 声明时就会开辟存储空间
	var arr [3]int
	arr[0] = 2
	arr[1] = 4
	arr[2] = 6
	fmt.Println(arr) // [2 4 6]

	// 切片声明后不能直接使用, 只有通过make或语法糖创建之后才会开辟空间,才能使用
	var s []int
	s[0] = 2 // 编译报错
	s[1] = 4
	s[2] = 6
	fmt.Println(s)
}

// 字符串的底层是[]byte数组, 所以字符也支持切片相关操作
func action4() {
	str := "abcdefg"
	// 通过字符串生成切片
	s1 := str[3:]
	fmt.Println(s1) // defg

	s2 := make([]byte, 10)
	// 将字符串拷贝到切片中
	copy(s2, str)
	//转换成ASCII码了
	fmt.Println(s2) //[97 98 99 100 101 102 103 0 0 0]
}

/*三种创建方式*/
func method() {
	method1()
	method2()
	method3()
}

/*其他使用*/
func use() {
	use1()
	use2()
	use3()
	use4()
	use5()
	use6()
	use7()
	use8()
}

/*注意点*/
func action() {
	action1()
	action2()
	action3()
	action4()
}

func main() {
	// method()
	// pointIndex()
	// use()
	// action()
}
