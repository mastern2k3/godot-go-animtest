package main

import (
	"log"

	gd "github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot"
)

type Simple struct {
	godot.Node2D
}

func NewSimple() godot.Class {
	return &Simple{}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(godot.Log)

	gd.EnableDebug()

	godot.AutoRegister(NewSimple)
}

func main() {
}

func (h *Simple) X_Ready() {

	anim := h.GetNode(gd.NewNodePath("anim")).(godot.AnimationPlayerImplementer)
	// anim.Play("wobble", 1.0, -1.0, false)
	arr := gd.NewArray()
	arr.PushBack(gd.NewVariantString("wobble"))
	arr.PushBack(gd.NewVariantReal(1))
	arr.PushBack(gd.NewVariantReal(-1))
	arr.PushBack(gd.NewVariantBool(false))
	anim.Callv("play", arr)

	// len := anim.GetCurrentAnimationLength()
	len := anim.Get("current_animation_length")

	godot.Log.Warning("len:", len.AsReal())
}
