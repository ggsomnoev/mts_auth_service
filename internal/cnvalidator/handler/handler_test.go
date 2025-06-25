package handler_test

import (
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ggsomnoev/mts-auth-service/internal/cnvalidator/handler"
)

var _ = Describe("TLS Auth Handler", func() {
	var (
		e                *echo.Echo
		trustedClientCNs []string
	)

	BeforeEach(func() {
		e = echo.New()
		trustedClientCNs = []string{"trusted-cn-1", "trusted-cn-2"}
		handler.RegisterHandlers(e, trustedClientCNs)
	})

	Describe("/auth", func() {
		var (
			req      *http.Request
			recorder *httptest.ResponseRecorder
			cert     *x509.Certificate
		)

		BeforeEach(func() {
			req = httptest.NewRequest(http.MethodGet, "/auth", nil)
			recorder = httptest.NewRecorder()

			if cert != nil {
				req.TLS = &tls.ConnectionState{
					PeerCertificates: []*x509.Certificate{cert},
				}
			}
		})

		JustBeforeEach(func() {
			e.ServeHTTP(recorder, req)
		})

		When("no TLS client certificate is provided", func() {
			It("returns 401", func() {
				Expect(recorder.Code).To(Equal(401))
				Expect(recorder.Body.String()).To(ContainSubstring("no TLS client certificate"))
			})
		})

		When("an invalid TLS state is provided", func() {
			BeforeEach(func() {
				req.TLS = &tls.ConnectionState{
					PeerCertificates: []*x509.Certificate{},
				}
			})

			It("returns 401", func() {
				Expect(recorder.Code).To(Equal(401))
				Expect(recorder.Body.String()).To(ContainSubstring("no TLS client certificate"))
			})
		})

		When("client cert CN is valid", func() {
			BeforeEach(func() {
				cert = &x509.Certificate{
					Subject: pkix.Name{CommonName: "trusted-cn-1"},
				}
				req.TLS = &tls.ConnectionState{
					PeerCertificates: []*x509.Certificate{cert},
				}
			})

			It("returns 200 and correct CN", func() {
				Expect(recorder.Code).To(Equal(200))
				Expect(recorder.Body.String()).To(ContainSubstring("Authentication and authorization successful"))
				Expect(recorder.Body.String()).To(ContainSubstring("trusted-cn-1"))
			})
		})

		When("client cert CN is invalid", func() {
			BeforeEach(func() {
				cert = &x509.Certificate{
					Subject: pkix.Name{CommonName: "invalid-cn"},
				}
				req.TLS = &tls.ConnectionState{
					PeerCertificates: []*x509.Certificate{cert},
				}
			})

			It("returns 403 and shows invalid CN", func() {
				Expect(recorder.Code).To(Equal(403))
				Expect(recorder.Body.String()).To(ContainSubstring("invalid certificate CN"))
				Expect(recorder.Body.String()).To(ContainSubstring("invalid-cn"))
			})
		})
	})
})
