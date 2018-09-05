Simple unit test function to make testing a bit easier.

Instead of the verbose

~~~~~~~~~~~~~~~~
func TestME (t *testing.T) {
     if s != expectedStr {
          t.Error("strings don't match")
     }
}
~~~~~~~~~~~~~~~~


Use TEST_STR
~~~~~~~~~~~~~~~~
func TestME (t *testing.T) {

    ut.TEST_STR(t, s, expectedStr)
}

~~~~~~~~~~~~~~~~


Other functions include TEST(t, bool), TEST_EQ(t, a, b)
