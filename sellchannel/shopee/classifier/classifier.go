// Package classifier maps Shopee push notification codes and order statuses
// to canonical event types used across all sell-channels.
package classifier

import "fmt"

// Canonical event type constants — shared across sell-channels.
const (
	OrderCreated         = "ORDER_CREATED"
	OrderProcessed       = "ORDER_PROCESSED"
	OrderReadyToShip     = "ORDER_READY_TO_SHIP"
	OrderShipped         = "ORDER_SHIPPED"
	OrderCompleted       = "ORDER_COMPLETED"
	OrderCancelled       = "ORDER_CANCELLED"
	OrderCancelRequested = "ORDER_CANCEL_REQUESTED"
	OrderStatusUpdated   = "ORDER_STATUS_UPDATED"
	OrderTrackingUpdated = "ORDER_TRACKING_UPDATED"
	OrderItemUpdated     = "ORDER_ITEM_UPDATED"
	OrderBookingUpdated  = "ORDER_BOOKING_UPDATED"
	ProductStatusUpdated = "PRODUCT_STATUS_UPDATED"
	ProductAttrUpdated   = "PRODUCT_ATTR_UPDATED"
)

// Classify maps a Shopee push code + status to a canonical EventType string.
//
// Shopee push codes: https://open.shopee.com/documents/v2/v2.push.order_status
//
//	3  = ORDER_STATUS_UPDATE   (status field carries the lifecycle state)
//	4  = ORDER_TRACKING_NUMBER_UPDATE
//	5  = ORDER_ITEM_STATUS_UPDATE
//	6  = ORDER_BOOKING_STATUS_UPDATE
//	15 = PRODUCT_ITEM_UPDATED
//	16 = PRODUCT_ITEM_ATTRIBUTE_UPDATED
func Classify(code int, status string) string {
	switch code {
	case 3:
		switch status {
		case "UNPAID":
			return OrderCreated
		case "PROCESSED":
			return OrderProcessed
		case "READY_TO_SHIP":
			return OrderReadyToShip
		case "SHIPPED":
			return OrderShipped
		case "COMPLETED":
			return OrderCompleted
		case "CANCELLED":
			return OrderCancelled
		case "IN_CANCEL":
			return OrderCancelRequested
		default:
			return OrderStatusUpdated
		}
	case 4:
		return OrderTrackingUpdated
	case 5:
		return OrderItemUpdated
	case 6:
		return OrderBookingUpdated
	case 15:
		return ProductStatusUpdated
	case 16:
		return ProductAttrUpdated
	default:
		return fmt.Sprintf("UNKNOWN_%d", code)
	}
}
