package mob_smssdk

import (
	"fmt"
	"testing"
)

func TestMobSmsSdkClient_Verify(t *testing.T) {
	client := NewMobSmsSdk("1111111111111")
	err := client.Verify("13900000000", "86", "1234")
	fmt.Printf("err: %v\n", err)
}