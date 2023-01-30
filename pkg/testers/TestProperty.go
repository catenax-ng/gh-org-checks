package testers

type TestProperty struct {
	testName string
}

func NewTestProperty(testName string) TestProperty {
	return TestProperty{testName: testName}
}

func (property TestProperty) GetTestName() string {
	return property.testName
}
