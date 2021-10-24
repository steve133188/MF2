package main

import (
	"context"
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type attrs struct {
QrStr string `json:"qr_str"`
}

var htmlContent string
var a attrs
var atts  map[string]string
var obj = qrcodeTerminal.New()
var nodes = []*cdp.Node{}

func main()  {
	mac_agent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36"
	//win_agent := `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`
	opts :=append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless",false),chromedp.UserAgent(mac_agent),chromedp.Flag("hide-scrollbars", false),)


	ctx , cancel :=context.WithCancel(context.Background())
	defer cancel()

	allocCtx , cancel := chromedp.NewExecAllocator(context.Background() , opts...)
	//allocCtx , cancel := chromedp.NewExecAllocator(context.Background())
	defer cancel()
	ctx , cancel = chromedp.NewContext(allocCtx , chromedp.WithLogf(log.Printf))
	defer cancel()



	wa_url :="https://web.whatsapp.com/"

	//sel :="//*[@class='_2UwZ_']"
	err := chromedp.Run(ctx ,wa(wa_url))


	//if len(atts) !=0{
	//	//go qrterminal.Generate(atts["data-ref"], qrterminal.L, os.Stdout)
	//	func() {
	//		obj := qrcodeTerminal.New()
	//		obj.Get(atts["data-ref"]).Print()
	//	}()
	//}
	fmt.Println(atts["data-ref"])
	fmt.Println(htmlContent)
	if err != nil{
	log.Fatal(err)
	}
	defer chromedp.Cancel(ctx)
	
}

func WaitQrTask() chromedp.Tasks{
	return chromedp.Tasks{

	}
}


func saveCookies () chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {

		if err =chromedp.WaitVisible("div[class=\"_2UwZ_\"]").Do(ctx) ;err != nil {
			return
		}
		cookies , err := network.GetAllCookies().Do(ctx)
		if err != nil {
			return
		}
		cookiesData , err := network.GetAllCookiesReturns{Cookies: cookies}.MarshalJSON()
		if err != nil{
			return
		}
		if err = ioutil.WriteFile("cookies.tmp", cookiesData , 0755);err != nil{
			return
		}
		return
	}
}

func getCookies() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		if _ ,_err :=os.Stat("cookies.tmp");os.IsNotExist(_err){
			return
		}

		cookiesData ,err := ioutil.ReadFile("cookies.tmp")
		if err != nil {
			return
		}

		cookiesParams :=network.SetCookiesParams{}
		if err = cookiesParams.UnmarshalJSON(cookiesData);err != nil{
			return
		}

		return network.SetCookies(cookiesParams.Cookies).Do(ctx)
	}
}

func checkLoginStatus()chromedp.ActionFunc{
	return func(ctx context.Context) (err error) {
		var url string
		if err = chromedp.Evaluate(`windocw.location.href` , &url).Do(ctx);err !=nil{
			return
		}
		if chromedp.Attributes("div[class=\"_2UwZ_\"]",&atts,chromedp.ByQuery) ==nil{
			log.Println("Log in  already")
			chromedp.Stop()
		}
		return
	}
}

func wa (url string)chromedp.Tasks{
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("div[class=\"_2UwZ_\"]"),
		chromedp.Sleep(time.Second),
		chromedp.Attributes("div[class=\"_2UwZ_\"]",&atts,chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context ) error {
			n:=  atts["data-ref"]
			obj.Get(n).Print()

			for i := 1;  i<=30 ; i++ {
					//chromedp.EvaluateAsDevTools(`document.querySelector("header[class='_1G3Wr']")`, &res).Do(ctx)
					//if res["className"] != ""{
					//	return nil
					//}
					err :=chromedp.Attributes("div[class=\"_2UwZ_\"]",&atts,chromedp.ByQuery).Do(ctx)
					if err != nil {
						return err
					}
					fmt.Println(atts["data-ref"])
					fmt.Println(len(atts))
					//if atts["data-ref"] == ""{
					//	return nil
					//}
					if len(atts["data-ref"]) < 2{
						fmt.Println("login")
						break
					}
					time.Sleep(time.Second * 1)
			}
			return nil
		}),
		chromedp.WaitVisible("#side > header" , chromedp.ByQuery),
		//chromedp.OuterHTML("body" , &htmlContent ,chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			//ticker := time.NewTicker(1 * time.Second)
			//quit := make(chan struct{})
			//for {
			//	select {
			//	case <- ticker.C:
			//		chromedp.Nodes("#pane-side > div:nth-child(3) > div > div > div:nth-child(7)",&nodes , chromedp.ByQuery).Do(ctx)
			//		fmt.Println(nodes)
			//	case <- quit:
			//		ticker.Stop()
			//		return nil
			//	}
			//}

			err := chromedp.Nodes("#pane-side > div:nth-child(3) > div > div > div:nth-child(7)",&nodes , chromedp.ByQuery).Do(ctx)
			if err !=nil{
				return err
			}
			fmt.Println(nodes)
			if len(nodes) == 0{
				fmt.Println("err")
			}

			return nil
		}),
		//chromedp.Stop(),
		//chromedp.Sleep(20 *time.Second),

	}
}