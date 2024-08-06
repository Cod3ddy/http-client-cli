package main

import (
	"testing"

	cn "github.com/Cod3ddy/http-client-cli/cmd/common"
)

func TestConvertToJSON(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{"{username:Shadowyhands,password:mdskipper}", "{\"password\":\"mdskipper\",\"username\":\"Shadowyhands\"}"},
		{"{id:qwerty,gender:male}", "{\"gender\":\"male\",\"id\":\"qwerty\"}"},
	}

	for _, c := range cases {
		got, err:= cn.ConvertToJSON(c.Input)
		if err != nil{
			t.Error(err)
		}

		if got != c.Expected{
			t.Errorf("ConvertToJson(%q) = %q want %q", c.Input, got, c.Expected)
		}

	}
}
