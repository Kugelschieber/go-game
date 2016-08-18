package goga

func GetSpriteRenderer() *SpriteRenderer {
	renderer, ok := GetSystemByName(sprite_renderer_name).(*SpriteRenderer)

	if !ok {
		panic("Could not obtain sprite renderer")
	}

	return renderer
}

func GetModelRenderer() *ModelRenderer {
	renderer, ok := GetSystemByName(model_renderer_name).(*ModelRenderer)

	if !ok {
		panic("Could not obtain model renderer")
	}

	return renderer
}

func GetCulling2DSystem() *Culling2D {
	system, ok := GetSystemByName(culling_2d_name).(*Culling2D)

	if !ok {
		panic("Could not obtain culling system")
	}

	return system
}

func GetKeyframeRenderer() *KeyframeRenderer {
	renderer, ok := GetSystemByName(keyframe_sprite_renderer_name).(*KeyframeRenderer)

	if !ok {
		panic("Could not obtain keyframe renderer")
	}

	return renderer
}

func GetTextRenderer() *TextRenderer {
	renderer, ok := GetSystemByName(text_renderer_name).(*TextRenderer)

	if !ok {
		panic("Could not obtain text renderer")
	}

	return renderer
}
