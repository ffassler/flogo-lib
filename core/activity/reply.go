package activity

// ReplyHandler is used to reply back to whoever started the flow instance
type ReplyHandler interface {

	// Reply is used to reply with the results of the instance execution
	Reply(replyCode int, replyData map[string]interface{}, err error)
}
