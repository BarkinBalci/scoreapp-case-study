# Aday Görevi — User Score Service (Golang)

Bu dosya adaya verilecektir.

## 🎯 Amaç
Küçük bir backend servisi geliştirmeniz bekleniyor. Servis, bir kullanıcının aksiyonlarına göre skor hesaplar, bu skoru kaydeder ve bir HTTP endpoint üzerinden tetiklenebilir olur. Amaç, **Clean Architecture**, **bağımlılıkların arayüzlerle yönetilmesi** ve **mock kullanılarak test yazılması** konularındaki becerilerinizi ölçmektir.

---

## 📌 Gereksinimler

### Fonksiyonel Gereksinimler
- Endpoint: `POST /scores/calculate?user_id=<id>`
- `ActionService` arayüzü üzerinden kullanıcı aksiyonlarını getir.
- Aşağıdaki kurallara göre skor hesapla:
  - login → +1
  - challenge_completed → +10 × amount
  - quiz_answer → +2 × amount
- Hesaplanan skoru `ScoreRepository` arayüzü ile kaydet.
- JSON response döndür.

---

## 📦 Clean Architecture Yapısı
Aşağıdaki gibi bir yapı önerilir:

```
scoreapp/
  cmd/api/main.go
  domain/
  usecase/
  infrastructure/repository/
  interfaces/http/
```

Her katman bağımsız olmalı, iş kuralları HTTP’den etkilenmemeli.

---

## 🧪 Test Gereksinimleri
- `ScoreCalculator` için unit test yazılmalı.
- ActionService mocklanmalı.
- ScoreRepository mocklanmalı.
- Pozitif ve negatif senaryolar test edilmeli.

---

## 📁 Domain Model Örneği
```go
package domain

type UserAction struct {
    Type string
    Amount int
}

type UserScore struct {
    UserID string
    Score int
}
```

---

## 🧩 Usecase Örneği
```go
package usecase

import "scoreapp/domain"

type ActionService interface {
    GetActions(userID string) ([]domain.UserAction, error)
}

type ScoreRepository interface {
    Save(score domain.UserScore) error
}

// ...
```

---

## 🚀 Teslimat Beklentileri
- Derlenebilir bir Go projesi
- Birim testleri içeren `*_test.go` dosyaları
- Temiz mimariye uygun klasörleme
- Mock kullanılan testler

