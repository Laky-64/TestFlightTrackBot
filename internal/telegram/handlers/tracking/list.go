package tracking

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	"github.com/Laky-64/TestFlightTrackBot/internal/telegram/core"
)

func List(ctx *core.UpdateContext, message types.Message) error {
	fmt.Println(ctx.DB.ChatLinkStore.TrackedList(message.Chat.ID))
	return nil
}
