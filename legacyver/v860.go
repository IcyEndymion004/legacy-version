package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/akmalfairuz/legacy-version/mapping"
)

func New860(dragonflyMapping bool) *Protocol {
	itemMapping := mapping.NewItemMapping(requiredItemList859, ItemVersion859)
	blockTranslator := lookupOrCreateBlockTranslator(860, BlockVersion859, blockStateData859)
	return &Protocol{
		ver:             "1.21.124",
		id:              proto.ID860,
		blockTranslator: blockTranslator,
		itemTranslator:  NewItemTranslator(itemMapping, itemMappingLatest(dragonflyMapping), blockTranslator.BlockMapping(), blockMappingLatest),
	}
}
