## Request API Document Postaman

## http://localhost:7000/consumers methode POST

<br/>
{
<br/>
  "nik": "1234567890123456",
  <br/>
  "full_name": "John Doe",
  <br/>
  "legal_name": "Jonathan Doe",
  <br/>
  "place_of_birth": "Jakarta",
  <br/>
  "date_of_birth":"3034-04-03",
  <br/>
  "salary": 7500000,
  <br/>
  "ktp_photo": "https://example.com/ktp.jpg",
  <br/>
  "selfie_photo": "https://example.com/selfie.jpg"
 <br/>
}


## limit methode POST
<br/>
{
<br/>
  "consumer_id": 4 ,
<br/>
  "tenor": 1,
<br/>
  "amount": 300000
<br/>
}

## Transaction methode POST
<br/>
{
<br/>
  "consumer_id": 3,
<br/>
  "contract_no": "12333456",
<br/>
  "otr": 300000000,
<br/>
  "admin_fee": 3000000,
<br/>
  "installment": 3000000,
<br/>
  "interest": 40,
<br/>
  "asset_name": "mobil"
<br/>
}


