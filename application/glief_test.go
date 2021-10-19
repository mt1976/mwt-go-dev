package application

import (
	"log"
	"testing"
)

func TestGlief(t *testing.T) {

	g, e := GLIEF_leiLookup("GB00BL68HJ26")
	if e != nil {
		log.Println(e.Error())
	}
	t.Log(g)
	log.Printf("LEI=%q", g)

}

func TestGliefFail(t *testing.T) {

	g, e := GLIEF_leiLookup("PO00BL68poop")
	if e != nil {
		log.Println(e.Error())
	}
	t.Log(g)
	log.Printf("LEI=%q", g)

}
