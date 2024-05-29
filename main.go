package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// ChromeDP context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// zaman limiti önemli
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// screenshot değişken
	var buf []byte

	err := chromedp.Run(ctx, fullScreenshot("https://www.google.com", 90, &buf))
	if err != nil {
		fmt.Println("Hata oluştu:", err)
		return
	}

	// Ekran görüntüsünü dosyaya yaz
	if err = os.WriteFile("screenshot.png", buf, 0644); err != nil {
		fmt.Println("Hata oluştu:", err)
		return
	}

	fmt.Println("Ekran görüntüsü alındı ve screenshot.png dosyasına kaydedildi.")
}

// fullScreenshot takeoic url
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second), // yüklenmeyi bekle
		chromedp.FullScreenshot(res, quality),
	}
}
