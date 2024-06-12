// package convert // to test exported and imported packages should be same as go file
package humannumbers_test //To test only public packages (exported functions) is prevents the private functions
import (
	"testing"

	"github.com/umashankar-reddy-r/humannumbers"
)


func numbetTest(t *testing.T, numberToTest int, expectedResult string){
	result, err:= humannumbers.NumToWords(numberToTest)
	if err!=nil {
		t.Fatalf("valid Value %v resulted in error %v", numberToTest, err)
	}
	if result!=expectedResult {
		t.Fatalf("expected %v got %v", expectedResult, result)
	}
}
func TestOutOfRange(t *testing.T){

	_, err := humannumbers.NumToWords(-1)
	if err == nil {
		t.Fatal("out of range did not return error")
	}
}

func TestInRange(t *testing.T){
	numbetTest(t,10,"ten")
	
}

func TestZero(t *testing.T){
numbetTest(t,0,"zero")
}

func TestUnit(t *testing.T){
	numbetTest(t,13,"thirteen")
}

func TestHundred(t *testing.T){
	numbetTest(t,100,"hundred")
}

func TestTwenty(t *testing.T){
	numbetTest(t,20,"twenty")
}

