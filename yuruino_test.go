package main

import "testing"

func TestCommunityName(t *testing.T) {
    const expected = "Yuruhachi!"
    if communityName() != expected {
        t.Error("wrong community name!")
    }
}
