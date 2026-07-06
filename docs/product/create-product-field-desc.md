# Create Product - Partial Field Descriptions

| Field Name | Require | Field Description |
|---|---|---|
| Images | Optional | Uploads product images; `Image` is an array holding up to 8 items. Placing `Images` inside a SKU sets a variant-level image; placing it outside the SKU applies at the product (SPU) level. |
| name | Mandatory | Product name, max 255 characters. |
| AssociatedSku | Optional | Links new SKUs to an existing product; the referenced SKU must already be visible in Seller Center. |
| description | Optional | Max 25,000 characters. HTML allowed, but only Lazada-hosted image URLs (no external URLs). |
| short_description | Optional | Plain text only — embedded images/URLs are stripped automatically; only `<ul><li>` or `<ol><li>` formatting allowed. |
| brand | Mandatory | Deprecated — use `brand_id` instead. |
| brand_id | Mandatory | Use the `brand_id` value returned by the GetBrandByPages API. |
| video | Optional | Supply the video ID from GetVideo/CompleteCreateVideo after upload; status must be AUDIT_SUCCESS. |
| gift_wrapping | Optional | Whitelisted feature (contact PSC to enable). Enum: Yes/No. |
| name_engravement | Optional | Whitelisted feature (contact PSC to enable). Enum: Yes/No. |
| preorder_enable | Optional | Turns pre-order mode on/off. Enum: Yes/No. |
| preorder_days | Optional | Estimated processing days when pre-order is on; allowed min/max range is seller-specific, set by Lazada operations (contact PSC to change). |
| disableAutoFillAttribute | Optional | Set `true` to disable the default auto-fill mechanism that otherwise populates unused attributes automatically. Enum: true/false (default false). |
| SellerSku | Mandatory | Seller-defined identifier, unique within the same item. |
| price | Mandatory | Standard retail price, shown when no active special price applies. |
| special_price | Optional | Discounted sale price; stays active indefinitely if no from/to dates are set. |
| special_from_date | Optional | Start of the special-price window; required if `special_price` is set. |
| special_to_date | Optional | End of the special-price window; required if `special_price` is set. |
| package_height | Mandatory | Up to 2 decimals, unit: cm. |
| package_length | Mandatory | Up to 2 decimals, unit: cm. |
| package_width | Mandatory | Up to 2 decimals, unit: cm. |
| package_weight | Mandatory | Up to 2 decimals, unit: kg. |
| package_content | Optional | Description of package contents. |
| saleProp | Optional | Groups sales-related SKU attributes (as opposed to non-sales attributes). |
| color_family | Optional | The category's standard variant attribute (name may vary, e.g. "size"). Optional for single-SKU listings, required for multi-SKU listings; a custom variant can be created if the category has none. |
| color_thumbnail | Optional | Variant thumbnail tag, usable only when the SKU variant is a standard sales attribute. |


What else needs to be done before creating a product？
Get the category tree

Use this API to obtain the catalog library information of each country.

API Usage Documentation

1、GetCategorySuggestion
Get product's category suggestion by product title.

We strongly recommend you to include this step for the following reasons:

a. This feature helps you identify the most suitable categories of your listings from Lazada's category tree

b. Mis-categorized products are subject to deactivation according to Lazada's policy

API Usage Documentation

2、Get the attributes available for the target category
After getting the category ID, use the category ID to call the GetCategoryAttributes API to get the attributes that can be filled in a category.

API Usage Documentation

3、Get the brand id (if not, ignore)
Call the GetBrandByPages API to get the list of lazada brands and extract the "brand_id" field of the brand to be used and use it when creating the product.

Note: The brand list is different from country to country.

4、Upload images
External image URLs or local images cannot be used directly when creating or updating products, please refer to this document to convert images to Lazada image URLs.

API Usage Documentation

5、Upload video (if no video, please ignore)
Please refer to this document to upload a video and get the video id.

API Usage Documentation