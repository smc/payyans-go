package payyans

// import (
// 	"main/normalizer"
// 	"testing"
// )

// func testHelper(t *testing.T, lhs, rhs string) {
// 	lhs, err := normalizer.Normalize(lhs)
// 	if err != nil {
// 		t.Errorf("normalize function failed %s", err)
// 	}
// 	if lhs != rhs {
// 		t.Errorf("Failed, got %s expected %s", lhs, rhs)
// 	}
// }

// func TestNormalizerSingleLineMl(t *testing.T) {
// 	testHelper(t, `പൂമ്പാററ`, `പൂമ്പാറ്റ`)
// 	// TODO: make the tests cases legit.
// 	testHelper(t, `അവിൽ`, `അവില്‍`)
// 	testHelper(t, `രമണൻ`, `രമണന്‍`)
// 	testHelper(t, `അവൾ`, `അവള്‍`)
// 	testHelper(t, `ശ്രാവൺ`, `ശ്രാവണ്‍`)
// 	testHelper(t, `അവിൽപൊതി`, `അവില്‍പൊതി`)

// }

// func TestNormalizerMultiLineMl(t *testing.T) {
// 	input := `കുഞ്ചൻ നമ്പ്യാർ
// 			ചെണ്ടമേളം`
// 	expected := `കുഞ്ചന്‍ നമ്പ്യാര്‍
// 			ചെണ്ടമേളം`
// 	testHelper(t, input, expected)

// }
