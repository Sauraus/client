// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package engine

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/keybase/client/go/libkb"
)

func TestSaltPackSignDeviceRequired(t *testing.T) {
	tc := SetupEngineTest(t, "sign")
	defer tc.Cleanup()

	ctx := &Context{}
	eng := NewSaltPackSign(nil, tc.G)
	err := RunEngine(eng, ctx)
	if err == nil {
		t.Fatal("sign not logged in returned no error")
	}
	if _, ok := err.(libkb.DeviceRequiredError); !ok {
		t.Errorf("error type: %T, expected DeviceRequiredError", err)
	}
}

func TestSaltPackSignVerify(t *testing.T) {
	tc := SetupEngineTest(t, "sign")
	defer tc.Cleanup()

	fu := CreateAndSignupFakeUser(tc, "sign")
	_ = fu

	for _, test := range signTests {
		var sink bytes.Buffer

		earg := &SaltPackSignArg{
			Sink:   libkb.NopWriteCloser{W: &sink},
			Source: ioutil.NopCloser(bytes.NewBufferString(test.input)),
		}

		eng := NewSaltPackSign(earg, tc.G)
		ctx := &Context{}

		if err := RunEngine(eng, ctx); err != nil {
			t.Errorf("%s: run error: %s", test.name, err)
			continue
		}

		sig := sink.String()

		if len(sig) == 0 {
			t.Errorf("%s: empty sig", test.name)
		}

		/*
			_, err = key.VerifyString(sig, []byte(test.input))
			if err != nil {
				t.Errorf("%s: verify error: %s", test.name, err)
				continue
			}
		*/
	}
}