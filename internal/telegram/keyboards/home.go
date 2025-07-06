package keyboards

import (
	"github.com/GoBotApiOfficial/gobotapi/types"
	"github.com/Laky-64/TestFlightTrackBot/internal/translator"
)

func Home(t *translator.Translator) *types.ReplyKeyboardMarkup {
	return &types.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]types.KeyboardButton{
			{
				{
					Text: t.T(translator.StartTracking),
				},
				{
					Text: t.T(translator.TrackingList),
				},
			},
		},
	}
}
