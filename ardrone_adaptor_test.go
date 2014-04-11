package gobotArdrone

import (
	"github.com/hybridgroup/go-ardrone/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArdroneAdaptor", func() {
	var (
		adaptor *ArdroneAdaptor
	)

	BeforeEach(func() {
		adaptor = new(ArdroneAdaptor)
		ardroneConnect = func() (*ardrone.Client, error) {
			return nil, nil
		}
	})

	It("Must be able to Finalize", func() {
		Expect(adaptor.Finalize()).To(Equal(true))
	})
	It("Must be able to Connect", func() {
		Expect(adaptor.Connect()).To(Equal(true))
	})
	It("Must be able to Disconnect", func() {
		Expect(adaptor.Disconnect()).To(Equal(true))
	})
	It("Must be able to Reconnect", func() {
		Expect(adaptor.Reconnect()).To(Equal(true))
	})
})
