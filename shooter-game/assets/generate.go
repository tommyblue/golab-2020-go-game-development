//go:generate file2byteslice -input ./hud.png  -output hud.go -package assets -var hudBytes
//go:generate file2byteslice -input ./objects.png  -output objects.go -package assets -var objectsBytes
//go:generate file2byteslice -input ./stall.png  -output stall.go -package assets -var stallBytes
package assets
