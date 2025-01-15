<div align="center">
  <h1>Açık Kaynak Adres Kayit Sistemi</h1>
</div>
<p align="center"> 
 Nufus ve Vatandaşlık İşleri Genel Müdürlüğümüzün Herkese Açık Erişim İzni Dahilinde Otomasyon Adres Şeması Sistemi Projesidir.
</p>

Servisi hazır hale getirmek için go.dev adresinden go kurulumu yapmanız gerekmektedir. 

Kurulum tamamlandıktan sonra, dosyada bir düzenleme yapmaya ihtiyacınız yoktur.
Mahalle, Sokak, Daire, Bina, Açık Adres bilgileri istek olarak dahil edilmiştir.

YASAL UYARI: Projemde her hangi bir kişisel veri bulunmamakla beraber, paylaşılan kodlar herkese açık bir adres şema programına ilişkin otomasyon kodlarını içermektedir. Genel işleyiş amacı arz/talep ilişkisi neticesinde haritalama ve DASK sigortalarında UAVT kodu tespit için hazırlanmıştır.

Projeyi Başltamak için:

```
go mod tidy
go run main.go
```

Nasıl Kullanırım?

```go
http.ListenAndServe(":1337", Router.AppRoutes())
```
Bu İfade bir HTTP sunucusunu başlatır ve belirtilen bağlantı noktasından gelen istekleri dinlemeye başlar. (http://localhost:1337/)
localhost:1337/hizmet-dokumu adresinden sunucu endpoint istek örneklemelerine ve içeriklerine ulaşılabilmektedir. Ayrıca kısayoldan olması amacıyla aşağıda istekler bulunmaktadır.


Kısaca GET isteği Örneklemeleri:

```
GetirIlListe            http://localhost:1337/AdresKayitSistemi/GetirIlListe ; Parametre: Yok
GetirIlceListe          http://localhost:1337/AdresKayitSistemi/GetirIlListe ; Parametre: IlNo (GetirIlListe)
GetirMahalleListe       http://localhost:1337/AdresKayitSistemi/GetirMahalleListe?IlceNo=0000 ; Parametre: IlceNo (GetirIlceListe)
GetirCsbmListe          http://localhost:1337/AdresKayitSistemi/GetirCsbmListe?MahalleNo=0000 ; Parametre: MahalleNo (GetirMahalleListe)
GetirBinaListe          http://localhost:1337/AdresKayitSistemi/GetirBinaListe?MahalleKayitNo=0000&CsbmNo=0000 ; Parametre: MahalleKayitNo (GetirMahalleListe), CsbmNo (GetirCsbmListe)
GetirBagimsizBolum      http://localhost:1337/AdresKayitSistemi/GetirBagimsizBolum?MahalleKayitNo=0000&BinaKimlikNo=0000 ; Parametre: MahalleKayitNo (GetirMahalleListe), BinaKimlikNo (GetirBinaListe)
AcikAdres               http://localhost:1337/AdresKayitSistemi/AcikAdres?AdresNo=0000 ; Parametre: AdresNo (GetirBagimsizBolum)
```
