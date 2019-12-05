package main_test

import (
	"testing"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jaytaylor/mockery-example/mocks"
	"github.com/stretchr/testify/mock"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsS3API Test function Suite")

}

var _ = Describe("AwsS3API", func() {
    var (
        mockS3 *mocks.S3API
    )

    BeforeEach(func() {
	    mockS3 = &mocks.S3API{}
    })

    Describe("S3 test", func() {
        Context("in some context", func() {
            It("should be like this....", func() {
				expected := []*s3.CommonPrefix{
					&s3.CommonPrefix{
						Prefix: aws.String("2017-01-01"),
					},
				}
	            mockResultFn := func(input *s3.ListObjectsInput) *s3.ListObjectsOutput {
		            output := &s3.ListObjectsOutput{}
		            output.SetCommonPrefixes(expected)
		            return output
	            }

	            // NB: .Return(...) must return the same signature as the method being mocked.
	            //     In this case it's (*s3.ListObjectsOutput, error).
	            mockS3.On("ListObjects", mock.MatchedBy(func(input *s3.ListObjectsInput) bool {
		            return input.Delimiter != nil && *input.Delimiter == "/" && input.Prefix == nil
	            })).Return(mockResultFn, nil)

				listingInput := &s3.ListObjectsInput{
					Bucket:    aws.String("foo"),
					Delimiter: aws.String("/"),
				}

				listingOutput, err := mockS3.ListObjects(listingInput)
				if err != nil {
					panic(err)
				}

				for i, x := range listingOutput.CommonPrefixes {
                    Expect(&x).To(Equal(&expected[i]))
				}
			})
        })
    })
})
