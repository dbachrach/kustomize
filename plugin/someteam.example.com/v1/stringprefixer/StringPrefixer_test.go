// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	"sigs.k8s.io/kustomize/v3/pluglib"

	"sigs.k8s.io/kustomize/v3/pkg/kusttest"
)

func TestStringPrefixerPlugin(t *testing.T) {
	tc := pluglib.NewEnvForTest(t).Set()
	defer tc.Reset()

	tc.BuildGoPlugin(
		"someteam.example.com", "v1", "StringPrefixer")
	th := kusttest_test.NewKustTestPluginHarness(t, "/app")

	m := th.LoadAndRunTransformer(`
apiVersion: someteam.example.com/v1
kind: StringPrefixer
metadata:
  name: wowsa
`,
		`apiVersion: apps/v1
kind: MeatBall
metadata:
  name: meatball
`)

	th.AssertActualEqualsExpected(m, `
apiVersion: apps/v1
kind: MeatBall
metadata:
  name: wowsa-meatball
`)
}
