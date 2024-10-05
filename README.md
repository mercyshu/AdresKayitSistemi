<div align="center">
  <h1>T.C. Nufus İşleri GM'ye ilişkin Adres Kayit Sistemi Otomasyonu</h1>
</div>
<p align="center">
 Adres verilerinin tek merkezden standartlara uygun bir şekilde yönetilmesi ve paylaşılması hedeflenen projedir.
</p>

Servisi hazır hale getirmek için go.dev adresinden go kurulumu yapmanız gerekmektedir. 

Kurulum tamamlandıktan sonra, dosyada bir düzenleme yapmaya ihtiyacınız yoktur.
Tüm Türkiye Mahalle, Sokak, Daire, Bina, Açık Adres bilgileri istek olarak dahil edilmiştir.

Projeyi Başltamak için:

```
go mod tidy
go run main.go
```

Nasıl Kullanırım?

```go
http.ListenAndServe(":1384", Router.AppRoutes())
```
Bu İfade bir HTTP sunucusunu başlatır ve belirtilen bağlantı noktasından gelen istekleri dinlemeye başlar. (http://localhost:1384/)


Kısaca GET isteği Örneklemeleri:

```
GetirIlListe            http://localhost:1384/AdresKayitSistemi/GetirIlListe ; Parametre: Yok
GetirIlceListe          http://localhost:1384/AdresKayitSistemi/GetirIlListe ; Parametre: IlNo (GetirIlListe)
GetirMahalleListe       http://localhost:1384/AdresKayitSistemi/GetirMahalleListe?IlceNo=0000 ; Parametre: IlceNo (GetirIlceListe)
GetirCsbmListe          http://localhost:1384/AdresKayitSistemi/GetirCsbmListe?MahalleNo=0000 ; Parametre: MahalleNo (GetirMahalleListe)
GetirBinaListe          http://localhost:1384/AdresKayitSistemi/GetirBinaListe?MahalleKayitNo=0000&CsbmNo=0000 ; Parametre: MahalleKayitNo (GetirMahalleListe), CsbmNo (GetirCsbmListe)
GetirBagimsizBolum      http://localhost:1384/AdresKayitSistemi/GetirBagimsizBolum?MahalleKayitNo=0000&BinaKimlikNo=0000 ; Parametre: MahalleKayitNo (GetirMahalleListe), BinaKimlikNo (GetirBinaListe)
AcikAdres               http://localhost:1384/AdresKayitSistemi/AcikAdres?AdresNo=0000 ; Parametre: AdresNo (GetirBagimsizBolum)
```
