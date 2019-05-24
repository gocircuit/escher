package asyncassertion_test

import (
	"errors"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega/internal/asyncassertion"
	. "github.com/onsi/ginkgo"
	"time"
)

var _ = Describe("Async Assertion", func() {
	var (
		failureMessage string
		callerSkip     int
	)

	var fakeFailHandler = func(message string, skip ...int) {
		failureMessage = message
		callerSkip = skip[0]
	}

	BeforeEach(func() {
		failureMessage = ""
		callerSkip = 0
	})

	Describe("Eventually", func() {
		Context("the positive case", func() {
			It("should poll the function and matcher", func() {
				arr := []int{}
				a := New(AsyncAssertionTypeEventually, func() []int {
					arr = append(arr, 1)
					return arr
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

				a.Should(HaveLen(10))

				Ω(arr).Should(HaveLen(10))
				Ω(failureMessage).Should(BeZero())
			})

			It("should continue when the matcher errors", func() {
				var arr = []int{}
				a := New(AsyncAssertionTypeEventually, func() interface{} {
					arr = append(arr, 1)
					if len(arr) == 4 {
						return 0 //this should cause the matcher to error
					}
					return arr
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

				a.Should(HaveLen(4), "My description %d", 2)

				Ω(failureMessage).Should(ContainSubstring("Timed out after"))
				Ω(failureMessage).Should(ContainSubstring("My description 2"))
				Ω(callerSkip).Should(Equal(4))
			})

			It("should be able to timeout", func() {
				arr := []int{}
				a := New(AsyncAssertionTypeEventually, func() []int {
					arr = append(arr, 1)
					return arr
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

				a.Should(HaveLen(11), "My description %d", 2)

				Ω(arr).Should(HaveLen(10))
				Ω(failureMessage).Should(ContainSubstring("Timed out after"))
				Ω(failureMessage).Should(ContainSubstring("<[]int | len:10"), "Should pass the correct value to the matcher message formatter.")
				Ω(failureMessage).Should(ContainSubstring("My description 2"))
				Ω(callerSkip).Should(Equal(4))
			})
		})

		Context("the negative case", func() {
			It("should poll the function and matcher", func() {
				counter := 0
				arr := []int{}
				a := New(AsyncAssertionTypeEventually, func() []int {
					counter += 1
					if counter >= 10 {
						arr = append(arr, 1)
					}
					return arr
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

				a.ShouldNot(HaveLen(0))

				Ω(arr).Should(HaveLen(1))
				Ω(failureMessage).Should(BeZero())
			})

			It("should timeout when the matcher errors", func() {
				a := New(AsyncAssertionTypeEventually, func() interface{} {
					return 0 //this should cause the matcher to error
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

				a.ShouldNot(HaveLen(0), "My description %d", 2)

				Ω(failureMessage).Should(ContainSubstring("Timed out after"))
				Ω(failureMessage).Should(ContainSubstring("Error:"))
				Ω(failureMessage).Should(ContainSubstring("My description 2"))
				Ω(callerSkip).Should(Equal(4))
			})

			It("should be able to timeout", func() {
				a := New(AsyncAssertionTypeEventually, func() []int {
					return []int{}
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

				a.ShouldNot(HaveLen(0), "My description %d", 2)

				Ω(failureMessage).Should(ContainSubstring("Timed out after"))
				Ω(failureMessage).Should(ContainSubstring("<[]int | len:0"), "Should pass the correct value to the matcher message formatter.")
				Ω(failureMessage).Should(ContainSubstring("My description 2"))
				Ω(callerSkip).Should(Equal(4))
			})
		})

		Context("with a function that returns multiple values", func() {
			It("should eventually succeed if the additional arguments are nil", func() {
				i := 0
				Eventually(func() (int, error) {
					i++
					return i, nil
				}).Should(Equal(10))
			})

			It("should eventually timeout if the additional arguments are not nil", func() {
				i := 0
				a := New(AsyncAssertionTypeEventually, func() (int, error) {
					i++
					return i, errors.New("bam")
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)
				a.Should(Equal(2))

				Ω(failureMessage).Should(ContainSubstring("Timed out after"))
				Ω(failureMessage).Should(ContainSubstring("Error:"))
				Ω(failureMessage).Should(ContainSubstring("bam"))
				Ω(callerSkip).Should(Equal(4))
			})
		})
	})

	Describe("Consistently", func() {
		Describe("The positive case", func() {
			Context("when the matcher consistently passes for the duration", func() {
				It("should pass", func() {
					calls := 0
					a := New(AsyncAssertionTypeConsistently, func() string {
						calls++
						return "foo"
					}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

					a.Should(Equal("foo"))
					Ω(calls).Should(Equal(10))
					Ω(failureMessage).Should(BeZero())
				})
			})

			Context("when the matcher fails at some point", func() {
				It("should fail", func() {
					calls := 0
					a := New(AsyncAssertionTypeConsistently, func() interface{} {
						calls++
						if calls > 9 {
							return "bar"
						}
						return "foo"
					}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

					a.Should(Equal("foo"))
					Ω(failureMessage).Should(ContainSubstring("to equal"))
					Ω(callerSkip).Should(Equal(4))
				})
			})

			Context("when the matcher errors at some point", func() {
				It("should fail", func() {
					calls := 0
					a := New(AsyncAssertionTypeConsistently, func() interface{} {
						calls++
						if calls > 5 {
							return 3
						}
						return []int{1, 2, 3}
					}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

					a.Should(HaveLen(3))
					Ω(failureMessage).Should(ContainSubstring("HaveLen matcher expects"))
					Ω(callerSkip).Should(Equal(4))
				})
			})
		})

		Describe("The negative case", func() {
			Context("when the matcher consistently passes for the duration", func() {
				It("should pass", func() {
					c := make(chan bool)
					a := New(AsyncAssertionTypeConsistently, c, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

					a.ShouldNot(Receive())
					Ω(failureMessage).Should(BeZero())
				})
			})

			Context("when the matcher fails at some point", func() {
				It("should fail", func() {
					c := make(chan bool)
					go func() {
						time.Sleep(time.Duration(100 * time.Millisecond))
						c <- true
					}()

					a := New(AsyncAssertionTypeConsistently, c, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

					a.ShouldNot(Receive())
					Ω(failureMessage).Should(ContainSubstring("not to receive anything"))
				})
			})

			Context("when the matcher errors at some point", func() {
				It("should fail", func() {
					calls := 0
					a := New(AsyncAssertionTypeConsistently, func() interface{} {
						calls++
						return calls
					}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)

					a.ShouldNot(BeNumerically(">", 5))
					Ω(failureMessage).Should(ContainSubstring("not to be >"))
					Ω(callerSkip).Should(Equal(4))
				})
			})
		})

		Context("with a function that returns multiple values", func() {
			It("should consistently succeed if the additional arguments are nil", func() {
				i := 2
				Consistently(func() (int, error) {
					i++
					return i, nil
				}).Should(BeNumerically(">=", 2))
			})

			It("should eventually timeout if the additional arguments are not nil", func() {
				i := 2
				a := New(AsyncAssertionTypeEventually, func() (int, error) {
					i++
					return i, errors.New("bam")
				}, fakeFailHandler, time.Duration(0.2*float64(time.Second)), time.Duration(0.02*float64(time.Second)), 1)
				a.Should(BeNumerically(">=", 2))

				Ω(failureMessage).Should(ContainSubstring("Error:"))
				Ω(failureMessage).Should(ContainSubstring("bam"))
				Ω(callerSkip).Should(Equal(4))
			})
		})
	})

	Context("when passed a function with the wrong # or arguments & returns", func() {
		It("should panic", func() {
			Ω(func() {
				New(AsyncAssertionTypeEventually, func() {}, fakeFailHandler, 0, 0, 1)
			}).Should(Panic())

			Ω(func() {
				New(AsyncAssertionTypeEventually, func(a string) int { return 0 }, fakeFailHandler, 0, 0, 1)
			}).Should(Panic())

			Ω(func() {
				New(AsyncAssertionTypeEventually, func() int { return 0 }, fakeFailHandler, 0, 0, 1)
			}).ShouldNot(Panic())

			Ω(func() {
				New(AsyncAssertionTypeEventually, func() (int, error) { return 0, nil }, fakeFailHandler, 0, 0, 1)
			}).ShouldNot(Panic())
		})
	})

	Describe("bailing early", func() {
		Context("when actual is a value", func() {
			It("Eventually should bail out and fail early if the matcher says to", func() {
				c := make(chan bool)
				close(c)

				t := time.Now()
				failures := InterceptGomegaFailures(func() {
					Eventually(c, 0.1).Should(Receive())
				})
				Ω(time.Since(t)).Should(BeNumerically("<", 90*time.Millisecond))

				Ω(failures).Should(HaveLen(1))
			})
		})

		Context("when actual is a function", func() {
			It("should never bail early", func() {
				c := make(chan bool)
				close(c)

				t := time.Now()
				failures := InterceptGomegaFailures(func() {
					Eventually(func() chan bool {
						return c
					}, 0.1).Should(Receive())
				})
				Ω(time.Since(t)).Should(BeNumerically(">=", 90*time.Millisecond))

				Ω(failures).Should(HaveLen(1))
			})
		})
	})
})
