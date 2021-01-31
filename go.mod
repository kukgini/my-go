module github.com/kukgini/my-go

go 1.15

require (
	github.com/aws/aws-sdk-go v1.35.7
	github.com/hyperledger/aries-framework-go v0.1.5 // indirect
	github.com/jaytaylor/mockery-example v0.0.0-20170323165341-bf04a9147d8e
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.7.1
	github.com/stretchr/testify v1.6.1
	goji.io v2.0.2+incompatible
)

replace (
	github.com/kilic/bls12-381 => github.com/trustbloc/bls12-381 v0.0.0-20201104214312-31de2a204df8
	github.com/piprate/json-gold => github.com/trustbloc/json-gold v0.3.1-0.20200414173446-30d742ee949e
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6
)
