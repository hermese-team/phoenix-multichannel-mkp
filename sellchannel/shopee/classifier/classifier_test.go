package classifier

import "testing"

func TestClassify(t *testing.T) {
	cases := []struct {
		code   int
		status string
		want   string
	}{
		// code=3 (ORDER_STATUS_UPDATE)
		{3, "UNPAID", OrderCreated},
		{3, "PROCESSED", OrderProcessed},
		{3, "READY_TO_SHIP", OrderReadyToShip},
		{3, "SHIPPED", OrderShipped},
		{3, "COMPLETED", OrderCompleted},
		{3, "CANCELLED", OrderCancelled},
		{3, "IN_CANCEL", OrderCancelRequested},
		{3, "SOME_UNKNOWN", OrderStatusUpdated},
		// code=0: poll-discovered — must classify by status, same as code=3
		{0, "READY_TO_SHIP", OrderReadyToShip},
		{0, "SHIPPED", OrderShipped},
		{0, "CANCELLED", OrderCancelled},
		{0, "SOME_UNKNOWN", OrderStatusUpdated},
		// other push codes
		{4, "", OrderTrackingUpdated},
		{5, "", OrderItemUpdated},
		{6, "", OrderBookingUpdated},
		{15, "", ProductStatusUpdated},
		{16, "", ProductAttrUpdated},
	}

	for _, c := range cases {
		got := Classify(c.code, c.status)
		if got != c.want {
			t.Errorf("Classify(%d, %q) = %q, want %q", c.code, c.status, got, c.want)
		}
	}
}
