package components

import (
	"image"

	"github.com/kongbong/invadersGo/ecs"

	"github.com/disintegration/gift"
)

type Sprite interface {
	ecs.Component
	GetSrc() image.Image
	GetSize() image.Rectangle
	GetFilter() *gift.GIFT
	GetFilterA() *gift.GIFT
	GetFilterE() *gift.GIFT
}

func NewSprite(src image.Image, sprite, alt, explode image.Rectangle) Sprite {
	return &implSprite{
		src:     src,
		size:    sprite,
		filter:  gift.New(gift.Crop(sprite)),
		filterA: gift.New(gift.Crop(alt)),
		filterE: gift.New(gift.Crop(explode)),
	}
}

// Sprite represents a sprite in the game
type implSprite struct {
	src     image.Image
	size    image.Rectangle // the sprite size
	filter  *gift.GIFT      // normal filter used to draw the sprite
	filterA *gift.GIFT      // alternate filter used to draw the sprite
	filterE *gift.GIFT      // exploded filter used to draw the sprite
}

func (s *implSprite) GetType() int {
	return CompTypeSprite
}

func (s *implSprite) GetSrc() image.Image {
	return s.src
}

func (s *implSprite) GetSize() image.Rectangle {
	return s.size
}

func (s *implSprite) GetFilter() *gift.GIFT {
	return s.filter
}

func (s *implSprite) GetFilterA() *gift.GIFT {
	return s.filterA
}

func (s *implSprite) GetFilterE() *gift.GIFT {
	return s.filterE
}
