# QR Kod Oluşturma Servisi

Bu basit Go programı, belirtilen veriye dayalı olarak bir QR kodu oluşturur ve istemcilere geri döndürür.

## Nasıl Çalışır

1. **Kurulum**: Projeyi klonlayın ve Go'nun ve Gin paketinin yüklü olduğundan emin olun.
2. **Bağımlılıklar**: `go get` komutuyla `github.com/gin-gonic/gin` ve `github.com/skip2/go-qrcode` paketlerini indirin.
3. **Sunucuyu Başlatın**: `go run main.go` komutuyla sunucuyu başlatın.
4. **Test Edin**: Tarayıcıdan veya bir API test aracından `http://localhost:8080/qr?data=YOUR_DATA&size=YOUR_SIZE` adresine bir GET isteği yaparak QR kodunu oluşturun.

## API Dökümantasyonu

### `/qr` Endpoint'i

- **Method**: GET
- **Parametreler**:
  - `data`: QR kodunda kodlanacak veri.
  - `size`: QR kodunun boyutu (piksel cinsinden).
- **Cevap**: Oluşturulan QR kodunu içeren bir PNG resmi döndürür.

Örnek istek:
GET http://localhost:8080/qr?data=https://www.example.com&size=200


![Açıklama](https://github.com/omerfdev/qrAPI/blob/main/qrAPI.png)
