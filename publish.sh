# Make sure to commit and push first

VERSION=${1:-v0.1.0}
git tag ${VERSION}
git push origin ${VERSION}
GOPROXY=proxy.golang.org go list -m github.com/rivernews/GoTools@${VERSION}
