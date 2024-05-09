package main

import (
	"errors"
	"github.com/Kagami/go-face"
)

type Face struct {
	Cords      Cords           `json:"cords"`
	Descriptor face.Descriptor `json:"descriptor"`
}

type Cords struct {
	Left   int `json:"left"`
	Top    int `json:"top"`
	Right  int `json:"right"`
	Bottom int `json:"bottom"`
}

var recognizer *face.Recognizer

func recognizeFaces(img []byte) ([]Face, error) {
	recognizedFaces, err := recognizer.Recognize(img)

	if err != nil {
		err := errors.New("wrong file")
		return nil, err
	}

	var faces []Face

	for _, f := range recognizedFaces {
		rect := f.Rectangle
		faces = append(faces, Face{
			Cords:      Cords{Left: rect.Min.X, Top: rect.Min.Y, Right: rect.Max.X, Bottom: rect.Max.Y},
			Descriptor: f.Descriptor,
		})
	}

	return faces, nil

}
