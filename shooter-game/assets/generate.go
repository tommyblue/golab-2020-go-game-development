//go:generate file2byteslice -input ./hud.png  -output hud.go -package assets -var hudBytes
//go:generate file2byteslice -input ./objects.png  -output objects.go -package assets -var objectsBytes
//go:generate file2byteslice -input ./stall.png  -output stall.go -package assets -var stallBytes
//go:generate file2byteslice -input ./hit.wav  -output hit_sound.go -package assets -var HitSound
//go:generate file2byteslice -input ./miss.wav  -output miss_sound.go -package assets -var MissSound
//go:generate file2byteslice -input ./ragtime.ogg  -output ragtime_sound.go -package assets -var RagtimeSound
package assets
