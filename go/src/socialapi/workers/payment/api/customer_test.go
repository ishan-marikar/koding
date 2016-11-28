package api

import (
	"encoding/json"
	"koding/db/mongodb/modelhelper"
	"socialapi/rest"
	"socialapi/workers/common/tests"
	"socialapi/workers/payment"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stripe/stripe-go"
)

func TestCustomer(t *testing.T) {
	Convey("Given a user", t, func() {
		withTestServer(t, func(endpoint string) {
			withStubData(endpoint, func(username, groupName, sessionID string) {
				Convey("Then Group should have customer id", func() {
					group, err := modelhelper.GetGroup(groupName)
					tests.ResultedWithNoErrorCheck(group, err)

					So(group.Payment.Customer.ID, ShouldNotBeBlank)
					Convey("We should be able to get the customer", func() {
						getURL := endpoint + EndpointCustomerGet

						res, err := rest.DoRequestWithAuth("GET", getURL, nil, sessionID)
						So(err, ShouldBeNil)
						So(res, ShouldNotBeNil)

						v := &stripe.Customer{}
						err = json.Unmarshal(res, v)
						So(err, ShouldBeNil)

						So(v.Deleted, ShouldEqual, false)
						So(v.Desc, ShouldContainSubstring, groupName)
						So(len(v.Meta), ShouldBeGreaterThanOrEqualTo, 2)
						So(v.Meta["groupName"], ShouldEqual, groupName)
						So(v.Meta["username"], ShouldEqual, username)

						Convey("After adding credit card to the user", func() {
							addCreditCardToUserWithChecks(endpoint, sessionID)

							res, err = rest.DoRequestWithAuth("GET", getURL, nil, sessionID)
							So(err, ShouldBeNil)
							So(res, ShouldNotBeNil)

							Convey("Customer should have CC assigned", func() {
								v = &stripe.Customer{}
								err = json.Unmarshal(res, v)
								So(err, ShouldBeNil)

								So(v.DefaultSource, ShouldNotBeNil)
								So(v.DefaultSource.Deleted, ShouldBeFalse)
								So(v.DefaultSource.ID, ShouldNotBeEmpty)
							})
						})
					})
				})
			})
		})
	})
}

func TestCouponApply(t *testing.T) {
	Convey("Given a user", t, func() {
		withTestServer(t, func(endpoint string) {
			withStubData(endpoint, func(username, groupName, sessionID string) {
				withTestCoupon(func(couponID string) {
					Convey("After adding coupon to the user", func() {

						updateURL := endpoint + EndpointCustomerUpdate

						cp := &stripe.CustomerParams{
							Coupon: couponID,
						}

						req, err := json.Marshal(cp)
						So(err, ShouldBeNil)
						So(req, ShouldNotBeNil)

						res, err := rest.DoRequestWithAuth("POST", updateURL, req, sessionID)
						So(err, ShouldBeNil)
						So(res, ShouldNotBeNil)

						v := &stripe.Customer{}
						err = json.Unmarshal(res, v)
						So(err, ShouldBeNil)

						Convey("Customer should have coupon assigned", func() {
							So(v.Discount, ShouldNotBeNil)
							So(v.Discount.Coupon.ID, ShouldEqual, couponID)
							So(v.Discount.Coupon.Valid, ShouldBeTrue)
							So(v.Discount.Coupon.Deleted, ShouldBeFalse)
						})
					})
				})
			})
		})
	})
}

func TestInfoPlan(t *testing.T) {
	Convey("Given a user", t, func() {
		withTestServer(t, func(endpoint string) {
			withStubData(endpoint, func(username, groupName, sessionID string) {
				withTestPlan(func(planID string) {
					addCreditCardToUserWithChecks(endpoint, sessionID)
					withSubscription(endpoint, groupName, sessionID, planID, func(subscriptionID string) {
						Convey("We should be able to get info", func() {
							infoURL := endpoint + EndpointInfo
							res, err := rest.DoRequestWithAuth("GET", infoURL, nil, sessionID)
							tests.ResultedWithNoErrorCheck(res, err)

							v := &payment.Usage{}
							err = json.Unmarshal(res, v)
							So(err, ShouldBeNil)

							So(v.ExpectedPlan.ID, ShouldEqual, planID)
						})
					})
				})
			})
		})
	})
}
