
1.1作为练习函数和循环的简单途径，用牛顿法实现开方函数。
package main

import (
"fmt"
"math"
)

func Sqrt(x , z float64) float64 {
//	c := z
	var c = z
	z = z - (z*z - x)/(2*z)
	if math.Abs(c-z) > 0.00001{
		return Sqrt(x,z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2,1))
}
练习2:slices
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var x = make([]([]uint8),dy)
	for i,_ := range x{
		x[i] = make([]uint8,dx)
		for j,_ := range x[i]{
			x[i][j] =uint8 (dx+dy)/2
		}
	}
	return x
}

func main() {
	pic.Show(Pic)
}

练习3.map exercise
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _,word := range strings.Fields(s){
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
4函数闭包练习，Fibonacci数列
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum , x , y := 0 , 0 , 1
	return func() int{
		sum , x , y = x , y , x+y
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
5.Stringers练习
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string{
		return fmt.Sprintf("%v.%v.%v.%v",ip[0],ip[1],ip[2],ip[3])
}
func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
6.Error 练习
package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v.",float64(e))
}
func Sqrt(x , z float64) (float64, error) {
//	c := z
	if x >= 0{
		var c = z
		z = z - (z*z - x)/(2*z)
		if math.Abs(c-z) > 0.00001{
			return Sqrt(x,z)
		}
		return z,nil
	}else{
		return 0,ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2,1))
	fmt.Println(Sqrt(-2,1))
}
7.Reader 练习
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (i int, e error) {
	for x:=0 ; x<len(b) ; x++ {
        b[x] = 'A'
    }
    return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
8.rot13Reader 练习
package main

import (
	"io"
	"os"
	//"fmt"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (ro *rot13Reader) Read(by []byte) (n int, err error){
	n , err = ro.r.Read(by)//题目要求从io.Reader里读取数据
	//fmt.Print(by)
	for i,b := range by{
      switch {
		case 'A' <= b && b <= 'M':
			b = b + 13
		case 'M' < b && b <= 'Z':
			b = b - 13
		case 'a' <= b && b <= 'm':
			b = b + 13
		case 'm' < b && b <= 'z':
			b = b - 13
		}
	    by[i] = b
    }
	
	return n , err
    
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
9.HTTP Handlers练习
package main

import (
	"log"
	"fmt"
	"net/http"
)
type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s%s%s", s.Greeting, s.Punct, s.Who)
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", s)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
10.Images 练习
package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct{
	Width, Height int
}
    // ColorModel returns the Image's color model.
func (im *Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (im *Image)  Bounds() image.Rectangle{
	return image.Rect(0, 0, im.Width,im.Height)
}
       
func (im *Image) At(x, y int) color.Color{
	return color.RGBA{ 128+uint8(x), 128+uint8(y), 255, 255}
}
 
func main() {
	m := &Image{100,100}
	pic.ShowImage(m)
}
管道练习：
//管道在函数定义时value就可以赋进去，所以可以在go后定义的goroutine线程里完成管道的填充。
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
12. Equivalent Binary Trees 练习
package main

import(
	"golang.org/x/tour/tree"
	"fmt"
	)
// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree){
		if t == nil{
			return	
		}else{
			walk(t.Left)
			ch <- t.Value
			walk(t.Right)
		}
	}
	walk(t)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1,c1)
	go Walk(t2,c2)
	
	for {
		v1 , ok1 := <- c1
		v2 , ok2 := <- c2
		fmt.Println(v1,v2)
		if !(v1 == v2 && ok1 == ok2) {
			return false
		} else if ok1 == false && ok2 == false{
			return true
		}
	}
}

func main() {
	
	t1 := tree.New(1)
	t2 := tree.New(2)
	t3 := tree.New(1)
//New（）生成的树为二叉搜索树，这个方法只能解决二叉搜索树的等价分析问题。
	fmt.Println(Same(t1,t3))
	fmt.Println(Same(t1,t2))
}
