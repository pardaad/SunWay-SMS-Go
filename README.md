# SunWay-Go
SunWay SMS gateway package for Golang

 وب سرویس سان وی اس ام اس
 یک وب سرویس کامل برای ارسال و دریافت پیامک برای برنامه نویسان زبان گو


نحوه استفاده
نمونه کد برای ارسال پیامک

```go
username := "username";
password := "password";
SpecialNumber := "3000...";
RecepientNumber := "09123456789";
MessageBody := "تست وب سرویس به زبان گو";
isFlash := false;
SendArray(username, password, RecepientNumber, SpecialNumber, MessageBody, IsFlashMessage)
```

### وب سرویس پیامک
متدهای وب سرویس:

#### ارسال پیام کوتاه
```go
SendArray(username, password, RecepientNumber, SpecialNumber, MessageBody, IsFlashMessage)
```

#### مشاهده موجودی پنل
```go
GetCredit(username, password)
```

#### دریافت وضعیت پیام های ارسال شده
```go
GetMessageStatus1(username, password, MessageID)
```

#### ارسال پیام زمانیندی شده
```go
SendArraySchedule(username, password, RecepientNumber, MessageBody, SpecialNumber, Year, Month, Day, Hour, Minute, IsFlashMessage)
```
