package commands

const (
	// Start is the command to start a conversation.
	Start string = "/start"
	// Accept is the command to accept an invitation.
	Accept string = "Принять приглашение"
	// Decline is the command to decline an invitation.
	Decline string = "К сожалению, не смогу"
	// Alone is the command to answer that a guest will come alone.
	Alone string = "Один/одна"
	// WithSomebody is the command to answer that a guest will come with someone.
	WithSomebody string = "С кем-то еще"
	// TransferNotNeeded is the command to answer that a guest does not need a transfer.
	TransferNotNeeded string = "Доберусь самостоятельно"
	// RostovTransferNeeded is the command to answer that a guest needs a transfer from Rostov.
	RostovTransferNeeded string = "Нужен из Ростова (и обратно)"
)
