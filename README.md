
# MSIB5 Hacktiv8 - Final Project 2: MyGram

Berikut adalah Final Project kedua dari Hacktiv8 yang disebut MyGram. Aplikasi ini memungkinkan pengguna untuk mengunggah foto serta memberikan komentar pada foto yang diunggah oleh pengguna lain. Aplikasi ini akan memiliki fitur CRUD..

## Nama
 - Alif Wildan Azzahran - GLNG-KS08-06

## User Login
```
Email : alifwildanaz@hotmail.com
Password : agent9029
```

## Endpoint
Di bawah ini merupakan semua endpoint yang dapat diakses di aplikasi ini.

### Users
 
 Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel Users:
 
 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//users/register] |
| POST | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//users/login] |
| PUT | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//users] |
| DELETE | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//users] |

###### Prosedur request users

POST Register User
 ```sh
{
    "age": integer,
    "email": "string",
    "password": "string",
    "username": "string"
}
```
#

POST Login User
 ```sh
{
    "email": "string",
    "password": "string"
}
```
#

PUT User,
Diperlukan:
- Bearer Token <br />
- Param userId
 ```sh
{
    "email": "string",
    "username": "string"
}
```
#
DELETE User
- Authorization: Bearer Token

> Note: Untuk method PUT dan DELETE diperlukan autorisasi yang memerlukan Bearer Token untuk dimasukkan terlebih dahulu. Bearer Token didapatkan melalui response pengguna saat melakukan login.
#


### Photos
  Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel Photos:

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//photos] |
| GET | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//photos] |
| PUT | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//photos/:id] |
| DELETE | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//photos/:id] |

###### Prosedur request photos

POST Photo,
Diperlukan:
- Bearer Token
 ```sh
{
    "title": "string",
    "caption": "string",
    "photo_url": "string"
}
```
#

GET Photo,
Diperlukan:
- Bearer Token

#

PUT Photo,
Diperlukan:
- Bearer Token  <br />
- Param photoId
 ```sh
{
    "title": "string",
    "caption": "string"
    "photo_url": "string"
}
```
#

DELETE Photo,
Diperlukan:
- Bearer Token  <br />
- Param PhotoId
> Note: Seluruh method diperlukan autorisasi, yang mana perlu memasukan Bearer Token terlebih dahulu. Bearer Token didapatkan melalui response pengguna saat melakukan login. Untuk method PUT dan DELETE hanya bisa dilakukan oleh pengguna yang mengunggah foto, sebagaimana hasil dari autorisasi menggunakan Bearer Token, dan juga perlu menyertakan parameter Id foto pada URL.
#


### Comments
  Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel Comments:

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//comments] |
| GET | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//comments] |
| PUT | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//comments/:id] |
| DELETE | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//comments/:id] |

###### Prosedur request comments

POST Comment,
Diperlukan:
- Bearer Token
 ```sh
{
    "message": "string",
    "photo_id": integer
}
```

#
GET Comment,
Diperlukan:
- Bearer Token

#
PUT Comment,
Diperlukan:
- Bearer Token  <br />
- Param commentId
 ```sh
{
    "message": "string"
}
```

#
DELETE Comment,
Diperlukan:
-Bearer Token  <br />
-Param commentId

> Note: Seluruh method diperlukan autorisasi, yang mana perlu memasukan Bearer Token terlebih dahulu. Bearer Token didapatkan melalui response pengguna saat melakukan login. Untuk method PUT dan DELETE hanya bisa dilakukan oleh pengguna yang mengunggah komentar, sesuai autorisasi dengan Bearer Token, dan perlu menyertakan parameter Id komentar pada URL. 
#


 ### SocialMedias
  Berikut ini merupakan endpoint-endpoint yang dapat diakses untuk tabel SocialMedias:

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//socialmedias] |
| GET | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//socialmedias] |
| PUT | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//socialmedias/:id] |
| DELETE | [https://fp2-msib5-hacktiv8-mygram.up.railway.app//socialmedias/:id] |

###### Prosedur request socialmedias

POST SocialMedia,
Diperlukan:
- Bearer Token
 ```sh
{
    "name": "string",
    "social_media_url": "string"
}
```
#
GET SocialMedia,
Diperlukan:
- Bearer Token

#
PUT SocialMedia
Diperlukan:
- Bearer Token  <br />
- Param socialMediaId
 ```sh
{
    "name": "string",
    "social_media_url": "string"
}
```
#
DELETE SocialMedia,
Diperlukan:
- Bearer Token  <br />
- Param socialMediaId


> Note: Seluruh method diperlukan autorisasi, yang mana perlu memasukan Bearer Token terlebih dahulu. Bearer Token didapatkan melalui response pengguna saat melakukan login. Untuk method PUT dan DELETE hanya bisa dilakukan oleh pengguna yang mengunggah sosial media, sesuai proses autorisasi dengan Bearer Token, pada URL dan perlu menyertakan parameter Id sosial media tersebut.
