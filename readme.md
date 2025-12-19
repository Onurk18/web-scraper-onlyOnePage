# Web Scraper - Go Projesi

Basit ve güvenilir bir web scraper uygulaması. Belirtilen web sitesinin HTML içeriğini çeker, ekran görüntüsü alır ve sayfadaki linkleri toplar.

## Gereksinimler

- Go 1.16 veya üzeri
- Google Chrome veya Chromium tarayıcı

## Kurulum

### 1. Go Modülünü Başlat

```bash
go mod init web-scraper
```

### 2. Gerekli Kütüphaneyi Yükle

```bash
go get github.com/chromedp/chromedp@v0.9.3
```

### 3. chromium yükle

```bash
sudo apt-get install -y chromium
```

## Kullanım

### Basit Kullanım

```bash
go run main.go https://www.example.com
```

### Test İçin Önerilen Siteler

```bash
# Basit site
go run main.go https://example.com

# Haber sitesi
go run main.go https://www.bbc.com

# E-ticaret
go run main.go https://www.amazon.com

# Sosyal medya
go run main.go https://www.reddit.com
```

## Çıktı Dosyaları

Program çalıştığında aşağıdaki dosyalar oluşturulur:

1. **site_data.html** - Sayfanın tam HTML içeriği
2. **screenshot.png** - Sayfanın ekran görüntüsü
3. **links.txt** - Sayfada bulunan tüm URL'ler

## Program Akışı

```
1. Komut satırından URL al
   ↓
2. Chrome tarayıcıyı başlat
   ↓
3. Sayfaya git ve yüklenmesini bekle
   ↓
4. HTML içeriğini çek → site_data.html
   ↓
5. Ekran görüntüsü al → screenshot.png
   ↓
6. Linkleri topla → links.txt
   ↓
7. Başarı mesajı göster
```

## Hata Durumları

Program aşağıdaki durumlarda hata verir:

- URL belirtilmezse
- Sayfa 30 saniyede yüklenmezse
- Bağlantı kurulamazsa
- Dosya yazma hatası olursa

## Örnek Çıktı

```
Hedef URL: https://www.example.com
✓ HTML içeriği kaydedildi: site_data.html
✓ Ekran görüntüsü kaydedildi: screenshot.png
✓ 25 adet link bulundu ve kaydedildi: links.txt

İşlem başarıyla tamamlandı!
```

## Kod Açıklaması

### Ana Bileşenler

1. **main()** - Program giriş noktası, tüm işlemleri yönetir
2. **chromedp.Run()** - Tarayıcı görevlerini sırayla çalıştırır
3. **removeDuplicates()** - Tekrar eden linkleri temizler

### Neden Chromedp?

- Gerçek tarayıcı gibi davranır
- JavaScript çalıştırabilir
- Modern sitelerde güvenilir çalışır
- Ekran görüntüsü alma özelliği var

## Sorun Giderme

**"Chrome bulunamadı" hatası:**
```bash
# Linux
sudo apt install chromium

# Mac
brew install chromium
```