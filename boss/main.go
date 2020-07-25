package main

import (
	"fmt"
	"github.com/fedesog/webdriver"
	"log"
	"os"
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
		//"goog:chromeOptions": map[string][]string{"args": {"--headless","--no-sandbox","--disable-gpu"}, "extensions": {}, "excludeSwitches": {"enable-automation"}},
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
	err = session.Url("https://www.zhipin.com/hangzhou/")
	if err != nil {
		log.Println(err)
	}


	el,err := session.FindElement(webdriver.XPath, `//*[@id="wrap"]//input[@name="query"]`)

	str,_ := el.GetAttribute("name")
	fmt.Println(str);
	el.SendKeys("golang")



	search,err := session.FindElement(webdriver.XPath,`//*[@id="wrap"]//form/button[@ka="search_box_index"]`)
	ka,_ := search.GetAttribute("ka")
	fmt.Println(ka);
	search.Click()
	time.Sleep(1*time.Second)
	xinzi,err := session.FindElement(webdriver.XPath,`//*[@id="filter-box"]/div/div[4]/div[3]/span/input`)
	ScreenShot(session,"1")

	time.Sleep(1*time.Second)
	err = session.Click(webdriver.LeftButton)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(1*time.Second)

	err = session.MoveTo(xinzi,5,30)
	if err != nil {
		log.Println(err)
	}
	xinzi6,err := session.FindElement(webdriver.XPath,`//*[@id="filter-box"]/div/div[4]/div[3]/span/div/ul/li[6]/a`)


	time.Sleep(1*time.Second)
	xinzi6.Click()
	ScreenShot(session,"2")

	ul,_ := session.FindElement(webdriver.XPath,`//*[@id="main"]/div/div[2]/ul`)
	lis,_ := ul.FindElements(webdriver.TagName,`li`)
	for key, value := range lis {
		//local,_ := value.GetLocation()
		//local,_ := value.GetLocationInView()
		href_urls,_ := value.FindElement(webdriver.CSS_Selector,`.primary-box .job-name > a`)

		url,_ := href_urls.GetAttribute("href")
		jobName,_ := href_urls.GetAttribute("title")

		addressS,_ := value.FindElement(webdriver.CSS_Selector,`.primary-box .job-area-wrapper > span`)
		address,_ := addressS.Text()

		companyS,_ := value.FindElement(webdriver.CSS_Selector,`.info-company h3.name > a`)
		company,_ := companyS.Text()

		salaryS,_ := value.FindElement(webdriver.CSS_Selector,`.primary-box .job-limit > span`)
		salary,_ := salaryS.Text()

		fmt.Println(key,url,jobName,address,company,salary);

	}


	time.Sleep(30 * time.Second)
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