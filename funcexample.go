package main

func testInt() (t int){
	t = 1
	return
}
func testString() (r string) {
	r = "OK"
	return
}

type Counter struct {
	n int
}

func (c *Counter) testReceiver() (n int) {
	c.n = 1
	return c.n
}

func main() {
	c := Counter{n : 1}
	testInt()
	testString()
	c.testReceiver()

}