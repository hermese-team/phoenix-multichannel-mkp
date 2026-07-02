# Lazada Open Platform API Reference

> Auto-generated from https://open.lazada.com/apps/doc/api

## Table of Contents

- [System API](#system-api) (4 APIs)
- [Seller API](#seller-api) (17 APIs)
- [Product API](#product-api) (28 APIs)
- [Cross Boarder Product API](#cross-boarder-product-api) (11 APIs)
- [Product Review API](#product-review-api) (3 APIs)
- [Store Decoration API](#store-decoration-api) (1 APIs)
- [Media Center API](#media-center-api) (6 APIs)
- [Flexicombo API](#flexicombo-api) (9 APIs)
- [Seller Voucher API](#seller-voucher-api) (9 APIs)
- [Free Shipping API](#free-shipping-api) (11 APIs)
- [Early Bird Price API](#early-bird-price-api) (4 APIs)
- [Order API](#order-api) (8 APIs)
- [Return and Refund API](#return-and-refund-api) (8 APIs)
- [Fulfillment API](#fulfillment-api) (10 APIs)
- [Logistics API](#logistics-api) (9 APIs)
- [FirstMile Bigbag(only for CN)](#firstmile-bigbagonly-for-cn) (9 APIs)
- [Finance API](#finance-api) (4 APIs)
- [Membership API](#membership-api) (10 APIs)
- [FBL API](#fbl-api) (51 APIs)
- [Instant Messaging API](#instant-messaging-api) (7 APIs)
- [Lazada Logistics API](#lazada-logistics-api) (29 APIs)
- [E-Tickets API](#e-tickets-api) (8 APIs)
- [LazPay API](#lazpay-api) (25 APIs)
- [Lazada Wallet Corporate Top-up API](#lazada-wallet-corporate-top-up-api) (5 APIs)
- [RedMart API](#redmart-api) (8 APIs)
- [Lazada DG API](#lazada-dg-api) (7 APIs)
- [Sponsored Solutions API](#sponsored-solutions-api) (28 APIs)
- [Service Market API](#service-market-api) (2 APIs)
- [Choice Customized API](#choice-customized-api) (12 APIs)
- [LazLike API](#lazlike-api) (13 APIs)
- [LazLive API](#lazlive-api) (1 APIs)
- [Logistics Station API](#logistics-station-api) (18 APIs)
- [LazCredit Risk API](#lazcredit-risk-api) (0 APIs)
- [Content API](#content-api) (7 APIs)
- [Store Flash Sale API](#store-flash-sale-api) (0 APIs)

---
## System API

### GenerateAccessToken
`GET/POST` `/auth/token/create`

**Description:** generate access_token for call api

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `code` | String | Yes | oauth code, get from app callback URL |
| `uuid` | String | No | This field is currently invalid,  do not use this field please |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `expires_in` | Number | Yes | The expiring time of the access token, in seconds |
| `account_id` | String | Yes | Account ID，Allow null. if(account_platform=seller_center) account_id=null |
| `country` | String | Yes | The country ID (sg:Singapore, my:Malaysia, ph:Philippines, th:Thailand, id:Indonesia, vn:Vietnam) |
| `country_user_info` | Object[] | Yes |  Country user details |
| `account_platform` | String | Yes | Account platform |
| `access_token` | String | Yes | Access token |
| `account` | String | Yes | User account(login user) |
| `refresh_expires_in` | String | Yes | The expiring time of th refresh token |
| `refresh_token` | String | Yes | Refresh token, used to refresh the token when “refresh_expires_in”>0. |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `MissingParameter` | the input parameter “sign” that is mandatory for processing this request is not supplied | 1 |
| `IncompleteSignature` | The request signature does not conform to lazop standards | 1 |
| `InvalidCode` |  Invalid authorization code | Possible causes, incorrect authorisation url; authorisation code more than half an hour old |
| `InvalidCode` | Invalid authorization code | 1、please check if your Code is from the callback URL;2、Please check if your Code has already been used, each Code can on |


### GenerateAccessTokenWithOpenId
`GET/POST` `/auth/token/createWithOpenId`

**Description:** generate access_token with openId for call api

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `code` | String | Yes | oauth code, get from app callback URL |
| `uuid` | String | No | This field is currently invalid,  do not use this field please |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `expires_in` | Number | Yes | The expiring time of the access token, in seconds |
| `account_id` | String | Yes | Account ID，Allow null. if(account_platform=seller_center) account_id=null |
| `country` | String | Yes | The country ID (sg:Singapore, my:Malaysia, ph:Philippines, th:Thailand, id:Indonesia, vn:Vietnam) |
| `country_user_info` | Object[] | Yes |  Country user details |
| `account_platform` | String | Yes | Account platform |
| `access_token` | String | Yes | Access token |
| `account` | String | Yes | User account(login user) |
| `refresh_expires_in` | String | Yes | The expiring time of th refresh token |
| `refresh_token` | String | Yes | Refresh token, used to refresh the token when “refresh_expires_in”>0. |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `MissingParameter` | the input parameter “sign” that is mandatory for processing this request is not supplied | 1 |
| `IncompleteSignature` | The request signature does not conform to lazop standards | 1 |
| `InvalidCode` |  Invalid authorization code | Possible causes, incorrect authorisation url; authorisation code more than half an hour old |


### RefreshAccessToken
`GET/POST` `/auth/token/refresh`

**Description:** refresh access_token, the endpoint is https://auth.lazada.com/rest

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `refresh_token` | String | Yes | refresh_token |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `expires_in` | Number | Yes | The expiring time of the access token, in seconds |
| `account_id` | String | Yes | Account ID，Allow null. if(account_platform=seller_center) account_id=null |
| `country` | String | Yes | The country ID (sg:Singapore, my:Malaysia, ph:Philippines, th:Thailand, id:Indonesia, vn:Vietnam) |
| `country_user_info_list` | Object[] | Yes | Country user details |
| `account_platform` | String | Yes | Account platform |
| `access_token` | String | Yes | Access token |
| `account` | String | Yes | User account(login user) |
| `refresh_expires_in` | Number | Yes | The expiring time of th refresh token |
| `refresh_token` | String | Yes | Refresh token, used to refresh the token when “refresh_expires_in”>0. |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalRefreshToken` | "The specified refresh token is invalid or expired" | "The specified refresh token is invalid or expired" |
| `AUTH_TYPE_UNSUPPORTED` | XXX can only be authorized by market, not support refresh | The APP has been uploaded to the service market and the validity period of the access token has been bound to the servic |
| `IllegalRefreshToken` | The specified refresh token is invalid or expired | Please have your seller re-login for authorization. |
| `AUTH_TYPE_UNSUPPORTED` | XXX can only be authorized by market, not support refresh | The APP has been uploaded to the service market and the validity period of the access token has been bound to the servic |
| `IllegalRefreshToken` | The specified refresh token is invalid or expired | Please have your seller re-login for authorization. |
| `AUTH_TYPE_UNSUPPORTED` | XXX can only be authorized by market, not support refresh | The APP has been uploaded to the service market and the validity period of the access token has been bound to the servic |
| `IllegalRefreshToken` | The specified refresh token is invalid or expired | Please have your seller re-login for authorization. |
| `IllegalRefreshToken` | The specified refresh token is invalid or expired | Please have your seller re-login for authorization. |
| `AUTH_TYPE_UNSUPPORTED` | appkey can only be authorized by market, not support refresh | The APP has been uploaded to the service market and the validity period of the access token has been bound to the servic |


### startExportByDataset
`GET/POST` `/fbi/download/startExportByDataset`

**Description:** Open the download operation

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `oldSystemId` | String | No | 222 |
| `useNewEngine` | String | No | true |
| `appName` | String | Yes | 1 |
| `secret` | String | Yes | 1 |
| `workId` | String | Yes | 1 |
| `datasetId` | String | Yes | 1 |
| `fileType` | String | Yes | 1 |
| `uploadType` | String | Yes | 1 |
| `dispatchUserInfo` | String[] | No | 1 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 1 |


---
## Seller API

_Seller Profile_

### BatchQueryFollowStatus
`GET/POST` `/shop/follow/status/batch/query`

**Description:** Query whether these customers follow this seller.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `buyer_ids` | String[] | Yes | buyerId array |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | Rensponse WrapperClass |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### GetPickUpStoreList
`GET/POST` `/rc/store/list/get`

**Description:** return the list of pick up store infomation for requested Seller

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | String | Yes | result |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### GetSeller
`GET` `/seller/get`

**Description:** Get seller information by current seller ID.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### GetSellerMetricsById
`GET` `/seller/metrics/get`

**Description:** Provide seller metrics data of the specific seller, like positive seller rating, ship on time rate and etc.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### GetSellerPerformance
`GET/POST` `/seller/performance/get`

**Description:** Provide the performance metrics of the current seller, such as positive seller rating, ship on time, etc.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `language` | String | No | Optional ISO 639-1 standard language code (default: en-US, supported languages: en-US, zh-CN, ms-MY, th-TH, vi-VN, id-ID). |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response payload. |
| `success` | Boolean | Yes | true for success, false for error. |
| `error_code` | String | Yes | Error code if success = false. |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### GetWarehouseBySellerId
`GET/POST` `/rc/warehouse/get`

**Description:** get warehouse by seller id

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | xxxx |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### QueryWarehouseDetailInfoBySellerId
`GET/POST` `/rc/warehouse/detail/get`

**Description:** query warehouse detail info by seller id

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | xxx |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### SellerCenterMsgList
`GET/POST` `/sellercenter/msg/list`

**Description:** seller center msg box

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `language` | String | No | Set the language for returned messages.(en/vn/id/sg/ph...) |
| `page` | String | No | Paged query. |
| `pageSize` | String | No | Paged query, with a maximum return of one hundred records. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### SellerPolicyFetch
`GET` `/seller/policy/fetch`

**Description:** Fetch seller policy information

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `locale` | String | Yes | locale |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | String | No | result status |
| `data` | String | No | data obj |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `UnknownRuntimeException` | The request has failed due to RPC runtime failure | Incorrect input venture |
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### SynchronizeSellerItemArConfig
`GET/POST` `/seller/ar/config/syn`

**Description:** synchronize seller item ar config

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `siteId` | String | Yes | site Id |
| `source` | String | Yes | ar config isv |
| `uid` | String | Yes | uid |
| `contents` | String | Yes | syn sku ar config info |
| `synDate` | String | Yes | synDate |
| `business` | String | No | business |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | success |
| `errorCode` | String | Yes | errorCode |
| `model` | Object | Yes | syn result |
| `errorMsg` | String | Yes | errorMsg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### getCountryInfo
`GET/POST` `/seller/cb/country/get`

**Description:** getCountryInfo

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `type` | String | Yes | scence description |
| `seller_country` | String | No | seller register country |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | returned data |
| `success` | String | No | if success |


### getSellerRegisterInfo 
`GET/POST` `/seller/cb/register/info`

**Description:** getSellerRegisterInfo 

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Object[] | Yes | * |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | * |
| `success` | String | No | * |


### getSubAddress
`GET/POST` `/seller/cb/country/location/get`

**Description:** get location info

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `location_id` | String | Yes | * |
| `level` | Number | Yes | * |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | * |
| `success` | Boolean | No | if success |


### paymentBinding
`GET/POST` `/seller/cb/payment/config`

**Description:** paymentBinding

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | String | Yes | I * |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | * |
| `success` | Boolean | No | * |


### queryBuyboxHuntingInfo
`GET/POST` `/hunting/buybox/get`

**Description:** SPU竞价接口

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `HuntingQueryParam` | Object | Yes | param |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### saveSellerWarehouseInfo
`GET/POST` `/rc/sellerWarehouse/saveWarehouseInfo`

**Description:** Api to create or edit the seller warehouse info except the "default"
dropshipping warehouse and the return warehouse.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `ownerType` | Number | Yes | the fixed value is 0 |
| `sellerId` | Number | Yes | seller id |
| `warehouseOwnerType` | String | Yes | the fixed value is SELLER |
| `warehouseContactDTO` | Object | Yes | address info |
| `siteId` | String | Yes | site id |
| `warehouseAddressInfoDTO` | Object | Yes | address info |
| `warehouseType` | Number | Yes | the fixed value is 200 |
| `ownerId` | Number | Yes | seller id |
| `warehouseName` | String | Yes | warehouse name |
| `currencyCode` | String | Yes | currency code |
| `resourceType` | Number | Yes | resourceType - the fixed value is 1. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### sellerFieldVerify
`GET/POST` `/seller/cb/register/fieldcheck`

**Description:** verify seller info field 

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Object[] | Yes | * |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | * |
| `success` | String | No | * |


---
## Product API

_Product Management_

### AdjustSellableQuantity
`POST` `/product/stock/sellable/adjust`

**Description:** Use this API to increase or decrease sellable quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | Please take demo as reference. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format | The request format is not valid. |
| `6` | E006: Unexpected internal error | Unexpected internal error. |
| `30` | E030: Empty Request | The request URL is not complete. |
| `204` | E204: Too many SKU in one request | The number of SKUs exceeds the limit. |
| `501` | E501: Update product failed | Failed to update the product price or stock. |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `1000` | Internal Application Error | Internal system error. |
| `212` | Sellable inventory cannot be negative | Please call GetProduct/GetProductItem API to check the current sellable inventory of SKU, the quantity of reduced sellab |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `4170` | During the Bday Mega campaign, there are restrictions for stock adjustments in effect between YYYY-MM-DD HH:MM:SS - YYYY-MM-DD HH:MM:SS. Sellers can increase stocks, but may not decrease stocks. | This SKU is participating in a special Campaign, so this SKU can't be updated to set stock less than current stock. |
| `212` | INV_NEGATIVE_SELLABLE | The updated Sellable Inventory Quantity cannot be negative, and the quantity of inventory reduced in the request cannot  |


### BatchUpdateSizeChart
`POST` `/size/chart/batch/update`

**Description:** 批量更新尺码表

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | product size chart  |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `4174` | E4174 | The size template corresponding to this product does not exist |
| `4175` | E4175 | The size chart image url incorrect |
| `4177` | E4177 | Empty Product Id or Size Chart |
| `4178` | E4178 | Invalid size chart format, Size Chart format must image url or template id |
| `4179` | E4179 | Cannot exceed the maximum size chart，maximum is 50 |
| `4180` | E4180 | The product category not support size chart |
| `4181` | E4181 | Update size chart all failed |
| `4182` | E4182 | only local seller and IntraAsean seller can set size chart |
| `4183` | E4183 | Update size chart part failed |
| `4185` | E4185 | The third-party ic service invocation is error |
| `4187` | E4187 | The size chart value is invalid,Please input correct template id or url |
| `4189` | E4189 | One product only can set one size chart |


### CreateProduct
`POST` `/product/create`

**Description:** Use this API to create a single new product.

Find more details below: https://open.lazada.com/apps/doc/doc?nodeId=30720&docId=120949

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | <a href='https://open.lazada.com/apps/doc/doc?nodeId=30720&docId=120949' target='_brank'>Parameter description</a> |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `1` | E001: Parameter %s is mandatory | The parameter is mandatory but not specified. |
| `5` | E005: Invalid Request Format | The request format is not valid. |
| `6` | E006: Unexpected internal error | Unexpected internal error. |
| `30` | E030: Empty Request | The request URL is not complete. |
| `201` | E201: %s Invalid CategoryId | The specified category ID is not valid. |
| `202` | E202: %s Invalid SPUId | The specified SPU ID is not valid. |
| `205` | E205: SPU does not exist | The specified SPU ID does not exist. |
| `206` | E206: Different category id in SPU and PrimaryCategory | The specified category ID is not consistent. |
| `500` | E500: Create product failed | Failed to create the product. |
| `502` | E502: Search SPU failed | Failed to search for the specified SPU. |
| `512` | E512: BIZ_CHECK_MANGROVE_RULE_QC | The request failed because the category was banned |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `1000` | Internal Application Error |  Internal system error. |
| `4104` | BIZ_CHECK_PRICE_PRECISION_INVALID | Price accuracy check failed |
| `4105` | BIZ_CHECK_SELLER_SKU_DUPLICATE | SellerSku repeat |
| `4106` | CHK_CATPROP_CPV_INPUT_SIZE_LIMIT | Item customization attributes exceeded the limit |
| `4107` | CHECK_CAT_PROP_INVALID_NUMBER | The category attribute value is invalid |
| `4108` | CHK_BASIC_REQUIRED | Basic attributes Mandatory verification |
| `4109` | CHK_SKU_PROPS_NOT_MATCH_SALE_PROP | Sku sales attributes do not match |
| `4110` | BIZ_CHECK_CAT_PROP_MANDATORY | Category attribute This parameter is mandatory |
| `4111` | CHK_CATPROP_CPV_TEXT_REPEAT | Category attribute content repeats |
| `4112` | CHK_SKU_PROPS_DUPLICATE | Duplicate Sku attributes |
| `4113` | CHK_SKU_PROPS_NOT_IDENTICAL | Sales attribute is not filled in |
| `4114` | BIZ_CHECK_PRICE_SAMPLE_NON_ZERO | The sample price is 0 |
| `4115` | CHK_CATPROP_CPV_NOT_ENUM | The CPV attribute is not one of the options provided by the category |
| `4116` | BIZ_CHECK_MAIN_IMAGE_DUPLICATE | Repeat check of master diagram |
| `4117` | BIZ_CHECK_SPECIAL_PRICE_FROM_DATE_AFTER_TO_DATE | Special offer date check |
| `4118` | BIZ_CHECK_PRICE_IS_ZERO | Price is not 0 check |
| `4119` | BIZ_CHECK_SPECIAL_PRICE_RATE_OUT_OF_RANGE | Special price range check |
| `4120` | CHK_CATPROP_CPV_MAX_LEGNTH | Verify the maximum CPV value of a category |
| `4121` | BIZ_CHECK_SPECIAL_PRICE_PRECISION_INVALID | Special accuracy check does not pass |
| `4122` | BIZ_CHECK_VIRTUAL_BUNDLE_SKU_SUB_OVER_LIMIT | virtual bundle sku relation skuc over limit |
| `4123` | BIZ_CHECK_MANGROVE_RULE | Restricted publication check |
| `4124` | BIZ_CHECK_MANGROVE_RULE_QC | MANGROVE rule verification |
| `4125` | THD_IC_F_IC_DOMAIN_PROPERTY_002 | IC Verification category Attribute This parameter is mandatory |
| `4126` | THD_IC_F_IC_INFRA_PRODUCT_036 | SellerSku repeat |
| `4127` | THD_IC_F_IC_SCENE_PUBLISH_012 | ProductId repeat |
| `4128` | THD_IC_F_IC_DOMAIN_ACTOR_006 | Seller lock cannot be edited |
| `4129` | BIZ_CHECK_PROP_SPECIAL_CHAR | Special characters are not allowed:   "^~<>/ |
| `4130` | BIZ_CHECK_OFFICIAL_STORE_BRAND_UNAUTHORIZED | Uncertified brand |
| `4131` | BIZ_CHECK_CAT_PROP_SENSITIVE_WORDS | description has sensitive words New brand |
| `4132` | Invalid Request Format | Invalid Request Format |
| `4133` | Invalid variation | Invalid variation |
| `4134` | Please select the last level category. | Please select the last level category. |
| `4135` | THD_IC_ERR | IC service error |
| `4136` | SELLER_SKU_NOT_FOUND | seller sku not found |
| `4137` | ITEM_NOT_FOUND | IC commodity query less than |
| `4138` | BIZ_CHECK_EXIST_OUTER_IMAGE | The image exists in the outer link |
| `4139` | BIZ_CHECK_MAIN_IMAGE_REQUIRE | Main image is require |
| `4140` | CHK_ENUM_PROP_VALUE_NOT_IN_OPTION | Class does not have this attribute |
| `4141` | THD_IC_ERR_F_IC_INFRA_PRODUCT_036 | SellerSku repeat |
| `4142` | THD_BRAND_ID_IS_NOT_VALID_IN_CATEGORY | The brand is invalid in the category |
| `4143` | BIZ_CHECK_SALEPROP_ATTRIBUTE_INVALID | The selling attributes are not defined in the variation |
| `4144` | BIZ_CHECK_SKU_NOT_CONTAIN_SALEPROP | The sku does not contain the saleProp tag |
| `4145` | BIZ_CHECK_SALEPROP_AND_OLD_PARAM_REPEAT | You can't put sales properties in both saleProp and sku |
| `4146` | BIZ_CHECK_SALEPROP_NOT_SUPPORT_THUMBNAIL | Thumbnails are not supported for this sale attribute |
| `4147` | THD_IC_ERR_F_IC_SERVICE_EDIT_002 | Concurrent product edit is not allowed |
| `4148` | BIZ_CHECK_ITEM_HAS_REACH_LIMIT | Seller's online product count has reach limit |
| `4149` | BIZ_CHECK_PACKAGE_DECIMAL_INVALID | Package attribute value is not valid |
| `4150` | SELLER_SKU_INVALID | seller sku is invalid |
| `4151` | BIZ_CHECK_MTEE_RULE_QC | Quality Check, Failed by Lazada Policy |
| `4152` | THD_INVENTORY_ERR_INV_PARAM_ILLEGAL | illegal parameter |
| `4153` | THD_IC_ERR_FC_IC_SKU_IMAGE_001 | If you upload a sku picture then all sku must be uploaded |
| `4154` | SYS_REQUEST_TOO_FAST | Slow down a bit! Too many opertions at once. Please try again later |
| `4155` | BIZ_CHECK_NO_EDIT_ITEM_LOCK | This product is currently locked so you are unable to do editing. Please go to Data Inisight > Policy Compliance to unlo |
| `4156` | C035: No brand cannot be selected | Can't set Brand name as No brand, Mandatory select Brand due to LazMall identity |
| `4157` | BIZ_CHECK_SPECIAL_PRICE_GREATER_THAN_PRICE | The discount price must be cheaper than the regular price |
| `4158` | THD_IC_ERR_F_DOMAIN_IMAGE_00_01_003 | Main image has duplicate. Please remove to continue |
| `4159` | IC_EXCEPTION | IC service is exception |
| `4160` | THD_IC_ERR_F_PRODUCT_00_15_004 | The max length of Multi title for locale  can not exceed 255 bytes |
| `4161` | VARIATION_CATEGORY_ATTRIBUTE_INVALID | This  variaition attribute  is not found in the category attribute library |
| `4162` | THD_IC_ERR_F_IC_ABILITY_PG | Product edit was blocked by tag |
| `4163` | BIZ_CHECK_RESTRICTED_CATEGORY | You are not authorised to sell this category |
| `4164` | BIZ_CHECK_MAX_PACKAGE_WEIGHT | The package weight exceeds 40kg, please make change or contact your category manager to apply DBS permission. |
| `4165` | BIZ_CHECK_MAX_PACKAGE_DIMENISIONS | The Package dimension(Length+Width+Height)  exceeds 300cm,  please make change or contact your category manager to  appl |
| `4166` | THD_IC_ERR_F_IC_INFRA_SPU_036 | eancode already exists |
| `4167` | THD_IC_ERR_F_IC_DOMAIN_PROPERTY_002 | The value of category property is required,property |
| `4168` | BIZ_CHECK_BRAND_PERMISSION_TIER_TWO | Tienes que solicitar dicha marca para poder utilizarla. |
| `4169` | CHK_IMAGE_MAX_ITEMS | Upload a maximum of 8 pictures |
| `500` | Create product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `SellerNotActive` | Seller not active,please check seller status | The seller's store status is inactive can not call the commodity API, you can call the GetSeller API and based on the St |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `500` | Create product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `5` | Invalid Request Format | Please check that the payload is documented and conforms to the XML formatting requirements. If you have URLs with “&” i |
| `4221` | BIZ_CHECK_MTEE_RISK_RULE_TRIGGER_MTEE_RISK_RULE_TRIGGER_ERROR | The price or content of this product violates the current national policy, please modify it according to the relevant po |
| `4139` | BIZ_CHECK_MAIN_IMAGE_REQUIRE:Main image is require | Please add at least one main or SKU image of the item. lazada does not allow uploading items without images. |
| `4130` | BIZ_CHECK_OFFICIAL_STORE_BRAND_UNAUTHORIZED | The store you requested is a Lazmall store, there are strict restrictions on the use of brands for this type of store, p |
| `4112` | CHK_SKU_PROPS_DUPLICATE | The values of variant attributes are duplicated between SKUs. Multi-SKU products need to set variant attributes and make |
| `209` | Invalid variation | The number of variants in the payload exceeds the upper limit or does not meet the requirements, please check the messag |


### DeactivateProduct
`POST` `/product/deactivate`

**Description:** Use this API to deactivate Product or SKUs corresponding to the product

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `apiRequestBody` | String | Yes | Parameter ItemId is mandatory, Skus is optional |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `E0001` | Parameter ItemId is mandatory | Parameter ItemId is mandatory |
| `E0002` | Product not exists | Product not exists |
| `E0003` | Seller Sku not exists | Seller Sku not exists |
| `E0004` | Product Status not online | Product Status not online |
| `E0006` | Unexpected internal error | Unexpected internal error |
| `E0004` | Product Status not online | The current item is already in the Inactive state and does not need to call this API. |
| `E0004` | Product Status not online | The current item is already in the Inactive state and does not need to call this API. |
| `E0004` | Product Status not online | The current item is already in the Inactive state and does not need to call this API. |
| `E0004` | Product Status not online | The current item is already in the Inactive state and does not need to call this API. |
| `E0004` | Product Status not online | The current item is already in the Inactive state and does not need to call this API. |
| `E0002` | Product:item id not exist | The item id in your request does not exist with the current store, please call GetProducts/GetProductItem API to check. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `4193` | The SellerSku parameter is no longer supported. Please update your parameter to use SkuId and try again | Seller sku field does not have uniqueness, so it cannot be used as a key field for editing products, please add SkuId fi |


### GetBrandByPages
`GET/POST` `/category/brands/query`

**Description:** Use this API to retrieve all product brands by page index in the system.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `startRow` | String | Yes | Number of brands to skip (i.e., an offset into the result set; together with the "limit" parameter, simple result set paging is possible; if you do page through results, note that the list of brands might change during paging). |
| `pageSize` | String | Yes | The maximum number of brands that can be returned. If you omit this parameter, the default of 40 is used. The Maximum is 200. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response data |
| `success` | Boolean | Yes | operation success or not  |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |


### GetCategoryAttributes
`GET` `/category/attributes/get`

**Description:** Use this API to get a list of attributes for a specified product category.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `primary_category_id` | String | Yes | identifiers of category code |
| `language_code` | String | No | Language code indicates the type of language you would like to translate. Please note not all languages are available in every region. For example, in Indonesia, only English and Indonesia are available.  If you are passing a language code which does not belong to your area, null value might receive. Please do make sure your language code is correct. Supported  language codes are listed as below: English:"en_US"  -  available in every area     Singapore:"en_SG" - available in Singapore    Thailand"th_TH" - available in Thailand     Indonesia:"id_ID" - available in Indonesia     Vietnam:"vi_VN" - available in Vietnam    Philippines: "fil_PH" - available in Philippines     Malaysia : "ms_MY" - available in Malaysia     Default(if null is passed): "en_US" |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `57` | E057: No attribute sets linked to that category. |  No attributes are linked to the specified category. |
| `4228` | Query category is not active | The category ID in the request is in Inactive state and cannot be used. Please call GetCategoryTree to query the latest  |
| `4227` | Query category is null | The category ID in the request does not exist in the current country, call the GetCategoryTree API to query the latest c |
| `4227` | Query category is null | The category ID in the request does not exist in the current country, call the GetCategoryTree API to query the latest c |
| `4228` | Query category is not active | The category ID in the request is in Inactive state and cannot be used. Please call GetCategoryTree to query the latest  |
| `4228` | Query category is not active | The category ID in the request is in Inactive state and cannot be used. Please call GetCategoryTree to query the latest  |
| `4227` | Query category is null | The category ID in the request does not exist in the current country, call the GetCategoryTree API to query the latest c |


### GetCategorySuggestion
`GET` `/product/category/suggestion/get`

**Description:** Get product's category suggestion by product title

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `product_name` | String | Yes | Product Name |
| `image_url` | String | Yes | image url |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `701` | E701: Empty category suggestion. | Empty category suggestion. |
| `1000` | Internal Application Error | Internal Application Error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `701` | Empty category suggestion. | The API is unable to provide suggestions, please change the product name and retry. |


### GetCategoryTree
`GET` `/category/tree/get`

**Description:** Use this API to retrieve the list of all product categories in the system.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `language_code` | String | No | Language code indicates the type of language you would like to translate. Please note not all languages are available in every region. For example, in Indonesia, only English and Indonesia are available.  If you are passing a language code which does not belong to your area, null value might receive. Please do make sure your language code is correct. Supported  language codes are listed as below: English:"en_US"  -  available in every area     Singapore:"en_SG" - available in Singapore    Thailand"th_TH" - available in Thailand     Indonesia:"id_ID" - available in Indonesia     Vietnam:"vi_VN" - available in Vietnam    Philippines: "fil_PH" - available in Philippines     Malaysia : "ms_MY" - available in Malaysia     Default(if null is passed): "en_US" |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | Response body |


### GetNextCascadeProp
`GET/POST` `/category/cascade/getNextCascadeProp`

**Description:** Use this API to query next cascade prop.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `categoryId` | Number | Yes | Category id |
| `cascadeId` | Number | Yes | Cascade id. Query from https://open.lazada.com/apps/doc/api?path=%2Fcategory%2Fattributes%2Fget |
| `path` | String | No | current cascade property path |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |


### GetPreQcRules
`GET/POST` `/product/seller/item/getPreQcRules`

**Description:** query pre qc rules

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `option` | Number | Yes | query qc option |
| `option_set` | Number[] | Yes | query qc rules option.[1] return item limit, [2] return restricted category id, [1,2] return both |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `values` | Object | Yes | response value |


### GetProductContentScore
`GET/POST` `/product/content/score/get`

**Description:** get product content score

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `item_id` | Number | Yes | Call this API; "Item Id" must be selected as the request parameter.  |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Result |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |


### GetProductItem
`GET` `/product/item/get`

**Description:** Get single product by ItemId or SellerSku.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `item_id` | Number | Yes | Call this API; "Item Id"  must be selected as the request parameter |
| `seller_sku` | String | No | The parameter has been deprecated and is no longer supported after November 15th, 2023. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `200` | E200: Empty SellerSku | Empty Item Id and Seller Sku. |
| `207` | E207: SKU not exist | Cannot find a Sku by the Seller Sku. |
| `208` | E208: Item not exist | Cannot find a Item by the Item Id. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `207` | SKU not exist | The item id used in the request does not exist on the current site.Please call GetProducts API to check if the sku you a |
| `207` | SKU not exist | The item id used in the request does not exist on the current site.Please call GetProducts API to check if the sku you a |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `207` | SKU not exist | The item id used in the request does not exist on the current site.Please call GetProducts API to check if the sku you a |
| `207` | SKU not exist | The item id used in the request does not exist on the current site.Please call GetProducts API to check if the sku you a |
| `207` | SKU not exist | The item id used in the request does not exist on the current site.Please call GetProducts API to check if the sku you a |
| `207` | SKU not exist | The item id used in the request does not exist on the current site.Please call GetProducts API to check if the sku you a |
| `EDIT_ITEM_NOT_BELONG_SELLER` | You are not authorized to edit the item. | The item id queried in the request does not belong to the current store, please call the GetProducts API to resynchroniz |
| `EDIT_ITEM_NOT_BELONG_SELLER` | You are not authorized to edit the item. | The item id queried in the request does not belong to the current store, please call the GetProducts API to resynchroniz |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `6` | Unexpected internal error | System fluctuations cause the query to fail please retry, if you encounter this error frequently when querying a particu |


### GetProducts
`GET` `/products/get`

**Description:** Use this API to get detailed information of the specified products.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `filter` | String | No | Returns the products with the status matching this parameter. Possible values are all, live, inactive, deleted, pending, rejected, sold-out. Mandatory. |
| `update_before` | String | No | Limits the returned product list to those updated before or on a specified date, given in ISO 8601 date format. Optional |
| `create_before` | String | No | Limits the returned products to those created before or on the specified date, given in ISO 8601 date format. Optional |
| `offset` | String | No | Deprecated(The number of Items you want to skip before you start counting),It is recommended to use date for scrolling query.The maximum offset is 10000 |
| `create_after` | String | No | Limits the returned products to those created after or on the specified date, given in ISO 8601 date format. Optional |
| `update_after` | String | No | Limits the returned products to those updated after or on the specified date, given in ISO 8601 date format. Optional |
| `limit` | String | No | The number of Items you would like to fetch from every response,The maximum is 50. |
| `options` | String | No | This value can be used to get more stock information. e.g., Options=1 means contain ReservedStock, RtsStock, PendingStock, RealTimeStock, FulfillmentBySellable. |
| `sku_seller_list` | String | No | Only products that have the Seller SKU in this list will be returned. Input should be a JSON array. For example, ["Apple 6S Gold", "Apple 6S Black"]. It only matches the whole words. A maximum of 100 SKUs can be returned. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format |  The request URL is not valid. |
| `6` | E006: Unexpected internal error | Unexpected internal error.  |
| `14` | E014: "%s" Invalid Offset |  The value for the offset parameter is not valid.  |
| `17` | E017: "%s" Invalid Date Format |  The date format is not valid. |
| `19` | E019: "%s" Invalid Limit |   The value for the limit parameter is not valid.  |
| `36` | E036: Invalid status filter |  The specified status filter is not valid. |
| `70` | E070: You have corrupt data in your sku seller list. |  Data in the SKU list are not valid. |
| `506` | E506: Get product failed |  Failed to get the product information. |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `SellerNotVerified` | Seller not verified,please check seller status | The seller's store opening process has not been completed, please log in to the Seller Center, check the store informati |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `19` | Invalid Limit | The limit field value is incorrect and should not exceed a maximum of 50. |


### GetQCAlertProducts
`GET` `/product/qc/alert/list`

**Description:** Getting seller's products that have been alerted by quality control.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `offset` | String | Yes | Number of QC alert products to skip |
| `limit` | String | Yes | The maximum number of QC alert products that can be returned.  |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | Response data list |


### GetResponse
`GET` `/image/response/get`

**Description:** Use this API to get the returned information from the system for the MigrateImages API.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `batch_id` | String | Yes | Request ID from the MigrateImages request |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format |  The format of the request URL is not valid. |
| `6` | E006: Unexpected internal error |  Unexpected internal error. |
| `302` | Not supported URL | The server is unable to download the image from the link, please check that the image link you added to the MigrateImage |
| `301` | Migrate Image Failed | The server is unable to download the image from the link, please check that the image link you added to the MigrateImage |
| `1000` | Internal Application Error | Please check that you are uploading a JPG or PGN image that meets the requirements, and if you are sure that there is no |


### GetSellerItemLimit
`GET` `/product/seller/item/limit`

**Description:** The platform will provide the product quantity limit information by this interface. The qps will be limited by seller, 10 qps per seller.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | The result of this request,true or false. |
| `errorCodes` | String[] | No | If the request failed, errorCodes will be returned. |
| `errorMsgs` | String[] | No | The error msg, may be null even though the result is failed. |
| `data` | Object | No | The data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `HOT_KEY_BLOCK_EXCEPTION` | hot key protect | 10 qps promised for each seller |
| `SELLER_SERVICE_FAIL` | inner service fail | inner system error, please retry |
| `ONLY_CB_SELLER_SUPPORTED` | For now, only cb seller supported | For local seller, we will support later. |
| `THIRD_SERVICE_ERROR` | inner service fail | inner system error, please retry |
| `SYS_ERROR` | inner service fail | inner system error, please retry |


### GetSizeChartTemplate
`GET` `/size/chart/template/get`

**Description:** 获取尺码模板列表

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `template_id` | Number | No | size chart template id |
| `template_name` | String | No | size chart name |
| `page_no` | Number | Yes | page no |
| `page_size` | Number | Yes | page size |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `4184` | E4184 | Size chart template id must be a number and greater than 0 |
| `4190` | E4190 | getSizeChartTemplate pageSize maximum value is 100 |
| `4176` | E4176 | Size chart list query fail |


### GetUnfilledAttributeItem
`GET/POST` `/product/unfilled/attribute/get`

**Description:** Get products without key attributes. (For cross boarder sellers Only)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `page_index` | Number | Yes | page_index |
| `attribute_tag` | String | Yes | The tag of attributes. Currently only has one value "key_prop"  属性标示。当前只支持key_prop |
| `page_size` | Number | Yes | The number of Products you would like to fetch from every response. The max number is 50.  返回的最大商品量。最大值50。商品级别 |
| `language_code` | String | Yes | Multi-language of category attributes that need to be returned |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | api |
| `total_products` | Number | Yes | The current product volume returned. Commodity level |
| `products` | Object[] | Yes | products |
| `error_msg` | String | Yes | error_msg |


### MigrateImage
`POST` `/image/migrate`

**Description:** Use this API to migrate a single image from an external site to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | Request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format |  The request URL is not valid. |
| `6` | E006: Unexpected internal error |  Unexpected internal error. |
| `30` | E030: Empty Request |  The request is not complete. |
| `301` | E301: Migrate Image Failed |  Failed to migrate the image. |
| `302` | E302: Not supported URL |  The image URL is not supported. |
| `303` | E303: The image is too large | The size of the migrated image exceeds the 1M limit. |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `1000` | Internal Application Error |  Internal system error. |
| `302` | Not supported URL | The server could not download the image from the link, please check that your link responds with an HTTP status code of  |
| `302` | Not supported URL | The server could not download the image from the link, please check that your link responds with an HTTP status code of  |
| `5` | Invalid Request Format | Please check that the payload is documented and conforms to the XML formatting requirements. If you have URLs with “&” i |
| `304` | Get Response Failed | Please check if the URL of the image you provided is externally accessible or if the HTTP status code of the response is |
| `303` | The image is too large | Please make sure your image size is less than 5000*5000px and file size is less than 3145728B. |
| `302` | Not supported URL | Please check if the http status code in response to the image link in the request is 200, and check if your image meets  |
| `1000` | Internal Application Error | Please check that you are uploading a JPG or PGN image that meets the requirements, and if you are sure that there is no |


### MigrateImages
`POST` `/images/migrate`

**Description:** Use this API to migrate multiple images from an external site to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB. A single call can migrate 8 images at most.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | Request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `batch_id` | String | Yes | The returned request ID is used by the GetResponse API to get the migrated image information. |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format |  The format of the request URL is not valid. |
| `6` | E006: Unexpected internal error |  Unexpected internal error. |
| `30` | E030: Empty Request |  The request is not complete. |
| `301` | E301: Migrate Image Failed |  Failed to migrate the images. |
| `302` | E302: Not supported URL |  The image URL is not supported. |
| `303` | E303: The image is too large |  The size of the migrated image exceeds the 1M limit. |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `1000` | Internal Application Error |  Internal system error. |


### ProductCheck
`GET/POST` `/product/pre/check`

**Description:** Use this API to check CB seller quantity limit of adding product .

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | String | Yes | <a href='https://open.lazada.com/apps/doc/doc?nodeId=10557&docId=108253' target='_brank'>Parameter description</a> |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |


### RemoveProduct
`POST` `/product/remove`

**Description:** Use this API to remove an existing product, some SKUs in one product, or all SKUs in one product. System supports a maximum number of 50 SellerSkus in one request.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `seller_sku_list` | String | No | sellerSku in a json list to be removed. System supports a maximum number of 50 sellerSku in one request.;for example: itemid: 1269656765 sellerSku: test00111 、test00222、test00333, then Param should be: ["test00111","test00222","test00333"]  |
| `sku_id_list` | String | No | Highest priority,skuId in a json list to be removed. System supports a maximum number of 50 skuId in one request.; for example: itemid: 1269656765 skuid: 5230534246, then Param should be: ["SkuId_1269656765_5230534246"]  |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format |  The request format is not valid. |
| `6` | E006: Unexpected internal error |  Unexpected internal error. |
| `30` | E030: Empty Request |  The request URL is not complete. |
| `204` | E204: Too many SKU in one request |  The number of SKUs exceeds the limit. |
| `503` | E503: Remove product failed |  Failed to remove the product. |
| `512` | E512: BIZ_CHECK_MANGROVE_RULE_QC | The request failed because the category was banned |
| `1000` | Internal Application Error |  Internal system error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `6` | Unexpected internal error | The seller_sku_list field has been deprecated, please use the sku_id_list field, if you still encounter this issue frequ |
| `503` | Remove product failed | This is a generalized error code, it is not possible to determine the specific problem based on this error code, please  |


### RemoveSku
`POST` `/product/sku/remove`

**Description:** Use this API to delete SKUs and sales attributes of corresponding products.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | String | Yes | <Request>     <Product>        <ItemId>1911687838</ItemId>      <variation><variation1><name>color_family</name> </variation1></variation>          <Skus>             <Sku>        <SellerSku>1911687838-1627269303789-1</SellerSku>                </Sku>         </Skus>     </Product> </Request> |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | Invalid Request Format | The request parameter is not formatted correctly, check that you are using the correct format against the RemoveSKU sect |


### SetImages
`POST` `/images/set`

**Description:** Use this API to set the images for an existing product by associating one or more image URLs with it.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | <a href='https://open.lazada.com/apps/doc/doc?nodeId=10557&docId=108254' target='_brank'>Parameter description</a> |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format |  The request format is not valid. |
| `6` | E006: Unexpected internal error |  Unexpected internal error. |
| `30` | E030: Empty Request |  The request URL is not complete. |
| `200` | E200: Empty SellerSku |  The Seller SKU is not specified. |
| `203` | E203: Too many images in one SKU |  The number of images exceeds the limit (8 images). |
| `204` | E204: Too many SKU in one request |  The number of SKUs exceeds the limit. |
| `504` | E504: Set product Image failed |  Failed to set images for the product. |
| `1000` | Internal Application Error |  Internal system error. |
| `504` | THD_IC_ERR_F_IC_ABILITY_PG_004:THD_IC_ERR_F_IC_ABILITY_PG_004 | The product is participating in a special Camapign that does not allow modification of images until the end of this Camp |


### UpdatePriceQuantity
`POST` `/product/price_quantity/update`

**Description:** Use this API to update the price and quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | <a href='https://open.lazada.com/apps/doc/doc?nodeId=42713&docId=121234' target='_brank'>Parameter description</a> |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format | The request format is not valid. |
| `6` | E006: Unexpected internal error | Unexpected internal error. |
| `30` | E030: Empty Request |  The request URL is not complete. |
| `204` | E204: Too many SKU in one request | The number of SKUs exceeds the limit. |
| `501` | E501: Update product failed |  Failed to update the product price or stock. |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `1000` | Internal Application Error |  Internal system error. |
| `4104` | BIZ_CHECK_PRICE_PRECISION_INVALID | Price accuracy check failed |
| `4105` | BIZ_CHECK_SELLER_SKU_DUPLICATE | SellerSku repeat |
| `4106` | CHK_CATPROP_CPV_INPUT_SIZE_LIMIT | Item customization attributes exceeded the limit |
| `4107` | CHECK_CAT_PROP_INVALID_NUMBER | The category attribute value is invalid |
| `4108` | CHK_BASIC_REQUIRED | Basic attributes Mandatory verification |
| `4109` | CHK_SKU_PROPS_NOT_MATCH_SALE_PROP | Sku sales attributes do not match |
| `4110` | BIZ_CHECK_CAT_PROP_MANDATORY | Category attribute This parameter is mandatory |
| `4111` | CHK_CATPROP_CPV_TEXT_REPEAT | Category attribute content repeats |
| `4112` | CHK_SKU_PROPS_DUPLICATE | Duplicate Sku attributes |
| `4113` | CHK_SKU_PROPS_NOT_IDENTICAL | Sales attribute is not filled in |
| `4114` | BIZ_CHECK_PRICE_SAMPLE_NON_ZERO | The sample price is 0 |
| `4115` | CHK_CATPROP_CPV_NOT_ENUM | The CPV attribute is not one of the options provided by the category |
| `4116` | BIZ_CHECK_MAIN_IMAGE_DUPLICATE | Repeat check of master diagram |
| `4117` | BIZ_CHECK_SPECIAL_PRICE_FROM_DATE_AFTER_TO_DATE | Special offer date check |
| `4118` | BIZ_CHECK_PRICE_IS_ZERO | Price is not 0 check |
| `4119` | BIZ_CHECK_SPECIAL_PRICE_RATE_OUT_OF_RANGE | Special price range check |
| `4120` | CHK_CATPROP_CPV_MAX_LEGNTH | Verify the maximum CPV value of a category |
| `4121` | BIZ_CHECK_SPECIAL_PRICE_PRECISION_INVALID | Special accuracy check does not pass |
| `4122` | BIZ_CHECK_VIRTUAL_BUNDLE_SKU_SUB_OVER_LIMIT | virtual bundle sku relation skuc over limit |
| `4123` | BIZ_CHECK_MANGROVE_RULE | Restricted publication check |
| `4124` | BIZ_CHECK_MANGROVE_RULE_QC | MANGROVE rule verification |
| `4125` | THD_IC_F_IC_DOMAIN_PROPERTY_002 | IC Verification category Attribute This parameter is mandatory |
| `4126` | THD_IC_F_IC_INFRA_PRODUCT_036 | SellerSku repeat |
| `4127` | THD_IC_F_IC_SCENE_PUBLISH_012 | ProductId repeat |
| `4128` | THD_IC_F_IC_DOMAIN_ACTOR_006 | Seller lock cannot be edited |
| `4129` | BIZ_CHECK_PROP_SPECIAL_CHAR | Containssymbol/characterthatisnotallowed:"<".Pleaseremovethenre-upload |
| `4130` | BIZ_CHECK_OFFICIAL_STORE_BRAND_UNAUTHORIZED | Uncertified brand |
| `4131` | BIZ_CHECK_CAT_PROP_SENSITIVE_WORDS | description has sensitive words New brand |
| `4132` | Invalid Request Format | Invalid Request Format |
| `4133` | Invalid variation | Invalid variation |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `4225` | Your product participated in semi-hosted program, please go to GSP to edit the product price/stock/package details information. | To modify the inventory of a Global Plus item call the AdjustSellableQuantity or UpdateSellableQuantity APIs. |
| `4225` | Your product participated in semi-hosted program, please go to GSP to edit the product price/stock/package details information. | To modify the inventory of a Global Plus item call the AdjustSellableQuantity or UpdateSellableQuantity APIs. |
| `4225` | Your product participated in semi-hosted program, please go to GSP to edit the product price/stock/package details information. | To modify the inventory of a Global Plus item call the AdjustSellableQuantity or UpdateSellableQuantity APIs. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `4225` | Your product participated in semi-hosted program, please go to GSP to edit the product price/stock/package details information. | To modify the inventory of a Global Plus item call the AdjustSellableQuantity or UpdateSellableQuantity APIs. |
| `513` | Internal call exception | A small number of occurrences are normal, if you want to avoid this error as much as possible, reduce the number of SKUs |
| `4225` | Your product participated in semi-hosted program, please go to GSP to edit the product price/stock/package details information. | To modify the inventory of a Global Plus item call the AdjustSellableQuantity or UpdateSellableQuantity APIs. |
| `4171` | The updated SKU quantity exceeds the maximum number 50, please do not update more than 50 SKUs at once | The number of SKUs included in a single request cannot exceed 50, and no more than 20 is recommended. |
| `4170` | During the Bday Mega campaign, there are restrictions for stock adjustments in effect between YYYY-MM-DD HH:MM:SS - YYYY-MM-DD HH:MM:SS. Sellers can increase stocks, but may not decrease stocks. | This SKU is participating in a special Campaign, so this SKU can't be updated to set stock less than current stock. |


### UpdateProduct
`POST` `/product/update`

**Description:** Use this API to update attributes or SKUs of an existing product. if need update inventory, offline, price, not recommended to use this API.
The iteration 25/6/2020 Updated for DBS changes. Refer to Input Parameters Payload

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | String | Yes | <a href='https://open.lazada.com/apps/doc/doc?nodeId=30715&docId=121228' target='_brank'>Parameter description</a> |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `1` | E001: Parameter %s is mandatory |  The parameter is mandatory but not specified. |
| `5` | E005: Invalid Request Format |  The request format is not valid. |
| `6` | E006: Unexpected internal error |  Unexpected internal error. |
| `30` | E030: Empty Request |  The request URL is not complete. |
| `201` | E201: %s Invalid CategoryId |  The specified category ID is not valid. |
| `202` | E202: %s Invalid SPUId |  The specified SPU ID is not valid. |
| `501` | E501: Update product failed |  Failed to update the product. |
| `512` | E512: BIZ_CHECK_MANGROVE_RULE_QC | The request failed because the category was banned |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `1000` | Internal Application Error |  Internal system error. |
| `4104` | BIZ_CHECK_PRICE_PRECISION_INVALID | Price accuracy check failed |
| `4105` | BIZ_CHECK_SELLER_SKU_DUPLICATE | SellerSku repeat |
| `4106` | CHK_CATPROP_CPV_INPUT_SIZE_LIMIT | Item customization attributes exceeded the limit |
| `4107` | CHECK_CAT_PROP_INVALID_NUMBER | The category attribute value is invalid |
| `4108` | CHK_BASIC_REQUIRED | Basic attributes Mandatory verification |
| `4109` | CHK_SKU_PROPS_NOT_MATCH_SALE_PROP | Sku sales attributes do not match |
| `4110` | BIZ_CHECK_CAT_PROP_MANDATORY | Category attribute This parameter is mandatory |
| `4111` | CHK_CATPROP_CPV_TEXT_REPEAT | Category attribute content repeats |
| `4112` | CHK_SKU_PROPS_DUPLICATE | Duplicate Sku attributes |
| `4113` | CHK_SKU_PROPS_NOT_IDENTICAL | Sales attribute is not filled in |
| `4114` | BIZ_CHECK_PRICE_SAMPLE_NON_ZERO | The sample price is 0 |
| `4115` | CHK_CATPROP_CPV_NOT_ENUM | The CPV attribute is not one of the options provided by the category |
| `4116` | BIZ_CHECK_MAIN_IMAGE_DUPLICATE | Repeat check of master diagram |
| `4117` | BIZ_CHECK_SPECIAL_PRICE_FROM_DATE_AFTER_TO_DATE | Special offer date check |
| `4118` | BIZ_CHECK_PRICE_IS_ZERO | Price is not 0 check |
| `4119` | BIZ_CHECK_SPECIAL_PRICE_RATE_OUT_OF_RANGE | Special price range check |
| `4120` | CHK_CATPROP_CPV_MAX_LEGNTH | Verify the maximum CPV value of a category |
| `4121` | BIZ_CHECK_SPECIAL_PRICE_PRECISION_INVALID | Special accuracy check does not pass |
| `4122` | BIZ_CHECK_VIRTUAL_BUNDLE_SKU_SUB_OVER_LIMIT | virtual bundle sku relation skuc over limit |
| `4123` | BIZ_CHECK_MANGROVE_RULE | Restricted publication check |
| `4124` | BIZ_CHECK_MANGROVE_RULE_QC | MANGROVE rule verification |
| `4125` | THD_IC_F_IC_DOMAIN_PROPERTY_002 | IC Verification category Attribute This parameter is mandatory |
| `4126` | THD_IC_F_IC_INFRA_PRODUCT_036 | SellerSku repeat |
| `4127` | THD_IC_F_IC_SCENE_PUBLISH_012 | ProductId repeat |
| `4128` | THD_IC_F_IC_DOMAIN_ACTOR_006 | Seller lock cannot be edited |
| `4129` | BIZ_CHECK_PROP_SPECIAL_CHAR | Containssymbol/characterthatisnotallowed:"<".Pleaseremovethenre-upload |
| `4130` | BIZ_CHECK_OFFICIAL_STORE_BRAND_UNAUTHORIZED | Uncertified brand |
| `4131` | BIZ_CHECK_CAT_PROP_SENSITIVE_WORDS | description has sensitive words New brand |
| `4132` | Invalid Request Format | Invalid Request Format |
| `4133` | Invalid variation | Invalid variation |
| `4134` | CHK_CATEGORY_ID_NOT_LEAF_CATEGORY | The category Id is Invalid |
| `4135` | THD_IC_ERR | IC service error reported |
| `4136` | SELLER_SKU_NOT_FOUND | Seller Sku is not found |
| `4137` | ITEM_NOT_FOUND | item not found |
| `4138` | BIZ_CHECK_EXIST_OUTER_IMAGE | The picture exists in the outer chain |
| `4139` | BIZ_CHECK_MAIN_IMAGE_REQUIRE | Main image is require |
| `4140` | CHK_ENUM_PROP_VALUE_NOT_IN_OPTION | Class does not have this attribute |
| `4141` | THD_IC_ERR_F_IC_INFRA_PRODUCT_036 | SellerSku repeat |
| `4142` | THD_BRAND_ID_IS_NOT_VALID_IN_CATEGORY | This brand is not valid in the category package |
| `4143` | BIZ_CHECK_SALEPROP_ATTRIBUTE_INVALID | The selling attributes are not defined in the variation |
| `4144` | BIZ_CHECK_SKU_NOT_CONTAIN_SALEPROP | The sku does not contain the saleProp tag |
| `4145` | BIZ_CHECK_SALEPROP_AND_OLD_PARAM_REPEAT | You can't put sales properties in both saleProp and sku |
| `4146` | BIZ_CHECK_SALEPROP_NOT_SUPPORT_THUMBNAIL | Thumbnails are not supported for this sale attribute |
| `10002` | Incorrect/missing/unavailable product attributes | Please check the details in the API response in order to confirm the properties that are causing the problem and the cau |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `10006` | the control price is not pass | Global Plus products have a price control logic: the price limit is: sku without postal price ≤ (pre-upgrade retail pric |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `10002` | System error update fail | Please check the details in the API response in order to confirm the properties that are causing the problem and the cau |
| `10006` | the control price is not pass | Global Plus products have a price control logic: the price limit is: sku without postal price ≤ (pre-upgrade retail pric |
| `10006` | the control price is not pass | Global Plus products have a price control logic: the price limit is: sku without postal price ≤ (pre-upgrade retail pric |
| `10006` | the control price is not pass | Global Plus products have a price control logic: the price limit is: sku without postal price ≤ (pre-upgrade retail pric |
| `10002` | System error update fail | Please check the details in the API response in order to confirm the properties that are causing the problem and the cau |
| `10002` | System error update fail | Please check the details in the API response in order to confirm the properties that are causing the problem and the cau |
| `10002` | System error update fail | Please check the details in the API response in order to confirm the properties that are causing the problem and the cau |
| `4137` | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProductItem API to query for the correct item id. | The item id in the request does not belong to the current store or country, please call the GetProduct/GetProductItem AP |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `4218` | Update product failed | The product has been penalized down to prohibit editing, if the seller has a problem with this product, please let the s |
| `4137` | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProductItem API to query for the correct item id. | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProdu |
| `4137` | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProductItem API to query for the correct item id. | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProdu |
| `10006` | the control price is not pass | Global Plus products have a price control logic: the price limit is: sku without postal price ≤ (pre-upgrade retail pric |
| `4137` | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProductItem API to query for the correct item id. | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProdu |
| `SellerNotActive` | Seller not active,please check seller status | The seller's store status is inactive can not call the commodity API, you can call the GetSeller API and based on the St |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `4218` | Update product failed | The product has been penalized down to prohibit editing, if the seller has a problem with this product, please let the s |
| `4216` | skuId is a mandatory field and must be filled in. | Sku id is a mandatory parameter when updating a product. |
| `4155` | Update product failed | The product is locked by the penalty does not support the update, please create a new product or appeal in the seller ce |
| `4152` | THD_INVENTORY_ERR_INV_PARAM_ILLEGAL:illegal parameter:quantity | Negative numbers are not allowed in the quantity field. |
| `4137` | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProductItem API to query for the correct item id. | The item id entered in the request does not exist on the current country and store, please call the GetProducts/GetProdu |
| `4115` | Attribute value that you input is not included in the dropdown list given. Please select from dropdown to avoid error | A value in a single or multiple select attribute does not exist in the Option provided by Lazada, call the GetCategoryAt |
| `4113` | CHK_SKU_PROPS_NOT_IDENTICAL | The custom variant attribute you are using is not declared in the variation tag, so please declare the variant attribute |
| `4108` | CHK_BASIC_REQUIRED | The current product may have unfilled mandatory attributes due to category update, please call GetCategoryAttributes API |
| `209` | Invalid variation | The number of variants in the payload exceeds the upper limit or does not meet the requirements, please check the messag |
| `1001` | The parameters are not in JSON format | Make sure that your payload is JSON compliant, that the attributes and structure are filled out correctly according to t |


### UpdateSellableQuantity
`GET/POST` `/product/stock/sellable/update`

**Description:** Use this API to update sellable quantity of one or more existing products. The maximum number of products that can be updated is 50, but 20 is recommended.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | String | Yes | Please take demo as reference. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `5` | E005: Invalid Request Format | The request format is not valid. |
| `6` | E006: Unexpected internal error | Unexpected internal error. |
| `30` | E030: Empty Request | The request URL is not complete. |
| `204` | E204: Too many SKU in one request | The number of SKUs exceeds the limit. |
| `501` | E501: Update product failed | Failed to update the product price or stock. |
| `901` | E901: The request is too frequent, or the requested functionality is temporarily disabled. | Failed to return the requested data due to high calling frequency or disabled functionality. Please try again later. |
| `1000` | Internal Application Error |  Internal system error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `501` | Update product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `4170` | During the Bday Mega campaign, there are restrictions for stock adjustments in effect between YYYY-MM-DD HH:MM:SS - YYYY-MM-DD HH:MM:SS. Sellers can increase stocks, but may not decrease stocks. | This SKU is participating in a special Campaign, so this SKU can't be updated to set stock less than current stock. |
| `901` | Limit service request speed in server side temporarily. | API level QPS limiting flow, please retry in the next second when you encounter this error. |
| `6` | Unexpected internal error | System fluctuation please retry, if you encounter this error frequently, please create a ticket to consult. |
| `513` | Internal call exception | A small number of occurrences are normal, if you want to avoid this error as much as possible, reduce the number of SKUs |
| `4170` | During the Bday Mega campaign, there are restrictions for stock adjustments in effect between YYYY-MM-DD HH:MM:SS - YYYY-MM-DD HH:MM:SS. Sellers can increase stocks, but may not decrease stocks. | This SKU is participating in a special Campaign, so this SKU can't be updated to set stock less than current stock. |


### UploadImage
`POST` `/image/upload`

**Description:** Use this API to upload a single image file to Lazada site. Allowed image formats are JPG and PNG. The maximum size of an image file is 1MB.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `image` | byte[] | Yes | Upload an image file |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `30` | E030: Empty Request |  The request is not complete. |
| `300` | E300: Upload Image Failed |  Failed to upload the image. |
| `303` | E303: The image is too large |  The size of the uploaded image exceeds the 1M limit. |
| `1000` | Internal Application Error |  Internal system error. |
| `302` | Not supported URL | The image field should be passed in as a stream rather than a string, check that you are not passing in data that is not |
| `302` | Not supported URL | The image field should be passed in as a stream rather than a string, check that you are not passing in data that is not |
| `302` | Not supported URL | The image field should be passed in as a stream rather than a string, check that you are not passing in data that is not |
| `302` | Not supported URL | The image field should be passed in as a stream rather than a string, check that you are not passing in data that is not |
| `302` | Not supported URL | The image field should be passed in as a stream rather than a string, check that you are not passing in data that is not |
| `303` | The image is too large | Please make sure your image size is less than 5000*5000px and file size is less than 3145728B. |
| `302` | Not supported URL | The image field should be passed in as a stream rather than a string, check that you are not passing in data that is not |
| `1000` | Internal Application Error | Please check that you are uploading a JPG or PGN image that meets the requirements, and if you are sure that there is no |


---
## Cross Boarder Product API

_API group for crossborder APIs (Cross Boarder Sellers Only)_

### CreateGlobalProduct
`POST` `/product/global/create`

**Description:** Use this API to create a single new global product to multiple Lazada sites. (For cross boarder sellers ONLY)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | <a href='https://open.lazada.com/apps/doc/doc?nodeId=30715&docId=121751' target='_brank'>Parameter description</a> |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `"500"` | "E500: Create product failed" | The product was not created, please check the detailed error message. |
| `ServiceTimeout` | The request has failed due to service timeout | The request has failed due to service timeout |
| `"6"` | "E006: Unexpected internal error" | There is internal error, please contact our tech support team for assistance. |
| `"5"` | "E005: Invalid Request Format" | There is something wrong in the request format, please check the detailed error message |
| `IllegalAccessToken` | The specified access token is invalid or expired | Your access token is either expired or invalid. Pleaes refresh your access token and contact our tech support team to re |
| `4136` | SYSTEM_BUSY | System is busy,try later |
| `4137` | SYSTEM_TIMEOUT | System is timeout,try later |
| `4138` | SYSTEM_EXCEPTION | System is under maintenance |
| `4139` | UNKNOWN_ERROR | System is upgrading... |
| `4140` | CATEGORY_CANNOT_FIND | The specified category cannot be found |
| `4141` | CATEGORY_NOT_PERMITTED | The category is not permitted |
| `4142` | CATEGORY_IS_INACTIVE | The category is inactive |
| `4143` | NO_TARGET_USER_BIND | Do not have access to target market |
| `4144` | LOCAL_CATEGORY_CANNOT_FIND | The specified local category cannot be found |
| `4145` | GLOBAL_PRODUCT_CANNOT_FIND | The global product cannot be found |
| `4146` | LOCAL_PRODUCT_CANNOT_FIND | Cannot find product at local venture: %s |
| `4147` | LOCAL_SKU_CANNOT_FIND | Cannot find any SKUs at local venture: %s |
| `4148` | LOCAL_PRODUCT_HAS_NOT_BEEN_SYNCED | This product has not been published to venture: %s |
| `4149` | LOCAL_SKU_HAS_NOT_BEEN_SYNCED | The SKUs is still synchronizing to venture: %s... |
| `4150` | PRODUCT_DOES_NOT_BELONG_TO_USER | Target product does not belong to the user account |
| `4151` | DAO_NOT_SUPPORT_BIZ_TYPE | The biz type hasn't been supported |
| `4152` | DAO_GLOBAL_PDT_NOT_FOUND | Can not find the global product in the data base |
| `4153` | DAO_GLOBAL_SKU_NOT_FOUND | Can not find the global SKUs in the data base |
| `4154` | DAO_LOCAL_ITEM_RELATION_NOT_FOUND | Can not find the local items in the data base |
| `4155` | DAO_LOCAL_SKU_RELATION_NOT_FOUND | Can not find the local skus in the data base |
| `4156` | NO_CREATE_PRODUCT_PERMISSION | You do not have permission to create product |
| `4157` | PRICE_GENERAL_TOO_LOW_ERROR | Retail price & Sale price at %s cannot be lower than %s %s |
| `4158` | PRICE_GENERAL_TOO_HIGH_ERROR | Retail price & Sale price at %s cannot be higher than %s %s |
| `4159` | PRICE_GENERAL_DISCOUNT_TOO_HIGH_ERROR | Sale price discount at %s should not be more than or equal to %s. Currently, price is %s %s, and sale price is %s %s |
| `4160` | PRICE_GENERAL_DISCOUNT_TOO_LOW_ERROR | Sale price discount at %s should not be less than or equal to %s. Currently, price is %s %s, and sale price is %s %s |
| `4161` | PB_SALE_PROP_RENDER_ILLEGAL_SALE_PROP | The product has illegal sale properties |
| `4162` | PB_SKU_DESC_RENDER_ILLEGAL_SKU_DESC | The product has illegal sku description properties |
| `4163` | PB_VENTURE_NO_VENTURE_SELECT | No venture has been selected |
| `4164` | PB_VENTURE_MY_NOT_PUBLISHED | Malaysia venture is a must for Cross-Border publishing |
| `4165` | PB_NAME_NAME_CANNOT_BE_NULL | Title cannot be empty |
| `4166` | PB_NAME_NAME_CANNOT_BE_TOO_LONG | Title cannot be longer than %d |
| `4167` | PB_NAME_NAME_TRAN_TOO_LONG | publish failed cause by title translation words overflow |
| `4168` | PB_BRAND_ILLEGAL | The brand is invalid |
| `4169` | PB_DETAIL_ATTRIBUTE_REQUIRED | This attribute cannot be empty |
| `4170` | PB_SHORTDESC_REQUIRED | Highlights cannot be empty |
| `4171` | PB_DETAIL_LENGTH_ERROR | This attribute cannot be longer than 255 characters |
| `4172` | PB_SKU_PROP_REQUIRED | This attribute cannot be empty |
| `4173` | PB_NO_PROPER_SKU | No proper sku found |
| `4174` | PB_NO_PC_DECO | PC decoration cant be empty |
| `4175` | PB_NO_WIRELESS_DECO | wireless decoration cant be empty |
| `4176` | PB_IMG_CANNOT_BE_EMPTY | Please upload at least one image for every SKU |
| `4177` | PB_IMG_URL_INVALID | the img url is invalid |
| `4178` | PB_IMG_CANNOT_FETCH | the img could not be fetch caused by network reason or firewall,push the img to where we can access and download |
| `4179` | PB_SALE_PROP_SUBMIT_SALE_PROP_CANNOT_BE_EMPTY | The sale property cannot be empty |
| `4180` | PB_SALE_PROP_SUBMIT_SALE_PROP_ILLEGAL_VAL | The sale property value is invalid |
| `4181` | PB_SALE_PROP_SUBMIT_SALE_PROP_INVALID_INPUT_VAL | The sale property value cannot contain illegal character \"*^~<>//\ |
| `4182` | PB_SALE_PROP_SUBMIT_SALE_PROP_TOO_LONG_INPUT_VAL | The sale property value cannot be longer than 255 characters |
| `4183` | PB_SALE_PROP_SUBMIT_TOO_MORE_SKU | The SKUs are too more for these sale properties |
| `4184` | PB_CURRENCY_CANNOT_BE_EMPTY | The SKU does not have legal currency |
| `4185` | PB_ORIGIN_PRICE_CANNOT_BE_EMPTY | The SKU's original price is empty |
| `4186` | PB_ORIGIN_SALE_PRICE_CANNOT_BE_EMPTY | The SKU's original sale price is empty |
| `4187` | PB_ORIGIN_SALE_PRICE_CANNOT_BE_HIGH | The SKU's sale price must be lower than original price |
| `4188` | PB_MARKET_PRICE_CANNOT_BE_EMPTY | The %s SKU's original price is empty |
| `4189` | PB_MARKET_SALE_PRICE_CANNOT_BE_EMPTY | The %s SKU's original sale price is empty |
| `4190` | PB_MARKET_SALE_PRICE_TOO_HIGH | The %s SKU's sale price must be lower than retail price |
| `4191` | PB_STOCK_CANNOT_BE_EMPTY | The %s SKU's stock cannot be empty |
| `4192` | PB_STOCK_INVALID | The %s SKU's stock is invalid |
| `4193` | PB_WARRANTY_INVALID | Warranty Period has not been selected while Warranty Type is not \"No warranty\" |
| `4194` | PB_SELLER_SKU_EXIST | Seller SKU exists |
| `4195` | PB_SELLER_SKU_LENGTH_ERROR | Seller SKU should be 1-50 characters |
| `4196` | PB_SELLER_SKU_DUPLICATE | Seller SKU duplicates |
| `4197` | PB_SELLER_SKU_INVALID | Seller SKU must consist of \"A-Z\", \"a-z\", \"0-9\", \"-\", \"_\" |
| `4198` | PB_SELLER_SKU_CANNOT_BE_REVISED | Seller SKU cannot be revised, it should be kept as: %s |
| `4199` | PB_PACKAGE_UNMATCHED | Package parameters should be all the same for one product |
| `4200` | IMAP_BRAND_NOT_MATCHED | Brand doesn't match at local venture |
| `4201` | IMAP_SALE_PROP_UNMATCHED | Sale properties unmatched |
| `4202` | IMAP_SALE_PROP_ERR_MATCHED | Sale properties error matched |
| `4203` | IMAP_DEST_SALE_PROP_IS_SPU | Dest sale property is SPU property |
| `4204` | IMAP_SALE_PROP_VAL_ERR_MATCHED | Sale properties values error matched |
| `4205` | INVALID_IMAGE_FORMAT | Invalid image format |
| `4206` | INVALID_IMAGE_DIMENSION | Image resolution shoule be from 330 * 330 to 5000 * 5000 |
| `4207` | IMPORT_SELLER_SKU_EMPTY | Import empty seller sku |
| `4208` | IMPORT_SELLER_SKU_INVALID | Import invalid seller sku |
| `4209` | INVALID_CATEGORY | Invalid category |
| `4210` | FAIL_TO_GET_CATEGORY_ID | Fail to get category id |
| `4211` | HAZMAT_WARN | HAZMAT_WARN |
| `4212` | PDT_LIMIT_REACH | Your product list has reached the limit at %s |
| `4213` | MIGRAGE_IMAGE_FAILED | Fail to migrate |
| `4214` | PG_NOT_PERMIT | QC rule checking failed |
| `4215` | DECO_CREATE_ERROR | Fail to create decoration |
| `4216` | DECO_NO_TARGET_USER_BIND | cant find target user |
| `4217` | DECO_SOURCE_QUERY_ERROR | Fail to query for source decoration |
| `4218` | DECO_TRANSLATE_ERROR | Decoration translation error |
| `4219` | DECO_SYNC_ERROR | Decoration sync error |
| `4220` | TRANSLATE_OVER_FLOW | Translate over flow,and will retry auto |
| `4221` | ITEM_NEVER_PUBLISH_SUCCESSED | item have not pulish success to market before |
| `4222` | NO_SKU_COULD_BE_UPDATE | no sku could be update |
| `4223` | PRODUCT_NUM_REACH_LIMITATION | Seller's online product quantity has reach limit in all site |
| `4224` | SELLER_STATUS_INVALID | Seller's account are inactive or not verified in all publish venture. |
| `4225` | PRICE_NOT_VALID | Please review product price to ensure accuracy. |
| `4226` | SELLER_PUNISHMENT_INVALID | Seller is under punishment of blocking edit venture. |
| `4227` | SKU_IMAGE_INVALID | If you upload a sku picture, then all sku must be uploaded. |
| `4228` | PRODUCT_IMAGE_INVALID | Product or sku Image is missing for live product |
| `4229` | PROHIBITED_BRAND | Product has problem with brand. |
| `4230` | PROHIBITED_KEYWORD | Product content exist keyword. |
| `4223` | Seller's online product quantity has reach limit in all publish venture. | The number of products with active status in GPS has exceeded the limit. Please call updateProductStatus API to drop the |
| `500` | Create product failed | This error code is an overview error code and cannot be used to determine the detailed cause of the error, please check  |
| `5` | Invalid Request Format | Please refer to the “CreateGlobalProduct payload and parameter description” document to check if the payload in your req |
| `4214` | Create product failed | This is a generalized error, this error indicates that there is a sales/logistics policy in your item that does not comp |
| `4194` | Seller SKU exists. | This seller sku already exists in the current store, please change to another seller sku to post the item. |
| `4178` | Fail to migrate image | Please do not include external image links in the payload, use the MigrateImage API to migrate the images to Lazada imag |
| `4169` | This attribute cannot be empty. | Mandatory attributes are not used in the Payload, call the GetCategoryAttributes API to check if you missed any attribut |
| `4159` | Create product failed | Local price limit. sale price and sepcial price can't be less than the specified percentage, please check the detail fie |
| `309` | Video id status is not audit success | Only videos that are in the AUDIT SUCCESSS state can be used in payload. |


### GetGlobalProductExtension
`GET` `/product/global/extension`

**Description:** Use this API to query the extension info of the specified global product. (CrossBoarderSellersOnly)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `global_item_ids` | Number[] | No | Batch size is limited to 50 |
| `item_ids` | Number[] | No |  Batch size is limited to 50, if global_Item_ids is present, this field will be ignored |
| `country` | String | No | country,if global_Item_ids is present, this field will be ignored |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | process result，If this is true, it doesn't mean that everything is processed successfully. It is necessary to judge that the item_err_code in packages is equal to 0 to determine that the processing is successful.  |
| `error_code` | String | No | exists when success is false |
| `error_msg` | String | No | exists when success is false |
| `data` | Object[] | No | resp body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `E1000` | Internal Application Error | Endpoint exception, please use MY endpoint for GSP related requests. |


### GetGlobalProductStatus
`GET` `/product/global/status/get`

**Description:** Use this API to query the status of the specified global product. It takes several minutes for the global product to be created on each site. (CrossBoarderSellersOnly)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `params` | Object | Yes | put the "sellerSku" as the key |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | Yes | result json type string |
| `success` | Boolean | Yes | success flag |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `"E0207"` | "E207: SKU not exist" | This SKU can not be found under your shop account. |
| `E0208` | Product not exist | The requested seller sku does not exist in the current store, please check the correctness of the seller sku. |
| `E1000` | Internal Application Error | Endpoint exception, please use MY endpoint for GSP related requests. |
| `E0208` | Product not exist | The requested seller sku does not exist in the current store, please check the correctness of the seller sku. |


### GetRecommendPrice
`GET/POST` `/product/global/semi/recommend/price/get`

**Description:** get recommend price

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | request data |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | data |
| `success` | Boolean | Yes | true |
| `error_code` | String | Yes | null |
| `error_msg` | String | Yes | null |


### GetUnfilledAttribute
`GET` `/product/global/unfilled/attribute/get`

**Description:** get the product which have attribute not filled （for cross boarder sellers Only）

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `offset` | Number | Yes | offset |
| `limit` | Number | Yes | pageSize |
| `attributeTag` | String | Yes | only support key_prop |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |
| `success` | Boolean | Yes | success or false |
| `error_detail` | String | Yes | error detail |
| `error_code` | String | Yes | error code |
| `errors` | String | Yes | errors |
| `error_msg` | String | Yes | error msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `19` | E019: Invalid Limit | The maximum value of limit is 50 |
| `306` | E306: attribute tag not allowed | attributeTag only enter "key_ prop" |


### GetUpgradableGlobalPlusProductList
`GET/POST` `/product/global/semi/avaible/get`

**Description:** get an upgradeable global plus product list 

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `type` | String | Yes | global |
| `country` | String | No | country |
| `pageNo` | String | Yes | page no |
| `pageSize` | String | Yes | page size |
| `currentIndex` | String | Yes | current index |
| `itemIds` | Number[] | No | itemId or productId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | data |
| `success` | Boolean | Yes | true |


### SemiProductUpdate
`POST` `/product/global/semi/update`

**Description:** SemiProductUpdate

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | String | Yes | request data |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | response data |
| `success` | Boolean | No | success or fail |
| `error_code` | String | No | error code |
| `error_msg` | String | No | error msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `10001` | Illegal parameters | 参数不合法 |
| `10002` | System error  | 系统异常 |
| `10003` | Item not found | 商品未找到 |
| `10004` | price need to be lower than the original price | 价格需低于零售价 |
| `10005` | 商品已升级 | 商品已升级 |
| `10006` | 商品校验失败，无法升级 | 商品校验失败，无法升级 |


### SemiProductUpgrade
`GET/POST` `/product/global/semi/upgrade`

**Description:** SemiProductUpgrade

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | request data |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response data |
| `success` | Boolean | Yes | success or fail |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error msg |


### UpdateGlobalProductAttribute
`POST` `/product/global/attribute/update`

**Description:** update global product attribute (For cross boarder sellers only)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `payload` | Payload | Yes | the content want to update |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | success or fail |
| `error_detail` | String | Yes | error detail |
| `error_code` | String | Yes | error code |
| `errors` | String | Yes | all errors |
| `error_msg` | String | Yes | error msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `501` | E501: Update product failed | Update product failed |


### deleteMerchantProduct
`POST` `/product/global/delete`

**Description:** Use this API to delete the product。(CrossBoarderSellersOnly)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `type` | String | Yes | Product Types |
| `country` | String | No | country,if type is "global", this field will be ignored |
| `product_id` | Number | Yes | When  type is "global", it is the global product ID, when type is "single",  product id is the IC product ID. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | body，deleteGspProductResult is true，mean update gsp product success。when deleteICProductResult is false，mean update IC product fail，and deleteIcProductFailResultList will show the reason |
| `success` | Boolean | No | process result，If this is true, it doesn't mean that everything is processed successfully |
| `error_code` | String | No | exists when success is false |
| `error_msg` | String | No | exists when success is false |


### updateProductStatus
`POST` `/product/global/update/status`

**Description:** product up shelf or down shelf，(CrossBoarderSellersOnly)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `type` | String | Yes | Product Types |
| `country` | String | No | country,if type is "global", this field will be ignored |
| `product_id` | Number | Yes | When  type is "global", it is the global product ID, when type is "single",  product id is the IC product ID. |
| `status` | String | Yes | update product type |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | body，updateGspProductResult is true，mean update gsp product success。when updateICProductResult is false，mean update IC product fail，and updateIcProductFailResultList will show the reason |
| `success` | Boolean | No | process result，If this is true, it doesn't mean that everything is processed successfully |
| `error_code` | String | No | exists when success is false |
| `error_msg` | String | No | exists when success is false |


---
## Product Review API

_For product review record and reply API_

### GetHistoryReviewIdList
`GET/POST` `/review/seller/history/list`

**Description:** Get history review id list for one seller(reviews within 3 months can be get)

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `item_id` | String | Yes | Product Item ID |
| `order_id` | Number | No | Order ID |
| `start_time` | Number | Yes | Start Time, timestamp in millisecond, this is the same with "create_time" in the response data of interface (/review/seller/list/v2)；The time range cannot exceed 7 days |
| `end_time` | Number | Yes | End Time, timestamp in millisecond, this is the same with "create_time" in the response data of interface (/review/seller/list/v2)；The time range cannot exceed 7 days |
| `current` | Number | Yes | The current pageNo, default value = 1, max value = 50 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response data |
| `success` | Boolean | Yes | success or fail |
| `error_code` | String | No | error code |
| `error_msg` | String | No | error msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `PARAMS_VALIDATE_ERROR` | NULL_SELLERID | Cannot recognize "seller_id" |
| `PARAMS_VALIDATE_ERROR` | NULL_ITEMID | Cannot recognize "item_id" |
| `PARAMS_VALIDATE_ERROR` | NULL_CURRENT | Cannot recognize "current" |
| `PARAMS_VALIDATE_ERROR` | CURRENT_ABOVE_LIMIT | "current" is above the limit, the max value is 50 |
| `PARAMS_VALIDATE_ERROR` | NULL_STARTTIME_OR_ENDTIME | Cannot recognize "start_time" or "end_time" |
| `PARAMS_VALIDATE_ERROR` | STARTTIME_OVER_LIMIT | Only support checking 90 days of history data |
| `PARAMS_VALIDATE_ERROR` | TIMESPAN_ABOVE_LIMIT | Only support checking 7days data at one time |
| `PARAMS_VALIDATE_ERROR` | WRONG_ORDER_ID | Cannot recognize "order_id" |
| `TRAFFIC_CONTROL` | TRAFFIC_CONTROL | Traffic control |
| `PARAMS_VALIDATE_ERROR` | PARAMS_VALIDATE_ERROR | start_time&end_time range cannot exceed 7 days. |
| `PARAMS_VALIDATE_ERROR` | PARAMS_VALIDATE_ERROR | start_time&end_time range cannot exceed 7 days. |
| `Mp3SellerApiLimit` | Mp3 Seller not support the api -apipath | MP3 sellers cannot call the current API, please readthis document for a list of APIs that can be called by MP3 sellers,  |
| `PARAMS_VALIDATE_ERROR` | PARAMS_VALIDATE_ERROR | start_time&end_time range cannot exceed 7 days. |


### GetReviewListByIdList
`GET` `/review/seller/list/v2`

**Description:** get review list by id list, need get id list first

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id_list` | Number[] | Yes | id list, maxLength = 10 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response data |
| `success` | Boolean | Yes | * |
| `error_code` | String | Yes | * |
| `error_msg` | String | Yes | * |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `PARAMS_VALIDATE_ERROR` | NULL_SELLERID | Cannot recognize sellerid |
| `PARAMS_VALIDATE_ERROR` | NULL_ID | id list is null |
| `TRAFFIC_CONTROL` | TRAFFIC_CONTROL | Traffic control |
| `Mp3SellerApiLimit` | Mp3 Seller not support the api - apipath | MP3 sellers cannot call the current API, please readthis document for a list of APIs that can be called by MP3 sellers,  |


### SubmitSellerReply
`GET` `/review/seller/reply/add`

**Description:** submit seller reply for customers review

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | review id that user wants to reply to. Can be obtain from GetProductReviewList |
| `content` | String | Yes | reply content in text, only support reply in text.max length = 500 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Boolean | No | reply success or fail |
| `success` | Boolean | No | reply success or fail |
| `error_code` | String | No | error code |
| `error_msg` | String | No | error msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `PARAMS_VALIDATE_ERROR` | NULL_SELLERID | Cannot recognize sellerid |
| `PARAMS_VALIDATE_ERROR` | NULL_ID | Cannot recognize id |
| `PARAMS_VALIDATE_ERROR` | NULL_CONTENT | Empty content |
| `PARAMS_VALIDATE_ERROR` | REPLY_ALREADY | Already replied. All reply needs go through quality control process. |
| `PARAMS_VALIDATE_ERROR` | NO_SUCH_REVIEW | No such review |
| `PARAMS_VALIDATE_ERROR` | REVIEW_STATUS_CANNOT_REPLY | Review status cannot be replied to, review's status may be changed because of being edited or reported   |
| `PARAMS_VALIDATE_ERROR` | REVIEW_TYPE_DONOT_SUPPORT_REPLY | Review type cannot be replied to, only reply to PRODUCT_REVIEW |
| `PARAMS_VALIDATE_ERROR` | REVIEW_INFO_DONOT_SUPPORT_REPLY | Review info cannot be replied to, review must have text content or images or video  |
| `PARAMS_VALIDATE_ERROR` | REVIEW_REPORTED_CANNOT_REPLY | Review been reported cannot be repied to |
| `PARAMS_VALIDATE_ERROR` | REPLY_CONTENT_TOO_LONG | Reply too long |
| `PARAMS_VALIDATE_ERROR` | BEYOND_REPLY_PERIOD | Reply over due |
| `TRAFFIC_CONTROL` | TRAFFIC_CONTROL | Traffic control |
| `PARAMS_VALIDATE_ERROR` | REPLY_ALREADY | This review has already been replied to and does not support multiple replies. |


---
## Store Decoration API

### GetStoreCustomPage
`GET/POST` `/store/custom/page/get`

**Description:** GetStoreCustomPagevice


**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `page` | String | Yes | page |
| `size` | String | Yes | size |
| `keyword` | String | No | Support keyword search |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | ellipsis |


---
## Media Center API

_For video upload/delete/create API_

### CompleteCreateVideo
`POST` `/media/video/block/commit`

**Description:** After uploading all blocks of the video file,  call CompleteCreateVideo to complete the video uploading process.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `uploadId` | String | Yes | return by calling InitCreateVideo |
| `parts` | String | Yes | a json string contains e_tag info of each block |
| `title` | String | Yes | the video title |
| `coverUrl` | String | Yes | the url of the video's cover image |
| `videoUsage` | String | No | the usage of video, "pro_main_video" represent prodcut main video, "im" represent chat video |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | whether the operation succeeds |
| `result_code` | String | Yes | error code when the operation fails |
| `video_id` | String | Yes | return video_id for further call |
| `result_message` | String | Yes | error message when the operation fails |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `ILLEGAL_PARAMETER` | detail message | illegal parameter |
| `FAIL_TO_BLOCK_COMPLETE` | detail message | fail to complete block upload |
| `FAIL_TO_VALIDATE` | detail message | fail to validate video |
| `FAIL_TO_ADD_VIDEO` | detail message | fail to add video |


### GetVideo
`GET` `/media/video/get`

**Description:** You call this action to get video info after uploading.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `videoId` | Number | Yes | the previous return value by calling CompleteCreateVideo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cover_url` | String | Yes | cover url of the video |
| `video_url` | String | Yes | url of the video |
| `success` | Boolean | Yes | whether the operation succeeds |
| `result_code` | String | Yes | error code when the operation fails |
| `state` | String | Yes | possible values: READY_FOR_TRANSCODE, TRANSCODING, TRANSCODE_FAILED, READY_FOR_AUDIT, AUDIT_FAILED, AUDIT_SUCCESS, DELETED |
| `title` | String | Yes | title of the video |
| `result_message` | String | Yes | error message when the operation fails |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `ILLEGAL_PARAMETER` | detail message | illegal parameter |
| `FAIL_TO_GET_SHOP_INFO` | detail message | fail to get shop info |
| `FAIL_TO_GET_VIDEO` | detail message | fail to get video |


### GetVideoQuota
`GET` `/media/video/quota/get`

**Description:** You call this api to get the capacity quota of seller.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `capacity_size` | Number | Yes | the max space of all video files |
| `used_size` | Number | Yes | current space taken |
| `success` | Boolean | Yes | whether the operation succeeds |
| `result_code` | String | Yes | error code when the operation fails |
| `result_message` | String | Yes | error message when the operation fails |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `ILLEGAL_PARAMETER` | more detail | illegal parameter |
| `FAIL_TO_GET_SHOP_INFO` | more detail | fail to get shop info |
| `FAIL_TO_GET_USER_CAPACITY` | more detail | fail to get capacity |


### InitCreateVideo
`POST` `/media/video/block/create`

**Description:** A seller starts to upload a video file

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `fileName` | String | Yes | local file name of vedio file |
| `fileBytes` | Number | Yes | video file's bytes, should be less than 100M |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `upload_id` | String | Yes | return upload_id for further operation  |
| `success` | Boolean | Yes | whether the operation succeeds |
| `result_code` | String | Yes | error code when the operation fails |
| `result_message` | String | Yes | error message when the operation fails |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `ILLEGAL_PARAMETER` | detail message | illegal parameter |
| `FAIL_TO_GET_SHOP_INFO` | detail message | fail to get shop info |
| `FAIL_TO_BLOCK_INIT` | detail message | fail to create block upload |


### RemoveVideo
`POST` `/media/video/remove`

**Description:** You can this api to delete a video file permanently.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `videoId` | Number | Yes | the previous return value by calling CompleteCreateVideo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | whether the operation succeeds |
| `result_code` | String | Yes | error code when the operation fails |
| `result_message` | String | Yes | error message when the operation fails |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `ILLEGAL_PARAMETER` | detail message | illegal parameter |
| `FAIL_TO_GET_SHOP_INFO` | detail message | fail to get shop info |
| `FAIL_TO_GET_VIDEO` | detail message | fail to get video |
| `FAIL_TO_DELETE_VIDEO` | detail message | fail to delete video |


### UploadVideoBlock
`POST` `/media/video/block/upload`

**Description:** The API is used to upload one block of origin video file. The video file can split into multiple files. For example, a 8MB video file can be split into three blocks. 3MB, 3MB and 2MB. These three blocks can be uploaded by calling UploadVideoBlock three times.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `uploadId` | String | Yes | return by calling InitCreateVideo |
| `blockNo` | String | Yes | the current block number, from 0 to N-1 |
| `blockCount` | String | Yes | total block count of file |
| `file` | byte[] | Yes | binary content of the current block |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | whether the operation succeeds |
| `result_code` | String | Yes | error code when the operation fails |
| `e_tag` | String | Yes | return e_tag for using in commit operation  |
| `result_message` | String | Yes | error message when the operation fails |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `ILLEGAL_PARAMETER` | detail message | illegal parameter |
| `FAIL_TO_UPLOAD_BLOCK` | detail message | fail to upload block |


---
## Flexicombo API

_The APIs for Flexicombo Tools_

### ActivateFlexiCombo
`POST` `/promotion/flexicombo/activate`

**Description:** activate flexi combo

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code  |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |
| `23` | E023: activate failed | activate failed |


### AddFlexiComboProducts
`POST` `/promotion/flexicombo/products/add`

**Description:** add flexi combo products

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | promotion id |
| `sku_ids` | Number[] | Yes | sku list that will be added to this flexi combo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | sku list that fail to add |
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |


### CreateFlexiCombo
`POST` `/promotion/flexicombo/create`

**Description:** create a  new promotion flexi combo

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `apply` | String | Yes | apply scope: ENTIRE_STORE / SPECIFIC_PRODUCTS |
| `sample_skus` | Object[] | No | sample list |
| `criteria_type` | String | Yes | AMOUNT / QUANTITY |
| `criteria_value` | String[] | Yes | criteria value list |
| `order_numbers` | Number | Yes | orders numbers that can use flexi combo |
| `name` | String | Yes | flexi combo name |
| `platform_channel` | String | No | platform channel, default is 1 |
| `gift_skus` | Object[] | No | gift list |
| `start_time` | Number | Yes | start time |
| `discount_type` | String | Yes | money / discount / freeGift / freeSample /  discountWithGift / moneyWithGift / discountWithSample / moneyWithSample |
| `end_time` | Number | Yes | end time |
| `discount_value` | String[] | Yes | discount value list |
| `stackable` | String | No | Stackable Discount，Ex. Buy 2SGD Save 1SGD, Buy 4SGD Save 2SGD, Buy 6SGD Save 3SGD, etc. |
| `gift_buy_limit_value` | String[] | No | buyer can choose gift/sample quantity limit value list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Number | Yes | flexi combo id |
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |
| `22` | E022: "%s" | validate param error |
| `23` | E023: "%s" | create error |


### DeactivateFlexiCombo
`POST` `/promotion/flexicombo/deactivate`

**Description:** deactivate flexi combo

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |
| `23` | E023: deactivate failed | deactivate failed |


### DeleteFlexiComboProducts
`POST` `/promotion/flexicombo/products/delete`

**Description:** delete flexi combo products

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | id |
| `sku_ids` | Number[] | Yes | sku list that will remove from flexi combo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |
| `23` | E023: remove sku from flexi combo failed | remove sku from flexi combo failed |


### GetFlexiComboDetails
`GET` `/promotion/flexicombo/details`

**Description:** get promotion flexi combo detail by id

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |


### ListFlexiCombo
`GET` `/promotion/flexicombo/list`

**Description:** list flexi combo

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cur_page` | Number | Yes | current page |
| `name` | String | No | name |
| `page_size` | Number | Yes | page size |
| `status` | String | No | NOT_START / ONGOING / SUSPEND / FINISH |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | success |
| `error_code` | String | Yes | error_code |
| `error_msg` | String | Yes | error_msg |
| `data` | Object | Yes | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |


### ListFlexiComboProducts
`GET` `/promotion/flexicombo/products/list`

**Description:** list flexi combo products

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cur_page` | Number | Yes | current page |
| `page_size` | Number | Yes | page size;Maximum value: 100; Minimum value: 10 |
| `id` | Number | Yes | flexi combo id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | data |
| `success` | Boolean | Yes | true/false |
| `error_code` | String | Yes | error_code |
| `error_msg` | String | Yes | error_msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |


### UpdateFlexiCombo
`POST` `/promotion/flexicombo/update`

**Description:** update flexi combo

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `apply` | String | Yes | apply scope: ENTIRE_SHOP / SPECIFIC_PRODUCTS |
| `sample_skus` | Object[] | No | sample list |
| `criteria_type` | String | Yes | AMOUNT / QUANTITY |
| `criteria_value` | String[] | Yes | criteria value list |
| `order_numbers` | Number | Yes | orders numbers that can use flexi combo |
| `name` | String | Yes | flexi combo name |
| `platform_channel` | String | No | platform channel |
| `gift_skus` | Object[] | No | gift list |
| `start_time` | Number | Yes | start time |
| `discount_type` | String | Yes | money / discount / freeGift / freeSample /  discountWithGift / moneyWithGift / discountWithSample / moneyWithSample |
| `id` | Number | Yes | flexi combo id |
| `end_time` | Number | Yes | end time |
| `discount_value` | String[] | Yes | discount value list |
| `stackable` | String | No | Stackable Discount，Ex. Buy 2SGD Save 1SGD, Buy 4SGD Save 2SGD, Buy 6SGD Save 3SGD, etc. |
| `gift_buy_limit_value` | String[] | No | buyer can choose gift/sample quantity limit value list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021: Internal System Error | Internal System Error |
| `22` | E022: "%s" | invalid param |
| `23` | E023: "%s" | update failed |


---
## Seller Voucher API

_for Seller Voucher API_

### SellerVoucheDeleteSelectedProductSKU
`POST` `/promotion/voucher/product/sku/remove`

**Description:** delete seller voucher promotion product sku

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucher_type` | String | Yes | voucher type COLLECTIBLE_VOUCHER / CODE_VOUCHER |
| `id` | Number | Yes | promotion ID |
| `sku_ids` | Number[] | Yes | sku ID list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code  |
| `error_msg` | String | Yes | error message |


### SellerVoucherActivate
`POST` `/promotion/voucher/activate`

**Description:** activate seller voucher promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucher_type` | String | Yes | voucher type COLLECTIBLE_VOUCHER / CODE_VOUCHER |
| `id` | Number | Yes | Promotion ID |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### SellerVoucherAddSelectedProductSKU
`POST` `/promotion/voucher/product/sku/add`

**Description:** add seller voucher promotion product sku

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucher_type` | String | Yes | voucher type COLLECTIBLE_VOUCHER / CODE_VOUCHER |
| `id` | Number | Yes | promotion ID |
| `sku_ids` | Number[] | Yes | sku ID list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | sku list that fail to add |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code  |
| `error_msg` | String | Yes | error message |


### SellerVoucherCreate
`POST` `/promotion/voucher/create`

**Description:** create a new seller voucher promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `criteria_over_money` | String | Yes | Discount details, if order value reaches set value, will money discount or percentage discount |
| `voucher_type` | String | Yes | Voucher type, just set COLLECTIBLE_VOUCHER |
| `apply` | String | Yes | apply scope: ENTIRE_SHOP / SPECIFIC_PRODUCTS |
| `collect_start` | Number | No | The time that customers can collect the voucher |
| `display_area` | String | Yes | The area that customers can see the voucher. REGULAR_CHANNEL/STORE_FOLLOWER/OFFLINE/LIVE_STREAM/CEM_SELLER |
| `period_end_time` | Number | Yes | The period end time that customers can use the voucher |
| `voucher_name` | String | Yes | Voucher name |
| `voucher_discount_type` | String | Yes | Discount type, MONEY_VALUE_OFF / PERCENTAGE_DISCOUNT_OFF  |
| `offering_money_value_off` | String | No | Discount details, if order value reaches criteria_over_money value, will discount money value |
| `period_start_time` | Number | Yes | The period start time that customers can use the voucher |
| `limit` | Number | Yes | Voucher limit per customer |
| `issued` | Number | Yes | Revision should be greater than the current setting |
| `max_discount_offering_money_value` | String | No | Discount details, if order value reaches criteria_over_money value, allow maximum discount per order, just support percentage discount off type |
| `offering_percentage_discount_off` | Number | No | Discount details, if order value reaches criteria_over_money value, will percentage discount off value |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Number | Yes | promotion ID |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `23` | E023: xxxx | business validation error |
| `21` | E023: Internal System Error | Internal System Error |
| `24` | E024: Parameter illegal | Parameter illegal |
| `25` | E025: UMP Exception | UMP Exception |
| `26` | E026: Seller Unauthorized | Seller Unauthorized |


### SellerVoucherDeactivate
`POST` `/promotion/voucher/deactivate`

**Description:** deactivate seller voucher promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucher_type` | String | Yes | voucher type COLLECTIBLE_VOUCHER / CODE_VOUCHER |
| `id` | Number | Yes | Promotion ID |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### SellerVoucherDetailQuery
`GET` `/promotion/voucher/get`

**Description:** get a seller voucher promotion detail

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucher_type` | String | Yes | voucher type COLLECTIBLE_VOUCHER / CODE_VOUCHER |
| `id` | Number | Yes | promotion ID |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code  |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `21` | E021 | The voucher_type field only supports the enumeration COLLECTIBLE_VOUCHER / CODE_VOUCHER. |


### SellerVoucherList
`GET` `/promotion/vouchers/get`

**Description:** query seller voucher promotion list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cur_page` | Number | No | current page |
| `voucher_type` | String | Yes | voucher type COLLECTIBLE_VOUCHER / CODE_VOUCHER |
| `name` | String | No | promotion name |
| `page_size` | Number | No | page size |
| `status` | String | No | NOT_START / ONGOING / SUSPEND / FINISH |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `Mp3SellerApiLimit` | Mp3 Seller not support the api -apipath | MP3 sellers cannot call the current API, please readthis document for a list of APIs that can be called by MP3 sellers,  |


### SellerVoucherSelectedProductList
`GET` `/promotion/voucher/products/get`

**Description:** query seller voucher selected products list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucher_type` | String | Yes | voucher type COLLECTIBLE_VOUCHER / CODE_VOUCHER |
| `id` | Number | Yes | Promotion ID |
| `cur_page` | Number | No | cur page |
| `page_size` | Number | No | page size |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code |
| `error_msg` | String | Yes | error message |


### SellerVoucherUpdate
`POST` `/promotion/voucher/update`

**Description:** update a existing seller voucher promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `max_discount_offering_money_value` | String | No | Discount details, if order value reaches criteria_over_money value, allow maximum discount per order, just support percentage discount off type |
| `offering_percentage_discount_off` | Number | No | Discount details, if order value reaches criteria_over_money value, will percentage discount off value |
| `id` | String | Yes | Promotion ID |
| `criteria_over_money` | String | Yes | Discount details, if order value reaches set value, will money discount or percentage discount |
| `voucher_type` | String | Yes | Voucher type, just set COLLECTIBLE_VOUCHER |
| `apply` | String | Yes | apply scope: ENTIRE_SHOP / SPECIFIC_PRODUCTS |
| `collect_start` | Number | No | The time that customers can collect the voucher |
| `display_area` | String | Yes | The area that customers can see the voucher. |
| `period_end_time` | Number | Yes | The period end time that customers can use the voucher |
| `voucher_name` | String | Yes | Voucher name |
| `voucher_discount_type` | String | Yes | Discount type |
| `offering_money_value_off` | String | Yes | Discount details, if order value reaches criteria_over_money value, will discount money value |
| `period_start_time` | Number | Yes | The period start time that customers can use the voucher |
| `limit` | Number | Yes | Voucher limit per customer |
| `issued` | Number | Yes | Revision should be greater than the current setting |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Number | Yes | promotion ID |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


---
## Free Shipping API

_for Free Shipping API_

### FreeShippingActivate
`POST` `/promotion/freeshipping/activate`

**Description:** activate free shipping promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | promotion id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingAddSelectedProductSKU
`POST` `/promotion/freeshipping/product/sku/add`

**Description:** add sku for free shipping promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | promotion id |
| `sku_ids` | Number[] | Yes | sku id list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | sku list that fail to add |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingCreate
`POST` `/promotion/freeshipping/create`

**Description:** create a new free shipping promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `budget_type` | String | Yes | UNLIMITED_BUDGET / LIMITED_BUDGET |
| `template_type` | String | No | template type, MANUALLY / CAMPAIGN / TEMPLATE |
| `apply` | String | Yes | apply scope: ENTIRE_SHOP / SPECIFIC_PRODUCTS / CAMPAIGN_PRODUCTS |
| `period_end_time` | Number | Yes | when specific period required, the period end time that this promotion takes effect (timestamp) |
| `template_code` | String | No | template code  |
| `category_name` | String | No | product category id |
| `budget_value` | String | No | when limited budget required |
| `promotion_name` | String | Yes | promotion name |
| `period_type` | String | Yes | LONG_TERM / SPECIAL_PERIOD |
| `region_type` | String | Yes | ALL_REGIONS / SPECIAL_REGIONS, when regions query api return empty just support ALL_REGIONS |
| `period_start_time` | Number | Yes | when specific period required, the period start time that this promotion takes effect (timestamp) |
| `campaign_tag` | String | No | when CAMPAIGN template type and CAMPAIGN_PRODUCTS apply type required |
| `region_value` | String[] | No | when SPECIAL_REGIONS  required, data from regions query api  |
| `delivery_option` | String | Yes | data from delivery options query list api |
| `tiers` | Object[] | Yes | promotion tier list |
| `discount_type` | String | Yes | shipping fee subsidy type,FULL_SUBSIDY/PARTIAL_SUBSIDY |
| `deal_criteria` | String | Yes | the criteria that customer can enjoy shipping fee subsidy, MONEY_VALUE_FROM_X/ITEM_QUANTITY_FROM_X/NO_CONDITION |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Number | Yes | promotion ID |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingDeactivate
`POST` `/promotion/freeshipping/deactivate`

**Description:** deactivate free shipping promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | promotion id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingDeleteSelectedProductSKU
`POST` `/promotion/freeshipping/product/sku/remove`

**Description:** delete sku for free shipping promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | promotion id |
| `sku_ids` | Number[] | Yes | sku id list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingDeliveryOptionsQuery
`GET` `/promotion/freeshipping/deliveryoptions/get`

**Description:** query free shipping promotion delivery options

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | response data |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingGet
`GET` `/promotion/freeshipping/get`

**Description:** get free shipping promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `id` | Number | Yes | promotion id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |
| `success` | Boolean | Yes | true / false |
| `error_code` | String | Yes | error code  |
| `error_msg` | String | Yes | error message |


### FreeShippingList
`GET` `/promotion/freeshippings/get`

**Description:** query free shipping promotion list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `curPage` | Number | No | current page |
| `name` | String | No | promotion name |
| `pageSize` | Number | No | page size |
| `status` | String | No | NOT_START / ONGOING / SUSPEND / FINISH |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingRegionsQuery
`GET` `/promotion/freeshipping/regions/get`

**Description:** query free shipping promotion regions

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | response data |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingSelectedProductList
`GET` `/promotion/freeshipping/products/get`

**Description:** query free shipping promotion selected product list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `curPage` | Number | No | current page |
| `pageSize` | Number | No | page size |
| `id` | Number | Yes | promotion id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response data |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


### FreeShippingUpdate
`POST` `/promotion/freeshipping/update`

**Description:** update free shipping promotion

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `budget_type` | String | Yes | UNLIMITED_BUDGET / LIMITED_BUDGET |
| `template_type` | String | Yes | template type, MANUALLY / CAMPAIGN / TEMPLATE |
| `apply` | String | Yes | apply scope: ENTIRE_SHOP / SPECIFIC_PRODUCTS / CAMPAIGN_PRODUCTS |
| `period_end_time` | Number | Yes | when specific period required, the period end time that this promotion takes effect (timestamp) |
| `template_code` | String | No | template code  |
| `category_name` | String | No | product category id |
| `budget_value` | String | No | when limited budget required |
| `promotion_name` | String | Yes | promotion name |
| `period_type` | String | Yes | LONG_TERM / SPECIAL_PERIOD |
| `region_type` | String | Yes | ALL_REGIONS / SPECIAL_REGIONS, when regions query api return empty just support ALL_REGIONS |
| `period_start_time` | Number | Yes | when specific period required, the period start time that this promotion takes effect (timestamp) |
| `campaign_tag` | String | No | when CAMPAIGN template type and CAMPAIGN_PRODUCTS apply type required |
| `region_value` | String[] | No | when SPECIAL_REGIONS  required, data from regions query api  |
| `id` | Number | Yes | promotion id |
| `delivery_option` | String | Yes | data from delivery options query list api |
| `discount_type` | String | Yes | shipping fee subsidy type,FULL_SUBSIDY/PARTIAL_SUBSIDY |
| `deal_criteria` | String | Yes | the criteria that customer can enjoy shipping fee subsidy, MONEY_VALUE_FROM_X/ITEM_QUANTITY_FROM_X/NO_CONDITION |
| `tiers` | Object[] | Yes | promotion tier list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Number | Yes | promotion ID |
| `success` | Boolean | Yes | true / false |
| `error_code` | Number | Yes | error code |
| `error_msg` | String | Yes | error message |


---
## Early Bird Price API

_Set early bird price for new products and get more sales.
_

### CreateEarlyBirdActivityV2
`POST` `/activity/early/bird/create/v2`

**Description:** early bird price activity create 

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `sku_list` | Object[] | Yes | sku list |
| `page_no` | Number | No | page no |
| `name` | String | No | activity name |
| `page_size` | Number | No | page_size |
| `id` | Number | No | activity id |
| `source` | String | No | source |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### EarlyBirdActivityAddSkusV2
`POST` `/activity/early/bird/addSkus/v2`

**Description:** add skus for early bird activity

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `sku_list` | Object[] | Yes | sku list |
| `page_no` | Number | No | page no |
| `name` | String | No | activity name |
| `page_size` | Number | No | page size |
| `id` | Number | Yes | activity id |
| `source` | String | No | source |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### EarlyBirdActivityDeactivateSkusV2
`POST` `/activity/early/bird/deactivateSkus/v2`

**Description:** deactivate Skus for early bird acivity

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `sku_list` | Object[] | Yes | sku list |
| `page_no` | Number | No | page no |
| `name` | String | No | activity name |
| `page_size` | Number | No | page size |
| `id` | Number | Yes | activity id |
| `source` | String | No | source |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### EarlyBirdActivityIsWhitelistSeller
`POST` `/activity/early/bird/isWhitelistSeller`

**Description:** is whitelist seller for early bird acivity

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


---
## Order API

_Order Information_

### GetDocument
`GET` `/order/document/get`

**Description:** Use this API to retrieve order-related documents, including invoices and shipping labels.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `doc_type` | String | Yes | Document types, including 'invoice', 'shippingLabel', or 'carrierManifest'. Mandatory. |
| `order_item_ids` | String | Yes | Identifier of the order item for which the caller wants to get a document. Mandatory. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `20` | E020: "%s" Invalid Order Item IDs | The specified order item ID is not valid. |
| `21` | E021: OMS Api Error Occurred | Internal system error. |
| `32` | E032: Document type "%s" is not valid | The specified document type is not valid. |
| `34` | E034: Order Item must be packed. Please call SetStatusToReadyToShip before | The current status of the order item is not valid. |
| `35` | E035: "%s" was not found |  The specified order item is not found. |
| `30012` | rts package not found | Order item ID status must be "packed" or "ready to ship" |
| `700040` | There are no packages that support printing! | Printing AWB is not supported for orders in Unpaid, pending, canceled status or SOF/DBS orders. |
| `700040` | There are no packages that support printing! | Printing AWB is not supported for orders in Unpaid, pending, canceled status or SOF/DBS orders. |
| `700040` | There are no packages that support printing! | Printing AWB is not supported for orders in Unpaid, pending, canceled status or SOF/DBS orders. |
| `700040` | There are no packages that support printing! | Printing AWB is not supported for orders in Unpaid, pending, canceled status or SOF/DBS orders. |
| `6` | For input string: "" | Make sure you enter an array and not a string in the order_item_ids parameter. |
| `50008` | ot support operation for sof order | SOF/DBS type orders do not support the call of this API to query Shipping label, this type of orders by the seller to co |


### GetMultipleOrderItems
`GET` `/orders/items/get`

**Description:** Use this API to get the item information of one or more orders.（No more than 50 at a time）

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_ids` | Number[] | Yes | Comma-separated list of order identifiers in square brackets.（No more than 50 at a time） |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `37` | E037: One or more order id in the list are incorrect | One or more order IDs specified are not valid. |
| `38` | E038: Too many orders were requested |  The number of orders exceeds the limit.  |
| `39` | E039: No orders were found |  The specified orders are not found. |
| `56` | E056: Invalid OrdersIdList format. Must use array format [1,2] |  The format of the order ID list is not valid. |


### GetOVOOrders
`GET/POST` `/orders/ovo/get`

**Description:** This interface is only applicable to the merchant side of the business and is used to set the maximum number of SKUs that certain merchants can sell per day

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `tradeOrderIds` | String | Yes | id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### GetOrder
`GET` `/order/get`

**Description:** Use this API to get the list of items for a single order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_id` | Number | Yes | The identifier that was assigned to the order by the Seller Center |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `16` | E016: "%s" Invalid Order ID |  The specified order ID is not valid. |
| `6` | E006: System Error | System Error |
| `16` | Invalid Order ID | The order number in the request does not exist in the current store, please call GetOrders API to synchronize the order  |
| `16` | Invalid Order ID | The order number in the request does not exist in the current store, please call GetOrders API to synchronize the order  |
| `16` | Invalid Order ID | The order number in the request does not exist in the current store, please call GetOrders API to synchronize the order  |


### GetOrderItems
`GET` `/order/items/get`

**Description:** Use this API to get the item information of an order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_id` | Number | Yes | The identifier that was assigned to the order by the Seller Center. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `16` | E016: "%s" Invalid Order ID |  The specified order ID is not valid. |
| `6` | E006: System Error | System Error |
| `16` | Invalid Order ID | The order number in the request does not exist in the current store, please call GetOrders API to synchronize the order  |
| `16` | Invalid Order ID | The order number in the request does not exist in the current store, please call GetOrders API to synchronize the order  |
| `16` | Invalid Order ID | The order number in the request does not exist in the current store, please call GetOrders API to synchronize the order  |


### GetOrders
`GET` `/orders/get`

**Description:** Use this API to get the list of items for a range of orders1..

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `update_before` | String | No | Limits the returned orders to those updated before or on the specified date, given in ISO 8601 date format. Optional. |
| `sort_direction` | String | No | Specify the sorting type. Possible values are ASC and DESC. |
| `offset` | Number | No | Number of orders to skip at the beginning of the list. |
| `limit` | Number | No | The maximum number of orders that can be returned. The supported maximum number is 100. |
| `update_after` | String | No | Limits the returned orders to those updated after or on the specified date, given in ISO 8601 date format. Either UpdatedAfter or CreatedAfter is mandatory. |
| `sort_by` | String | No | Allows to choose the sorting column. Possible values are created_at and updated_at. |
| `created_before` | String | No | Limits the returned orders to those updated before or on the specified date, given in ISO 8601 date format. Optional. |
| `created_after` | String | No | Limits the returned orders to those updated after or on the specified date, given in ISO 8601 date format. Either UpdatedAfter or CreatedAfter is mandatory. |
| `status` | String | No | When set, limits the returned set of orders to loose orders, which return only entries which fit the status provided. Possible values are unpaid, pending, canceled, ready_to_ship, delivered, returned, shipped , failed, topack,toship,shipping and lost |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `14` | E014: "%s" Invalid Offset |  The value for the offset parameter is not valid. |
| `17` | E017: "%s" Invalid Date Format |  The date format is not valid. |
| `19` | E019: "%s" Invalid Limit |  The value for the limit parameter is not valid. |
| `36` | E036: Invalid status filter |  The specified status filter is not valid. |
| `74` | E074: Invalid sort direction. |   The specified sort direction is not valid. |
| `75` | E075: Invalid sort filter. |   The specified sort filter is not valid. |
| `SellerNotVerified` | Seller not verified,please check seller status | The seller's store opening process has not been completed, please log in to the Seller Center, check the store informati |
| `SellerNotVerified` | Seller not verified,please check seller status | The seller's store opening process has not been completed, please log in to the Seller Center, check the store informati |
| `6` | Invalid status filter | The status field value is incorrect and only these enumerations are currently supported:unpaid, pending, packed, cancele |
| `17` | Invalid Date Format | The time format used in the request is incorrect, please make sure your time format meets this format requirement: YYYY- |


### OrderCancelValidate
`GET` `/order/reverse/cancel/validate`

**Description:** Seller can check whether the order can be canceled through this API and get corresponding reasons if not.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_id` | String | Yes | order id |
| `order_item_id_list` | String[] | Yes | all order items need to be cancel |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `102` | E0102: trade order line id is empty or invalid | E0102: trade order line id is empty or invalid |
| `106` | E0106: ROC internal error | E0106: ROC internal error |
| `115` | E0115: order id is null | E0115: order id is null |
| `116` | E0116: no seller id | E0116: no seller id |
| `117` | E0117: no user id | E0117: no user id |
| `118` | E0118: no user email | E0118: no user email |
| `122` | E0122: invalid trade order  | E0122: invalid trade order  |
| `123` | E0123: invalid trade order lines %s | E0123: invalid trade order lines %s |
| `124` | E0124: invalid seller id for this order line %s | E0124: invalid seller id for this order line %s |


### SetInvoiceNumber
`POST` `/order/invoice_number/set`

**Description:** Use this API to set the invoice number for the specified order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_item_id` | Number | Yes | Identifier of the order item. |
| `invoice_number` | String | Yes | The invoice number. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `20` | "E020: 59871357123 Invalid Order Item ID" | Order Item ID is incorrect, please verify |
| `34` | "E034: Order Item must be packed. Please call setStatusToReadyToShip before" | Canceled or pending orders are not allowed to call this API |


---
## Return and Refund API

_The APIs to manage return orders_

### GetReverseOrderDetail
`GET` `/order/reverse/return/detail/list`

**Description:** Get the detailed information for a specific reverse order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reverse_order_id` | Number | Yes | 0 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `105` | E0105: reverse order id is empty or invalid | E0105: reverse order id is empty or invalid |
| `106` | E0106: ROC internal error | E0106: ROC internal error |
| `116` | E0116: no seller id | E0116: no seller id |
| `117` | E0117: no user id | E0117: no user id |
| `118` | E0118: no user email | E0118: no user email |
| `Mp3SellerApiLimit` | Mp3 Seller not support the api -apipath | MP3 sellers cannot call the current API, please readthis document for a list of APIs that can be called by MP3 sellers,  |
| `106` | ROC internal error | The reverse ID entered in reverse_order_id does not exist in the current store or is incorrect, call the GetReverseOrder |


### GetReverseOrderHistoryList
`GET` `/order/reverse/return/history/list`

**Description:** Get the communication history of the reverse order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reverse_order_line_id` | Number | Yes | reverse order line id |
| `page_size` | Number | No | default 10 |
| `page_number` | Number | No | default 1 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | {} |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `103` | E0103: reverse order line id is empty when query reject reason | E0103: reverse order line id is empty when query reject reason |
| `106` | E0106: ROC internal error | E0106: ROC internal error |
| `116` | E0116: no seller id | E0116: no seller id |
| `117` | E0117: no user id | E0117: no user id |
| `118` | E0118: no user email | E0118: no user email |
| `120` | E0120: page size invalid | E0120: page size invalid |
| `121` | E0121: page number invalid | E0121: page number invalid |


### GetReverseOrderReasonList
`GET` `/order/reverse/reason/list`

**Description:** Get the list of reject reason. Need to be used in all refuse refund actions

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reverse_order_line_id` | Number | Yes | reverse order line,Can be understood as reverse order item id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `103` | E0103: reverse order line id is empty when query reject reason | E0103: reverse order line id is empty when query reject reason |
| `106` | E0106: ROC internal error | E0106: ROC internal error |
| `116` | E0116: no seller id | E0116: no seller id |
| `117` | E0117: no user id | E0117: no user id |
| `118` | E0118: no user email | E0118: no user email |
| `119` | E0119: cannot find any cancel reasons for these orders | E0119: cannot find any cancel reasons for these orders |


### GetReverseOrdersForSeller
`GET/POST` `/reverse/getreverseordersforseller`

**Description:** Use this API to get the list of items for a range of reverse orders.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `request_type_list` | String[] | No | request type |
| `ofc_status_list` | String[] | No | Limit the ofc status |
| `reverse_order_id` | Number | No | Specify reverse order id |
| `trade_order_id` | Number | No | Specify trade order id |
| `page_size` | Number | Yes | Page size, default 10 |
| `reverse_status_list` | String[] | No | Limit the reverse status. |
| `page_no` | Number | Yes | Page no |
| `return_to_type` | String | No | Return Type. Enum Values：[RTM, RTW]（ RTW: return to the lazada warehouse; RTM: return to the seller） |
| `dispute_in_progress` | Boolean | No | Is dispute in progress |
| `TradeOrderLineCreatedTimeRangeStart` | Number | No | timestamp in Milliseconds |
| `TradeOrderLineCreatedTimeRangeEnd` | Number | No | timestamp in Milliseconds |
| `ReverseOrderLineTimeRangeStart` | Number | No | timestamp in Milliseconds |
| `ReverseOrderLineTimeRangeEnd` | Number | No | timestamp in Milliseconds |
| `ReverseOrderLineModifiedTimeRangeStart` | Number | No | timestamp in Milliseconds |
| `ReverseOrderLineModifiedTimeRangeEnd` | Number | No | timestamp in Milliseconds |
| `QC_Decision` | String | No | warehouse qc decision, select one from the following: scrap/return_to_merchant/return_to_merchant_cb/return_to_customer/return_to_warehouse/not_returned |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `Mp3SellerApiLimit` | Mp3 Seller not support the api - apipath | MP3 sellers cannot call the current API, please readthis document for a list of APIs that can be called by MP3 sellers,  |


### InitReverseOrderCancel
`GET` `/order/reverse/cancel/create`

**Description:** Seller initiates a cancelation

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_item_id_list` | String[] | Yes | all order items need to be cancel |
| `order_id` | Number | Yes | order id |
| `reason_id` | String | Yes | reason id  |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `102` | E0102: trade order line id is empty or invalid | E0102: trade order line id is empty or invalid |
| `104` | E0104: reason is empty or invalid | E0104: reason is empty or invalid |
| `106` | E0106: ROC internal error | E0106: ROC internal error |
| `115` | E0115: order id is null | E0115: order id is null |
| `116` | E0116: no seller id | E0116: no seller id |
| `117` | E0117: no user id | E0117: no user id |
| `118` | E0118: no user email | E0118: no user email |
| `122` | E0122: invalid trade order | E0122: invalid trade order |
| `123` | E0123: invalid trade order lines %s | E0123: invalid trade order lines %s |
| `124` | E0124: invalid seller id for this order line %s | E0124: invalid seller id for this order line %s |


### InitReverseOrderCancelDecide
`GET` `/order/reverse/cancel/seller/decide`

**Description:** Seller initiates a cancelation

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reverse_order_id` | Number | Yes | The reverse order to be cancelled |
| `agree_cancel` | Boolean | Yes | decision |
| `reason_code` | Number | No | reason id  |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | null |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `116` | E0116: no seller id | E0116: no seller id |
| `105` | E0105: reverse order id is empty or invalid | E0105: reverse order id is empty or invalid |
| `131` | E0131: no decision for this reverse order | E0131: no decision for this reverse order |
| `106` | E0106: ROC internal error | E0106: ROC internal error |


### ReverseOrderOnlyRefundDecide
`GET` `/order/reverse/onlyrefund/seller/decide`

**Description:** Seller can use this API to operate only refund requests

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `action` | String | Yes | agreeRefund, startDispute |
| `reverse_order_id` | Number | Yes | reverse order id |
| `reverse_order_item_ids` | Number[] | Yes | reverse order item id list, currently list size can be only 1 |
| `comment` | String | No | comment, required if action is startDispute |
| `image_info_list` | Object[] | No | image info list, required if action is startDispute |
| `video_info_list` | Object[] | No | video info list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | null |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `116` | E0116: no seller id | E0116: no seller id |
| `118` | E0108: reason can't be empty if you want to refuse return or refund | E0108: reason can't be empty if you want to refuse return or refund |
| `100` | E0100: reverse order list is empty | E0100: reverse order list is empty |
| `125` | E0125: invalid reverse id | E0125: invalid reverse id |
| `112` | E0112: no reverse order found | E0112: no reverse order found |
| `133` | E0133: do not support batch operation | E0133: do not support batch operation |
| `126` | E0126: invalid reverse order lines | E0126: invalid reverse order lines |
| `114` | E0114: this reverse does not support this action | E0114: this reverse does not support this action |
| `107` | E0107: invalid action | E0107: invalid action |
| `109` | E0109: comment can't be empty if startDispute | E0109: comment can't be empty if startDispute |
| `110` | E0110: image can't be empty if startDispute | E0110: image can't be empty if startDispute |
| `106` | E0106: ROC internal error | E0106: ROC internal error |
| `113` | E0113: reverse order line have unknown status | E0113: reverse order line have unknown status |
| `114` | E0114: this reverse does not support this action | E0114: this reverse does not support this action |


### ReverseOrderReturnUpdate
`GET` `/order/reverse/return/update`

**Description:** Seller can use this API to action on return and refund related.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `action` | String | Yes | instantRefund;agreeReturn;refuseReturn;agreeRefund;refuseRefund;confirmDelivery |
| `reverse_order_id` | Number | Yes | reverse order id |
| `reverse_order_item_ids` | Number[] | Yes | reverse order item id list |
| `reason_id` | Number | No | reason id |
| `comment` | String | No | comment |
| `image_info` | Object[] | No | image_info |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `100` | E0100: reverse order list is empty | E0100: reverse order list is empty |
| `106` | E0106: ROC internal error | E0106: ROC internal error |
| `107` | E0107: invalid action | E0107: invalid action |
| `108` | E0108: reason can't be empty if you want to refuse return or refund | E0108: reason can't be empty if you want to refuse return or refund |
| `109` | E0109: comment can't be empty if you want to refuse return or refund | E0109: comment can't be empty if you want to refuse return or refund |
| `110` | E0110: image can't be empty if you want to refuse refund | E0110: image can't be empty if you want to refuse refund |
| `111` | E0111: do not support massive reverse order line operation if you want to refuse return or refund | E0111: do not support massive reverse order line operation if you want to refuse return or refund |
| `112` | E0112: no reverse order found | E0112: no reverse order found |
| `113` | E0113: reverse order line have unknown status | E0113: reverse order line have unknown status |
| `114` | E0114: this reverse does not support this action | E0114: this reverse does not support this action |
| `116` | E0116: no seller id | E0116: no seller id |
| `117` | E0117: no user id | E0117: no user id |
| `118` | E0118: no user email | E0118: no user email |
| `125` | E0125: invalid reverse id | E0125: invalid reverse id |
| `126` | E0126: invalid reverse order lines | E0126: invalid reverse order lines |
| `127` | E0127: invalid seller id for this reverse order line | E0127: invalid seller id for this reverse order line |


---
## Fulfillment API

### ConfirmCollectForDBS
`POST` `/order/package/sof/collect`

**Description:** Use this API to mark an sof order item as being collected.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `dbsCollectReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


### ConfirmDeliveryForDBS
`POST` `/order/package/sof/delivered`

**Description:** Use this API to mark an sof order item as being delivered.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `dbsDeliveryReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


### DeliverDigital
`POST` `/order/digital/delivered`

**Description:** Use this API to mark a digital order item as being delivered.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `digitalDeliveryReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


### FailedDeliveryForDBS
`POST` `/order/package/sof/failed_delivery`

**Description:** Use this API to mark an sof order item as being delivered failed

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `dbsFailedDeliveryReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


### GetShipmentProvider
`GET/POST` `/order/shipment/providers/get`

**Description:** Use this API to get the list of all active shipping providers, which is needed when working with the PackOrder API.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `getShipmentProvidersReq` | Object | Yes | req body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


### Pack
`POST` `/order/fulfill/pack`

**Description:** Use this API to mark an order item as being packed.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `6` | SYSTEM_ERROR | system is busy now ,pls try later |
| `1003` | E1003_3PL_ALLOCATION_FAIL | 3pl allocation failed |
| `40011` | RPC_ERROR | system is busy now ,pls try later |
| `700000` | PACKAGE_STATUS_NOT_ALLOW_TO_OP | current package status not allow to operation |
| `700001` | DBS_SHIPMENT_PROVIDER_CODE_NOT_EXITS | shipment provider code not exits |
| `700004` |  PARAM_ILLEGAL | param illegal |
| `700013` | OP_NOT_SUPPORT | operation is no support |
| `700016` | NOT_AVAILABLE_NTFS_3PL | seller not available 3pl , pls contact us to subscription 3pl |
| `700017` | PARAM_IS_NULL | param can't be null |
| `700018` |  PARAM_SIZE_ERROR | param size not match |
| `700019` | PARAM_MIN_ERROR | param min not match |
| `700020` | ORDER_ITEM_NOT_FOUND_OR_NOT_BELONG_DIGITAL | order item not found or not belong to digital |
| `700021` |  ORDER_NOT_FOUND | order not found |
| `700022` | BATCH_SIZE_OUT_OF_LIMIT | batch size out of limit |
| `700023` |  PICKUP_IN_STORE_NO_SUPPORT | pickup in store order no allow to operation |
| `700024` | GET_LOCK_FAILED | failed get lock,pls try later |
| `700025` | ORDER_ITEM_NOT_FOUND | order item not found |
| `700026` | FO_ITEM_NOT_ALLOW_TO_PACK | item current status not allow to pack |
| `700027` | NOT_SUPPORT_FBL_TO_PACK | Does not support FBL order to pack |
| `700028` | NOT_SUPPORT_PACK_UP_IN_STORE_TO_PACK | Does not support pickup_in_store order to pack |
| `700029` |  ITEM_MUST_BELONG_SAME_WAREHOUSE | item must belong same warehouse |
| `700030` |  NOT_SUPPORT_DG_SERVICE_TO_PACK | digital or service order not need  to pack |
| `700031` |  ITEM_NOT_READY_TO_FULFILL | item not ready to fulfill |
| `700032` | SELLER_NOT_FOUND | can't found seller |
| `700033` | TRANSFERRING_WAREHOUSE_PROVIDER | transferringWarehouseCode cannot found |


### PackageStatusUpdateForDBS
`POST` `/order/package/sof/status/update`

**Description:** DBS package status update.
This interface is only open to some stores

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `trackingNumber` | String | Yes | waybill no |
| `source` | String | Yes | OPENAPI |
| `carrierCode` | String | No | SF |
| `tag` | String | Yes | package no |
| `trackInfo` | Object | Yes | track info |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | api result |
| `module` | Object | No | content |
| `errorCode` | Object | No | error msesage |


### PrintAWB
`GET/POST` `/order/package/document/get`

**Description:** Use this API to retrieve order-related documents, only for shipping labels.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `getDocumentReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


### ReadyToShip
`POST` `/order/package/rts`

**Description:** Use this API to mark an order item as being ready to ship.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `readyToShipReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


### RecreatePackage
`GET/POST` `/order/package/repack`

**Description:** Use this API to mark a package item as being repack.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `rePackReq` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | resp body |


---
## Logistics API

_Logistics_

### AddOrUpdatePickupStop
`POST` `/logistics/tps/runsheets/stops`

**Description:** 3PL call TPS to update pickup stops

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stopId` | String | Yes | Stop ID |
| `sellerId` | String | Yes | Seller ID (Sent in pickup request) |
| `warehouseCode` | String | Yes | Warehouse code (Sent in pickup request) |
| `dopStationId` | String | No | DOP station code |
| `dopStationName` | String | No | DOP station name |
| `pickupType` | String | Yes | Type: Pickup/Drop-off |
| `status` | String | Yes | 1. planned: when stop is dispatched to courier\n 2. arrived: when driver arrived at stop and start pickup\n 3. finished: when driver finished pickup at the stop\n 4. skipped: when driver selected to skip the stop due to some reason\n 5. removed: when the stop has 0 RTS |
| `statusUpdateTime` | Number | Yes | actual process time when reaching the status |
| `dispatcherName` | String | No | Dispatcher name |
| `dispatcherContact` | String | No | Dispatcher phone number |
| `driverId` | String | No | Driver ID |
| `driverName` | String | Yes | Driver name |
| `driverContact` | String | No | Driver phone number |
| `eta` | Number | No | when the ETA is updated, need to update the data to Lazada side, scenario include:   1. without ETA  >> with ETA   2. with ETA >> without ETA   3. ETA change from A to B  |
| `successVolume` | String | No | Success count |
| `failedVolume` | String | No | Failed count |
| `failedVolumeList` | Object[] | No | Failed list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `success` | Boolean | Yes | Is success? |
| `errors` | Object[] | Yes | Error detail |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | No | Error code |


### Create3PLStation
`POST` `/logistics/tps/stations/create`

**Description:** TPS_CREATE_STATION_API
External partner call TPS to create station

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `externalCode` | String | Yes | Station code in 3PL system |
| `modifier` | String | No | Modifier name. if blank will use 3PL name |
| `name` | String | Yes | Station name in 3PL system |
| `functionCodes` | String[] | Yes | Station functions |
| `subTypes` | String[] | Yes | Y	Station subtypes (depends on function) enum:  DOP function: MDOP, DOP, OTC,IDOP CP function: COLLECTION_ON_POINT, MOBILE_COLLECTION_POINT, LOCKER Return function: CUSTOMER_RETURN |
| `codSupport` | Boolean | Yes | Support COD or not |
| `age` | Number | No | Number of days the station can keep packages for (used by LOP station tool). If not withdrawn by the customer within the age value, the package will be picked up from the station by a dedicated 3PL and brought to the warehouse. The package will be marked as failed delivery. Unit: Days |
| `firstMileTplSlugs` | String[] | Yes | Which 3PL can go and pick up the seller dropped-off parcel from the station  |
| `lastMileTplSlugs` | String[] | Yes | This is a list of logistics providers which can deliver packages to this station.  |
| `contact` | Object | Yes | Station contact information |
| `address` | Object | Yes | Station address |
| `timeZone` | String | No | Timezone (used to calculate the schedules) If not specified, use default country timezone format: (+/-)XX:XX |
| `schedules` | Object[] | No | Station schedules |
| `constraints` | Object[] | No | Function constraint |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `retryable` | Boolean | No | Is failed request retryable? |
| `errorMessage` | String | No | Error message |
| `errorCode` | String | No | Error code |
| `errors` | Object[] | No | Error list |


### GetOrderTrace
`GET/POST` `/logistic/order/trace`

**Description:** Query logistic detail for seller erp with seller id, order id and locale info. This api is only available in the state after ready to ship.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_id` | String | Yes | order id |
| `locale` | String | No | local |
| `ofcPackageIdList` | String[] | No | package id list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Result |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `INPUT_PARAM_VALID` | query trade failed | order id or ofcPackageIdList invalid |
| `LD_INVOKE_DOWNSTREAM_RESPONSE_BLANK` | LD_INVOKE_DOWNSTREAM_RESPONSE_BLANK | This order does not exist in the current country or store, please call the GetOrders API to check if you have entered th |
| `Dropshipping invalid` | input orderId: Own Warehouse invalid | The input parameters are incorrect, please check that the package id you entered in the ofcPackageIdList field is correc |
| `LD_INPUT_PARAM_VALID` | orderId is wrong | The order number does not exist in the current store or is incorrect, please check if the order number input format in y |


### ScanParcel
`GET/POST` `/dop/scan`

**Description:** DOP Scan Parcel

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cageNumber` | String | Yes | test |
| `trackingNumber` | String | Yes | test |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `trackingNumber` | String | No | test |


### StationDopScan
`GET/POST` `/stations/dop/scan`

**Description:** StationDopScan

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cageNumber` | String | Yes | test |
| `trackingNumber` | String | Yes | test |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | test |
| `data` | Object | No | test |
| `error` | Object | No | test |


### Update3PLStation
`POST` `/logistics/tps/stations/update`

**Description:** TPS_UPDATE_STATION_API
External partner call TPS to update station

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `externalCode` | String | Yes | Station code in 3PL system |
| `modifier` | String | No | Modifier name. if blank will use 3PL name |
| `enable` | Boolean | Yes | Enable or disable Station |
| `functionCodes` | String[] | Yes | Station functions |
| `subTypes` | String[] | Yes | Y	Station subtypes (depends on function) enum:  DOP function: MDOP, DOP, OTC,IDOP CP function: COLLECTION_ON_POINT, MOBILE_COLLECTION_POINT, LOCKER Return function: CUSTOMER_RETURN |
| `codSupport` | Boolean | Yes | Support COD or not |
| `age` | Number | No | Number of days the station can keep packages for (used by LOP station tool). If not withdrawn by the customer within the age value, the package will be picked up from the station by a dedicated 3PL and brought to the warehouse. The package will be marked as failed delivery. Unit: Days |
| `firstMileTplSlugs` | String[] | Yes | Which 3PL can go and pick up the seller dropped-off parcel from the station  |
| `lastMileTplSlugs` | String[] | Yes | This is a list of logistics providers which can deliver packages to this station.  |
| `contact` | Object | Yes | Station contact information |
| `address` | Object | Yes | Station address |
| `timeZone` | String | No | Timezone (used to calculate the schedules) If not specified, use default country timezone format: (+/-)XX:XX |
| `schedules` | Object[] | No | Station schedules |
| `constraints` | Object[] | No | Function constraint |
| `name` | String | No | Station name |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `retryable` | Boolean | No | Is failed request retryable? |
| `errorMessage` | String | No | Error message |
| `errorCode` | String | No | Error code |
| `errors` | Object[] | No | Error list |


### UpdatePickupTimeSlot
`POST` `/logistics/tps/sellers/pickup_timeslot`

**Description:** 3PL call TPS to update pickup timeslot

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `sellerId` | String | Yes | Seller ID (Sent in pickup request) |
| `warehouseCode` | String | Yes | Warehouse code (Sent in pickup request) |
| `pickupTimeslots` | String[] | Yes | Format: HH:mm, separate by comma |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `success` | Boolean | Yes | Is success? |
| `errors` | Object[] | Yes | Error detail |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | No | Error code |


### createConsolidationService
`POST` `/logistics/ldp/createConsolidationService`

**Description:** create Consolidation Service

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `unitCodes` | String[] | Yes | unit codes |
| `properties` | Object | Yes | prop |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | Yes | data |
| `success` | Boolean | Yes | is success |
| `errorCode` | String | Yes | error code |
| `errorMsg` | String | Yes | error Msg |


### updateLastMile
`POST` `/logistics/ldp/updateLastmile`

**Description:** 跨境场景，物流末端预报信息

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `unitCode` | String | Yes | unitCode |
| `shippingProviderCode` | String | Yes | shippingProviderCode |
| `trackingNumber` | String | Yes | trackingNumber |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | is success |
| `data` | String | No | data |
| `errorCode` | String | No | errorCode |
| `errorMsg` | String | No | errorMsg |


---
## FirstMile Bigbag(only for CN)

_大包接口_

### GetChannelcodeByFirstMileNo
`GET/POST` `/logistics/cngfc/fulfill/getchannelcode`

**Description:** get channelcode by first mile No

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `firstMileNos` | String[] | Yes | 首公里面单号 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | Token过期或输入有误 |


### GetLazadaBigbagPDFLable
`GET/POST` `/logistics/cnpms/bigbag/lable/getPdf`

**Description:** Get Lazada Bigbag PDF Lable

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userInfo` | Object | Yes | 用户信息 |
| `client` | String | Yes | ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称 |
| `orderCode` | String | No | 大包单号，即大包LP号，同handoverContentCode |
| `remark` | String | No | 备注 |
| `locale` | String | No | 多语言，默认zh_CN |
| `trackingNumber` | String | No | 大包运单号 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `P-088-0101-10-10-192` | query across account relation not found | 跨店铺组包授权关系不存在 |
| `P-088-0000-00-15-209` | handover content not found | 未找到指定的大包 |


### LazadaBigbagCancel
`GET/POST` `/logistics/cnpms/bigbag/cancel`

**Description:** Lazada Bigbag cancel

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userInfo` | Object | Yes | 用户信息 |
| `client` | String | Yes | ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称 |
| `orderCode` | String | No | 大包单号，即大包LP号，同handoverContentCode，orderCode、trackingNumber二者选其一 |
| `remark` | String | No | 备注 |
| `locale` | String | No | 多语言，默认zh_CN |
| `trackingNumber` | String | No | 大包运单号，orderCode、trackingNumber二者选其一 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `P-088-0101-10-10-191` | query across store account not found | 跨店铺组包账号不存在 |
| `P-088-0000-00-15-209` | handover content not found | 未找到指定的大包 |
| `P-088-0000-00-15-209` | handover content is not found | trackingNumber输入无效 |


### LazadaBigbagCollectionPoints
`GET/POST` `/logistics/cnpms/bigbag/querycollection`

**Description:** Lazada bigbag query collection points

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `pageSize` | String | No | 每页N条 |
| `currentPage` | String | No | 当前第N页 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | Token过期或输入有误 |


### LazadaBigbagCommit
`GET/POST` `/logistics/cnpms/bigbag/commit`

**Description:** Lazada bigbag commit

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userInfo` | Object | Yes | Lazada开放平台信息 |
| `orderCodeList` | String[] | Yes | 要创建交接单的小包编码集合，数量上限1000 |
| `weight` | String | Yes | 重量 |
| `client` | String | Yes | ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称 |
| `collectionInfo` | Object | No | 集货点信息 |
| `remark` | String | No | 备注 |
| `pickupInfo` | Object | Yes | 揽收信息 |
| `locale` | String | No | 多语言，默认zh_CN |
| `weightUnit` | String | Yes | 重量单位，克:g, 千克:kg，默认g |
| `type` | String | Yes | 类型：cainiao_pickup(菜鸟揽收)、self_post(自寄)、pickup_collection(集货) |
| `sellerTrackingNumber` | String | No | 商家定义的大包标签号，一般不传，需要将自有大包号作为菜鸟面单号时才传 |
| `returnInfo` | Object | Yes | 退件信息 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `P-088-0101-10-10-140` | all parcel order not found | 选择的所有小包都找不到，请核对后重试 |
| `P-088-0101-10-10-191` | query across store account not found | 跨店铺组包账号不存在 |
| `P-088-0000-00-15-170` | seller has stores that are not packaged across stores | 商家存在未跨店铺组包的店铺 |
| `InvalidParameter` | The specified parameter “null#addressId” is not valid | addressId是必填的 |
| `UnknownRuntimeException` | The request has failed due to RPC runtime failure | weight需要填整数 |
| `P-088-0000-00-15-231` | pick up collection point info missing | pickup_collection的条件下pickUpCode必填 |
| `P-088-0000-00-15-205` | param is null | self_post的条件下sellerTrackingNumber必填 |


### LazadaBigbagUpdate
`GET/POST` `/logistics/cnpms/bigbag/update`

**Description:** Lazada bigbag update

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userInfo` | Object | Yes | 用户信息 |
| `weight` | Number | Yes | 重量 |
| `locale` | String | No | 多语言，默认zh_CN |
| `orderCodeList` | String[] | Yes | 要创建交接单的小包编码集合，数量上限300 |
| `client` | String | Yes | ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称 |
| `orderCode` | String | No | 大包单号，即大包LP号，orderCode、trackingNumber二者选其一 |
| `trackingNumber` | String | No | 大包运单号，orderCode、trackingNumber二者选其一 |
| `weightUnit` | String | Yes | 重量单位，克:g, 千克:kg，默认g |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `P-088-0000-00-15-300` | handover content status not committed、awaiting_tracking_number or awaiting_pickup, can not update | 大包状态非已提交、等待分配运单号、待揽收，不能更新该大包 |
| `P-088-0000-00-15-209` | handover content not found | 未找到指定的大包 |
| `P-088-0101-10-10-140` | all parcel order not found | 选择的所有小包都找不到，请核对后重试 |
| `P-088-0101-10-10-191` | query across store account not found | 跨店铺组包账号不存在 |


### LazadaSellerAccountBind
`GET/POST` `/logistics/cnpms/account/bind`

**Description:** Lazada seller account bind for big bag pick up

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userInfo` | Object | Yes | 用户信息 |
| `client` | String | No | ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称 |
| `remark` | String | No | 备注 |
| `sellerList` | Object[] | Yes | 授权商家列表，最多一次传50 |
| `locale` | String | No | 多语言，默认zh_CN |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `P-088-0000-00-15-195` | query lzd merchant seller not found | 店铺信息不存在 |


### QueryAddressInformaiton
`GET/POST` `/logistics/cnpms/address/query`

**Description:** Query Address Informaiton

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `country` | String | Yes | 国家 |
| `zipCode` | String | No | 邮编 |
| `userInfo` | Object | Yes | 用户信息 |
| `city` | String | Yes | 市 |
| `remark` | String | No | 备注 |
| `locale` | String | No | 多语言，默认zh_CN |
| `province` | String | Yes | 省 |
| `street` | String | Yes | 街道 |
| `district` | String | Yes | 区/县 |
| `detailAddress` | String | Yes | 详细地址 |
| `client` | String | No | ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `P-088-0101-10-10-152` | address service result error | 地址解析失败，只能输入中国大陆地区 |
| `P-088-0000-00-15-213` | param country is null | 请输入中文地址，英文无效 |
| `P-088-0000-00-15-214` | param province is null | 请输入中文地址，英文无效 |
| `P-088-0000-00-15-215` | param city is null | 请输入中文地址，英文无效 |
| `P-088-0000-00-15-216` | param detailAddress is null | 请输入中文地址，英文无效 |
| `P-088-0000-00-15-217` | param country is not support | 请输入中文地址，英文无效 |
| `P-088-0000-00-15-218` | params is null | 请输入中文地址，英文无效 |


### QueryLazadaBigbagInfo
`GET/POST` `/logistics/cnpms/bigbag/query`

**Description:** Query Lazada Bigbag Info

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userInfo` | Object | Yes | 用户信息 |
| `client` | String | Yes | ISV名称，ISV：ISV-ISV英文或拼音名称、商家ERP：SELLER-商家英文或拼音名称 |
| `orderCode` | String | No | 大包单号，即大包LP号，同handoverContentCode，orderCode、trackingNumber二者选其一 |
| `remark` | String | No | 备注 |
| `locale` | String | No | 多语言，默认zh_CN |
| `trackingNumber` | String | No | 大包运单号，orderCode、trackingNumber二者选其一 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 同步响应结果 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `P-088-0101-10-10-191` | query across store account not found | 跨店铺组包账号不存在 |


---
## Finance API

_Finance_

### GetPayoutStatus
`GET` `/finance/payout/status/get`

**Description:** Get your transaction statements  created after the provided date

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `created_after` | String | Yes | Filter statements created after the provided date. Mandatory. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### QueryAccountTransactions
`POST` `/finance/transaction/accountTransactions/query`

**Description:** Query Account Transactions

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transaction_type` | String | No | transaction type,Enumeration values for(Deposit,Withdrawal,Payment,null) |
| `sub_transaction_type` | String | No | sub transaction type,Enumeration values for(Settlement,Failed Payment,Returned Payment,Auto Withdrawal,Manual Withdrawal,Sponsored Solutions Top-up,null) |
| `transaction_number` | String | No | transaction number |
| `page_size` | Number | Yes | page size |
| `start_time` | String | Yes | start time,format:yyyyMMdd |
| `end_time` | String | Yes | start time,format:yyyyMMdd |
| `page_num` | Number | Yes | page number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `msg` | String | Yes | error message |
| `data` | Object | Yes | result |
| `success` | Boolean | Yes | success:true,fail:false |
| `error_code` | String | Yes | error code  |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `IllegalAccessToken` | The specified access token is invalid or expired | access token is invalid or expired |


### QueryLogisticsFeeDetail
`GET/POST` `/lbs/slb/queryLogisticsFeeDetail`

**Description:** Api is provided for finance and seller to query logistics fee details from slb.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `seller_id` | String | Yes | identity of seller which should not be blank |
| `request_type` | String | Yes | type of request which is used to distinguish different systems(e.g. OPEN_API) |
| `trade_order_id` | String | No | identity of trade order |
| `trade_order_line_id` | String | No | item identity of trade order |
| `fee_type` | String | No | type of logistics fee |
| `biz_flow_type` | String | No | corresponding settlement scenario of request(e.g. LAZADA, LAZADA_3PV, default biz flow type is LAZADA) |
| `bill_start_time` | Number | No | timestamp of the time that bill started |
| `bill_end_time` | Number | No | timestamp of the time that bill ended |
| `page_no` | Number | No | number of page which default 1 |
| `page_size` | Number | No | size of page which default 20 |
| `total_records` | Number | No | total records that page included |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | response body |
| `success` | Boolean | Yes | response is success or not |
| `remark` | String | Yes | remark of response |


### QueryTransactionDetails
`GET` `/finance/transaction/details/get`

**Description:** API to query seller transaction details within specific date range.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `offset` | String | No | Number of transaction lines to skip at the beginning of the list. |
| `trans_type` | String | No | Transaction type ID. |
| `trade_order_id` | String | No | Order ID. |
| `limit` | String | No | Number of lines of transactions to be extracted. The supported maximum number is 500. |
| `start_time` | String | Yes | Starting date when transactions need to be extracted. |
| `end_time` | String | Yes | Ending date when transactions need to be extracted. |
| `trade_order_line_id` | String | No | Order Item ID. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `1000012` | endTime - startTime must should be less than 180 days | endTime - startTime must should be less than 180 days |
| `1000014` | Can not find that transactionType |  transaction type invalid |
| `1000012` | endTime - startTime must should be less than 180 days | Please make sure that the timeframe of your inquiry is within 180 days. |


---
## Membership API

_Membership APIs for Loyalty and User Account related services_

### GetLinkMember
`GET/POST` `/membership/linkmember/get`

**Description:** Query the linkmember relationship between buyers and sellers.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `seller_id` | String | Yes | seller id |
| `buyer_id` | String | Yes | buyer id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `LZD_MEMBER_USER_1011` | LZD_MEMBER_USER_1011 | The buyer id does not exist, call the PartnerTransaction API to query if you are using the correct buyer id. |


### GetLinkMember
`GET/POST` `/partner/get`

**Description:** Query the linkmember relationship between buyers and sellers.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `buyer_id` | String | Yes | buyer id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `LZD_MEMBER_USER_1011` | LZD_MEMBER_USER_1011 | The buyer id does not exist, call the PartnerTransaction API to query if you are using the correct buyer id. |


### GetLinkMemberList
`GET/POST` `/membership/linkmember/list`

**Description:** Query all linkmembers of the seller

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `page_num` | String | Yes | page number |
| `page_size` | String | Yes | page size |
| `seller_id` | String | Yes | seller id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### GetLinkMemberList
`GET/POST` `/partner/list`

**Description:** Query all linkmembers of the seller

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `page_num` | String | Yes | page number |
| `page_size` | String | Yes | page size |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### LinkMembership
`POST` `/membership/link`

**Description:** Used to push a new membership to Lazada for proactively linking memberships.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `p_uid` | String | Yes | A unique identifier of the member on partner side, generated and stored at partner side, that identifies that member and will be referenced by Lazada in further communications. |
| `member_name` | String | No | Name of member on partner side, to easier identify the membership on My Account pages |
| `tier` | String | No | Customer’s tier in partner side, shown as-is |
| `tier_expiry` | String | No | Expiry of the membership, shown as-is |
| `balance` | Number | No | Balance of the membership. |
| `valid_from` | String | No | Valid from of this balance in RFC RFC3339 format. Ignore if this is no validity period for the balance |
| `valid_to` | String | No | Valid to of this balance in RFC RFC3339 format. Ignore if this is no validity period for the balance |
| `linking_token` | String | Yes | Linking token. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | 1 |


### PartnerLink
`GET/POST` `/partner/link`

**Description:** Used to push a new membership to Lazada for proactively linking memberships.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `member_name` | String | No | Name of member on partner side, to easier identify the membership on My Account pages |
| `valid_from` | String | No | Valid from of this balance in RFC RFC3339 format. Ignore if this is no validity period for the balance |
| `linking_token` | String | Yes | Linking token. |
| `tier` | String | No | Customer’s tier in partner side, shown as-is |
| `balance` | Number | No | Balance of the membership. |
| `tier_expiry` | String | No | Expiry of the membership, shown as-is |
| `p_uid` | String | Yes | A unique identifier of the member on partner side, generated and stored at partner side, that identifies that member and will be referenced by Lazada in further communications. |
| `valid_to` | String | No | Valid to of this balance in RFC RFC3339 format. Ignore if this is no validity period for the balance |
| `from_source` | String | No | Where does this user come from.LAZADA or PARTNER |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | api result |


### PartnerTransaction
`GET/POST` `/partner/transaction`

**Description:** Using this interface, you can obtain the seller's transaction order based on the conditions, and also contain the membership information

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `status` | String | No | When set, limits the returned set of orders to loose orders, which return only entries which fit the status provided. Possible values are unpaid, pending, canceled, ready_to_ship, delivered, returned, shipped , failed, topack,toship,shipping and lost |
| `update_before` | String | No | Limits the returned orders to those updated before or on the specified date, given in ISO 8601 date format. Optional. |
| `sort_direction` | String | No | Specify the sorting type. Possible values are ASC and DESC. |
| `offset` | Number | No | Number of orders to skip at the beginning of the list. |
| `limit` | Number | No | The maximum number of orders that can be returned. The supported maximum number is 100. |
| `update_after` | String | No | Limits the returned orders to those updated after or on the specified date, given in ISO 8601 date format. Either UpdatedAfter or CreatedAfter is mandatory. |
| `sort_by` | String | No | Allows to choose the sorting column. Possible values are created_at and updated_at. |
| `created_before` | String | No | Limits the returned orders to those updated before or on the specified date, given in ISO 8601 date format. Optional. |
| `created_after` | String | No | Limits the returned orders to those updated after or on the specified date, given in ISO 8601 date format. Either UpdatedAfter or CreatedAfter is mandatory. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### PartnerUnlink
`GET/POST` `/partner/unlink`

**Description:** Used to remove a linked membership from Lazada. Please note that the link will not physically be removed, but deactivated.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `p_uid` | String | Yes | A unique identifier of the member on partner side, used to establish the link to partner ID. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | 1 |


### PartnerUpdate
`GET/POST` `/partner/update`

**Description:** Used to push membership bulk status updates to Lazada. Please note that this is not an incremental update, thus information left out that haven been in our system before, will be removed on our end.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `tier` | String | No | Customer’s tier in partner side, shown as-is |
| `balance` | Number | Yes | Balance of the membership. |
| `tier_expiry` | String | No | Expiry of the membership, shown as-is |
| `p_uid` | String | Yes | A unique identifier of the member on partner side, used to establish the link to partner ID. |
| `member_name` | String | No | Name of member on partner side, to easier identify the membership on My Account pages |
| `valid_from` | String | No | Valid from of this balance in RFC RFC3339 format. Ignore if this is no validity period for the balance |
| `status` | String | Yes | One of: ‘active’ – For activated members ‘inactive’ – For inactive member ‘pending’ – For members that are pending activation |
| `valid_to` | String | No | Valid to of this balance in RFC RFC3339 format. Ignore if this is no validity period for the balance |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | 1 |


### UpdatePartnerUserId
`GET/POST` `/partner/updatePartnerUserId`

**Description:** Used to update the partner user id to new partner user id

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `old_p_uid` | String | Yes | the current partner user id to match up with a user |
| `new_p_uid` | String | Yes | the new partner user id to be placed |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | api result |


---
## FBL API

_FBL API group_

### BuildFulfillmentSkuRelation
`POST` `/fbl/fulfillment_sku_relation/write`

**Description:** build the relation between platformSku and fulfillmentSku

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `site` | String | Yes | site |
| `item_id` | Number | Yes | itemId |
| `sku_id` | Number | Yes | skuId |
| `sc_item_id` | Number | No | fulfillmentSkuId |
| `fulfillment_sku` | String | No | fulfillmentSku |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result DTO |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `PARAM_ILLEGAL` | "sku not exists" | SKU incoming error please verify if it exists in the corresponding shop |


### CancelFulfillmentOrderForMCL
`POST` `/fbl/fulfillment_order/cancel`

**Description:** Cancel Fulfillment Order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_order_id` | String | Yes | Order level identifier for fulfilment order, unique for idempotence |
| `platform_name` | String | Yes | Trade platform name |
| `cancel_reason` | String | No | Cancelled reason |
| `items` | Object[] | Yes | Cancelled details |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success |
| `error_code` | String | No | Error code |
| `error_message` | String | No | Error message |


### CancelInboundReservation
`POST` `/fbl/inbound_reservation/cancel`

**Description:** cancel reservation order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reservation_order` | String | Yes | reservation order code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | success |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |


### CancelOutboundOrder
`POST` `/fbl/outbound_order/cancel`

**Description:** Cancel outbound order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `outbound_order_no` | String | Yes | Outbound order number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Cancel success or not. |
| `error_code` | String | No | Error code. |
| `error_message` | String | No | Error message. |


### CancelVasOrder4FBL
`GET/POST` `/fbl/vas/cancelVasOrder`

**Description:** 取消增值服务

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_name` | String | Yes | laz店铺所属的前台租户,例如: LAZADA_VN |
| `vas_order_no` | String | Yes | 增值服务单号 |
| `cancel_reason` | String | No | 取消原因 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | No | 取消结果 |


### CancelnBoundOrder
`POST` `/fbl/inbound_order/cancel`

**Description:** Cancel inbound order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `inbound_order_no` | String | Yes | Inbound order number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Cancel success or not. |
| `error_code` | String | No | Error code. |
| `error_message` | String | No | Error message. |


### CheckInboundReservationSlot
`GET` `/fbl/inbound_reservation/check`

**Description:** Check Available Reservation Slots for Inbound Order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `inbound_orders` | String | Yes | inbound order list  |
| `date` | String | Yes | date |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | success |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |
| `data` | Object | No | data |


### CreateFulfillmentOrderForMCL
`POST` `/fbl/fulfillment_order/create`

**Description:** Create Fulfillment Order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_payment_method` | String | Yes | Payment method, mainly check cod type |
| `remark` | String | No | Remark |
| `currency` | String | Yes | Currency |
| `items` | Object[] | Yes | Fulfillment order line list, contains no more than 300 items |
| `receiver` | Object | Yes | Receiver info |
| `platform_name` | String | Yes | Trade platform name |
| `fulfillment_finish_time` | String | No | Estimated warehouse outbound time in UTC |
| `platform_order_creation_time` | String | Yes | Trade order create time in UTC |
| `sales_order_number` | String | Yes | Sales order number from platform |
| `platform_order_id` | String | Yes | Unique order level identifier for fulfilment order |
| `out_order_creation_time` | String | No | Out fulfillment order create time in UTC |
| `is_platform_nominated_fleet` | Boolean | No | Whether platform nominated fleet |
| `seller_store_id` | String | No | seller store id |
| `seller_store_name` | String | No | seller store name |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success |
| `error_code` | String | No | Error code |
| `error_message` | String | No | Error message |


### CreateFulfillmentOrderForMCLV2PNF
`POST` `/fbl/fulfillment_order_pnf/create`

**Description:** Create Fulfillment Order for MCL2.0 PNF

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_payment_method` | String | Yes | Payment method, mainly check cod type |
| `remark` | String | No | Remark |
| `currency` | String | Yes | Currency |
| `items` | Object[] | Yes | Fulfillment order line list, contains no more than 300 items |
| `platform_name` | String | Yes | Trade platform name |
| `fulfillment_finish_time` | String | No | Estimated warehouse outbound time in UTC |
| `platform_order_creation_time` | String | Yes | Trade order create time in UTC |
| `sales_order_number` | String | Yes | Sales order number from platform |
| `platform_order_id` | String | Yes | Unique order level identifier for fulfilment order |
| `out_order_creation_time` | String | No | Out fulfillment order create time in UTC |
| `seller_store_id` | String | No | seller store id |
| `seller_store_name` | String | No | seller store name |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success |
| `error_code` | String | No | Error code |
| `error_message` | String | No | Error message |


### CreateFulfillmentSkuDecouple
`POST` `/fbl/fulfillment_sku/create`

**Description:** create fulfillment sku without product

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `fulfillment_sku_name` | String | Yes | title |
| `barcodes` | String[] | Yes | barcode list |
| `hygroscopic` | Boolean | Yes | true/false |
| `precious` | Boolean | Yes | true/false |
| `product_type` | String | Yes | food,liquid,danger,other |
| `temperature_requirement` | String | Yes | 1: normal temperature 4: refrigerated 6: frozen |
| `pic_urls` | String[] | Yes | at most 6 pictures url |
| `serial_number_flag` | Boolean | Yes | true/false |
| `shelf_life_flag` | Boolean | Yes | true/false |
| `shelf_life_days` | Number | No | required if shelf_life_day is life_mgnt |
| `reject_shelf_live` | Number | No | required if shelf_life_day is life_mgnt |
| `alert_shelf_live` | Number | No | required if shelf_life_day is life_mgnt |
| `offline_shelf_live` | Number | No | required if shelf_life_day is life_mgnt |
| `seller_sku` | String | Yes | erp sku code |
| `sale_price` | String | Yes | sale price |
| `length` | Number | No | length(mm) |
| `width` | Number | No | width(mm) |
| `height` | Number | No | height(mm) |
| `weight` | Number | No | weight(g) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |
| `success` | Boolean | No | is success |
| `error_code` | String | No | error_code |
| `error_message` | String | No | error_msg |


### CreateFulfillmentSkuForFBL
`POST` `/fbl/fulfillment_sku_fbl/create`

**Description:** create fulfillment sku for specified platform product

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `sku_id` | Number | Yes | platform sku sku_id |
| `barcodes` | String[] | Yes | barcode list |
| `hygroscopic` | Boolean | Yes | is product hygroscopic? |
| `product_type` | String | Yes | food / liquid / danger / other |
| `temperature_requirement` | String | Yes | "1": normal temperature  "4": refrigerated  "6": frozen |
| `serial_number_flag` | Boolean | Yes | is serial number management enabled? |
| `shelf_life_flag` | Boolean | Yes | is shelf life management enabled? |
| `shelf_life_days` | Number | No | days of shelf life, required if shelf_life_flag is true. |
| `reject_shelf_live` | Number | No | days to reject at inbound before expiry, required if shelf_life_flag is true. |
| `alert_shelf_live` | Number | No | days to alert before expiry, required if shelf_life_flag is true. |
| `offline_shelf_live` | Number | No | days to take offline before expiry, required if shelf_life_flag is true. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | request result |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |
| `data` | Object | No | data |


### CreateInboundOrder
`POST` `/fbl/inbound_order/create`

**Description:** Create inbound order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `warehouse_code` | String | Yes | Inbound warehouse code. |
| `delivery_type` | String | No | Delivery type,Enum: Dropoff / Pickup. |
| `seller_warehouse_code` | String | No | Seller warehouse code. Default value is seller's first sellerWarehouse, usually it's seller's address in asc. You can get the warehouse list by openApi listIcpWarehouse. |
| `estimate_time` | String | Yes | Estimated Arrival Time in UTC+0. format is "yyyy-MM-ddTHH:mm:ssZ". |
| `comment` | String | No | Inbound comment. |
| `reference_number` | String | No | Reference number. |
| `skus` | Object[] | Yes | List of inbound skus. Max list size is 100. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Create success or not. |
| `error_code` | String | No | Error code. |
| `error_message` | String | No | Error message. |
| `inbound_order_no` | String | No | Inbound order number. |


### CreateInboundReservation
`POST` `/fbl/inbound_reservation/create`

**Description:** create reservation order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `inbound_orders` | String[] | Yes | inbound order list |
| `slot` | String | Yes | reserve slot |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |
| `data` | Object | No | data |
| `success` | Boolean | No | is success |


### CreateOutBoundOrder
`POST` `/fbl/outbound_order/create`

**Description:** Create outbound order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reference_number` | String | No | Reference number. |
| `warehouse_code` | String | Yes | outbound warehouse code. |
| `delivery_type` | String | No | Delivery type,Enum: Dropoff / Pickup. |
| `seller_warehouse_code` | String | No | Seller warehouse code. Default value is seller's first sellerWarehouse, usually it's seller's address in asc. You can get the warehouse list by openApi listIcpWarehouse. |
| `estimate_time` | String | Yes | Estimated Time in UTC+0. format is "yyyy-MM-ddTHH:mm:ssZ". |
| `comment` | String | No | Outbound comment. |
| `inventory_type` | Number | Yes | Inventory type, 1 for good, 101 for defective. |
| `skus` | Object[] | Yes | List of outbound skus. Max list size is 100. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Create success or not. |
| `error_code` | String | No | Error code. |
| `error_message` | String | No | Error message. |
| `outbound_order_no` | String | No | Outbound order number |


### CreateProductReinboundOrderForMCL
`POST` `/fbl/product_reinbound/create`

**Description:** Create Product Reinbound Order on Failed Delivery for MCL

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_name` | String | Yes | Trade platform name |
| `sales_order_number` | String | Yes | Sales order number from platform |
| `platform_order_id` | String | Yes | Unique order level identifier for fulfilment order |
| `reinbound_order_id` | String | Yes | Package level identifier for product reinbound request, unique for idempotence |
| `tracking_number` | String | Yes | Tracking number for original package |
| `reason` | String | No | Failed delivery reason |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success |
| `error_code` | String | No | Error code |
| `error_message` | String | No | Error message |


### CreateVasOrder4FBL
`GET/POST` `/fbl/vas/createVasOrder`

**Description:** FBL增值服务创建

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_name` | String | Yes | laz店铺所属的前台租户,例如: LAZADA_VN |
| `idempotent_key` | String | Yes | 幂等码 |
| `service_provider_no` | String | No | 物流服务商单据号，比如：LBX |
| `target_order_no` | String | No | 服务目标单据号,比如：CO单号 |
| `target_order_type` | String | No | 服务对象类型：服务对象为入库单，则填写：CO；服务对象为品，则填写:GOODS; |
| `vas_code` | String | Yes | 增值服务Code：LABEL_PRINTING_PASTING_FOR_IB 打印并贴商品条码 LABEL_PRINTING_PASTING_FOR_ITEM 打印并贴商品条码 REPACKING_FOR_IB 重新包装 REPACKING_FOR_ITEM 重新包装 BUNDLING 绑定商品 LABEL_PRINTING_FOR_IB 打印商品条码 LABEL_PRINTING_FOR_ITEM 打印商品条码 LABEL_PASTING_FOR_IB 贴商品条码 LABEL_PASTING_FOR_ITEM 贴商品条码 SORTING 分类商品 INBOUND_QC 收货质检 |
| `warehouse_code` | String | Yes | 仓code |
| `lines` | Object[] | Yes | 明细行 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | No | 建单结果 |


### GetChannelStocksForMCL
`GET` `/fbl/channel_stocks/get`

**Description:** Query Channel Stocks

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_name` | String | Yes | Platform Name |
| `fulfillment_sku_id` | Number | Yes | Fulfillment Sku ID |
| `warehouse_code` | String | No | Warehouse Code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Success Flag |
| `error_code` | String | No | Error Code |
| `error_message` | String | No | Error Message |
| `data` | Object | No | Result Data |


### GetFulfillmentProductDetail
`GET` `/fbl/fulfillment_products/get`

**Description:** GET  fulfillment product Detail；Call Get Platform Products for fulfillment_sku first

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `per_page` | Number | No | Maximum number of results per page |
| `shelf_life_flag` | Boolean | No | Serial number flag. true or false |
| `marketplace` | String | Yes | Marketplace should be "LAZADA_MY","LAZADA_ID","LAZADA_VN","LAZADA_SG","LAZADA_TH","LAZADA_PH" |
| `fulfillment_sku` | String | No | Fulfillment SKU |
| `serial_number_flag` | Boolean | No | Serial number flag. true or false |
| `page` | Number | No | Page |
| `fulfillment_sku_name` | String | No | Fulfillment SKU Name used in Lazada fulfilment system |
| `barcode` | String | No | Barcode |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | List of products data |


### GetFulfillmentSkuListForMCL
`GET` `/fbl/fulfillment_sku_list/get`

**Description:** Get Fulfillment SKU List for LAZADA Partner

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `page` | Number | Yes | Page Index |
| `per_page` | String | Yes | Maximum number of results per page |
| `platform_name` | String | Yes | Platform name |
| `fulfillment_sku_name` | String | No | Fulfillment Sku Name |
| `seller_sku` | String | No | Seller Sku |
| `fulfillment_sku_code` | String | No | Fulfillment Sku Code |
| `barcode` | String | No | barcode |
| `fulfillment_sku_codes` | String | No | Fulfillment Sku Codes |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `error_message` | String | No | Error Message |
| `page` | Number | No | Page Index |
| `per_page` | Number | No | Maximum number of results per page |
| `total_count` | Number | No | Total Count |
| `data` | Object[] | No | Fulfillment sku list |
| `success` | Boolean | No | success flag |
| `error_code` | String | No | Error Code |


### GetFulfillmentSkuRelationByScItem
`GET/POST` `/fbl/fulfillment_sku_relation/get_by_sc_item`

**Description:** get the relation between platformSku and fulfillmentSku by scItem

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `site` | String | Yes | site |
| `sc_item_id` | Number | No | scItemId/fulfillment_sku_id |
| `fulfillment_sku` | String | No | fulfillment_sku |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result dto |


### GetFulfillmentSkuRelationBySku
`GET/POST` `/fbl/fulfillment_sku_relation/get_by_sku`

**Description:** get the relation between platformSku and fulfillmentSku by sku

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `site` | String | Yes | site |
| `item_id` | Number | Yes | itemId |
| `sku_id` | Number | Yes | skuId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result dto |


### GetFulfillmentSkuRelationsByScItems
`GET/POST` `/fbl/fulfillment_sku_relation/get_by_sc_items`

**Description:** get fulfillmentSku Relations By ScItems

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `biz_name` | String | Yes | bizName |
| `seller_ids` | Number[] | Yes | sellerIds |
| `sc_item_ids` | Number[] | No | scItemIds |
| `fulfillment_skus` | String[] | No | fulfillmentSkus |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### GetFulfillmentSkuRelationsBySkus
`GET/POST` `/fbl/fulfillment_sku_relation/get_by_skus`

**Description:** get fulfillmentSku Relations By Skus

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `site` | String | Yes | site |
| `item_sku` | Object | Yes | obj |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### GetIcpOrderFile
`GET` `/fbl/icp_order/file`

**Description:** Get Inbound/Outbound order print PDF file

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_number` | String | Yes | Inbound/Outbound order number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `error_code` | String | No | Error code. |
| `error_message` | String | No | Error message. |
| `data` | Object | No | File data |
| `success` | Boolean | No | Success or not. |


### GetInboundOrderDetail
`GET` `/fbl/inbound_order_detail/get`

**Description:** Use this API to get the Inbound Order Detail

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `inbound_order_no` | String | Yes | Inbound ouder number |
| `marketplace` | String | Yes | Enum Value:LAZADA_VN,LAZADA_SG,LAZADA_MY, LAZADA_ID,LAZADA_PH,LAZADA_TH |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Order detail |


### GetInboundOrderList
`GET` `/fbl/inbound_orders/get`

**Description:** Use this API to get inbound order list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `inbound_order_no` | String | No | Inbound order number, Multi orders split by ','. Max size is 100 |
| `creation_time_From` | String | No | Order's create time from |
| `creation_time_To` | String | No | Order's create time end |
| `inbound_warehouse` | String | No | Inbound warehouse name |
| `seller_sku` | String | No | seller sku name |
| `fulfillment_sku` | String | No | Fulfilment SKU code |
| `marketplace` | String | Yes | marketplace:LAZADA_VN,LAZADA_SG,LAZADA_MY, LAZADA_ID,LAZADA_PH,LAZADA_TH |
| `page` | String | No | Order list page index |
| `per_page` | String | No | Order list per page size, Max is 100 |
| `reservation_status` | String | No | ReservationStatus: PENDING_RESERVATION_ORDER_CREATE / RESERVATION_ORDER_CREATED / RESERVED /ARRIVED. PENDING_RESERVATION_ORDER_CREATE |
| `reservation_order` | String | No | Reservation Order number |
| `reference_number` | String | No | Reference number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Result |


### GetInboundReservationFile
`GET` `/fbl/inbound_reservation/file`

**Description:** get inbound reservation order file

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reservation_order` | String | Yes | reservation order code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | success |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |
| `data` | Object | No | data |


### GetInventoryChangedSKU
`GET` `/fbl/inventory_changed_sku/get`

**Description:** Use this API to get SKU list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `warehouse_code` | String | No | Warehouse code |
| `page` | Number | No | Sku list page index |
| `per_page` | Number | No | Sku list per page size |
| `market_place` | String | Yes | market place:LAZADA_VN,LAZADA_SG,LAZADA_MY, LAZADA_ID,LAZADA_PH,LAZADA_TH |
| `operate_Time_From` | String | No | Inventory operate time from. This param is Required |
| `operate_Time_To` | String | No | Inventory operate time to. This param is Required.We suggest that operate_time_to - operate_time_from < 6 months |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `per_page` | Number | No | Per page size |
| `page` | Number | No | Page index |
| `total_count` | Number | No | Total count of sku |
| `sku_list` | Object[] | No | Sku list |
| `success` | String | No | The api request success or not |
| `errMessage` | String | No | Error message when success=false |
| `errCode` | String | No | Error code when success=false |


### GetInventoryOccupyDetails
`GET` `/fbl/inventory_occupy_details/get`

**Description:** Use this API to get a sku's inventory occupy details

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `fulfillmentSku` | String | Yes | Fulfillment Sku Id |
| `storeCode` | String | Yes | Warehouse code |
| `marketplace` | String | Yes | market place:LAZADA_VN,LAZADA_SG,LAZADA_MY, LAZADA_ID,LAZADA_PH,LAZADA_TH |
| `pageNum` | Number | No | pageNum |
| `pageSize` | Number | No | pageSize |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `inventoryOccupyDetails` | Object[] | No | inventory occupy detail list |


### GetInventoryOperateLog
`GET` `/fbl/inventory_operate_log/get`

**Description:** Use this API to get a sku's inventory operate log

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `page` | Number | No | Operate log list page index |
| `per_page` | Number | No | Operate log list perpage size |
| `market_place` | String | Yes | market place:LAZADA_VN,LAZADA_SG,LAZADA_MY, LAZADA_ID,LAZADA_PH,LAZADA_TH |
| `operate_time_from` | String | No | Inventory operate time from, GMT+0.  |
| `operate_time_to` | String | No | Inventory operate time to, GMT+0. This param is Required. We suggest that operate_time_to - operate_time_from < 6 months |
| `warehouse_code` | String | No | Warehouse code |
| `fulfillment_sku_id` | String | No | Fulfillment Sku Id |
| `order_type_code` | String | No | Order Type Code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `inventory_operate_log` | Object[] | No | Inventory operate log |
| `success` | String | No | The api request success or not |
| `errMessage` | String | No | Error message when success=false |
| `errCode` | String | No | Error code when success=false |
| `page` | Number | No | Page index |
| `per_page` | Number | No | Per page size |
| `total_count` | Number | No | Total log count |


### GetOutboundOrderDetail
`GET` `/fbl/outbound_order_detail/get`

**Description:** Use this API to Get outbound order detail; shoud call GetOutboundOrderList for outbound_order_no first

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `outbound_order_no` | String | Yes | order number |
| `marketplace` | String | Yes | Enum Value:LAZADA_VN,LAZADA_SG,LAZADA_MY, LAZADA_ID,LAZADA_PH,LAZADA_TH |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Order detail |


### GetOutboundOrderList
`GET` `/fbl/outbound_orders/get`

**Description:** Use this API to get outbound order list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `outbound_order_no` | String | No | Outbound order number,Multi orders split by ','. Max size is 100 |
| `creation_time_from` | String | No | Order's create time from |
| `creation_time_to` | String | No | Order's create time end |
| `outbound_warehouse` | String | No | Outbound warehouse name |
| `seller_sku` | String | No | seller sku name |
| `fulfillment_sku` | String | No | Fulfilment SKU code |
| `marketplace` | String | Yes | marketplace:LAZADA_VN,LAZADA_SG,LAZADA_MY, LAZADA_ID,LAZADA_PH,LAZADA_TH |
| `page` | String | No | Order list page index |
| `per_page` | String | No | Order list per page size |
| `reference_number` | String | No | Reference number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Result |


### GetPlatformProductsV2
`GET` `/fbl/platform_products/get2`

**Description:** Search products list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `per_page` | Number | No | Maximum number of results Per Page |
| `seller_id` | Number | Yes | sellerId |
| `marketplace` | String | Yes | Marketplace |
| `seller_sku` | String | No | sellerSku |
| `platform_sku_name` | String | No | Platform SKU Name |
| `ready_for_inbound` | Boolean | No | Products that have binding stock in warsehouse |
| `platform_sku` | String | No | List of Platform SKU. Separate By Comma (,) |
| `page` | Number | No | Page Number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | Yes | List of products data |


### GetProductBatchList
`GET/POST` `/fbl/product_batch/query`

**Description:** query product batch list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `productBatchListRequest` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | response body |


### GetShipperInfo
`GET` `/fbl/shipper/get`

**Description:** Get Shipper Info for LAZADA Partner

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `error_message` | String | No | Error Message |
| `data` | Object | No | Result Data |
| `success` | Boolean | No | Whether Success |
| `error_code` | String | No | Error Code |


### GetStockRule
`GET` `/fbl/stock_rule/get`

**Description:** Get SKU stock rule by sku and warehouse

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `fulfillment_sku_ids` | String | No | fulfilment sku id list |
| `store_code` | String | Yes | warehouse code |
| `page` | String | No | page index, default: 1 |
| `per_page` | String | No | page size, default: 50 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | String | No | result |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |
| `page` | Number | No | page |
| `per_page` | Number | No | page size |
| `total_count` | Number | No | total count |
| `data` | Object[] | No | data list |


### GetVasOrderByNo4FBL
`GET/POST` `/fbl/vas/getVasOrderByNo`

**Description:** get vasOrder by orderNo

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_name` | String | Yes | laz店铺所属的前台租户,例如: LAZADA_VN |
| `vas_order_code` | String | Yes | 增值服务单号 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | No | 增值服务信息 |


### GetWarehouseListForMCL
`GET` `/fbl/warehouses/get`

**Description:** Get Warehouse List By Country And Multi-Channel

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `country_code` | String | Yes | CountryCode |
| `page` | Number | Yes | PageIndex |
| `per_page` | Number | Yes | Maximum number of results per page |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Success flag |
| `error_code` | String | No | Error Code |
| `error_message` | String | No | Error Message |
| `page` | Number | No | Page Index |
| `per_page` | Number | No | Maximum number of results per page |
| `total_count` | Number | No | Total count |
| `total_page` | Number | No | Total page |
| `data` | Object[] | No | Warehouse list |


### GetWarehouseStock
`GET` `/fbl/stocks/get`

**Description:** Get SKU list and stock by warehouse code

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `seller_sku` | String | No | Seller SKU |
| `marketplace` | String | Yes | Marketplace should be "LAZADA_MY","LAZADA_ID","LAZADA_VN","LAZADA_SG","LAZADA_TH","LAZADA_PH" |
| `fulfilment_sku` | String | No | List of shop SKU, Comma separated list in square brackets |
| `store_code` | String | No | Warehouse Code  List：https://www.yuque.com/u1990121/kb/exh5go#B4gg |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | Response body |


### GetWarehouseStockV3
`GET` `/fbl/stocks/getV3`

**Description:** Get SKU list and stock by warehouse code, this version separates pending inbound and stock in transit in return json.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `seller_sku` | String | No | Seller SKU, required when fulfilment_sku is empty |
| `marketplace` | String | Yes | Marketplace should be "LAZADA_MY","LAZADA_ID","LAZADA_VN","LAZADA_SG","LAZADA_TH","LAZADA_PH" |
| `fulfilment_sku` | String | No | List of shop SKU, Comma separated list in square brackets, required when seller_sku is empty |
| `store_code` | String | No | Warehouse Code  List：https://www.yuque.com/u1990121/kb/exh5go#B4gg |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | Response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `HttpConnectError` | Request failed, due to [%s] | The connection timed out or failed and needs to be retried. |


### ListIcpWarehouse
`GET` `/fbl/icp_warehouse/list`

**Description:** List warehouses for create InboundOrder and outboundOrder

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `warehouse_type` | String | Yes | Warehouse type. Enum: Inbound / outbound / Seller |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Success or not. |
| `error_code` | String | No | Error code. |
| `error_message` | String | No | Error message. |
| `data` | Object[] | No | warehouse list |


### QueryFulfillmentOrderForMCL
`GET` `/fbl/fulfillment_order_list/get`

**Description:** Query list of Fulfillment Orders by shipper

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platform_order_id` | String | No | Order level identifier for fulfilment order, unique for idempotence |
| `platform_name` | String | Yes | Trade platform name |
| `per_page` | Number | Yes | Page size |
| `page` | Number | Yes | Page index |
| `sales_order_number` | String | No | Sales order number from platform |
| `status` | String | No | Status |
| `create_start_time` | String | Yes | Order create time lower bound |
| `create_end_time` | String | Yes | Order create time upper bound |
| `delivery_type` | String | No | Delivery type |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success |
| `error_code` | String | No | Error code |
| `error_message` | String | No | Error message |
| `per_page` | Number | No | Page size |
| `page` | Number | No | Page index |
| `total_count` | Number | No | Total count |
| `data` | Object[] | No | Result order list |


### QueryInboundBatch
`GET/POST` `/fbl/inbound_batch/query`

**Description:** query inbound batch

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `query_request` | Object | Yes | request body |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | response body |


### QueryInboundReservationOrder
`GET` `/fbl/inbound_reservation/get`

**Description:** get inbound reservation order

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reservation_order` | String | No | reservation order |
| `inbound_order` | String | No | Inbound Order ID, required when  reservation order is not present. if  reservation order is present, use reservation order first |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | success |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |
| `data` | Object | No | data |


### QueryReverseOrderForMCL
`GET` `/fbl/reverse_order/get`

**Description:** Query Reverse Order for MCL

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `sales_order_number` | String | Yes | Sales order number from platform |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Whether Success |
| `error_message` | String | No | Error Message |
| `data` | Object[] | No | Result Data |


### RemoveFulfillmentSkuRelation
`POST` `/fbl/fulfillment_sku_relation/remove`

**Description:** remove the relation between platformSku and fulfillmentSku

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `site` | String | Yes | site |
| `item_id` | Number | Yes | itemId |
| `sku_id` | Number | Yes | skuId |
| `sc_item_id` | Number | No | fulfillmentSkuId |
| `fulfillment_sku` | String | No | fulfillmentSku |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result DTO |


### ReturnCancellation
`POST` `/fbl/returns/cancel`

**Description:** Return order cancellation

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `return_id` | String | Yes | return id created during return order creation |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | is success |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |


### ReturnOrderCreation
`POST` `/fbl/returns/create`

**Description:** Api to create customer returns

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `tracking` | Object | Yes | tracking |
| `platform_name` | String | Yes | Platform Name |
| `platform_order_creation_time` | String | Yes | Sales order creation time of platform side Datetime format: 2017-11-17T10:14:13.185Z |
| `return_comment` | String | Yes | Customer comments accompanying the return order, will be used as reference during quality check |
| `return_delivery_type` | String | Yes | Return delivery type (always return_by_customer) |
| `return_order_number` | String | Yes | Return order number from platform; must be unique |
| `sales_order_number` | String | Yes | Sales order number accompanying the original fulfilment order request |
| `currency` | String | Yes | Currency |
| `customer` | Object | Yes | customer info |
| `platform_order_id` | String | Yes | Return order id - unique order level Identifier used to send return order and item status notification events |
| `parcel` | Object | Yes | parcel |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | result |
| `success` | Boolean | No | is success |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |


### SetStockRule
`POST` `/fbl/stock_rule/set`

**Description:** set channel ratio by sku and warehouse

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `skus` | Object[] | Yes | skus |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | success |
| `error_code` | String | No | error code |
| `error_message` | String | No | error message |


### UpdateFulfillmentSkuDecouple
`POST` `/fbl/fulfillment_sku/update`

**Description:** update fulfillment sku without product

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `barcodes` | String[] | No | barcode list |
| `hygroscopic` | Boolean | No | true/false |
| `precious` | Boolean | No | true/false |
| `product_type` | String | No | food,liquid,danger,other |
| `temperature_requirement` | String | No | 1: normal temperature 4: refrigerated 6: frozen |
| `pic_urls` | String[] | No | at most 6 pictures url |
| `serial_number_flag` | Boolean | No | true/false |
| `shelf_life_flag` | Boolean | No | true/false |
| `shelf_life_days` | Number | No | required if shelf_life_day is life_mgnt |
| `reject_shelf_live` | Number | No | required if shelf_life_day is life_mgnt |
| `alert_shelf_live` | Number | No | required if shelf_life_day is life_mgnt |
| `offline_shelf_live` | Number | No | required if shelf_life_day is life_mgnt |
| `sale_price` | String | No | sale price |
| `fulfillment_sku_id` | Number | Yes | fulfillment_sku_id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | is success |
| `error_code` | String | No | error_code |
| `error_message` | String | No | error_msg |
| `data` | Boolean | No | is success |


### UploadWaybill
`GET/POST` `/fbl/waybill/upload`

**Description:** Use this API to upload a waybill pdf to Lazada site. The maximum size of an pdf file is 1MB.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `waybill` | byte[] | Yes | waybill pdf |
| `package_code` | String | Yes | package code |
| `tracking_number` | String | Yes | tracking number |
| `extends_field` | String | No | extend fields |
| `store_code` | String | Yes | warehouse_code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | whether success |
| `error_message` | String | Yes | error message |
| `error_code` | String | Yes | error code |


---
## Instant Messaging API

_Instant Messaging APIs_

### GetMessages
`GET` `/im/message/list`

**Description:** Get message list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `session_id` | String | Yes | session id |
| `start_time` | Number | Yes | when request the first page pls input current timestamp，get the next page pls input previous page response field  next_start_time |
| `page_size` | Number | Yes | page size |
| `last_message_id` | String | No | previous page output param [last_message_id];it could be null when get the first page, get the next page pls input previous page response field  last_message_id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `err_code` | String | Yes | error code |
| `data` | Object | Yes | json |
| `success` | Boolean | Yes | result true or false |
| `err_message` | String | Yes | error message |


### GetSessionDetail
`GET` `/im/session/get`

**Description:** get session detail by sessionid

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `session_id` | String | Yes | session id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `err_code` | String | Yes | error code 0=success |
| `data` | Object | Yes | json |
| `success` | Boolean | Yes | result true or false |
| `err_message` | String | Yes | error message |


### GetSessionList
`GET` `/im/session/list`

**Description:** query seller session list

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `last_session_id` | String | No | previous page output param [last_session_id];The last session id on this page, it needs to be passed in as an input parameter when pulling the next page |
| `start_time` | String | Yes | next page start time;when pull first page pls input current timestamp， when pull next page pls input last page response field next_start_time |
| `page_size` | String | Yes | page size |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | result true or false |
| `err_message` | String | Yes | error message |
| `err_code` | String | Yes | error code, 0=success |
| `data` | Object | Yes | json |


### MessageRecall
`GET/POST` `/im/message/recall`

**Description:** message recall

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `session_id` | String | Yes | session id;conversation id |
| `message_id` | String | Yes | the id of message that need to be recalled;1）Cannot be recalled more than two minutes since the message has been sent 2）system message could not be  recalled |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `err_code` | String | No | error code 0=success |
| `success` | Boolean | No | true or false |
| `err_message` | String | No | error message |


### OpenSession
`GET/POST` `/im/session/open`

**Description:** open a new conversation

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `order_id` | Number | Yes | orderId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `session_id` | String | No | unique id of conversation |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `-22` | order out of day limit: 30 | Order timeout, only order IDs created within 30 days can be used to create a session |


### ReadSession
`POST` `/im/session/read`

**Description:** session read

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `session_id` | String | Yes | session id;unique id of a conversation |
| `last_read_message_id` | String | Yes | last message id of user readed |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `err_code` | String | Yes | error code 0=success |
| `success` | Boolean | Yes | true or false |
| `err_message` | String | Yes | error message |


### SendMessage
`POST` `/im/message/send`

**Description:** send message

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `session_id` | String | Yes | conversation id |
| `template_id` | String | Yes | message template id, 1: normal text message   3: picture message 4: emoji message 10006: item message 10007:  order message 10008: voucher message 10010: invite buyers to follow the store 6: video message, use this API to upload video (The video duration is greater than 3s and less than 180s) |
| `txt` | String | No | template_id=1 required |
| `img_url` | String | No | template_id=3 required |
| `width` | Number | No | template_id=3/6 required |
| `height` | Number | No | template_id=3/6 required |
| `item_id` | String | No | template_id=10006 required |
| `order_id` | String | No | template_id=10007 required |
| `promotion_id` | String | No | template_id=10008 required |
| `video_id` | String | No | template_id=6 required |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `err_code` | String | Yes | error code 0=success |
| `data` | Object | Yes | json |
| `success` | Boolean | Yes | true or false |
| `err_message` | String | Yes | error message |


---
## Lazada Logistics API

### CreateCustomerAccountRelationshipByOTP
`GET/POST` `/logistics/epis/customers/external_relationships_bundle`

**Description:** Create customer account relationship for external by OTP

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `externalSellerId` | String | Yes | External seller ID |
| `platformName` | String | Yes | Platform name |
| `otp` | String | Yes | Bundle code generated in Lazada Logistics Website |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Request success or not |
| `retryable` | Boolean | No | Is fail request retryable |
| `traceId` | String | No | Trace ID for debugging |
| `errorMessage` | String | No | Error code |
| `errorCode` | String | No | Error message |
| `errors` | Object[] | No | Error detail |


### CreateCustomerAccountRelationshipForExternal
`POST` `/logistics/epis/customers/external_relationships`

**Description:** External partner calls LAZADA to create account relationship

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `externalSellerId` | String | Yes | External seller ID |
| `platformName` | String | Yes | Platform name |
| `customerId` | String | Yes | Customer ID sent generated by Lazada |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is fail request retryable |
| `traceId` | String | Yes | Trace ID for debugging |
| `success` | Boolean | Yes | Request success or not |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Error field |


### CreateOrUpdateCustomerWarehouse
`POST` `/logistics/epis/customers/warehouses`

**Description:** External partner calls LAZADA to create or update warehouses

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `externalSellerId` | String | Yes | External seller ID |
| `platformName` | String | Yes | Platform name |
| `warehouseCode` | String | Yes | External warehouse code |
| `warehouseName` | String | Yes | Warehouse name |
| `contactName` | String | Yes | Warehouse contact name |
| `phone` | String | Yes | Warehouse contact phone number. If no country phone prefix, EPIS will append the country prefix of current country |
| `email` | String | No | Warehouse contact email |
| `type` | String | Yes | Enum: NORMAL / RETURN |
| `address` | Object | Yes | Warehouse address |
| `solutionCodes` | String[] | Yes | List of Lazada solution codes. Enum [LAZADA_STANDARD_VN, LAZADA_BULKY_VN] |
| `configuration` | Object | No | Warehouse configuration |
| `dropshippingInfo` | Object | No | drop shipping info |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is fail request retryable |
| `traceId` | String | Yes | Trace ID for debugging |
| `success` | Boolean | Yes | Request success or not |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Error field |
| `data` | Object | No | Response |


### EPIS Send package information
`HSF` `/logistics/epis/external/packages`

**Description:** Sync package info to OneLink

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | No | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |
| `trackingNumber` | String | No | Tracking number generated by EPIS |
| `packageType` | String | No | Package type. Enum: [Sales_order, Customer_return] |
| `externalOrderId` | String | No | External order ID |
| `lazadaOrderId` | String | No | Lazada internal order ID |
| `platformOrderCreationTime` | Number | No | Unix timestamp in milliseconds. Default: Current timestamp when EPIS receives the request |
| `deliveryOption` | String | No | Delivery service type. Enum [standard, economy] |
| `dangerousGood` | Boolean | No | Is dangerous good.  Boolean true/false |
| `items` | Object | No | Item list |
| `partner` | Object | No | Platform information |
| `origin` | Object | No | Origin contact information |
| `destination` | Object | No | Destination contact information |
| `payment` | Object | No | Payment info |
| `dimWeight` | Object | No | Package level dimweight |
| `options` | Object | No | Package extra options |
| `rawPayload` | String | Yes | Raw payload in json |
| `lexMsgType` | String | Yes | Lex message type |
| `eposSellerId` | String | No | OneLinkAccountId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `retryable` | Boolean | No | Is failed request retryable? |
| `traceId` | String | No | Trace ID for debugging |
| `errorCode` | String | No | Error code |
| `errorMessage` | String | No | Error message |


### EpisGetDeliveryOptions
`GET` `/logistics/epis/service/delivery_options`

**Description:** External partner call EPIS to get delivery options for package

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `fromLocation` | Object | No | Origin geo location |
| `toLocation` | Object | No | Destination geo location |
| `shipper` | Object | Yes | Shipper information |
| `dimWeight` | Object | Yes | Package level dimweight |
| `origin` | Object | Yes | Origin info |
| `destination` | Object | Yes | Destination info |
| `payment` | Object | Yes | Payment info |
| `packageType` | String | No | Package type. [Sales_order, Customer_return] |
| `deliveryOption` | String | No | Delivery service type. Enum [standard, economy] |
| `externalOrderId` | String | No | Order Id from external |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object[] | No | Response data |
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Detail errors |


### EpisPackageCancellation
`POST` `/logistics/epis/packages/cancel`

**Description:** External partner call EPIS to cancel package

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reason` | String | Yes | Cancellation reason (Free text) |
| `packageCode` | String | Yes | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Error details |


### EpisPackageCancellationStatusUpdate
`HSF` `/logistics/epis/external/packages/cancel_status_update`

**Description:** EPIS update package cancellation status to external

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | Yes | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |
| `externalOrderId` | String | Yes | The external order id |
| `success` | String | Yes | Cancel request success or not |
| `reason` | String | No | When failed to cancel this field may include the reason |
| `authorization` | String | No | Authorization header |
| `eposSellerId` | String | No | Onelink account |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `retryable` | Boolean | No | Is failed request retryable? |
| `traceId` | String | No | Trace ID for debugging |
| `errorCode` | String | No | Error code |
| `errorMessage` | String | No | Error message |


### EpisPackageCancellationStatusUpdateV2
`HSF` `/logistics/epis/external/v2/packages/cancel_status_update`

**Description:** EPIS update package cancellation status to external

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | Yes | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |
| `externalOrderId` | String | Yes | The external order id |
| `success` | Boolean | Yes | Cancel request success or not |
| `reason` | String | No | When failed to cancel this field may include the reason |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `retryable` | Boolean | No | Is failed request retryable? |
| `traceId` | String | No | Trace ID for debugging |
| `errorCode` | String | No | Error code |
| `errorMessage` | String | No | Error message |


### EpisPackageCancellationV3
`POST` `/logistics/epis/packages/cancel/v3`

**Description:** External partner call EPIS to cancel FFM + DEL package

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `reason` | String | Yes | Cancellation reason (Free text) |
| `packageCode` | String | Yes | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |
| `logisticsOrderId` | String | Yes | logisticsOrderId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Error details |


### EpisPackageConsignment
`POST` `/logistics/epis/packages/consign`

**Description:** External partner call EPIS to consign package to get the tracking number and be able to print AWB after consign

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `multiParcel` | Object | No | Multi-parcel information |
| `dangerousGood` | Boolean | Yes | Is dangerous good.  Boolean true/false |
| `shipper` | Object | Yes | Shipper information |
| `dimWeight` | Object | Yes | Package level dimweight |
| `origin` | Object | Yes | Origin info |
| `destination` | Object | Yes | Destination info |
| `payment` | Object | Yes | Payment info |
| `externalOrderId` | String | Yes | External order id (uniquely identify partner's order). If Lazada receives mulitiple requests to create multiple orders with same externalOrderId then only the first arrived order information is recorded. All subsequent requests are treated as duplicated regardless the order information is changed or not. Therefore you can repush order, but cannot modify order information once it is already processed by Lazada |
| `platformOrderCreationTime` | Number | No | Unix timestamp in milliseconds. Default: Current timestamp when EPIS receives the request |
| `packageType` | String | No | Package type. Enum: [Sales_order, Customer_return] |
| `deliveryOption` | String | No | Delivery service type. Enum [standard, economy, point_to_point] |
| `items` | Object[] | Yes | Item list |
| `options` | Object | No | Package options |
| `exchangeOrder` | Object | No | Exchange order |
| `planInfo` | Object | No | Partner can send tracking number to LEX |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Detail errors |
| `data` | Object | No | Response data |


### EpisPackageConsignmentV2
`POST` `/logistics/epis/packages/consign/v2`

**Description:** External partner call EPIS to consign FFM + DEL package to get the tracking number and be able to print AWB after consign

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `dangerousGood` | Boolean | Yes | Is dangerous good.  Boolean true/false |
| `shipper` | Object | Yes | Shipper information |
| `dimWeight` | Object | Yes | Package level dimweight |
| `origin` | Object | Yes | Origin info |
| `destination` | Object | No | Destination info |
| `payment` | Object | Yes | Payment info |
| `externalOrderId` | String | Yes | External order id (uniquely identify partner's order). If Lazada receives mulitiple requests to create multiple orders with same externalOrderId then only the first arrived order information is recorded. All subsequent requests are treated as duplicated regardless the order information is changed or not. Therefore you can repush order, but cannot modify order information once it is already processed by Lazada |
| `platformOrderCreationTime` | Number | No | Unix timestamp in milliseconds. Default: Current timestamp when EPIS receives the request |
| `packageType` | String | No | Package type. Enum: [Sales_order, Customer_return] |
| `deliveryOption` | String | No | Delivery service type. Enum [standard, economy, point_to_point] |
| `items` | Object[] | Yes | Item list |
| `options` | Object | No | Package options |
| `exchangeOrder` | Object | No | Exchange order |
| `planInfo` | Object | No | Partner can send tracking number to LEX |
| `packageServices` | String[] | Yes | package services: FULFILLMENT,DELIVERY |
| `fulfillmentInfo` | Object | No | fulfillment info |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Detail errors |
| `data` | Object | No | Response data |


### EpisPackageCreation
`POST` `/logistics/epis/packages`

**Description:** External partner call EPIS to create package

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `dangerousGood` | Boolean | Yes | Is dangerous good.  Boolean true/false |
| `shipper` | Object | Yes | Shipper information |
| `dimWeight` | Object | Yes | Package level dimweight |
| `origin` | Object | Yes | Origin info |
| `destination` | Object | Yes | Destination info |
| `payment` | Object | Yes | Payment info |
| `externalOrderId` | String | Yes | External order id (uniquely identify partner's order). If Lazada receives mulitiple requests to create multiple orders with same externalOrderId then only the first arrived order information is recorded. All subsequent requests are treated as duplicated regardless the order information is changed or not. Therefore you can repush order, but cannot modify order information once it is already processed by Lazada |
| `platformOrderCreationTime` | Number | No | Unix timestamp in milliseconds. Default: Current timestamp when EPIS receives the request |
| `packageType` | String | No | Package type. Enum: [Sales_order, Customer_return] |
| `deliveryOption` | String | No | Delivery service type. Enum [standard, economy, point_to_point] |
| `items` | Object[] | Yes | Item list |
| `options` | Object | No | Package options |
| `exchangeOrder` | Object | No | Exchange order |
| `planInfo` | Object | No | Partner can send tracking number to LEX |
| `multiParcel` | Object | No | Multi-parcel information |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Detail errors |
| `data` | Object | No | Response data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `NO_ROUTE_ERROR` | No suitable route found | origin-destination route is out of coverage |
| `PARTIAL_DELIVERY_NOT_AVAILABLE` | Partial Delivery is not available because out of Lex coverage | Partial Delivery is not available because out of Lex coverage |
| `INVALID_ADDRESS_ID` | Provided address ID is not valid | Provided address ID is not valid |
| `BAD_REQUEST` | Invalid request | Check error message for details |
| `PARTNER_NOT_FOUND` | No partner matches with provided information | Need to register partner information with Lazada before calling API |
| `INTERNAL_SYSTEM_ERROR` | Internal system error. Please try again | Internal system error. |


### EpisPackageDimweightUpdate
`HSF` `/logistics/epis/external/packages/dim_weight`

**Description:** EPIS update package dim weight to external

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | Yes | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |
| `externalOrderId` | String | Yes | The external order id |
| `processTime` | Number | Yes | Dimweight update timestamp |
| `dimWeight` | Object | Yes | Dimension & Weight |
| `rawPayload` | String | No | Raw payload |
| `lexMsgType` | String | No | Message type |
| `authorization` | String | No | Authorization header |
| `sendForBilling` | Boolean | No | is send for billing? |
| `eposSellerId` | String | No | Onelink account |
| `xQueryA` | String | No | Onelink account |
| `xQueryC` | String | No | Onelink account |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `retryable` | Boolean | No | Is failed request retryable? |
| `traceId` | String | No | Trace ID for debugging |
| `errorCode` | String | No | Error code |
| `errorMessage` | String | No | Error message |


### EpisPackageInfoUpdate
`POST` `/logistics/epis/packages/update`

**Description:** External partner call EPIS to update package info after RTS

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | Yes | Package code |
| `receiverName` | String | Yes | Receiver name |
| `receiverPhone` | String | Yes | Receiver phone number |
| `totalAmount` | String | No | Payment total amount |
| `insuranceAmount` | String | No | Payment insurance amount |
| `deliveryNote` | String | No | Delivery note |
| `receiverAddress` | Object | No | Receiver address |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | No | Is failed request retryable? |
| `traceId` | String | No | trace id for debug |
| `success` | Boolean | No | Is success? |
| `errorMessage` | String | No | Error message |
| `errorCode` | String | No | Error code |
| `errors` | Object[] | No | Error detail |
| `data` | Object | No | update result |


### EpisPackagePrintAwb
`GET` `/logistics/epis/packages/awb`

**Description:** External partner call LAZADA to print AWB

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | Yes | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |
| `type` | String | Yes | Type of AWB output. Enum [pdf] |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is fail request retryable |
| `traceId` | String | Yes | Trace ID for debugging |
| `data` | Object | Yes | AWB data |
| `success` | Boolean | Yes | Request success or not |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Error field |


### EpisPackageReAttempt
`POST` `/logistics/epis/packages/reattempt`

**Description:** Send re-attempt package request

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | Yes | Package code |
| `reAttemptDateTime` | Number | No | Re attempt time |
| `sellerNote` | String | No | Seller note |
| `feedbackType` | String | Yes | REATTEMPT or RETURN |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | No | Is failed request retryable? |
| `traceId` | String | No | trace id for debug |
| `success` | String | No | is success? |
| `errorMessage` | String | No | Error message |
| `errorCode` | String | No | Error code |


### EpisPackageReadyToBeShipped
`POST` `/logistics/epis/packages/rts`

**Description:** External partner calls EPIS to mark a package as ready to be shipped

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `trackingNumber` | String | Yes | Unique tracking number generated by EPIS after consignment |
| `paidEstimatedShippingFee` | String | No | paidEstimatedShippingFee |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Detail errors |
| `data` | Object | No | Response data |


### EpisPackageStatusUpdate
`HSF` `/logistics/epis/external/packages/statuses`

**Description:** EPIS update package status to external

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `packageCode` | String | No | Unique package identifier generated by EPIS. Used to communicate between EPIS and external systems |
| `externalOrderId` | String | Yes | The external order id. |
| `status` | String | Yes | Lazada level 2 status |
| `processTime` | Number | Yes | Status update timestamp |
| `location` | String | No | Status update location |
| `reasonCode` | String | No | Reason code when fail delivery |
| `geoLocation` | Object | No | Geo location |
| `properties` | Object | No | Status update properties |
| `authorization` | String | No | Authorization header |
| `eposSellerId` | String | No | Onelink account |
| `logisticOrderId` | String | No | logisticsOrderId - required when use fulfillment services |
| `trackingNumber` | String | No | Tracking number |
| `xQueryA` | String | No | Extra parameter |
| `xQueryC` | String | No | Extra parameter |
| `xQueryCompanyId` | String | No | Extra parameter |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `retryable` | Boolean | No | Is failed request retryable? |
| `traceId` | String | No | Trace ID for debugging |
| `errorCode` | String | No | Error code |
| `errorMessage` | String | No | Error message |


### EpisUploadAwbFulfillment
`POST` `/logistics/epis/fulfillment/upload_awb`

**Description:** External partner call EPIS to upload awb for fulfillment

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `logisticsOrderId` | String | Yes | logisticsOrderId |
| `trackingNumber` | String | No | trackingNumber |
| `waybill` | byte[] | No | Waybill content |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Error details |


### EpisXspaceCreate
`POST` `/logistics/epis/xspace/create`

**Description:** Create Xspace case

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `caseTemplateId` | Number | No | case template id |
| `categoryId` | Number | No | cat id |
| `subject` | String | Yes | subject |
| `description` | String | Yes | description |
| `sellerName` | String | No | sellerName |
| `sellerEmail` | String | No | sellerEmail |
| `sellerPhoneNo` | String | No | sellerPhoneNo |
| `buyerName` | String | No | buyerName |
| `buyerEmail` | String | No | buyerEmail |
| `trackingNumber` | String | No | trackingNumber |
| `orderId` | String | No | orderId |
| `casePriority` | String | No | casePriority |
| `attachments` | String | No | attachments |
| `attributes` | String | No | attributes |
| `platformName` | String | No | platformName |
| `externalSellerId` | String | No | externalSellerId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | No | retryable |
| `success` | Boolean | No | success or not |
| `traceId` | String | No | traceId |
| `errorMessage` | String | No | errorMessage |
| `errorCode` | String | No | errorCode |
| `data` | Object | No | response data |


### EpisXspaceGetDetail
`POST` `/logistics/epis/xspace/detail`

**Description:** Get Xspace case detail

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `caseId` | Number | No | case id |
| `platformName` | String | No | platformName |
| `externalSellerId` | String | No | externalSellerId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | No | retryable |
| `success` | Boolean | No | success or not |
| `traceId` | String | No | traceId |
| `errorMessage` | String | No | errorMessage |
| `errorCode` | String | No | errorCode |
| `data` | Object | No | response data |


### EpisXspaceQuery
`GET/POST` `/logistics/epis/xspace/query`

**Description:** Query Xspace case

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `caseIds` | String[] | No | caseIds |
| `trackingNumbers` | String[] | No | trackingNumbers |
| `createTimeFrom` | String | No | createTimeFrom |
| `createTimeTo` | String | No | createTimeTo |
| `pageSize` | String | No | pageSize |
| `pageNo` | String | No | pageNo |
| `sortBy` | String | No | sortBy |
| `sortOrder` | String | No | sortOrder |
| `statuses` | String[] | No | statuses |
| `platformName` | String | No | platformName |
| `externalSellerId` | String | No | externalSellerId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | No | retryable |
| `success` | Boolean | No | success |
| `traceId` | String | No | traceId |
| `errorCode` | String | No | errorCode |
| `errorMessage` | String | No | errorMessage |
| `data` | Object | No | data |


### EpisXspaceRateTicket
`GET/POST` `/logistics/epis/xspace/rate`

**Description:** Rate Xspace ticket

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `platformName` | String | Yes | platformName |
| `externalSellerId` | String | Yes | externalSellerId |
| `caseId` | Number | Yes | caseId |
| `ratingStar` | Number | Yes | ratingStar |
| `ratingReasons` | String[] | No | ratingReasons |
| `ratingRemark` | String | No | ratingRemark |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | No | retryable |
| `success` | Boolean | No | success |
| `traceId` | String | No | traceId |
| `errorCode` | String | No | errorCode |
| `errorMessage` | String | No | errorMessage |


### EpisXspaceTicketStatusUpdate
`HSF` `/logistics/epis/xspace/ticket/status`

**Description:** EPIS update xspace status to external

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `caseId` | String | Yes | caseId |
| `subject` | String | Yes | subject |
| `description` | String | No | description |
| `status` | String | Yes | status |
| `processedTime` | String | No | processedTime |
| `lexMsgType` | String | Yes | lexMsgType |
| `externalSellerId` | String | Yes | externalSellerId |
| `eposSellerId` | String | No | OneLinkAccountId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | success |
| `retryable` | Boolean | No | retryable |
| `traceId` | String | No | traceId |
| `errorCode` | String | No | errorCode |
| `errorMessage` | String | No | errorMessage |


### EstimateShippingFee
`GET/POST` `/logistics/epis/estimate_shipping_fee`

**Description:** Estimate shipping fee

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `externalSellerId` | String | Yes | External seller ID |
| `platformName` | String | Yes | Platform where seller order comes from |
| `fromAddressId` | String | No | Lazada last level address R-code |
| `toAddressId` | String | No | Lazada last level address R-code |
| `chargeFactor` | Object | Yes | Charge factors |
| `fromLocation` | Object | No | Geo location |
| `toLocation` | Object | No | Geo location |
| `packageCode` | String | No | package code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `data` | Object[] | Yes | Rating response |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Detail errors |


### GetShippingFee
`GET/POST` `/logistics/epis/get_shipping_fee`

**Description:** Estimate package shipping fee (Estimated & Actual)

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `externalSellerId` | String | Yes | External seller ID |
| `platformName` | String | Yes | Platform where seller order comes from |
| `trackingNumber` | String | Yes | Lazada tracking number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `retryable` | Boolean | Yes | Is failed request retryable? |
| `traceId` | String | Yes | Trace id for debugging |
| `data` | Object | Yes | Package fee response |
| `success` | Boolean | Yes | Is success? |
| `errorMessage` | String | Yes | Error message |
| `errorCode` | String | Yes | Error code |
| `errors` | Object[] | Yes | Detail errors |


### MY - Pickupp addJobsToManifest
`HSF` `/v2/public/integrations/lazada`

**Description:** Add job to manifest request

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `manifestId` | Number | Yes | manifest id |
| `stops` | Object[] | Yes | na |
| `traceId` | String | Yes | na |
| `authToken` | String | Yes | auth |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |


### MY - Pickupp createManifest
`HSF` `/public/integrations/lazada`

**Description:** Call pickupp create manifest

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `manifestId` | Number | Yes | manifest id |
| `locationId` | String | Yes | node id |
| `subConId` | Number | Yes | sub-con virtual courier id |
| `fileUrl` | String | Yes | link to download manifest json file |
| `traceId` | String | Yes | trace id to track request |
| `authToken` | String | Yes | Authorization header |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |


### MY - Pickupp removeJobFromManifest
`HSF` `/public/integrations/lazada/cancel`

**Description:** Remove job from manifest

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `manifestId` | Number | Yes | na |
| `trackingNumber` | String | Yes | trackingNumber |
| `traceId` | String | Yes | traceId |
| `authToken` | String | Yes | auth |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |


---
## E-Tickets API

_E-Tickets API for digital voucher and e-ticketing service_

### GetOrderItemsFromBarCode
`GET/POST` `/eticket/code/query`

**Description:** E-Ticcket certificate query Open API

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `code` | String | Yes | certificate code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `100` | E100: Param Invalid, "%s" | Param invalid |
| `200` | E200: Certificate Not Exist | Certificate not exist |
| `201` | E201: Certificate Not Unique | More that one certificate matched |
| `202` | E202: Certificate Can Not Distinguish | Can't distinguish the business type of this code |


### GlobalEticketMerchantMaAvailable
`GET/POST` `/eticket/ma/available`

**Description:** the callback interface before consume  code

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `biz_type` | Number | Yes | biz type |
| `code` | String | Yes | waiting consume code |
| `serial_num` | String | Yes | consume serialVersionUID |
| `pos_id` | String | No | consume tools no |
| `outer_id` | String | Yes | order id |
| `consume_num` | Number | Yes | consume num |
| `consume_store_id` | String | Yes | consume store id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resp_body` | Object | Yes | response |
| `ret_code` | String | Yes | sub code |
| `ret_msg` | String | Yes | sub info |


### GlobalEticketMerchantMaConsume
`POST` `/eticket/ma/consume`

**Description:** consume ma

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `biz_type` | Number | Yes | biz type |
| `serial_num` | String | Yes | consume serialVersionUID |
| `pos_id` | String | No | consume tools no |
| `outer_id` | String | Yes | order id |
| `consume_num` | Number | Yes | consume num |
| `code` | String | Yes | waiting consume code |
| `consume_store_id` | String | Yes | consume store id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resp_body` | Object | Yes | response |
| `ret_code` | String | Yes | sub code |
| `ret_msg` | String | Yes | sub code info |


### GlobalEticketMerchantMaFailsend
`POST` `/eticket/ma/failsend`

**Description:** the callback interface when send code failed

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `biz_type` | Number | Yes | biz type |
| `sub_code` | String | Yes | fail reason code |
| `outer_id` | String | Yes | order id |
| `sub_msg` | String | Yes | fail reason desc |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resp_body` | Object | Yes | response body |
| `ret_code` | String | Yes | result code |
| `ret_msg` | String | Yes | result info |


### GlobalEticketMerchantMaQuery
`GET/POST` `/eticket/ma/query`

**Description:** the callback interface that query lazada platform ma

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `code` | String | Yes | code |
| `seller_id` | Number | Yes | sellerId |
| `store_id` | Number | No | storeId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resp_body` | Object | Yes | response |
| `ret_code` | String | Yes | ret code |
| `ret_msg` | String | Yes | ret msg |


### GlobalEticketMerchantMaQueryTbMa
`GET/POST` `/eticket/ma/queryTbMa`

**Description:** the callback interface that query tb ma

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `code` | String | Yes | code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resp_body` | Object | Yes | response |
| `ret_code` | String | Yes | sub code |
| `ret_msg` | String | Yes | sub code info |


### GlobalEticketMerchantMaSend
`GET/POST` `/eticket/ma/send`

**Description:** the callback interface when merchant send code successful

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `biz_type` | Number | Yes | biz type |
| `isv_ma_list` | Object[] | Yes | ma list |
| `outer_id` | String | Yes | order id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resp_body` | Object | Yes | response |
| `ret_code` | String | Yes | sub code |
| `ret_msg` | String | Yes | sub code info |


### RedeemOrderItems
`GET/POST` `/eticket/code/consume`

**Description:** Certificate Consume Open API

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `biz_type` | Number | Yes | biz type |
| `code` | String | Yes | certificate code |
| `outer_id` | String | Yes | outer id |
| `serial_num` | String | Yes | consume serial number |
| `consume_num` | Number | Yes | consume num |
| `store_id` | String | No | store id |
| `pos_id` | String | No | pos id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | response body |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `100` | E100: Param Invalid, "%s" | Param invalid |
| `101` | E101: Redemption Operator Invalid | The certificate not belongs to the seller |
| `200` | E200: Certificate Not Exist | Certificate not exist |
| `202` | E202: Certificate Can Not Distinguish | Can't distinguish the business type of this code |
| `203` | E203: Certificate Order Not Exist | No matched certificate of the outerId |
| `301` | E301: Certificate Not Available | Certificate status available, can't redeem |


---
## LazPay API

_Payment information_

### ConsultPayment
`GET/POST` `/lazadapay/v1/debit/consult_payment`

**Description:** The interface is used for consult pay view. Will return pay view info including balance, coupon, credit card etc. If we have no available coupon, we will return pay method view with an empty list of coupon. 

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `serviceCode` | String | Yes | Indentifier for service |
| `payFrom` | Object | Yes | Where is the money to be received, the receivable details, including the user and payment amount information |
| `payTos` | Object[] | No | Details payable, including sellers and amount |
| `orderGroup` | Object | No | Multi Orders Information |
| `envInfo` | String | No | Environment info from buyer |
| `payOptions` | String[] | No | pay simulate when payOptions is not null |
| `productExt` | String | No | Additional Info for payment product |
| `additionalInfo` | String | No | Additional Info |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `responseMessage` | String | Yes | Response code |
| `responseCode` | String | Yes | Response message |
| `errorCode` | String | No | Error Code |
| `additionalInfo` | String | Yes | Additional Info |
| `payOptions` | Object[] | Yes | Available payment option to user |


### CreateSubscriptionToFusion
`POST` `/insurance/subscription/create`

**Description:** Create User Subscription To Fusion

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `subscriptionStatus` | String | Yes | Subscription Status |
| `subscribeTime` | Number | No | Subscribe Time |
| `unsubscribeTime` | Number | No | Unsubscribe Time |
| `subscribeSource` | String | No | Subscribe Source |
| `unsubscribeSource` | String | No | Unsubscribe Source |
| `userToken` | String | Yes | User Id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `subscriptionStatus` | String | No | Subscription Status |
| `subscribeTime` | Number | No | Subscribe Time |
| `unsubscribeTime` | Number | No | Unsubscribe Time |


### DGUtiityPreCreateOrder
`GET/POST` `/digital/service/createorder`

**Description:** This API provides an open interface for partner users to create DG orders

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `miniToken` | String | Yes | mini token |
| `miniappId` | String | Yes | minapp id |
| `paymentRequestId` | String | Yes | partner order id |
| `extendInfo` | String | No | extend message |
| `signature` | String | No | md5 signture |
| `value` | String | Yes | price |
| `currency` | String | Yes | currency |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | true or false |
| `resultCode` | String | No | result code |
| `resultMsg` | String | No | result message |
| `tradeNo` | String | No | trade no |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `00` | sucess | sueccss |


### DGUtilityPreGetPaymentStatus
`GET/POST` `/digital/service/getPaymentStatus`

**Description:** get payment status

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `paymentRequestId` | String | Yes | paymentRequestId |
| `miniappId` | String | Yes | miniappId |
| `signature` | String | Yes | signature |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | true/false |
| `resultCode` | String | No | resultCode |
| `resultMsg` | String | No | resultMsg |


### DGUtilityPreUpdateFulfillemtStatus
`GET/POST` `/digital/service/updateFulfillemtStatus`

**Description:** update fulfillemt status

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `paymentRequestId` | String | Yes | paymentRequestId |
| `miniappId` | String | Yes | miniappId |
| `signature` | String | Yes | signature |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | true/false |
| `resultCode` | String | No | resultCode |
| `resultMsg` | String | No | resultMsg |


### DigitalAlterOrderStatus
`GET/POST` `/digital/order/alterStatus`

**Description:** Change Lazada Digital Order Status

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `requestId` | String | Yes | Reuqest id. |
| `transactionId` | Number | Yes | Third Party's orderId. |
| `sellerId` | Number | No | Seller id. |
| `cancelCode` | Number | No | If not null, then will do alarm in DG. |
| `cancelMsg` | String | No | Sent with the cancelCode. |
| `userToken` | String | Yes | Lazada user token. |
| `serviceName` | String | Yes | Lazada user token. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `traceId` | String | No | Lazada traceId. |
| `transactionId` | Number | No | Third Party's orderId. |
| `orderStatus` | String | No | If have, then order final status. |
| `paymentStatus` | String | No | Lazada order payment status. |
| `resultCode` | Number | No | Result code from Lazada. |


### DigitalCreateOrder
`GET/POST` `/digital/order/create`

**Description:** Create Digital Virtual Order

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `requestId` | String | Yes | Request id. |
| `itemPrice` | Number | Yes | Item price. |
| `currency` | String | Yes | Currency. |
| `transactionId` | Number | Yes | Third party's transactionId. |
| `sellerId` | Number | No | Seller id. |
| `userToken` | String | Yes | Token for Lazada User. |
| `serviceName` | String | Yes | Service name. |
| `skuId` | Number | Yes | Lazada sku id. |
| `itemId` | Number | Yes | Lazada item id. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transactionId` | Number | No | Third party's transactionId. |
| `paymentLink` | String | No | PaymentLink. |
| `resultCode` | Number | No | ResultCode. |
| `tradeOrderLineId` | String | No | Lazada's tradeOrderLine id. |
| `traceId` | String | No | Lazada's traceId. |


### DigitalQueryOrder
`GET/POST` `/digital/order/getStatus`

**Description:** Query Lazada Digital Order Status

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `requestId` | String | Yes | Reuqest id. |
| `transactionId` | Number | Yes | Third Party's transactionId. |
| `sellerId` | Number | No | Seller id. |
| `serviceName` | String | Yes | Service name. |
| `userToken` | String | Yes | Lazada user token. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transactionId` | Number | No | Third Party's transactionId. |
| `orderStatus` | String | No | If have, then order final status. |
| `paymentStatus` | String | No | Lazada order payment status. |
| `resultCode` | Number | No | Result code from Lazada. |
| `traceId` | String | No | Lazada traceId. |


### GetSubscriptionToFusion
`GET/POST` `/insurance/subscription/getSubscription`

**Description:** Get User Subscription To Fusion

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userToken` | String | Yes | User Id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `subscriptionStatus` | String | No | Subscription Status |
| `subscribeTime` | Number | No | Subscribe Time |
| `unsubscribeTime` | Number | No | Unsubscribe Time |


### InsuranceAlterOrderStatus
`GET/POST` `/insurance/order/alterStatus`

**Description:** Change Lazada Insurance Order Status

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `requestId` | String | Yes | Reuqest id. |
| `transactionId` | Number | Yes | Fusion's orderId. |
| `sellerId` | Number | No | Seller id. |
| `cancelCode` | Number | No | If not null, then will do alarm in DG. |
| `cancelMsg` | String | No | Sent with the cancelCode. |
| `userToken` | String | Yes | Lazada user token. |
| `serviceName` | String | Yes | Service name. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transactionId` | Number | No | Fusion's orderId. |
| `orderStatus` | String | No | If have, then order final status. |
| `paymentStatus` | String | No | Lazada order payment status |
| `resultCode` | Number | No | Result code from Lazada. |
| `traceId` | String | No | Lazada traceId. |


### InsuranceCreateOrder
`GET/POST` `/insurance/order/create`

**Description:** Lazada Insurance Create Order

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `requestId` | String | Yes | Request ID, unique for each request.aRequest ID, unique for each request.Fusion's product ID. |
| `productCode` | String | Yes | Fusion's product ID. |
| `itemPrice` | Number | Yes | Price that user need to pay. (Totally price) |
| `sstFee` | Number | Yes | SST amount. |
| `stampDuty` | Number | Yes | Stamp Duty amont. |
| `currency` | String | Yes | Currency Type. |
| `transactionId` | Number | Yes | Fusion's order ID. |
| `sellerId` | Number | No | Seller ID. |
| `serviceName` | String | Yes | Service name. |
| `userToken` | String | Yes | Token for Lazada User. |
| `orderExistTime` | String | No | Lazada order persit time. |
| `subProductCode` | String | No | Road tax's product code. |
| `subItemPrice` | String | No | Road tax's item price. (Totally price) |
| `subServiceFee` | String | No | Road tax's service fee. |
| `subTransactionId` | String | No | Road tax's transactionId. |
| `insuranceType` | String | No | Marketplace insurance type. |
| `partnerCode` | String | No | Traffic source. |
| `plateNo` | String | No | Car plate no. |
| `planCode` | String | No | planCode |
| `subPlanCode` | String | No | subPlanCode |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `tradeOrderLineId` | String | No | Lazada tradeOrderLine ID. |
| `transactionId` | Number | No | Fusion's order ID. |
| `paymentLink` | String | No | Lazada Independent Paymen Link. |
| `resultCode` | Number | No | Result code from Lazada. |
| `traceId` | String | No | Lazada traceId. |


### InsuranceGetPromotions
`GET/POST` `/insurance/promotion/getPromotions`

**Description:** get lazada marketplace  ump promotions

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | Yes | 主体信息 |
| `userToken` | String | Yes | userToken |
| `serviceName` | String | Yes | serviceName |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `traceId` | String | No | traceId |
| `data` | String | No | data |
| `resultCode` | Number | No | resultCode |
| `resultMessage` | String | No | message |


### InsuranceQueryOrder
`GET/POST` `/insurance/order/getStatus`

**Description:** Query Lazada Insurance Order Status

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `requestId` | String | Yes | Reuqest id. |
| `transactionId` | Number | Yes | Fusion's transactionId. |
| `sellerId` | Number | No | Seller id. |
| `serviceName` | String | Yes | Service name. |
| `userToken` | String | Yes | Lazada user token. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transactionId` | Number | No | Fusion's transactionId. |
| `orderStatus` | String | No | If have, then order final status. |
| `paymentStatus` | String | No | Lazada order payment status |
| `resultCode` | Number | No | Result code from Lazada. |
| `traceId` | String | No | Lazada traceId. |


### LazPayPaymentNotify
`HSF` `/lazpay/v1/payment/notify`

**Description:** Payment Notify

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `paymentRequestId` | String | Yes | paymentRequestId |
| `paymentId` | String | Yes | paymentId |
| `paymentAmount` | String | Yes | paymentAmount |
| `paymentStatus` | String | Yes | paymentStatus |
| `paymentApplyTime` | Number | Yes | paymentApplyTime |
| `paymentFinishTime` | Number | No | paymentFinishTime |
| `productCode` | String | No | productCode |
| `merchantInfo` | String | No | merchantInfo |
| `promotionInfo` | String | No | promotionInfo |
| `userPaymentAmount` | String | No | userPaymentAmount |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### LazadaCFOInvoiceRpaCallback
`GET/POST` `/rpa/id/tax/callback`

**Description:** Call RPA and return the official invoice

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `country` | String | Yes | Country |
| `batch_id` | String | Yes | Batch ID |
| `status` | String | Yes | status |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `is_success` | Boolean | No | true or false |
| `res_code` | String | No | if success,it is null |
| `content` | String | No | if success,it is null |
| `res_msg` | String | No | Error message |


### OpenServiceBalanceQuery
`GET/POST` `/wallet/open/service/balance/query`

**Description:** Open Service Account Balance Info Query

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `date_time` | Number | Yes | date time |
| `available_amount` | String | Yes | amount |
| `available_amount_cent` | Number | Yes | cent |
| `currency` | String | Yes | currency |


### OpenServiceKycQuery
`GET/POST` `/wallet/open/service/kyc/query`

**Description:** Open Service User KYC Info Query

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `need_cert_info` | Boolean | No | True means need KYC Info photo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `phone` | String | No | phone number |
| `prefix` | String | No | phone number prefix |
| `userId` | String | No | open platform user id |
| `birthday` | String | No | birthday, format is yyyy-MM-dd |
| `full_name` | String | No | full name |
| `cert_front_image` | String | No | certificate front image |
| `cert_type` | String | No | certificate type |
| `full_kyc_status` | Boolean | No | whether user has passed full kyc or not |
| `kyc_jump_url` | String | No | redirect url to let user finish kyc in lazada app |
| `extend_info` | String | No | extend infos |


### OpenServiceWithdrawApply
`GET/POST` `/wallet/open/service/withdraw`

**Description:** Open Service Withdraw Apply

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `withdraw_request_id` | String | Yes | ISV withdraw request id |
| `withdrawable` | Boolean | Yes | withdrawable feature |
| `withdraw_amount` | String | Yes | withdraw amount，precise to two decimal places. |
| `user_id` | String | Yes | LazOp user id |
| `need_verify_full_kyc` | Boolean | No | whether full kyc validation is needed in lazada, default false. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `withdraw_request_id` | String | No | ISV withdraw request id |
| `withdraw_id` | String | No | Lazada withdraw id |
| `withdraw_amount` | String | No | withdraw amount，precise to two decimal places. |
| `withdrawable` | String | No | withdrawable feature |
| `currency` | String | No | currency |
| `partner_deposit` | String | No | The available balance of ISV |


### OpenServiceWithdrawQuery
`GET/POST` `/wallet/open/service/withdraw/query`

**Description:** Open Service Withdraw Query

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `withdraw_request_id` | String | Yes | ISV withdraw request id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `withdraw_request_id` | String | No | ISV withdraw request id |
| `withdraw_id` | String | No | Lazada withdraw id |
| `withdraw_amount` | String | No | withdraw amount，precise to two decimal places. |
| `withdrawable` | String | No | withdrawable feature |
| `currency` | String | No | currency |
| `partner_deposit` | String | No | The available balance of ISV |


### Reconciliation
`GET/POST` `/wallet/open/service/reconciliation`

**Description:** Reconciliation

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `date` | String | Yes | A date in the format of "yyyy-mm-dd" |
| `business_type` | String | Yes | withdraw |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `res` | String | No | The reconciliation file encoded by base64, user needs to decode it into a readable csv file. |


### collectBenefit
`GET/POST` `/insurance/promotion/collectBenefit`

**Description:** collect lazada marketplace benefit

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | Yes | 主体信息 |
| `userToken` | String | Yes | userToken |
| `serviceName` | String | Yes | serviceName |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `trace_id` | String | Yes | trace |
| `resultCode` | Number | Yes | resultCode |
| `resultMessage` | String | Yes | resultMessage |
| `data` | String | No | data |


### insuranceRealTimeCDP
`GET/POST` `/insurance/syncCDP`

**Description:** 用户完成操作后，实时更新CDP人群

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `userToken` | String | Yes | Token for Lazada User. |
| `bizCode` | String | Yes | business code |
| `serviceName` | String | Yes | business type |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | String | No | 接口成功=true，接口失败=true， 系统失败=false |
| `resultCode` | String | No | 业务成功=SUCCESS， 业务失败=SUCCESS ，系统失败=SYSTEM_ERROR |
| `resultMessage` | String | No | 接口成功=Success ，接口失败=Success ，系统失败=System Error |
| `data` | Boolean | No | 业务成功=true，业务失败=false，系统失败=false |
| `redirectUrl` | String | No | 无 |


### queryAddonOrder
`GET/POST` `/insurance/addon/orders/query`

**Description:** list user  addon order detail 

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `pageNum` | Number | Yes | pageNum |
| `pageSize` | Number | Yes | pageSize |
| `userToken` | String | Yes | userToken |
| `orderStatus` | String | No | orderStatus |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `redirectUrl` | String | Yes | redirectUrl |
| `resultCode` | String | Yes | resultCode |
| `data` | Object | Yes | data |
| `success` | Boolean | Yes | success |
| `resultMessage` | String | Yes |  resultMessage |


### queryBenefit
`GET/POST` `/insurance/promotion/queryBenefit`

**Description:** get lazada marketplace benefit

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | String | Yes | 主体信息 |
| `userToken` | String | Yes | userToken |
| `serviceName` | String | Yes | serviceName |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `trace_id` | String | Yes | trace |
| `resultCode` | Number | Yes | resultCode |
| `resultMessage` | String | Yes | resultMessage |
| `data` | String | No | data |


### redeemMpVoucher
`GET/POST` `/insurance/voucher/redeemVoucher`

**Description:** 商城险域外voucher核销

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucherCode` | String | Yes | voucherCode |
| `userToken` | String | Yes | userToken |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `voucherTemplateId` | String | No | voucherTemplateId |
| `traceId` | String | No | traceId |
| `resultCode` | String | No | resultCode |
| `resultMessage` | String | No | resultMessage |
| `brokerName` | String | No | MSIG |


---
## Lazada Wallet Corporate Top-up API

_Lazada Wallet Corporate Top-up_

### DirectTransferQuery
`GET/POST` `/wallet/transfer/query`

**Description:** Direct Transfer - Query

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transfer_order_id` | String | Yes | ISV transfer order id, length <= 32 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `amount` | String | No | Transfer amount，precise to two decimal places. |
| `account_number` | String | No | Email or phone number, accepted phone number starts with (PH: +638, +639, 08, 09, 638, 639) |
| `transfer_order_id` | String | No | ISV transfer order id, length <= 32 |
| `transfer_request_id` | String | No | Lazada transfer order id |
| `deposit` | String | No | The available balance of ISV |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `TRANSFER_ERROR_MSG_RESPONSED_FAILED` | Error happens when transferring，please contact lazada team | Error happens when transferring，please contact lazada team |
| `OPEN_DIRECT_TRANSFER_INTERNAL_FAIL` | Direct transfer internal error, please retry or contact lazada tech team. | Direct transfer internal error, please retry or contact lazada tech team. |
| `TRANSFER_ERROR_MSG_AMOUNT_INVALID` | Amount is invalid | Amount is invalid |
| `APP_KEY_INVALID` | App key is invalid, please contact lazada tech team. | App key is invalid, please contact lazada tech team. |
| `USER_IS_NOT_LOGGED_IN` | The user is not logged in | The user is not logged in |
| `PROCEED_TRANSFER_EXCEPTION` | Internal error, please retry or contact lazada tech team. | Internal error, please retry or contact lazada tech team. |
| `OPEN_API_CALL_EXCEED_LIMIT` | Open Api call times exceeds: apiName_limitType | Open Api call times exceeds: apiName_limitType |
| `TRANSFER_ERROR_NATION_NOT_IN_LIST` | The current user's region does not have permission to access, please contact the lazada tech team. | The current user's region does not have permission to access, please contact the lazada tech team. |
| `USER_BALANCE_NOT_ENOUGH` | The available deposit is not enough for the transfer. | The available deposit is not enough for the transfer. |
| `TRANSFER_AMOUNT_EXCEED_LIMIT` | The transfer amount has exceeded the limit. | The transfer amount has exceeded the limit. |


### DirectTransferRequest
`GET/POST` `/wallet/transfer/request`

**Description:** Direct Transfer - Request to transfer

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `amount` | String | Yes | Transfer amount，precise to two decimal places. |
| `transfer_order_id` | String | Yes | ISV transfer order id，length <= 32 |
| `account_number` | String | Yes | Phone number or email address，accepted phone number starts with (PH : +639, +638, 08, 09, 638, 639) |
| `withdrawable` | Boolean | No | The funds type for transfers. Set true for funds that can be withdrawn and false for funds that cannot be withdrawn. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `account_number` | String | No | The email or phone number of user to be transferred to |
| `transfer_order_id` | String | No | ISV input transfer order id |
| `transfer_request_id` | String | No | Lazada transfer order id |
| `amount` | String | No | The amount to transfer |
| `deposit` | String | No | The available balance of ISV |
| `withdrawable` | Boolean | No | The funds type for transfers. Set true for funds that can be withdrawn and false for funds that cannot be withdrawn. |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `OPEN_DIRECT_TRANSFER_LOCK_CONFLICT` | Direct transfer request is already being processed，please wait for a moment and check status | Direct transfer request is already being processed，please wait for a moment and check status |
| `TRANSFER_ERROR_MSG_RESPONSED_FAILED` | Error happens when transferring，please contact lazada team | Error happens when transferring, please contact lazada team |
| `TRANSFER_VALUE_UNMATCHED` | Transfer amount does not match | Transfer amount does not match, please enter same amount |
| `TRANSFER_USER_UNMATCHED` | User to be transferred not match | User to be transferred not match, please use same account |
| `TRANSFER_ERROR_ACCOUNT_NUMBER_INVALID` | Account number is invalid | Please check and re-enter your account number |
| `OPEN_DIRECT_TRANSFER_INTERNAL_FAIL` | Direct transfer internal error, please retry or contact lazada tech team. | Direct transfer internal error, please retry or contact lazada tech team. |
| `TRANSFER_ERROR_TRANSFER_ORDER_ID_INVALID` | Transfer order ID is invalid | Please check and re-enter your transfer order ID |
| `TRANSFER_ERROR_MSG_AMOUNT_INVALID` | Amount is invalid | Please check and re-enter your amount |
| `APP_KEY_INVALID` | App key is invalid, please contact lazada tech team. | App key is invalid, please contact lazada tech team. |
| `USER_IS_NOT_LOGGED_IN` | The user is not logged in | Please log in  your account  |
| `PROCEED_TRANSFER_EXCEPTION` | Internal error, please retry or contact lazada tech team. | Internal error, please retry or contact lazada tech team. |
| `OPEN_API_CALL_EXCEED_LIMIT` | Open Api call times exceeds: apiName_limitType | Open Api call times exceeds, please contact lazada tech team or retry later |
| `BIZ_DEGRADATION_ERROR` | The service is not available now | The service is not available now, please retry or contact lazada tech team |
| `TRANSFER_ERROR_MSG_WALLET_INACTIVATED` | The transfer account has not activated the wallet | The transfer account has not activated the wallet, please activate your wallet |
| `TRANSFER_ERROR_MSG_USER_NOT_FOUND` | User to be transferred not found. | User to be transferred not found, please check your account or contact the lazada tech team |
| `USER_BALANCE_NOT_ENOUGH` | The available deposit is not enough for the transfer. | The available deposit is not enough for the transfer, please top up or reduce the transfer amount |
| `TRANSFER_AMOUNT_EXCEED_LIMIT` | The transfer amount has exceeded the limit. | The transfer amount has exceeded the limit, please reduce the transfer amount |
| `TRANSFER_IS_CORPORATE_USER_ERROR` | The recipient account is corporate user. | The recipient account is corporate user, please change the recipient account |
| `TRANSFER_ERROR_NATION_NOT_IN_LIST` | The current user's region does not have permission to access, please contact the lazada tech team. | The current user's region does not have permission to access, please contact the lazada tech team |
| `RISK_REJECT` | The transfer recipient's account status is abnormal, please check | The transfer recipient's account status is abnormal, please check |
| `TRANSFER_WITHDRAWABLE_UNMATCHED` | Transfer withdrawable does not match. | Transfer withdrawable does not match. |


### GiftCodeQuery
`GET/POST` `/wallet/giftcode/query`

**Description:** Gift Code - Query

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `page` | Number | Yes | The page to query, page should > 0 and < the total pages, default value is 1 if this parameter is null.  |
| `transfer_order_id` | String | Yes | Transfer order Id on the ISV side, length <= 32 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `records` | String[] | No | The list of gift codes, need to finish unmask verification firstly. |
| `total_page` | Number | No | The total page number of the code list |
| `current_page` | Number | No | The current queried page of the code list |
| `page_size` | Number | No | The default max number of codes contained in one page. |
| `transfer_order_id` | String | No | Transfer order Id on the ISV side, length <= 32 |
| `total_number` | String | No | The amount of created gift code, precise to two decimal places |
| `create_status` | String | No | The create status of the gift code |
| `deposit` | String | No | The available balance of ISV |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `GIFT_CODE_LOCK_CONFLICT` | Gift code is already being created，please wait for a moment and check the batch list | Gift code is already being created，please wait for a moment and check the batch list |
| `OPEN_API_CALL_EXCEED_LIMIT` | Open Api call times exceeds: apiName_limitType | Open Api call times exceeds: apiName_limitType |
| `PROCEED_TRANSFER_EXCEPTION` | Internal error, please retry or contact lazada tech team. | Internal error, please retry or contact lazada tech team. |
| `USER_IS_NOT_LOGGED_IN` | The user is not logged in | The user is not logged in |
| `APP_KEY_INVALID` | App key is invalid, please contact lazada tech team. | App key is invalid, please contact lazada tech team. |
| `TRANSFER_ERROR_TRANSFER_ORDER_ID_INVALID` | Transfer order ID is invalid | Transfer order ID is invalid |
| `GIFT_CODE_QUERY_EMPTY` | There are no such gift code | There are no such gift code |


### GiftCodeRequest
`GET/POST` `/wallet/giftcode/request`

**Description:** Gift Code - Request

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `amount` | String | Yes | The amount of each gift code, precise to two decimal places |
| `quantity` | Number | Yes | The quantity of gift codes to be created |
| `transfer_order_id` | String | Yes | ISV transfer order id，length <= 32 |
| `end_timestamp` | Number | Yes | End timestamp，13 bits |
| `start_timestamp` | Number | Yes | Start timestamp，13 bits |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transfer_order_id` | String | No | ISV transfer order id |
| `total_number` | Number | No | Total gift code quantity |
| `create_status` | String | No | Create status of gift code |
| `deposit` | String | No | The available balance of ISV |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `OPEN_API_TIMESTAMP_INVALID` | The input timestamp is invalid | The input timestamp is invalid |
| `BIZ_DEGRADATION_ERROR` | The service is not available now | The service is not available now |
| `OPEN_API_CALL_EXCEED_LIMIT` | Open Api call times exceeds: apiName_limitType | Open Api call times exceeds: apiName_limitType |
| `PROCEED_TRANSFER_EXCEPTION` | Internal error, please contact lazada tech team | Internal error, please contact lazada tech team |
| `USER_IS_NOT_LOGGED_IN` | The user is not logged in | The user is not logged in |
| `APP_KEY_INVALID` | App key is invalid, please contact lazada tech team. | App key is invalid, please contact lazada tech team. |
| `TRANSFER_ERROR_TRANSFER_ORDER_ID_INVALID` | Transfer order ID is invalid | Transfer order ID is invalid |
| `TRANSFER_ERROR_MSG_AMOUNT_INVALID` | Amount is invalid | Amount is invalid |
| `TRANSFER_ERROR_MSG_QUANTITY_INVALID` | The quantity of gift code is invalid | The quantity of gift code is invalid, only under test case. |
| `GIFT_CODE_LOCK_CONFLICT` | Gift code is already being created，please wait for a moment and check the batch list. | Gift code is already being created，please wait for a moment and check the batch list. |
| `BATCH_CREATE_ERROR` | Error happens when creating gift code. Please Retry. | Error happens when creating gift code. |
| `BALANCE_ACCOUNT_NOT_ENOUGH` | Balance account is not enough. | Balance account is not enough. |
| `TRANSFER_ERROR_NATION_NOT_IN_LIST` | The current user's region does not have permission to access, please contact the lazada tech team. | The current user's region does not have permission to access, please contact the lazada tech team. |


### Reconciliation
`GET/POST` `/wallet/open/reconciliation`

**Description:** Corporate TopUp - Reconciliation

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `date` | String | Yes | A date in the format of "yyyy-mm-dd" |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `res` | String | No | The reconciliation file encoded by base64, user needs to decode it into a readable csv file. |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `RECONCILIATION_INPUT_DATE_INVALID` |  Invalid input format of local date. |  Invalid input format of local date. |
| `ECONCILIATION_CSV_ERROR_FAILED` | Error happens when creating reconciliation file. | Error happens when creating reconciliation file. |
| `BIZ_DEGRADATION_ERROR` | The service is not available now. | The service is not available now. |


---
## RedMart API

_API for RedMart_

### RssGetOnePickupJob
`GET/POST` `/rss/pickup-job/get`

**Description:** Get details of a pickup job

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | store id |
| `pickupJobId` | Number | Yes | pickup job id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### RssGetPickupJobs
`GET/POST` `/rss/pickup-jobs/get`

**Description:** Retrieve RSS pickup jobs based on time range and status filter.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | Store id |
| `from` | Number | Yes | Epoch millis of job date from |
| `till` | Number | Yes | Epoch millis of job date till |
| `statuses` | String | No | Job statuses filter. Possible job statuses are "pending", "arrived", "pickedup", "cancelled" and "failed". Concatenate statuses of interest with "," to make queries with multiple job status filter. Leave this field blank or null to query without filtering. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | Result  |


### RssGetPickupLocations
`GET` `/rss/pickupLocations/get`

**Description:** rss get pickupLocations by storeId

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | store id |
| `page` | Number | Yes | page |
| `pageSize` | Number | Yes | pageSize |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### RssGetProduct
`GET` `/rss/product/get`

**Description:** get rss product by storeId and productId

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | store id |
| `productId` | Number | Yes | the RPC of the Product  |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | reuslt |


### RssGetProducts
`GET` `/rss/products/get`

**Description:** rss get products paged by storeId and pickupLocationId

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | store id |
| `pickupLocationIds` | Number[] | No | pickup location ids |
| `page` | Number | Yes | page |
| `pageSize` | Number | Yes | page size |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### RssGetStockLot
`GET` `/rss/stockLot/get`

**Description:** rss get stockLot

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | store id |
| `pickupLocationId` | Number | Yes | pickupLocation id |
| `productId` | Number | Yes | the RPC of the Product |
| `stockLotId` | String | Yes | Identifier of the requested Stock Lot. For now always hardcoded to "0" (please note the String type, do not always expect it to be a number !) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### RssGetStockLots
`GET` `/rss/stockLots/get`

**Description:** rss get stockLots

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | store id |
| `pickupLocationId` | Number | Yes | pickupLocation id |
| `productId` | Number | Yes | product id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


### RssUpdateStockLot
`GET/POST` `/rss/stockLot/update`

**Description:** rss update stockLot

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `storeId` | Number | Yes | store id |
| `pickupLocationId` | Number | Yes | The unique id of the pickup location where the product is stored |
| `productId` | Number | Yes | the RPC of the Product (so the RedMart-specific code, not the merchant-specific code) |
| `stockLotId` | String | Yes | Identifier of the requested Stock Lot. For now always hardcoded to "0" (please note the String type, do not always expect it to be a number !) |
| `stockLotUpdateDTO` | Object | Yes | stockLot update DTO |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


---
## Lazada DG API

_lazada digital good  API_

### InstallServiceCallBack
`GET/POST` `/digital/install/servicecallback`

**Description:** Install the service callback interface

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `orderNo` | String | Yes | service provider company orderId |
| `thirdOrderNo` | String | Yes | LZD orderLineId |
| `type` | String | Yes | type = 1 (mean install sevice finish)   type = 2(mean install update). type =3 (mean cancel install service) |
| `servicePrice` | String | No | install service price |
| `serviceDate` | String | No | install service date |
| `jobStatus` | String | No | The installation status of the external company |
| `jobReason` | String | No | Reasons for success or failure |
| `extendInfo` | String | No | extendInfo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resultCode` | String | No | result code |
| `resultMsg` | String | No | result message |
| `transactionId` | String | No | LZD orderLineId |
| `extendInfo` | String | No | extendInfo |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `00` | transaction success | transaction success |
| `01` | cancel success | cancel success |
| `02` | update serviceDate success | update serviceDate success |
| `99` | fail | fail |
| `11` | orderNo is empty | orderNo is empty |
| `12` | thirdOrderNo is empty | thirdOrderNo is empty |
| `13` | type is empty | type is empty |
| `14` | type not exist | type not exist |
| `15` | servicePrice is empty | servicePrice is empty |
| `16` | jobStatus is empty | jobStatus is empty |
| `17` | serviceDate is empty | serviceDate is empty |
| `21` | order processing | order processing |
| `31` | parse extendInfo to map fail | parse extendInfo to map fail |
| `32` | date format is wrong | date format is wrong |


### InstallServiceCallBack
`GET/POST` `/digital/test/install/servicecallback`

**Description:** Install the service callback interface

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `orderNo` | String | Yes | service provider company orderId |
| `thirdOrderNo` | String | Yes | LZD orderLineId |
| `type` | String | Yes | type = 1 (mean install sevice finish)   type = 2(mean install update). type =3 (mean cancel install service) |
| `servicePrice` | String | No | install service price |
| `serviceDate` | String | No | install service date |
| `jobStatus` | String | Yes | The installation status of the external company |
| `jobReason` | String | No | Reasons for success or failure |
| `extendInfo` | String | No | extendInfo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resultCode` | String | No | result code |
| `resultMsg` | String | No | result message |
| `transactionId` | String | No | LZD orderLineId |
| `extendInfo` | String | No | extendInfo |


### InstallServiceCallBackForTest
`GET/POST` `/digital/install/test/servicecallback`

**Description:** Install the service callback interface

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `orderNo` | String | Yes | service provider company orderId |
| `thirdOrderNo` | String | Yes | LZD orderLineId |
| `type` | String | Yes | type = 1 (mean install sevice finish)   type = 2(mean install update). type =3 (mean cancel install service) |
| `servicePrice` | String | No | install service price |
| `serviceDate` | String | No | install service date |
| `jobStatus` | String | Yes | The installation status of the external company |
| `jobReason` | String | No | Reasons for success or failure |
| `extendInfo` | String | No | extendInfo |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `resultCode` | String | No | result code |
| `resultMsg` | String | No | result message |
| `transactionId` | String | No | LZD orderLineId |
| `extendInfo` | String | No | extendInfo |


### InuranceNotication
`GET/POST` `/digital/insurance/notification`

**Description:** Third party insurance company callback interface


**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `orderNo` | String | Yes | Insurance company order number |
| `thirdOrderNo` | String | Yes | lazada orderId |
| `premium` | String | Yes | premium |
| `ePolicyLink` | String | Yes | ePolicy Link |
| `policyNo` | String | Yes | Policy No |
| `underwritingStatus` | String | Yes | Order Status |
| `underwritingReason` | String | No | Order Message |
| `expirationDate` | String | Yes | expirationDate |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `errorCode` | String | No | 错误码 |
| `errorMsg` | String | No | 错误信息 |
| `transactionId` | String | No | 交易Id |
| `extendInfo` | String | No | 拓展信息 |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `11` | policyNo is empty | policyNo is empty |
| `12` | orderNo is empty | orderNo is empty |
| `13` | thirdOrderNo is empty | lazada orderId |
| `14` | ePolicyLink is empty | Insurance information link |
| `15` | underwritingStatus is empty | Insurance status |
| `16` | premium is empty | premium is empty |
| `21` | order processing | The order is already being processed |
| `00` | success | success |
| `99` | fail | fail |


### InuranceNotication
`GET/POST` `/digital/insurance/test/notificationcopy`

**Description:** Third party insurance company callback interface


**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `orderNo` | String | Yes | Insurance company order number |
| `thirdOrderNo` | String | Yes | lazada orderId |
| `premium` | String | Yes | premium |
| `ePolicyLink` | String | Yes | ePolicy Link |
| `policyNo` | String | Yes | Policy No |
| `underwritingStatus` | String | Yes | Order Status |
| `underwritingReason` | String | No | Order Message |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `errorCode` | String | No | 错误码 |
| `errorMsg` | String | No | 错误信息 |
| `transactionId` | String | No | 交易Id |
| `extendInfo` | String | No | 拓展信息 |


### InuranceNotifyLapse
`GET/POST` `/digital/insurance/notificationlapse`

**Description:** Insurance company push the callback notification to partners once the policy has been cancelled successfully

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `orderNo` | String | Yes | 1234 |
| `thirdOrderNo` | String | Yes | 12344 |
| `policyNo` | String | Yes | 1234 |
| `lapseTime` | String | Yes | 1234 |
| `lapseType` | String | Yes | enum： expiration: policy expired. end: the customer has used up the sum insured amount, policy end. |
| `message` | String | No | expire |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `transactionId` | String | No | LZD orderLineId |
| `extendInfo` | String | No | extendInfo |
| `errorCode` | String | No | result code |
| `errorMsg` | String | No | result message |


### digitalServiceCdkCodeReceived
`POST` `/digital/service/cdkCodeReceived`

**Description:** 接受码商发码请求，给用户发送码。

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `tb_order_id` | String | Yes | 淘天主订单号 |
| `cdk_name` | String | No | 商品名称 |
| `cdk_code_items` | Object[] | Yes | CDK码对象 |
| `tb_order_line_id` | String | Yes | 淘天子订单号 |
| `valid_from` | String | No | 有效期起始时间，YYYY-MM-DD格式 |
| `cdk_code_number` | String | Yes | CDK码数量 |
| `valid_end` | String | No | 有效期结束时间，YYYY-MM-DD格式 |
| `terms_use` | String | No | 使用规则/条款等说明 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result_code` | String | Yes | 响应状态码 |
| `result_msg` | String | Yes | 响应描述信息 |


---
## Sponsored Solutions API

_The API to createCampaign,updateCampaign._

### addAdgroupBatch
`GET/POST` `/sponsor/solutions/adgroup/addAdgroupBatch`

**Description:** Do add adgroup for one campaign.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignId` | Number | Yes | Campaign id which you want to add into. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `adgroupViewDTOList` | Object[] | Yes | Adgroup list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Boolean | Yes | The detail result, for this api is boolean. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### addSolution
`POST` `/sponsor/solutions/addSolution`

**Description:** Add sponsor solution

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `autoKeyword` | Number | No | Let Lazada automatically set keyword for your products.1:manual(I want to select keywords manually for my product selection.);2:auto(Let Lazada optimize the keywords relating to your products in real time to maximize the campaigns' performance). |
| `endDate` | String | Yes | Campaign end date. |
| `platform` | Number[] | Yes | Placements determine where shoppers will see your promoted products.3:Search Result Page;4:Just For You Page |
| `autoCreative` | Number | Yes | Lazada automatically set creatives for your products.1:ON;0:OFF. |
| `campaignObjective` | Number | Yes | Your campaign objective helps determine your bidding strategy - Traffic objective helps you to increase the number of clicks to your store, while sales objective helps to increase your store’s sales.1:Traffic;2:Sales. |
| `campaignType` | Number | Yes | Unlock different ways to bids, select products, and keywords with campaign types.1:Standard;2:Smart. |
| `campaignModel` | Number | Yes | Fine granularity to distinguish solutions. |
| `maxBid` | String | Yes | Max bid determines the highest amount that you're willing to pay for a click on your promoted product.String type, -1 means no limit. |
| `autoItemSelect` | Number | Yes | The way the product be selected.1:manual(I want to select products manually from my store.);2:auto(Let Lazada optimize the products within the campaigns in real-time to maximize the campaigns' performance) |
| `dayBudget` | String | Yes | Budget indicates the maximum amount you’re willing to pay each day. |
| `campaignName` | String | Yes | Campaign name. |
| `startDate` | String | Yes | Campaign start date. |
| `adgroupViewDTOlistWithFeed` | Object[] | Yes | Adgroup list. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | System result for this api call. |
| `result` | Object | No | The detail result, for this api is boolean. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### clickserver
`GET/POST` `/gproject/ads/aidc/click`

**Description:** aidc click server interface

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cpcClickDO` | Object | No | cookie section |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Result |


### deleteAdgroupBatch
`GET/POST` `/sponsor/solutions/adgroup/deleteAdgroupBatch`

**Description:** Delete adgroup batch.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `adgroupIdList` | Number[] | Yes | Adgroup id list. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Boolean | Yes | The detail result, for this api is boolean. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### deleteCampaign
`GET/POST` `/sponsor/solutions/campaign/deleteCampaign`

**Description:** Delete campaign.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignIdList` | Number[] | Yes | Campaign id list. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Number | Yes | The detail result, for this api is deleted count. |
| `success` | Boolean | Yes | The detail result, for this api is boolean. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### getAccountSignInfo
`GET` `/sponsor/solutions/account/getAccountSignInfo`

**Description:** Get seller account sign status.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The detail result, for this api is boolean. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### getAutoTopUpOptionOneConfig
`GET/POST` `/sponsor/solutions/wallet/getAutoTopUpOptionOneConfig`

**Description:** Get auto top up option one config.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The detail result, for this api is configuration. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### getCampaign
`GET/POST` `/sponsor/solutions/campaign/getCampaign`

**Description:** Get campaign list with bizCode by seller.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `bizCode` | String | Yes | Discovery:sponsoredSearch |
| `campaignId` | Number | Yes | 123 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The detail result, for this api is campaign detail info. |
| `success` | String | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### getCampaignCount
`GET/POST` `/sponsor/solutions/campaign/getCampaignCount`

**Description:** Get campaign count with bizCode for each solution type.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Number | Yes | The detail result, for this api is campaign count. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### getDiscoveryReportAdgroup
`GET/POST` `/sponsor/solutions/report/getDiscoveryReportAdgroup`

**Description:** Get sponsored discovery report adgroup level

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignType` | String | No | Campaign Type,1 standard 2 automated |
| `campaignName` | String | No | Campaign Name, frazzy search |
| `campaignId` | String | No | Campaign Id |
| `adgroupName` | String | No | Adgroup Name |
| `adgroupId` | String | No | Adgroup Id |
| `itemId` | String | No | Item Id |
| `useRtTable` | Boolean | No | It means that if endDate have selected today, and you need realtime data,then set useRtTable=true If useRtTable=false,it will not search realtime data |
| `sort` | String | No | sort column,we have provide some index to sort |
| `pageNo` | String | Yes | Page No，default 1,max=100 |
| `pageSize` | String | Yes | Page No, default 10, max=100 |
| `order` | String | No | ASC or DESC, other String is invalid |
| `startDate` | String | Yes | start date, format like yyyy-MM-dd |
| `endDate` | String | Yes | end date , date, format like yyyy-MM-dd |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Result Details |


### getDiscoveryReportAudience
`GET/POST` `/sponsor/solutions/report/getDiscoveryReportAudience`

**Description:** Get sponsored discovery report audience level

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignName` | String | No | Campaign Name |
| `campaignId` | Number | No | Campaign Id |
| `audienceGroup` | Number | No | Audienct type  1:15 days Visitors 2:Similar Product Visitors 3:Store Awareness Audience 4:Store Interest Audience 5:DMP Crow Audience 6:Gender 7:Age |
| `sort` | String | No | sort column,we have provide some index to sort |
| `order` | String | No | ASC or DESC, other String is invalid |
| `pageNo` | Number | Yes | Page No，default 1,max=100 |
| `pageSize` | Number | Yes | Page No, default 10, max=100 |
| `startDate` | String | Yes | start date, format like yyyy-MM-dd |
| `endDate` | String | Yes | end date , date, format like yyyy-MM-dd |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Details |


### getDiscoveryReportCampaign
`GET/POST` `/sponsor/solutions/report/getDiscoveryReportCampaign`

**Description:** Get sponsored discovery report campaign level

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignId` | Number | No | Campaign Id |
| `useRtTable` | Boolean | No | It means that if endDate have selected today, and you need realtime data,then set useRtTable=true If useRtTable=false,it will not search realtime data |
| `sort` | String | No | sort column,we have provide some index to sort |
| `order` | String | No | ASC or DESC, other String is invalid |
| `startDate` | String | Yes | start date, format like yyyy-MM-dd |
| `endDate` | String | Yes | end date , date, format like yyyy-MM-dd |
| `pageNo` | String | Yes | Page No，default 1,max=100 |
| `pageSize` | String | Yes | Page No, default 10, max=100 |
| `campaignType` | Number | No | Campaign type, 1 Manual  2 Automated |
| `productType` | String | No | Placement , N Sponsored Search, J Sponsored Product |
| `campaignName` | String | No | campaign name，fuzzy search |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The Details |


### getDiscoveryReportKeyword
`GET/POST` `/sponsor/solutions/report/getDiscoveryReportKeyword`

**Description:** Get sponsored discovery report keyword level

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `adgroupName` | String | No | Adgroup Name |
| `adgroupId` | String | No | Adgroup Id |
| `keyword` | String | No | Keyword |
| `useRtTable` | Boolean | No | It means that if endDate have selected today, and you need realtime data,then set useRtTable=true If useRtTable=false,it will not search realtime data |
| `sort` | String | No | sort column,we have provide some index to sort |
| `order` | String | No | ASC or DESC, other String is invalid |
| `pageNo` | String | Yes | Page No，default 1,max=100 |
| `pageSize` | String | Yes | Page No, default 10, max=100 |
| `startDate` | String | Yes | start date, format like yyyy-MM-dd |
| `endDate` | String | Yes | end date , date, format like yyyy-MM-dd |
| `campaignName` | String | No | Campaign Name |
| `campaignId` | String | No | Campaign Id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Result Details |


### getLatestSignInfo
`GET` `/sponsor/solutions/account/getLatestSignInfo`

**Description:** Get the latest url of sign(T&C).

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The T&C url. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### getReportCampaignOnFIrstSlot
`GET/POST` `/sponsor/solutions/report/getReportCampaignOnPrePlacement`

**Description:** Get sponsored discovery report campaign first slot

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `sort` | String | No | sort column,we have provide some index to sort |
| `order` | String | No | ASC or DESC, other String is invalid |
| `pageNo` | Number | Yes | Page No，default 1,max=100 |
| `pageSize` | Number | Yes | Page Size, default 10, max=100 |
| `startDate` | String | Yes | start date, format like yyyy-MM-dd |
| `endDate` | String | Yes | end date , date, format like yyyy-MM-dd |
| `campaignName` | String | No | Campaign Name |
| `campaignId` | Number | No | campagnId |
| `productType` | String | No | Product Type, N:Sponsored Search(All)  F:Firsh Search Slot |
| `useRtTable` | Boolean | No | It means that if endDate have selected today, and you need realtime data,then set useRtTable=true If useRtTable=false,it will not search realtime data |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | Details |


### getReportOverview
`GET/POST` `/sponsor/solutions/report/getReportOverview`

**Description:** Get report overview.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `lastStartDate` | String | Yes | - |
| `endDate` | String | Yes | - |
| `useRtTable` | Boolean | Yes | - |
| `bizCode` | String | Yes | - |
| `lastEndDate` | String | Yes | - |
| `startDate` | String | Yes | - |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The detail data. |
| `success` | String | Yes | System result for this api call. |
| `analyseTraceId` | String | Yes | If the api call failed, you could find us with this. |
| `errorMsg` | String | Yes |  If the api call failed, this field will show the detail reason. |


### getReportOverviewMetric
`GET/POST` `/sponsor/solutions/report/getReportOverviewMetric`

**Description:** get report overview metric

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `metricType` | Number | Yes | The type pf metric.1:spend;2:impressions;3:clicks;4:ctr;5:units sold;6:revenue;7:cpc;8:roi;9:store order;10:store a2c;11:product order. |
| `endDate` | String | Yes | End date. |
| `useRtTable` | Boolean | Yes | If you need to search data for today, then use true, otherwise false. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `startDate` | String | Yes | Start date. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The detail result, for this api is metric data. |
| `success` | String | Yes | System result for this api call. |
| `analyseTraceId` | String | Yes |  If the api call failed, you could find us with this. |
| `errorMsg` | String | Yes | If the api call failed, this field will show the detail reason. |


### listCategory
`GET/POST` `/sponsor/solutions/category/listCategory`

**Description:** list category

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `parentId` | Number | No | The category parent id. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object[] | Yes | - |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### listKeywordByAdgroup
`GET/POST` `/sponsor/solutions/keyword/listKeywordByAdgroup`

**Description:** List keyword by adgroup.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignObjective` | Number | Yes | Your campaign objective helps determine your bidding strategy - Traffic objective helps you to increase the number of clicks to your store, while sales objective helps to increase your store’s sales.1:Traffic;2:Sales. |
| `campaignType` | Number | Yes | Unlock different ways to bids, select products, and keywords with campaign types. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `itemId` | Number | Yes | Product id. |
| `adgroupId` | Number | Yes | Adgroup id. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object[] | Yes | The detail result, for this api is keyword detail. |
| `success` | String | Yes | System result for this api call. |
| `analyseTraceId` | String | Yes | If the api call failed, you could find us with this. |
| `totalCount` | Number | Yes | Total count of keyword. |
| `errorMsg` | String | Yes | If the api call failed, this field will show the detail reason. |


### listKeywordByItem
`GET/POST` `/sponsor/solutions/keyword/listKeywordByItem`

**Description:** List keyword by item.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignObjective` | Number | Yes | Your campaign objective helps determine your bidding strategy - Traffic objective helps you to increase the number of clicks to your store, while sales objective helps to increase your store’s sales.1:Traffic;2:Sales. |
| `campaignType` | Number | Yes | Unlock different ways to bids, select products, and keywords with campaign types. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `itemId` | Number | Yes | Product id. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object[] | Yes |  The detail result, for this api is keyword detail. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### modifyAutoTopUpOptionOneConfig
`GET/POST` `/sponsor/solutions/wallet/modifyAutoTopUpOptionOneConfig`

**Description:** Modify auto top up option one config.1. each country has differect tax rate
2. we have minimum and maximam top-up amount limitation.For SG, min=5, max = 8,333,333,330;for PH, min=100,Max=17,895,600;for TH, min=100,max=8,333,333,300;for VN, min=50,000,max=833,333,300,000;for MY,min=10,max=8,333,333,330;for ID,min=25,000,max=8,333,333,000.the api timeout is 3s, max qps is 100, make sure do not over these num, especially qps, otherwise you may be blacklisted or limited request count for a while.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `status` | Number | Yes | The option one status.1:ON;0:OFF. |
| `limitAmount` | String | Yes | If balance is lower than this value, auto topUp operation will be done. |
| `topupAmount` | String | Yes | The amount of topUp for each auto topUp. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Boolean | Yes | The detail result, for this api is boolean. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### searchAdgroupList
`GET/POST` `/sponsor/solutions/adgroup/searchAdgroupList`

**Description:** Search adgroup with bizCode by seller.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `pageSize` | Number | Yes | Page size. |
| `endDate` | String | Yes | Campaign end date. |
| `campaignId` | Number | Yes | Campaign id. |
| `pageNo` | Number | Yes | Page number. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `adgroupName` | String | No | Adgroup name for fuzzy search. |
| `startDate` | String | Yes | Campaign start date. |
| `onlineStatus` | Number | No | The campaign online status.1:Online;0:Offline;9:deleted. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | The detail result, for this api is boolean. |
| `errorMsg` | String | Yes | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | Yes | If the api call failed, you could find us with this.If the api call failed, you could find us with this. |
| `totalCount` | Number | No | The count of adgorup. |
| `result` | Object[] | No | The detail result, for this api is adgroup detail list. |


### searchCampaignList
`GET/POST` `/sponsor/solutions/campaign/searchCampaignList`

**Description:** Search campaign list with bizCode for sellers.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `onlineStatus` | Number | No | The campaign online status.1:Online;0:Offline;9:deleted. |
| `startDate` | String | Yes | Campaign start date. |
| `endDate` | String | Yes | Campaign end date. |
| `pageNo` | String | Yes | Page number. |
| `pageSize` | String | Yes | Page size. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | Yes | System result for this api call. |
| `totalCount` | Number | Yes | Campaign total count. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |
| `result` | Object[] | Yes | The detail campaign list. |


### searchKeyword
`GET/POST` `/sponsor/solutions/keyword/searchKeyword`

**Description:** Search keyword with specific word.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignObjective` | Number | Yes | Your campaign objective helps determine your bidding strategy - Traffic objective helps you to increase the number of clicks to your store, while sales objective helps to increase your store’s sales.1:Traffic;2:Sales. |
| `campaignType` | Number | Yes | Unlock different ways to bids, select products, and keywords with campaign types. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `itemQuery` | String | Yes | The word you do not want to put in the result. |
| `itemId` | Number | Yes | Product id. |
| `searchWord` | String | Yes | The specific word. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object[] | Yes | The keyword detail. |
| `success` | Boolean | Yes | System result for this api call. |
| `analyseTraceId` | String | Yes | If the api call failed, you could find us with this. |
| `totalCount` | Number | Yes | Total count of keyword. |
| `errorMsg` | String | Yes | If the api call failed, this field will show the detail reason. |


### searchProductWithPage
`GET/POST` `/sponsor/solutions/product/searchProductWithPage`

**Description:** Search product.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `brandName` | String | No | Prodct brand name. |
| `campaignType` | Number | Yes | Unlock different ways to bids, select products, and keywords with campaign types. |
| `pageSize` | Number | Yes | Page size. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `placementList` | Number[] | Yes | Placements determine where shoppers will see your promoted products.3:Search Result Page;4:Just For You Page |
| `productName` | String | No | Product name to fuzzy search. |
| `campaignObjectLive` | Number | Yes | Your campaign objective helps determine your bidding strategy - Traffic objective helps you to increase the number of clicks to your store, while sales objective helps to increase your store’s sales.1:Traffic;2:Sales. |
| `eligible` | Number | Yes | Only search product which is eligible/ineligible.1:eligible;0:ineligible. |
| `pageNo` | Number | Yes | Page number. |
| `sellerSku` | String | No | Product sellerSku. |
| `maxCpc` | String | Yes | Max bid determines the highest amount that you're willing to pay for a click on your promoted product.-1 means no limit. |
| `categoryId` | Number | No | Input category id to exact search. |
| `itemIdBlackList` | Number[] | No | Input item id which you do not want put into result. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object[] | Yes |  The detail result, for this api is product detail info. |
| `success` | Boolean | Yes | System result for this api call. |
| `analyseTraceId` | String | Yes | If the api call failed, you could find us with this. |
| `totalCount` | Number | Yes | Total count of product. |
| `errorMsg` | String | Yes | If the api call failed, this field will show the detail reason. |


### sign
`GET/POST` `/sponsor/solutions/account/sign`

**Description:** Description: Do sign for seller. Seller or agencies can use this api to sign up the t&c.
Timeout Period： the api timeout is 10s, max qps is 300, make sure do not over these num, especially qps, otherwise you may be blacklisted or limited request count for a while.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The detail result, for this api is boolean. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### updateAdgroupBatch
`GET/POST` `/sponsor/solutions/adgroup/updateAdgroupBatch`

**Description:** Update adgroup batch.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `adgroupViewDTOList` | Object[] | Yes | Adgroup list |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Boolean | Yes | The detail result, for this api is boolean. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


### updateCampaign
`GET/POST` `/sponsor/solutions/campaign/updateCampaign`

**Description:** Update campaign with status field.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `campaignId` | Number | Yes | Campaign id. |
| `campaignName` | String | No | Campaign name. |
| `startDate` | String | No | Campaign start date. |
| `endDate` | String | No | Campaign end date. |
| `dayBudget` | String | No | Budget indicates the maximum amount you’re willing to pay each day. |
| `bizCode` | String | Yes | Decided to choose which advertisement solution.SD:sponsoredSearch. |
| `switchStatus` | Number | No | Campaign swtich status.1:Online;0:Offline. |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | The campaign's detail which you just updated. |
| `success` | Boolean | Yes | System result for this api call. |
| `errorMsg` | String | No | If the api call failed, this field will show the detail reason. |
| `analyseTraceId` | String | No | If the api call failed, you could find us with this. |


---
## Service Market API

_Lazada Service Market API_

### ServiceMarketAppKeyOrderQuery
`GET/POST` `/service/market/order/query`

**Description:** Query user order list for specific App on Service Market

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `endCreated` | String | No | order create time range end |
| `bizType` | Number | No | biz type |
| `bizOrderId` | Number | No | bi order id |
| `orderId` | Number | No | order_id |
| `pageNo` | Number | Yes | page no |
| `itemCode` | String | No | service market item code |
| `pageSize` | Number | Yes | page size |
| `startCreated` | String | No | order create time range start |
| `articleCode` | String | Yes | service market article code |
| `shortCode` | String | No | seller short code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | null |


### ServiceMarketAppKeySubQuery
`GET/POST` `/service/market/subs/query`

**Description:** Query user subscription info for specific App on Service Market

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `articleCode` | String | Yes | Service Market article code |
| `shortCode` | String | Yes | seller short code |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


---
## Choice Customized API

### BatchDeliverJitPurchaseOrder
`GET/POST` `/jit/purchase_order/batch_pickup_deliver`

**Description:** Batch Pickup Deliver Jit Purchase Order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `purchaseOrderNoList` | String[] | Yes | 采购单号列表，最大100个。{["POJ1001","POJ1002"]} |
| `shipperAreaCode` | String | Yes | 揽收联系人地址区域，如：CN： 当前支持CN，VN，TH，PH，ID，MY一共6个地区。必填。 |
| `shipperAddressId` | Number | Yes | 揽收联系人地址id。必填。 |
| `shipperAddressDetail` | String | Yes | 揽收详细地址。必填。 |
| `shipperMobilePhone` | String | Yes | 揽收联系人电话。必填。 |
| `shipperName` | String | Yes | 揽收联系人姓名。必填。 |
| `estimatedPickupDate` | String | No | 预约揽收日期 {yyyy-MM-dd}。非必填 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `INVALID_STATUS_FORBIDDEN_PICK_UP` | INVALID_STATUS_FORBIDDEN_PICK_UP | This API can only be called if the order is in “Ready To Ship (biz_status = 20)” status, please call QueryListJitPurchas |


### EditChoiceSkuStock
`POST` `/choice/stock/edit`

**Description:** batch update choice jit product stock by skuId

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `item_id` | Number | Yes | item id |
| `site` | String | Yes | The country site of the queried Product |
| `sku_edit_stock` | String | Yes | Key：sku_id Value: sellable stock |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | update result json |
| `success` | Boolean | No | success flag |
| `error_code` | String | No | error code |
| `error_msg` | String | No | error msg |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `E0208` | Product not exist | The item id in the request does not exist in the current store or the CHOICE item has not yet been reviewed by Lazada, u |
| `E1002` | not jit product | Non-JIT items do not support inventory modification, please call GetChoiceProducts or GetChoiceProductItem API to query  |
| `E1001` | not jit seller | Seller authorization is not a choice authorization, please ask the seller to re-authorize and select the 'country - choi |
| `E0208` | Product not exist | The item id in the request does not exist in the current store or the CHOICE item has not yet been reviewed by Lazada, u |


### GetChoiceProductItem
`GET` `/choice/product/item/get`

**Description:** Get single product by ItemId or SellerSku.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `item_id` | Number | No | Call this API; Either "Item Id" or "Seller Sku" must be selected as the request parameter |
| `seller_sku` | String | No | Call this API; Either "Item Id" or "Seller Sku" must be selected as the request parameter |
| `site` | String | Yes | The country site of the queried Product |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response body |


### GetChoiceProducts
`GET/POST` `/choice/products/get`

**Description:** Use this API to get detailed information of the specified products.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `filter` | String | No | Returns the products with the status matching this parameter. Possible values are all, live, inactive, deleted, pending, rejected, sold-out. Mandatory. |
| `update_before` | String | No | Limits the returned product list to those updated before or on a specified date, given in ISO 8601 date format. Optional |
| `create_before` | String | No | Limits the returned products to those created before or on the specified date, given in ISO 8601 date format. Optional |
| `offset` | String | No | Deprecated(The number of Items you want to skip before you start counting),It is recommended to use date for scrolling query.The maximum offset is 10000 |
| `create_after` | String | No | Limits the returned products to those created after or on the specified date, given in ISO 8601 date format. Optional |
| `update_after` | String | No | Limits the returned products to those updated after or on the specified date, given in ISO 8601 date format. Optional |
| `limit` | String | No | The number of Items you would like to fetch from every response,The maximum is 50. |
| `options` | String | No | This value can be used to get more stock information. e.g., Options=1 means contain ReservedStock, RtsStock, PendingStock, RealTimeStock, FulfillmentBySellable. |
| `sku_seller_list` | String | No | Only products that have the Seller SKU in this list will be returned. Input should be a JSON array. For example, ["Apple 6S Gold", "Apple 6S Black"]. It only matches the whole words. A maximum of 100 SKUs can be returned. |
| `site` | String | Yes | The country site of the queried Product |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response body |


### GetChoiceSeller
`GET/POST` `/choice/seller/get`

**Description:** Get choice seller information by seller ID and site

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `site` | String | Yes | The country site of the queried merchant |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | Response data |


### GetChoiceSkuItemRelationBySku
`GET/POST` `/choice/sku_item_relation/get_by_sku`

**Description:** get the relation between platformSku and item by sku

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `item_id` | String | Yes | itemId |
| `sku_id` | String | Yes | skuId |
| `site` | String | Yes | The country site of the queried Product item |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | Yes | Response |


### PackageJitPurchaseOrder
`POST` `/jit/purchase_order/package`

**Description:** Package Jit Purchase Order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `purchase_order_no_list` | String[] | Yes | 采购单列表，最大100个。{["POJ1001","POJ1002"]} |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### PrintJitPurchaseOrderAndItem
`GET/POST` `/jit/purchase_order/print`

**Description:** Print Jit Purchase Order And Item.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `purchase_order_no_list` | String[] | Yes | 采购单号列表，最大20个。{["POJ1001","POJ1002"]} |
| `print_order` | Boolean | Yes | 是否打印PO单。{true/false} |
| `print_barcode` | String | Yes | 是否打印货品barcode。{true/false} |
| `pdf_size` | String | Yes | pdf样式。{A4/6030/100150} |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### PrintPickuoOrder
`GET/POST` `/pickup_order/print`

**Description:** Print Pickuo Order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `pickup_order_no` | String | Yes | 揽收单号 |
| `pdf_size` | String | Yes | pdf格式枚举类型。A4纸大小样式、100*100大小样式。{PICKUP_A4/PICKUP_1010} |
| `box_number` | String | Yes | 装箱数量。（最大值 100） |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### QueryListJitPurchaseOrder
`GET/POST` `/jit/purchase_order/query_list`

**Description:** Query List Jit Purchase Order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `gmt_create_begin` | String | No | 单据创建开始时间，建单时间范围(即end-begin)需要在90天内。{yyyy-MM-dd HH:mm:ss} |
| `gmt_create_end` | String | No | 单据创建结束时间，建单时间范围(即end-begin)需要在90天内。{yyyy-MM-dd HH:mm:ss} |
| `purchase_order_no_list` | String[] | No | 采购单列表，最大20个。{["POJ1001","POJ1002"]} |
| `logistics_no_list` | String[] | No | 物流单列表，最大10个。{["LBX1001","LBX1002"]} |
| `order_status` | String | No | 单据状态 10:待打包; 20:待发货; 22:待收货; 25:已到仓; 40:已完成; -100610:超时关闭; -100:买家取消；不传则返回所有状态的采购单； |
| `page_index` | Number | No | 当前页，默认1。 |
| `page_size` | Number | No | 分页大小，最大50个，默认20。 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### QueryListPurchaseItem
`GET/POST` `/jit/purchase_order/query_list_purchase_item`

**Description:** Query List Purchase Item.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `purchase_order_no` | String | Yes | JIT采购单号 |
| `page_index` | Number | No | 当前页，默认1。 |
| `page_size` | Number | No | 分页大小，最大200个，默认20。 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### QueryPickupOrder
`GET/POST` `/pickup_order/query`

**Description:** Query Pickup Order.

**Auth:** Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | Yes | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `pickup_order_no` | String | Yes | 揽收单号 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


---
## LazLike API

_LazLike Content API_

### MCNQueryTagInfoByName
`GET/POST` `/content/mcn/content/queryTagInfosByName`

**Description:** MCNQueryTagInfoByName

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `tagNames` | String | Yes | The tag name you want to query, multiple tags are split according to, |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `api_result` | Object | No | result |


### McnContentCancelSchedulePublish
`POST` `/content/mcn/content/cancelScheduled`

**Description:** McnContentCancelSchedulePublish

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `contentId` | Number | Yes | Content ID that needs to be canceled scheduled release |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `api_result` | Object | No | result of api |


### McnContentCompleteCreateVideo
`POST` `/content/mcn/video/block/commit`

**Description:** After uploading all blocks of the video file, call McnContentCompleteCreateVideo to complete the video uploading process.


**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `uploadId` | String | Yes | come from the result of McnContentInitCreateVideo |
| `parts` | String | Yes | a json string contains e_tag info of each block |
| `title` | String | Yes | the video title |
| `coverUrl` | String | No | optional. cover Image of video，return by calling McnContentUploadImage |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result of api |


### McnContentCreate
`POST` `/content/mcn/content/create`

**Description:** create content

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `kolUserId` | Number | No | buyer account of kol |
| `contentType` | String | Yes | should be 'video' for video content |
| `description` | String | Yes | text part |
| `imageList` | String | No | image urls splitted by comma  |
| `itemList` | String | No | itemId list splitted by comma  |
| `videoId` | Number | No | return by calling McnContentCompleteCreateVideo |
| `categoryId` | Number | No | category id |
| `tags` | String | No | contents brief tags |
| `voiceLang` | String | Yes | language of voice |
| `subtitleLang` | String | Yes | language of subtitle |
| `descriptionLang` | String | No | language of description |
| `publishTimeMillis` | Number | No | Content release time, if it is to be released immediately, you can not pass it or pass 0. If you want to publish it regularly, pass > a timestamp of the current time, milliseconds and must be an hour. |
| `shopId` | Number | No | shop account |
| `proxyFlag` | Boolean | No | proxy flag |
| `title` | String | No | title |
| `extraTagIds` | String | No | Additional tags that need to be added, such as fashion tags and sale tags |
| `channel` | String | No | mcn_aigc or mcn_content |
| `bizType` | String | No | LazMall or LazLive |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result of api |


### McnContentInitCreateVideo
`GET/POST` `/content/mcn/video/block/create`

**Description:** Initial an upload video process, this API will return the corresponding UploadID

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `kolUserId` | Number | Yes | buyer account of kol |
| `fileName` | String | Yes | local filename, should be less than 20 chars |
| `fileBytes` | Number | Yes | video file's bytes, should be less than 100M |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result of api |


### McnContentListCategory
`GET` `/content/mcn/category/list`

**Description:** list mcn content categories

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result of api |


### McnContentPropertyTagList
`GET/POST` `/content/mcn/property/list`

**Description:** list mcn content property tags

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | whether the operation succeeds |
| `resultMessage` | String | No | error code provided when the operation fails |
| `resultCode` | String | No | error message provided when the operation fails |
| `tagList` | Object[] | No | result |


### McnContentReplySchedulePublish
`POST` `/content/mcn/content/replySchedulePublish`

**Description:** McnContentReplySchedulePublish

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `contentId` | Number | Yes | contentId |
| `publishTimeMillis` | Number | Yes | Resume scheduled publishing time |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `api_result` | Object | No | result of api |


### McnContentUploadImage
`POST` `/content/mcn/image/upload`

**Description:** upload image

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `kolUserId` | Number | Yes | kol user id |
| `image` | byte[] | Yes | file content |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result of api |


### McnContentUploadVideoBlock
`POST` `/content/mcn/video/block/upload`

**Description:** upload one block of video file

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `uploadId` | String | Yes | upload id |
| `blockNo` | Number | Yes | block number |
| `blockCount` | Number | Yes | block count |
| `file` | byte[] | Yes | block content |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result of api |


### McnProductValidator
`GET` `/content/mcn/product/validate`

**Description:** Identify high risk products

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `lazOpAppKey` | String | No | appKey |
| `itemIdList` | String | Yes | 商品id，多个用英文逗号隔开 |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result of api |


### McnSimilarProductSearch
`GET/POST` `/content/mcn/similar/product/search`

**Description:** 相似商品搜索接口

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `kolUserId` | Number | No | user id |
| `imageUrlList` | String | No | image url list |
| `shopId` | Number | No | shop id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `productList` | Object[] | No | product info |
| `confidentialityStatement` | String | No | confidentiality statement |
| `success` | Boolean | No | whether the operation succeeds |
| `result_code` | String | No | error code provided when the operation fails |
| `result_message` | String | No | error message provided when the operation fails |


### queryContentReviewRecords
`GET/POST` `/content/mcn/content/queryReviewRecords`

**Description:** Query content audit records. Currently, querying records with audit results of low (block) is supported.The number of query contents is limited to 500 (adjustable).

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `contentIds` | String | Yes | 内容id |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | No | result |


---
## LazLive API

_The API to highlight product in live room_

### HighlightProduct
`GET/POST` `/lazlive/product/highlight`

**Description:** highlight product

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `highLightRequest` | Object | Yes | Request parameters |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `data` | Object | No | data |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `BIZ_INVALID_ARGUMENT` | Please check whether the input parameter "action" is correct | 1 |
| `BIZ_USER_NOT_PERMITTED` | No permission | 1 |
| `BIZ_LIVE_NOT_FOUND` | The live room does not exist | 1 |
| `BIZ_INVALID_PRODUCT` | Invalid product | 1 |
| `BIZ_NOT_LIVE_PRODUCT` | It is not a product of the live room | 1 |
| `SYSTEM_ERROR` | We are experiencing a surge in traffic. Please try again. If you continue to get this message, try again later | 1 |


---
## Logistics Station API

_API to integrate with Lazada logistics station including Drop-off point and Collection point_

### CageValidation
`POST` `/logistics/station/cages/validate`

**Description:** Validate if a cage is valid

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cageNumber` | String | Yes | Cage number |
| `stationCode` | String | Yes | Station code/ID |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Boolean | No | Validate cage result success or not? |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `CAGE_NOT_FOUND` | Cage not found: {cageNumber} | Cage not found |
| `STATION_NOT_ACTIVE` | Station [{stationCode}] is not active | Station is not active |
| `UNEXPECTED_ERROR` | NullpointerException | Mostly the stacktrace of unexpected error |


### ConfirmInbound
`POST` `/logistics/station/v1/confirm-inbound`

**Description:** Confirm inbound. Call this API to inbound the scanned parcel and finish the inbound process

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `cageNumber` | String | No | Cage number. If cage number is present, it will be validated. In case missing cage number, the system will choose default cage number |
| `trackingNumbers` | String[] | Yes | List of tracking number |
| `serviceType` | String | Yes | Accept values: SELLER_DROPOFF, CUSTOMER_DROPOFF (Customer return), CUSTOMER_COLLECTION (Collection point) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Boolean | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### ConfirmParcelCollection
`POST` `/logistics/station/v1/cp/confirm-parcel-collection`

**Description:** Confirm customer collects or rejects parcel. This API is used after ValidateOTP success.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `trackingNumber` | String | Yes | Tracking number of parcel |
| `otp` | String | Yes | The parcel OTP is used for collecting parcel |
| `action` | String | Yes | Accept values: COLLECT, REJECT |
| `rejectCode` | String | No | Reject reason code, required in case REJECT action |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Boolean | No | Validate OTP result is success or not? |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### CreateScannedParcel
`POST` `/logistics/station/v1/scanned-parcels/create`

**Description:** Create a scanned parcel. Call this API when scanning the tracking number on the parcel.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `cageNumber` | String | No | Cage number. If cage number is present, it will be validated. In case missing cage number, the system will choose default cage number |
| `trackingNumber` | String | Yes | Tracking number of parcel |
| `serviceType` | String | Yes | Accept values: SELLER_DROPOFF, CUSTOMER_DROPOFF (Customer return), CUSTOMER_COLLECTION (Collection point) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### DeleteScannedParcel
`POST` `/logistics/station/v1/scanned-parcels/delete`

**Description:** Delete scanned parcels by tracking number. This API is required when user deletes the scanned parcels

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `trackingNumbers` | String[] | Yes | List of tracking numbers |
| `serviceType` | String | Yes | Accept values: SELLER_DROPOFF, CUSTOMER_DROPOFF (Customer return), CUSTOMER_COLLECTION (Collection point) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Boolean | No | Delete success or not? |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### DopConfirmInbound
`POST` `/logistics/station/dop/confirm-inbound`

**Description:** DOP confirm inbound

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationCode` | String | Yes | Station code/ID |
| `scannedParcels` | Object[] | Yes | List scanned parcels |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | String | No | Confirm inbound success or not? |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `DOP_RESERVED_PARCEL_NOT_FOUND` | No parcel info for tracking number {trackingNumber}. Please scan again or manually input the tracking number. | Cannot find parcel info with provided tracking number |
| `CAGE_NOT_FOUND` | Cage not found: {cageNumber} | Cage not found |
| `STATION_NOT_ACTIVE` | Station [{stationCode}] is not active | Station is not active |
| `PARCEL_ALREADY_INBOUND` | Parcel has already inbounded: {trackingNumber} | Parcel has already inbounded |
| `STATION_IS_NOT_DOP` | Station {stationCode} is not a DOP. You can not drop-off here. | Station is not DOP type |
| `DOP_PARCEL_STATUS_NOT_WHITELIST` | Parcel is not at correct status to dropoff, parcel {trackingNumber} is now {status} | Invalid status to inbound |
| `CANNOT_INBOUND_CANCELLED_TASK` | Tracking number {trackingNumber} is cancelled. Please remove out of list | Parcel is cancelled |
| `DOP_MERCHANT_MDOP` | Seller is a MDOP, your parcel cannot be dropped-off to any station. DOP Merchant={sellerName}, and TN={trackingNumber} | The seller of the parcel is MDOP |
| `DUPLICATE_REQUEST` | Your request is processing | Client submit duplicate request at the same time |
| `UNEXPECTED_ERROR` | NullpointerException | Mostly the stacktrace of unexpected error |


### DopCreateScannedParcel
`POST` `/logistics/station/dop/scanned-parcels`

**Description:** DOP create scanned parcel

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationCode` | String | Yes | Station code/ID |
| `cageNumber` | String | Yes | Cage number |
| `trackingNumber` | String | Yes | Tracking number of parcel |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `DOP_RESERVED_PARCEL_NOT_FOUND` | No parcel info for tracking number {trackingNumber}. Please scan again or manually input the tracking number. | Cannot find parcel info with provided tracking number |
| `CAGE_NOT_FOUND` | Cage not found: {cageNumber} | Cage not found |
| `STATION_NOT_ACTIVE` | Station [{stationCode}] is not active | Station is not active |
| `PARCEL_ALREADY_INBOUND` | Parcel has already inbounded: {trackingNumber} | Parcel has already inbounded |
| `STATION_IS_NOT_DOP` | Station {stationCode} is not a DOP. You can not drop-off here. | Station is not DOP type |
| `DOP_PARCEL_STATUS_NOT_WHITELIST` | Parcel is not at correct status to dropoff, parcel {trackingNumber} is now {status} | Invalid status to inbound |
| `CANNOT_INBOUND_CANCELLED_TASK` | Tracking number {trackingNumber} is cancelled. Please remove out of list | Parcel is cancelled |
| `DOP_MERCHANT_MDOP` | Seller is a MDOP, your parcel cannot be dropped-off to any station. DOP Merchant={sellerName}, and TN={trackingNumber} | The seller of the parcel is MDOP |
| `DUPLICATE_REQUEST` | Your request is processing | Client submit duplicate request at the same time |
| `UNEXPECTED_ERROR` | NullpointerException | Mostly the stacktrace of unexpected error |


### DopDeleteScannedParcel
`POST` `/logistics/station/dop/scanned-parcels/delete`

**Description:** DOP delete scanned parcel

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationCode` | String | Yes | Station code/ID |
| `trackingNumbers` | String[] | Yes | List scanned tracking number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | String | No | Delete success or not? |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `PARCEL_NOT_FOUND` | Parcel not found: {trackingNumber1,trackingNumber2} | Parcel not found |
| `UNEXPECTED_ERROR` | NullpointerException | Mostly the stacktrace of unexpected error |


### DopGetInboundedParcel
`GET` `/logistics/station/dop/inbounded-parcels/list`

**Description:** DOP get list scanned parcel

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationCode` | String | Yes | Station code/ID |
| `trackingNumbers` | String[] | Yes | List inbounded tracking number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object[] | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `UNEXPECTED_ERROR` | NullpointerException | Mostly the stacktrace of unexpected error |


### DopGetScannedParcel
`GET` `/logistics/station/dop/scanned-parcels/list`

**Description:** DOP get list scanned parcel

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationCode` | String | Yes | Station code/ID |
| `cageNumber` | String | No | Cage number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object[] | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |

**Error Codes:**

| Code | Message | Solution |
|------|---------|---------|
| `UNEXPECTED_ERROR` | NullpointerException | Mostly the stacktrace of unexpected error |


### GetCpScheduledPuParcel
`GET` `/logistics/station/v1/cp/scheduled-pu-parcels/list`

**Description:** Get a list of parcels that are scheduled to be picked up for return to seller. These parcels are expired (no collection from customer), SLA breached or customer rejected. This API is used to help the agent prepare parcels before seller comes.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object[] | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### GetInboundedParcel
`GET` `/logistics/station/v1/inbounded-parcels/list`

**Description:** Get a list of inbounded parcels by a list of tracking numbers. This API is used for checking the status of inbounded parcels such as parcels picked up by LEX, picked up by 3PL, or collected by a customer.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `trackingNumbers` | String[] | Yes | List of tracking number |
| `serviceType` | String | Yes | Accept values: SELLER_DROPOFF, CUSTOMER_DROPOFF (Customer return), CUSTOMER_COLLECTION (Collection point) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object[] | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### GetListAccessStation
`GET` `/logistics/station/list`

**Description:** Get list access station by APP

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object[] | No | List access station |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### GetMetaData
`GET` `/logistics/station/v1/metadata`

**Description:** Get metadata such as reject reasons, etc

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### GetScannedParcel
`GET` `/logistics/station/v1/scanned-parcels/list`

**Description:** Get a list of scanned parcels. This API is often used for synchronization purposes such as: user refreshes the page, partner system can call this API to get the list of scanned parcels again. This API is not required to call during operations.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `cageNumber` | String | No | Cage number |
| `serviceType` | String | Yes | Accept values: SELLER_DROPOFF, CUSTOMER_DROPOFF (Customer return), CUSTOMER_COLLECTION (Collection point) |
| `stationId` | String | Yes | Station ID in partner system |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object[] | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### SearchCustomerReturnParcel
`GET` `/logistics/station/v1/dop/cr-parcels/search`

**Description:** Search customer return parcel by at least 4 letters text. This API is to improve user experience, user can search for the tracking number instead of typing the full tracking number.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `searchText` | String | Yes | Search tracking number text at least 4 letters |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Object[] | No | Response data |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### ValidateCage
`POST` `/logistics/station/v1/cages/validate`

**Description:** Validate if a cage is valid. This API is often called before starting inbound but it's not required.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `cageNumber` | String | Yes | Cage number, Lazada will provides cage number |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Boolean | No | Validate cage result success or not? |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


### ValidateOTP
`POST` `/logistics/station/v1/cp/validate-otp`

**Description:** Validate if OTP of parcel is valid or not. This API is used for checking OTP before confirming collection.

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `stationId` | String | Yes | Station ID in partner system |
| `trackingNumber` | String | Yes | Tracking number of parcel |
| `otp` | String | Yes | The parcel OTP is used for collecting parcel |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `success` | Boolean | No | Is success? |
| `data` | Boolean | No | Validate OTP result is success or not? |
| `errorCode` | String | No | Error code |
| `errorMsg` | String | No | Error message |
| `traceId` | String | No | Trace id for debugging |


---
## LazCredit Risk API

_lazada credit risk_

_No APIs listed._

---
## Content API

_The Content APIs are content-related APIs that are only available to specific users. Make sure that communication has been completed prior to applying permission to visit them._

### cancelTask
`POST` `/content/ai/cancelTask`

**Description:** cancel tasks

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `task_ids` | String[] | Yes | task_ids |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### changeFace
`POST` `/content/ai/changeFace`

**Description:** change face using lazada AI algorithm

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `raw_image_url` | String | Yes | raw_image_url |
| `model_code` | String | Yes | model_code |
| `batch_size` | Number | No | batch_size |
| `ratio` | String | No | ratio |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### changeProductBackground
`POST` `/content/ai/changeProductBackground`

**Description:** change product background using lazada AI algorithm

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `product_image_url` | String | Yes | image url  |
| `background_code` | String | Yes | background code |
| `batch_size` | Number | Yes | batch size |
| `ratio` | String | No | ratio |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### fixHand
`POST` `/content/ai/fixHand`

**Description:** fixHand using lazada AI algorithm

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `raw_image_url` | String | Yes | raw_image_url |
| `batch_size` | Number | No | batch size |
| `base_ref` | Boolean | No | base_ref |
| `model_reference_image_url` | String | Yes | model_reference_image_url |
| `ratio` | String | No | ratio |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### getTaskStatus
`GET` `/content/ai/getTaskStatus`

**Description:** get task status

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `task_id` | String | Yes | taskId |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### productImageMatch
`GET/POST` `/content/ai/productImageMatch`

**Description:** match product image

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `match_num` | Number | Yes | match num |
| `image_url` | String | Yes | image url |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


### tryOnCloth
`POST` `/content/ai/tryOnCloth`

**Description:** try on cloth using lazada AI algorithm

**Auth:** No Authorization Required

**Common Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `app_key` | String | Yes | Unique app ID issued by LAZADA Open Platform console when you apply for an app category |
| `timestamp` | String | Yes | The time stamp of the request e.g. 1517820392000 (which translates to 5 February 2018 08:46:32) with less than 7200s difference from UTC time |
| `access_token` | String | No | API interface call credentials |
| `sign_method` | String | Yes | The HMAC hash algorithm you are using to calculate your signature |
| `sign` | String | Yes | Part of the authentication process that is used for identifying and verifying who is sending a request (click <a target='_blank' href='https://open.lazada.com/apps/doc/doc?nodeId=10450&docId=108068'>here</a> for details) |

**Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `keep_model` | Boolean | No | keep_model |
| `additional_cloth_image_url` | String | No | additional_cloth_image_url |
| `cloth_image_url` | String | Yes | cloth_image_url |
| `type` | String | Yes | type |
| `batch_size` | Number | No | batch_size |
| `model_reference_image_url` | String | No | model_reference_image_url |
| `ratio` | String | No | ratio |

**Response Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `result` | Object | Yes | result |


---
## Store Flash Sale API

_店铺闪购API_

_No APIs listed._
