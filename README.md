# 🖥️ Go Monitoring Sistemi

Bu proje, sistem kaynaklarını izleyen, belirli eşik değerleri aşıldığında log dosyasına yazan ve ilk uyarıda e-posta gönderen bir mini DevOps uygulamasıdır. Ayrıca logları gösteren basit bir web arayüzü içerir.

## 📦 Proje İçeriği

- `main.go` – CPU, RAM, Disk ve Proses sayısını izler ve loglar  
- `sendmail.go` – İlk uyarıda e-posta gönderir (kendi bilgilerinizi girmeniz gerekir)  
- `web.go` – Logları web arayüzü üzerinden gösterir  
- `logs/` – Log dosyalarının tutulduğu klasör  
- `go.mod` – Go modül dosyası  

## 🛠️ Kurulum ve Çalıştırma (Windows Kullanıcıları için)

### Go'yu kurun
Go dilini [https://go.dev/dl](https://go.dev/dl) adresinden indirip bilgisayarınıza kurun.

### Projeyi indirin
Bu repoyu ZIP olarak indirip klasöre çıkarın veya Git ile klonlayın:

```bash
git clone https://github.com/kullaniciadi/go-monitoring.git
```

### Mail adresinizi tanımlayın
`sendmail.go` dosyasını açın ve aşağıdaki alanları kendi bilgilerinizle doldurun:

```go
from := "seninmailin@gmail.com"
password := "uygulama_sifren"
to := "uyarinin_gidecegi_mail@gmail.com"
```

### Uygulamayı derleyin
Terminali (CMD veya PowerShell) açın ve proje klasörüne gidin:

```bash
cd proje-klasoru
go build .
```

### Çalıştırın
Derleme sonrası klasörde oluşan `.exe` dosyasını çift tıklayarak veya terminal üzerinden çalıştırarak uygulamayı başlatabilirsiniz.

## 🌐 Web Arayüzü (Opsiyonel)
Uygulama çalışmaya başladıktan sonra aşağıdaki adrese giderek logları görebilirsiniz:

```arduino
http://localhost:8080
```

## 📧 Not
E-posta gönderebilmek için uygulama şifresi kullanmalısınız. Gmail kullanıcıları için:

[https://myaccount.google.com/apppasswords](https://myaccount.google.com/apppasswords) adresinden uygulama şifresi oluşturabilirsiniz.

## 💬 Katkı ve Geri Bildirim
Her türlü geri bildirim ve katkıya açığım. Forklayarak kendi versiyonunuzu oluşturabilir veya pull request göndererek katkıda bulunabilirsiniz.
