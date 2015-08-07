package main

func stateAbbr(fips string) (abbr string) {
	states := make(map[string]string)
	states["01"] = "AL"
	states["11"] = "DC"
	return states[fips]
}
