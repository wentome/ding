# ding
```
func main() {
	ding := NewDing("https://oapi.dingtalk.com/robot/send",
		"your access_token",
		"your secret")
	res, err := ding.SendSignMsg("ding from golang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
```
