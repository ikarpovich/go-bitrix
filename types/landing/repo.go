package landing

type RepoRegisterRequest struct {
	Code string `json:"code" url:"code"`
	Fields BlockFields `json:"fields" url:"fields"`
	Manifest BlockManifest `json:"manifest" url:"manifest"`
}

type BlockFields struct {
	Name string `json:"name" url:"NAME"`
	Description string `json:"description" url:"DESCRIPTION"`
	Sections string `json:"sections" url:"SECTIONS"`
	Preview string `json:"preview" url:"PREVIEW"`
	Content string `json:"content" url:"CONTENT"`
}

type BlockManifest struct {
	Assets BlockManifestAssets `json:"assets" url:"assets"`
	Nodes map[string]BlockManifestNode `json:"nodes" url:"nodes"`
	Style map[string]BlockManifestStyle `json:"style" url:"style"`
	Attrs map[string][]BlockManifestAttr `json:"attrs" url:"attrs"`
}

type BlockManifestAssets struct {
	Css []string `json:"css" url:"css"`
	Js []string `json:"js" url:"js"`
}

type BlockManifestNode struct {
	Name string `json:"name" url:"name"`
	Type string `json:"type" url:"type"`
}

const (
	BlockManifestNodeTypeText  string = "text"
	BlockManifestNodeTypeLink  string = "link"
	BlockManifestNodeTypeImage string = "image"
)

type BlockManifestStyle struct {
	Name string `json:"name" url:"name"`
	Type string `json:"type" url:"type"`
}

const (
	BlockManifestStyleTypeText  string = "typo"
	BlockManifestStyleTypeLink  string = "box"
	BlockManifestStyleTypeImage string = "button"
)

type BlockManifestAttr struct {
	Name string `json:"name" url:"name"`
	Attribute string `json:"attribute" url:"attribute"`
	Items map[string]string `json:"items" url:"items"`
	Type string `json:"type" url:"type"`
	Value interface{} `json:"value" url:"value"`
	Placeholder string `json:"placeholder" url:"placeholder"`
}

type BlockManifestAttrValueImage struct {
	Src string `json:"src" url:"src"`
	Alt string `json:"alt" url:"alt"`
}

type BlockManifestAttrValueIcon struct {
	ClassList []string `json:"classList" url:"classList"`
}

type BlockManifestAttrValueLink struct {
	Text string `json:"text" url:"text"`
	Href string `json:"href" url:"href"`
	Target string `json:"target" url:"target"`
}

type BlockManifestAttrValueRange struct {
	From string `json:"from" url:"from"`
	To string `json:"to" url:"to"`
}

type BlockManifestAttrImage struct {
	BlockManifestAttr
	Value BlockManifestAttrValueImage `json:"value" url:"value"`
}

type BlockManifestAttrIcon struct {
	BlockManifestAttr
	Value BlockManifestAttrValueIcon `json:"value" url:"value"`
}

type BlockManifestAttrLink struct {
	BlockManifestAttr
	Value BlockManifestAttrValueLink `json:"value" url:"value"`
}

type BlockManifestAttrRange struct {
	BlockManifestAttr
	Value BlockManifestAttrValueRange `json:"value" url:"value"`
}

type BlockManifestAttrText struct {
	BlockManifestAttr
	Value string `json:"value" url:"value"`
}

type BlockManifestAttrTextArray struct {
	BlockManifestAttr
	Value []string `json:"value" url:"value"`
}

type BlockManifestAttrPalette struct {
	BlockManifestAttr
	Value string `json:"value" url:"value"`
	Property string `json:"property" url:"property"`
	StylePath string `json:"stylePath" url:"stylePath"`
	PseudoElement string `json:"pseudo-element" url:"pseudo-element"`
	PseudoClass string `json:"pseudo-class" url:"pseudo-class"`
}

type BlockManifestAttrUrl struct {
	BlockManifestAttr
	Value string `json:"value" url:"value"`
	Property string `json:"disableBlocks" url:"disableBlocks"`
	StylePath string `json:"disableCustomURL" url:"disableCustomURL"`
}

const (
	BlockManifestAttrTypeText		string = "text"
	BlockManifestAttrTypeImage		string = "image"
	BlockManifestAttrTypeIcon		string = "icon"
	BlockManifestAttrTypeLink		string = "link"
	BlockManifestAttrTypeRangeSlier		string = "range-slider"
	BlockManifestAttrTypeDropdown		string = "dropdown"
	BlockManifestAttrTypeCheckbox		string = "checkbox"
	BlockManifestAttrTypeMultiselect	string = "multiselect"
	BlockManifestAttrTypeSlider		string = "slider"
	BlockManifestAttrTypePalette		string = "palette"
	BlockManifestAttrTypeSortableList	string = "sortable-list"
	BlockManifestAttrTypePosition		string = "position"
	BlockManifestAttrTypeUrl		string = "url"
)