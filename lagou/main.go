package main

import (
	"fmt"
	"github.com/fedesog/webdriver"
	"log"
	"os"
	"strings"
	"time"
)


// 获取boss 首页golang 招聘信息
func main()  {
	chromeDriver := webdriver.NewChromeDriver("../chromedriver83.0.4103.39")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{
		"Platform":           "Linux",
		"goog:chromeOptions": map[string][]string{"args": {"--headless","--no-sandbox","--disable-gpu"}, "extensions": {}, "excludeSwitches": {"enable-automation"}},
		"browserName":        "chrome",
		"version":            "",
		"platform":           "ANY",
	}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	//err = session.SetTimeouts("page load", 10000)
	err = session.SetTimeoutsImplicitWait( 10000) // 隐是等待

	if err != nil {
		log.Println(err)
	}
	//err = session.Url("https://www.helloweba.net/demo/2017/unlock/")
	err = session.Url("https://www.lagou.com/hangzhou/")
	if err != nil {
		log.Println(err)
	}


	el,err := session.FindElement(webdriver.ID, `search_input`)

	str,_ := el.GetAttribute("placeholder")
	fmt.Println(str);
	el.SendKeys("golang")



	search,err := session.FindElement(webdriver.ID,`search_button`)
	ka,_ := search.GetAttribute("value")
	fmt.Println(ka);
	search.Click()
	time.Sleep(3*time.Second)

	alert,err := session.FindElement(webdriver.XPath,`//div[text()="给也不要"]`)
	if selectBool,err := alert.IsEnabled();err == nil && selectBool {
		alert.Click()
	}
	time.Sleep(1*time.Second)
	xinzi,err := session.FindElement(webdriver.XPath,`//*[@id="order"]/li/div[1]/a[2]`)
	ScreenShot(session,"1_lagou")
	xinzi.Click()

	is := 1
	for i := 1; i <= 5 ;i++ {
		time.Sleep(1*time.Second)

		li,_ := session.FindElements(webdriver.CSS_Selector,`#s_position_list > ul > li`)

		for _, value := range li {
			company,_ := value.GetAttribute("data-company")
			job,_ := value.GetAttribute("data-positionname")
			url,_ := value.GetAttribute("href")

			moneys,_ := value.FindElement(webdriver.CSS_Selector,`span.money`)
			money,_ := moneys.Text()
			//*[@id="s_position_list"]/ul/li[1]/div[1]/div[1]/div[2]/div/text()

			years,_ := value.FindElement(webdriver.CSS_Selector,` .p_bot > div.li_b_l`)
			yearStr,_ := years.Text()

			yearArr := strings.Split(yearStr," ")
			var (
				year,education string
			)
			for k, v := range yearArr {
				if k == 1{
					year = v
				}
				if k == 3 {
					education = v
				}
				//fmt.Println(k,v);
			}
			fmt.Println(is,company,job,money,year,education,url);
			is++

		}

		time.Sleep(2*time.Second)
		pageStr := fmt.Sprintf(`//*[@id="s_position_list"]//div/span[@action="%s"]`,"next")
		//fmt.Println(pageStr);
		pageEl,_ := session.FindElement(webdriver.XPath,pageStr)
		pageEl.Click()

	}




	time.Sleep(60 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}


func ScreenShot(w *webdriver.Session,key string)  {
	bs, err := w.Screenshot()
	if err != nil {
		log.Fatal("获取二维码图片失败!")
	}
	f, err := os.OpenFile("/Users/xuzan/go/src/goBoss/web_test_driver/" + key + ".png", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	f.Write(bs)
	defer f.Close()

}