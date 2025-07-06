package tracking

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	"github.com/Laky-64/TestFlightTrackBot/internal/telegram/core"
)

func TrackLink(ctx *core.UpdateContext, message types.Message) error {
	ctx.DB.PendingStore.Remove(message.Chat.ID)
	if linkID, err := ctx.DB.LinkStore.Add(message.Text); err != nil {
		return err
	} else {
		return addTrackingLink(ctx, message.Chat.ID, linkID)
	}
}

func SearchLink(ctx *core.UpdateContext, message types.Message) error {
	fmt.Println(ctx.DB.AppStore.FindByName(message.Text))
	return nil
}
