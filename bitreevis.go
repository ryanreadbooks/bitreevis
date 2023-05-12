package bitreevis

// VisAsSvg visualize the binary tree with given root in a svg graphic.
// The svg graphic is saved with the given filename.
func VisAsSvg(root BiNode, filename string, opt *RenderOption) error {
	// convert into inner placeable node
	pRoot := NewPlaceableTreeFromBiNode(root)
	// perform layout
	pRoot = PerformLayout(pRoot, opt.SiblingSeparation, opt.NodeRadius, opt.LevelSeparation)
	// do rendering
	renderer := NewSvgRenderer()

	result := renderer.Render(pRoot, opt)
	err := result.Error()
	if err != nil {
		return err
	}

	// save svg graphic
	return result.Save(filename)
}
