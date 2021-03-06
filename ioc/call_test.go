package ioc_test

import (
	. "github.com/l3eegbee/pigs/ioc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IOC call", func() {

	var (
		container *Container
	)

	BeforeEach(func() {
		container = NewContainer()
	})

	Describe("with simple injection", func() {

		It("should be able to get itself", func() {

			var injectedContainer *Container

			Expect(container.CallInjected(func(injected struct {
				Container *Container `inject:"ApplicationContainer"`
			}) {
				injectedContainer = injected.Container
			})).Should(Succeed())

			Expect(injectedContainer).To(Equal(container))

		})

		It("should inject a simple component", func() {

			container.PutFactory(SimpleStructFactory("A"), "A")

			var injectedA *SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A *SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).To(Equal(&SimpleStruct{"A"}))

		})

		It("should select not nil component", func() {

			container.PutFactory(func() *SimpleStruct {
				return nil
			}, "NIL", "A")

			container.PutFactory(SimpleStructFactory("A"), "NOT_NIL", "A")

			var injectedA *SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A *SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).To(Equal(&SimpleStruct{"A"}))

		})

		It("should inject test component if provided", func() {

			container.PutFactory(SimpleStructFactory("A"), "A")
			container.TestPutFactory(SimpleStructFactory("TEST"), "TEST", "A")

			var injectedA *SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A *SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).To(Equal(&SimpleStruct{"TEST"}))

		})

		It("should return an error if no component is provided", func() {

			var injectedA *SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A *SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(HaveOccurred())

		})

		It("should return an error if too many components are provided", func() {

			container.PutFactory(SimpleStructFactory("A1"), "A1", "A")
			container.PutFactory(SimpleStructFactory("A2"), "A2", "A")

			var injectedA *SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A *SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(HaveOccurred())

		})

		It("should inject an interface", func() {

			container.PutFactory(SimpleStructFactory("A"), "A")

			var injectedA SomethingDoer

			Expect(container.CallInjected(func(injected struct {
				A SomethingDoer
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).To(Equal(&SimpleStruct{"A"}))

		})

		It("should restore 'core' configuration after a TestClear", func() {

			container.PutFactory(SimpleStructFactory("A"), "A")
			container.TestPutFactory(SimpleStructFactory("TEST"), "TEST", "A")

			var injectedA *SimpleStruct

			// test

			Expect(container.CallInjected(func(injected struct {
				A *SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).To(Equal(&SimpleStruct{"TEST"}))

			// restore

			container.ClearTests()

			// core

			Expect(container.CallInjected(func(injected struct {
				A *SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).To(Equal(&SimpleStruct{"A"}))

		})

	})

	Describe("with slice injection", func() {

		It("should inject a slice", func() {

			container.PutFactory(SimpleStructFactory("A1"), "A1", "A")
			container.PutFactory(SimpleStructFactory("A2"), "A2", "A")
			container.PutFactory(SimpleStructFactory("A3"), "A3", "A")

			var injectedA []*SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A []*SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).Should(ConsistOf(
				&SimpleStruct{"A1"},
				&SimpleStruct{"A2"},
				&SimpleStruct{"A3"}))

		})

		It("should inject a slice without nil components", func() {

			container.PutFactory(func() *SimpleStruct {
				return nil
			}, "NIL", "A")

			container.PutFactory(SimpleStructFactory("A1"), "A1", "A")
			container.PutFactory(SimpleStructFactory("A2"), "A2", "A")
			container.PutFactory(SimpleStructFactory("A3"), "A3", "A")

			var injectedA []*SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A []*SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).Should(ConsistOf(
				&SimpleStruct{"A1"},
				&SimpleStruct{"A2"},
				&SimpleStruct{"A3"}))

		})

		It("should inject a slice of interface", func() {

			container.PutFactory(SimpleStructFactory("A1"), "A1", "A")
			container.PutFactory(SimpleStructFactory("A2"), "A2", "A")
			container.PutFactory(SimpleStructFactory("A3"), "A3", "A")

			var injectedA []SomethingDoer

			Expect(container.CallInjected(func(injected struct {
				A []SomethingDoer
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).Should(ConsistOf(
				&SimpleStruct{"A1"},
				&SimpleStruct{"A2"},
				&SimpleStruct{"A3"}))

		})

	})

	Describe("with map injection", func() {

		It("should inject a map", func() {

			container.PutFactory(SimpleStructFactory("A1"), "A1", "A")
			container.PutFactory(SimpleStructFactory("A2"), "A2", "A")
			container.PutFactory(SimpleStructFactory("A3"), "A3", "A")

			var injectedA map[string]*SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A map[string]*SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).Should(HaveLen(3))
			Expect(injectedA).Should(HaveKeyWithValue("A1", &SimpleStruct{"A1"}))
			Expect(injectedA).Should(HaveKeyWithValue("A2", &SimpleStruct{"A2"}))
			Expect(injectedA).Should(HaveKeyWithValue("A3", &SimpleStruct{"A3"}))

		})

		It("should inject a map without nil components", func() {

			container.PutFactory(func() *SimpleStruct {
				return nil
			}, "NIL", "A")

			container.PutFactory(SimpleStructFactory("A1"), "A1", "A")
			container.PutFactory(SimpleStructFactory("A2"), "A2", "A")
			container.PutFactory(SimpleStructFactory("A3"), "A3", "A")

			var injectedA map[string]*SimpleStruct

			Expect(container.CallInjected(func(injected struct {
				A map[string]*SimpleStruct
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).Should(HaveLen(3))
			Expect(injectedA).Should(HaveKeyWithValue("A1", &SimpleStruct{"A1"}))
			Expect(injectedA).Should(HaveKeyWithValue("A2", &SimpleStruct{"A2"}))
			Expect(injectedA).Should(HaveKeyWithValue("A3", &SimpleStruct{"A3"}))

		})

		It("should inject a map of interface", func() {

			container.PutFactory(SimpleStructFactory("A1"), "A1", "A")
			container.PutFactory(SimpleStructFactory("A2"), "A2", "A")
			container.PutFactory(SimpleStructFactory("A3"), "A3", "A")

			var injectedA map[string]SomethingDoer

			Expect(container.CallInjected(func(injected struct {
				A map[string]SomethingDoer
			}) {
				injectedA = injected.A
			})).Should(Succeed())

			Expect(injectedA).Should(HaveLen(3))
			Expect(injectedA).Should(HaveKeyWithValue("A1", &SimpleStruct{"A1"}))
			Expect(injectedA).Should(HaveKeyWithValue("A2", &SimpleStruct{"A2"}))
			Expect(injectedA).Should(HaveKeyWithValue("A3", &SimpleStruct{"A3"}))

		})

	})

})
