package main

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {
		hero1 := NewHero(&theWasp{})
		hero1.Talk()

		hero2 := NewHero(&antMan{})
		hero2.Talk()
	})
}
