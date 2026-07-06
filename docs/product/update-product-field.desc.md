Here's the "Updateproduct API Common Fields Description" table in markdown:

| Field Name | Require | Removable | Field Description |
|---|---|---|---|
| PrimaryCategory | Optional | No | Fill in this field when calling the API to update the category of the product |
| AssociatedSku | Optional | N/A | This field is not a product attribute, but a Key attribute. When you want to add a new SKU to an existing product using the Updateproduct API, you need to fill in this field with a sku id that already exists for the target product. |
| Images | Optional | No | The Images attribute is used to update images for this product. The Image field in the Images property must be of the array type, with a maximum of 8 images passed into each Image field. Putting Images in the sku attribute indicates that the image is a variant image, and putting it outside indicates an SPU image. Images field cannot delete the image, only replace the old one with the new one. Must use lazada URL images, external URLs are not supported. |
| name | Optional | No | Product Title |
| description | Optional | Yes | Maximum 25000 characters. HTML tags are allowed. Only Lazada image URLs are allowed, no external URLs. |
| d_description | Optional | N/A | Delete long description. |
| d_description_ms | Optional | N/A | Delete long description of existing Malay version. |
| d_description_en | Optional | N/A | Delete the long description of the English version that already exists. |
| short_description | Optional | Yes | Only text. Images and URLs will be automatically filtered; Short descriptions can only use `<ul><li>` or `<ol><li>` |
| d_short_description_en | Optional | N/A | Delete the short description of the English version that already exists. |
| gift_wrapping | Optional | No | Note: Whitelist function, if you need this function please contact PSC for more information. Enum: Yes/No |
| name_engravement | Optional | No | Note: Whitelist function, if you need this function please contact PSC for more information. Enum: Yes/No |
| brand | Optional | No | The name obtained by calling the GetBrandByPages API |
| brand_id | Optional | No | The brand id obtained by calling the GetBrandByPages API |
| video | Optional | Yes | After uploading the video, fill in the video id of the GetVideo or CompleteCreateVideo response, and the video status must be AUDIT_SUCCESS. Add `"remove_video": "Yes"` to the payload to remove the product video |
| SkuId | Mandatory | No | Sku id created by Lazada at the time of product creation and used to locate the Sku or product that needs to be updated when the product is updated. |
| SellerSku | Optional | No | Customizable by the seller, unique in the store. Does not support using Updateproduct API to modify seller sku (before 15 November 2023). |
| price | Optional | No | The retail price of the product, which will be displayed if the "special price" field is not filled or expired |
| special_price | Optional | No | The actual sales price of the product, if the "special_from_date" and "special_to_date" fields are not filled in, then the validity time is Long Time (Permanent validity). |
| special_from_date | Optional | No | Special price start time. If this parameter is passed, then "special_price" is mandatory. |
| special_to_date | Optional | No | Special price end time. If this parameter is passed, then "special_price" is mandatory. |
| Status | Optional | No | One of the following values: 'active', 'inactive' or 'deleted'. Optional. The default value is 'active' |
| package_height | Mandatory | No | Maximum two decimal places. Unit: cm |
| package_length | Mandatory | No | Maximum two decimal places. Unit: cm |
| package_width | Mandatory | No | Maximum two decimal places. Unit: cm |
| package_weight | Mandatory | No | Maximum two decimal places. Unit: kg |
| saleProp | Optional | No | Updateproduct can update the option of a variant or add a new variant. If you need to delete the variant, you need to call the Removesku api |
| color_thumbnail | Optional | No | Variant thumbnail tags, only available when the SKU variant is a standard sales attribute (except size). |

Notes
The list is only part of the general attributes, each category of goods may have different product attributes can be used, the product available attributes to obtain please refer to this document.
Please make sure the SKU variant attributes are the same when updating products.
SKU images do not support partial additions. If a SKU adds a SKU image, then other SKUs of this item also need to add images, and vice versa.



Notes

●Except for Sku Id, all other product attributes are not mandatory.

●Attributes must be preserved but can be left out of any field.

●Updateproduct API can only add variants and modify options, not delete variant attributes.

●The Updateproduct api needs to be requested using the POST method, so please put the parameters in the http body, otherwise the request may fail because the URL is too long.