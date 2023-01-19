package testers

type TestProperty struct {
	testName string
}

func (property TestProperty) GetTestName() string {
	return property.testName
}
