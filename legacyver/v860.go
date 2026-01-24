package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/akmalfairuz/legacy-version/mapping"
)

const (
	// ItemVersion860 ...
	ItemVersion860 = 251
	// BlockVersion860 ...
	BlockVersion860 int32 = (1 << 24) | (21 << 16) | (120 << 8)
)

var (
	//go:embed data/required_item_list_859.json
	requiredItemList860 []byte
	//go:embed data/block_states_859.nbt
	blockStateData860 []byte
)

func New860(dragonflyMapping bool) *Protocol {
	itemMapping := mapping.NewItemMapping(requiredItemList860, ItemVersion860)
	blockTranslator := lookupOrCreateBlockTranslator(860, BlockVersion860, blockStateData860)
	return &Protocol{
		ver:             "1.21.124",
		id:              proto.ID860,
		blockTranslator: blockTranslator,
		itemTranslator:  NewItemTranslator(itemMapping, itemMappingLatest(dragonflyMapping), blockTranslator.BlockMapping(), blockMappingLatest),
	}
}
