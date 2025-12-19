package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {

	url := os.Args[1]
	fmt.Printf("Hedef URL: %s\n", url)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Hata loglarını gizle
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(func(string, ...interface{}) {}))
	defer cancel()

	// Timeout ekle (30 saniye)
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string
	var links []string

	// Chromedp görevlerini tanımla
	err := chromedp.Run(ctx,
		// Sayfayı aç
		chromedp.Navigate(url),

		// Sayfanın yüklenmesini bekle
		chromedp.Sleep(2*time.Second),

		// HTML içeriğini al
		chromedp.OuterHTML("html", &htmlContent),

		// Tüm linkleri topla
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &links),

		// Ekran görüntüsü al
		chromedp.FullScreenshot(&screenBuf, 90),
	)

	if err != nil {
		log.Fatalf("Hata oluştu: %v", err)
	}

	// HTML içeriğini dosyaya kaydet
	htmlFileName := "site_data.html"
	err = os.WriteFile(htmlFileName, []byte(htmlContent), 0644)
	if err != nil {
		log.Fatalf("HTML kaydedilemedi: %v", err)
	}
	fmt.Printf("✓ HTML içeriği kaydedildi: %s\n", htmlFileName)

	// Ekran görüntüsünü kaydet
	screenshotFileName := "screenshot.png"
	err = os.WriteFile(screenshotFileName, screenBuf, 0644)
	if err != nil {
		log.Fatalf("Ekran görüntüsü kaydedilemedi: %v", err)
	}
	fmt.Printf("✓ Ekran görüntüsü kaydedildi: %s\n", screenshotFileName)

	// Linkleri dosyaya kaydet
	linksFileName := "links.txt"
	uniqueLinks := removeDuplicates(links)
	linksContent := strings.Join(uniqueLinks, "\n")
	err = os.WriteFile(linksFileName, []byte(linksContent), 0644)
	if err != nil {
		log.Fatalf("Linkler kaydedilemedi: %v", err)
	}
	fmt.Printf("✓ %d adet link bulundu ve kaydedildi: %s\n", len(uniqueLinks), linksFileName)

	fmt.Println("\nİşlem başarıyla tamamlandı!")
}

// Tekrar eden linkleri kaldır
func removeDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, item := range slice {
		if item != "" && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

var screenBuf []byte
