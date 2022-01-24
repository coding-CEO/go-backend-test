package domainUtils

//TODO: add proper safe domains here
var AllowedDomains = []string{"http://localhost:3000"}

func IsAllowedDomain(domain string) bool {
	for _, allowedDomain := range(AllowedDomains) {
		if allowedDomain == domain {
			return true
		}
	}
	return false
}