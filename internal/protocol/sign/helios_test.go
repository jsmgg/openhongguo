package sign

import "testing"

// TestHeliosMatchesCapturedDevice pins x-helios against a captured openHarmony
// 7.4.4.32 playback request. Decoding the header yields the random prefix
// 0x1162b475 and the plaintext "1794874377-279266862-8662" (07 07 ... padding),
// which fixes both the licence id and the padding scheme.
//
// The vector was checked live on 2026-11-17: re-signing that request with a
// fresh random prefix and the same timestamp returned the full playback model,
// while the same timestamp with channelID (1588093228) returned an empty body.
func TestHeliosMatchesCapturedDevice(t *testing.T) {
	const want = "dbRiEbe/CaSUZcKSKjDMti2rF0vKMO9EtDbK6KGUAR1HMweQ"
	if got := assembleHelios(1794874377, 0x1162b475); got != want {
		t.Fatalf("x-helios mismatch\n got %s\nwant %s", got, want)
	}
}
